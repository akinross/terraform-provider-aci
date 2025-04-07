package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"
)

const (
	pubhupDevnetBaseUrl         = "https://pubhub.devnetcloud.com/media/model-doc-latest/docs"
	metaPath                    = "./meta"
	definitionsPath             = "./definitions"
	migrationSchemaDumpFileName = "schema-git-commit-e21fb3e5.json"
)

var metaDataMap map[string]*MetaData
var globalMetaDefinition GlobalMetaDefinition
var migrationDefinition map[string]interface{}

type MetaData struct {
	ClassName          string
	Comment            string
	Name               string
	ResourceClassName  string
	ResourceName       string
	ResourceNameNested string
	RnFormat           string
	Package            string
	Platform           string
	Versions           string
	IncludedChildren   []Child
	ContainedBy        []string
	Children           []string
	DnFormats          []string
	IdentifiedBy       []string
	SingleNested       bool
	AllowDelete        bool
	Migration          bool
	Definition         MetaDefinition
	Properties         map[string]*Property
	RequiredProperties []string
	OptionalProperties []string
	ReadOnlyProperties []string
	// Addtional HAS booleans shoudl be re-evaluated and added here
}

type Child struct {
	ClassName     string
	PointsToClass string
	Relational    bool
}

type Property struct {
	Name          string
	AttributeName string
	Required      bool
	Optional      bool
	Computed      bool
	ValueType     string
	CustomType    bool
	DefaultValue  string
	Documentation string
	Validators    []Validator
	ValidValues   []string
	Versions      string
	PointsToClass string
	TestValues    TestValues
	Migrations    map[int]MigrationValue
	// Addtional HAS booleans shoudl be re-evaluated and added here
}

type MigrationValue struct {
	Type     string
	Required bool
	Optional bool
	Computed bool
}

type TestValues struct {
	Initial []string
	Changed []string
}

type Validator struct {
	Min       float64
	Max       float64
	RegexList []RegexStatement
}

type RegexStatement struct {
	Regex string
	Type  string
}

func main() {
	initLogger()
	initMeta()
}

// Utils

// Split the class name into name space and name
func splitClassName(className string) (string, string) {
	log.Printf("Splitting class name '%s' for name space separation", className)
	for index, character := range className {
		if unicode.IsUpper(character) {
			return className[:index], className[index:]
		}
	}
	log.Fatal("Splitting of className failed. Invalid class name: " + className)
	panic("Splitting of className failed. Invalid class name: " + className)
}

func setNestedResourceName(metaData *MetaData) {

	log.Printf("Setting nested resource name for class: %s, %d", metaData.ClassName, len(metaData.IdentifiedBy))
	if len(metaData.IdentifiedBy) != 0 {
		for _, pluralNotation := range globalMetaDefinition.PluralResourceNaming {
			if strings.HasSuffix(metaData.ResourceName, pluralNotation.Single) {
				metaData.ResourceNameNested = strings.Replace(metaData.ResourceName, pluralNotation.Single, pluralNotation.Plural, 1)
			}
		}
	} else {
		metaData.ResourceNameNested = metaData.ResourceName
	}

	m := regexp.MustCompile("from_(.*)_to_")
	if m.MatchString(metaData.ResourceNameNested) {
		metaData.ResourceNameNested = m.ReplaceAllString(metaData.ResourceNameNested, "to_")
	}
}

// Reused from https://github.com/buxizhizhoum/inflection/blob/master/inflection.go#L8 to avoid importing the whole package
func Underscore(s string) string {
	for _, reStr := range []string{`([A-Z]+)([A-Z][a-z])`, `([a-z\d])([A-Z])`} {
		re := regexp.MustCompile(reStr)
		s = re.ReplaceAllString(s, "${1}_${2}")
	}
	return strings.ToLower(s)
}

func cleanDocumentationString(docString string) string {
	// Not all comments end with a dot, this is added to ensure the comment is correctly formatted.
	if len(docString) > 0 && docString[len(docString)-1:] != "." {
		docString = fmt.Sprintf("%s.", docString)
	}
	return docString
}

// File Operations

func getFileNames(path string) []string {
	var names []string
	entries, err := os.ReadDir(path)
	if err == nil && len(entries) > 0 {
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
	}
	return names
}

// Set the logger implementation

func initLogger() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	os.Remove("generator.log")
	f, err := os.OpenFile("generator.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(io.MultiWriter(os.Stdout, f))
}

// Meta Retrieval

func initMeta() {
	// Initialize the global class data map that contains all the information required for generating the provider
	metaDataMap = make(map[string]*MetaData)

	// Initialize the global meta definition for settings and generic overwrites
	globalMetaDefinition = getGlobalMetaDefinition(fmt.Sprintf("%s/global.yaml", definitionsPath))
	migrationDefinition = getMigrationDefinition(fmt.Sprintf("%s/%s", definitionsPath, migrationSchemaDumpFileName))

	// Retrieve new versions of the metadata when the environment 'RE_GEN_CLASSES' variable is set
	refreshMetadata()

	// Retrieve the metadata for the classes specified in the environment 'GEN_CLASSES' variable
	classNames := strings.Split(os.Getenv("GEN_CLASSES"), ",")
	if classNames[0] != "" {
		log.Printf("Retrieving meta files for classes: %s", classNames)
		retrieveMetadata(classNames)
	}

	// Load the metadata for all the classes in the meta folder
	loadMetaData()
	generateMigrationData()
	generateTestData()

}

func refreshMetadata() {
	classes, err := strconv.ParseBool(os.Getenv("RE_GEN_CLASSES"))
	if err == nil && classes {
		fileNames := getFileNames(metaPath)
		log.Printf("Refreshing metadata for %d classes.", len(fileNames))
		retrieveMetadata(fileNames)
	}
}

func retrieveMetadata(classNames []string) {

	var url string
	host := os.Getenv("GEN_HOST")

	for _, className := range classNames {

		if strings.HasSuffix(className, ".json") {
			className = strings.Replace(className, ".json", "", 1)
			log.Printf("Retrieving metadata for class: %s", className)
		}

		// When the metadata for a class is already retrieved during refresh, skip it
		if _, ok := metaDataMap[className]; ok {
			log.Printf("Metadata for class '%s' already retrieved during refresh. Skipping...", className)
			continue
		}

		packageName, name := splitClassName(className)
		initMetaData(className, packageName, name)

		if host == "" {
			url = fmt.Sprintf("%s/doc/jsonmeta/%s/%s.json", pubhupDevnetBaseUrl, packageName, name)
		} else {
			url = fmt.Sprintf("https://%s/doc/jsonmeta/%s/%s.json", host, packageName, name)
		}

		log.Printf("Retrieving meta data from '%s'", url)

		client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		res, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		outputFile, err := os.Create(fmt.Sprintf("%s/%s.json", metaPath, className))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Writing meta data to '%s'", outputFile.Name())

		outputFile.Write(resBody)
	}
}

func loadMetaData() {
	files, err := os.ReadDir(metaPath)
	if err != nil {
		log.Fatal("Error when reading file: ", err)
	}
	for _, file := range files {
		log.Printf("Loading metadata for class: %s", file.Name())
		fileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", metaPath, file.Name()))
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		// TODO investigate Unmarshalling in to struct
		var classInfo map[string]interface{}
		err = json.Unmarshal(fileContent, &classInfo)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		setMetaData(file.Name(), classInfo)
	}
}

func setMetaData(fileName string, classInfo map[string]interface{}) {
	packageName, name := splitClassName(strings.Replace(fileName, ".json", "", 1))
	className := fmt.Sprintf("%s%s", packageName, name)

	var classDetails map[string]interface{}
	if details, ok := classInfo[fmt.Sprintf("%s:%s", packageName, name)]; ok {
		classDetails = details.(map[string]interface{})
	}

	if _, ok := metaDataMap[className]; !ok {
		initMetaData(className, packageName, name)
	}

	metaData := metaDataMap[className]

	if metaData.Definition.ResourceName != "" {
		metaData.ResourceName = metaData.Definition.ResourceName
	} else if label, ok := classDetails["label"]; ok {
		metaData.ResourceName = Underscore(label.(string))
	} else {
		log.Fatal("Label is not found in meta file and resource name not provided as overwrite.")
	}

	if rnFormat, ok := classDetails["rnFormat"]; ok {
		if metaData.Definition.RnPrepend == "" {
			metaData.RnFormat = rnFormat.(string)
		} else {
			metaData.RnFormat = fmt.Sprintf("%s/%s", metaData.Definition.RnPrepend, rnFormat.(string))
		}
	}

	if len(metaData.Definition.DnFormats) != 0 {
		metaData.DnFormats = metaData.Definition.DnFormats
	} else if dnFormats, ok := classDetails["dnFormats"]; ok {
		for _, dnFormat := range dnFormats.([]interface{}) {
			if !slices.Contains(metaData.DnFormats, dnFormat.(string)) {
				metaData.DnFormats = append(metaData.DnFormats, dnFormat.(string))
			}
		}
	} else {
		log.Fatal("DnFormat is not found in meta file and not provided as overwrite.")
	}

	if identifiedBy, ok := classDetails["identifiedBy"]; ok {
		for _, identifier := range identifiedBy.([]interface{}) {
			if !slices.Contains(metaData.IdentifiedBy, identifier.(string)) {
				metaData.IdentifiedBy = append(metaData.IdentifiedBy, identifier.(string))
			}
		}
		if len(metaData.IdentifiedBy) == 0 {
			metaData.SingleNested = true
		}
	}

	setNestedResourceName(metaData)

	if isCreatableDeletable, ok := classDetails["isCreatableDeletable"].(string); ok && isCreatableDeletable == "never" {
		metaData.AllowDelete = false
	}

	if len(metaData.Definition.ContainedBy) > 0 {
		for _, containsByClass := range metaData.Definition.ContainedBy {
			if !slices.Contains(metaData.ContainedBy, containsByClass) && !slices.Contains(globalMetaDefinition.ExcludedClasses, containsByClass) {
				metaData.ContainedBy = append(metaData.ContainedBy, containsByClass)
			}
		}
	} else if containsByClasses, ok := classDetails["containedBy"]; ok {
		for containsByClass := range containsByClasses.(map[string]interface{}) {
			className := strings.ReplaceAll(containsByClass, ":", "")
			if !slices.Contains(metaData.ContainedBy, className) && !slices.Contains(globalMetaDefinition.ExcludedClasses, className) {
				metaData.ContainedBy = append(metaData.ContainedBy, className)
			}
		}
	}

	if containsClasses, ok := classDetails["contains"]; ok {
		for containedClass := range containsClasses.(map[string]interface{}) {
			if !slices.Contains(metaData.Children, containedClass) {
				metaData.Children = append(metaData.Children, strings.ReplaceAll(containedClass, ":", ""))
			}
		}
	}

	if relationalClasses, ok := classDetails["relationTo"]; ok {
		for relationalClass, pointerClass := range relationalClasses.(map[string]interface{}) {
			metaData.IncludedChildren = append(
				metaData.IncludedChildren,
				Child{
					ClassName:     strings.ReplaceAll(relationalClass, ":", ""),
					PointsToClass: strings.ReplaceAll(pointerClass.(string), ":", ""),
					Relational:    true,
				},
			)
		}
	}

	metaData.Platform = "both"
	if platform, ok := classDetails["platformFlavors"]; ok && len(platform.([]interface{})) != 0 {
		if platform.([]interface{})[0].(string) == "apic" {
			metaData.Platform = "apic"
		} else if platform.([]interface{})[0].(string) == "capic" {
			metaData.Platform = "cloud"
		}
	}

	for _, child := range globalMetaDefinition.AlwaysIncludeAsChildren {
		if slices.Contains(metaData.Children, child) {
			metaData.IncludedChildren = append(
				metaData.IncludedChildren,
				Child{
					ClassName: child,
				},
			)
		}
	}

	if versions, ok := classDetails["versions"]; ok {
		metaData.Versions = versions.(string)
	} else {
		log.Printf("Versions is not found for class '%s' in meta file and not provided as overwrite.", className)
	}

	if properties, ok := classDetails["properties"]; ok {

		relationInfo := make(map[string]interface{})
		if relation, ok := classDetails["relationInfo"]; ok {
			relationInfo = relation.(map[string]interface{})
		}
		setProperties(className, properties.(map[string]interface{}), relationInfo)
	}

	log.Printf("Metadata loaded successfully from '%s'.", fileName)
}

func initMetaData(className, packageName, name string) {
	metaDataMap[className] = &MetaData{
		ClassName:         className,
		ResourceClassName: fmt.Sprintf("%s%s", strings.ToUpper(className[:1]), className[1:]),
		Package:           packageName,
		Name:              name,
		Definition:        getMetaDefinition(fmt.Sprintf("%s/%s.yaml", definitionsPath, className)),
		Properties:        make(map[string]*Property, 0),
	}
}

func setProperties(className string, properties map[string]interface{}, relationInfo map[string]interface{}) {

	log.Printf("MetaDataMap resourceName: %v", metaDataMap[className].ResourceName)
	for propertyName, v := range properties {

		propertyOverwriteValues := metaDataMap[className].Definition.Properties[propertyName]

		propertyValue := v.(map[string]interface{})

		if (propertyValue["isConfigurable"] == true || propertyOverwriteValues.ReadOnly) &&
			!propertyOverwriteValues.Exclude &&
			!slices.Contains(globalMetaDefinition.IgnoreProperties, propertyName) {

			property := Property{
				Name: fmt.Sprintf("%s%s", strings.ToUpper(propertyName[0:1]), propertyName[1:]),
			}

			// Named relationships are typically defined by the prefix 'tn', className and suffix 'Name'
			// This conditional determines if it is targetting a class name and deduces the class name from the property name
			// This is not retrieved from relationInfo as only 1 is available in the meta file where sometimes a class can have multiple pointers to classes
			if strings.HasPrefix(propertyName, "tn") && strings.HasSuffix(propertyName, "Name") {
				// Get the class name and covert to start with lower case for resource name lookup.
				// - Example: tnVzOOBBrCPName -> VzOOBBrCP -> vzOOBBrCP
				className := propertyName[2 : len(propertyName)-4]
				property.PointsToClass = strings.ToLower(className[:1]) + className[1:]
			} else if propertyName == "tDn" {
				// This is a special case where the property name is tDn and the class name is in the relationInfo
				if relationClassName, ok := relationInfo["toMo"]; ok {
					property.PointsToClass = strings.ReplaceAll(relationClassName.(string), ":", "")
				} else {
					// No need for overwrite yet could be added in the future
					log.Fatal("tDn property is not found in relationInfo.")
				}
			}

			if globalMetaDefinition.Overwrites.PropertyName[propertyName] != "" {
				property.AttributeName = globalMetaDefinition.Overwrites.PropertyName[propertyName]
			} else if propertyOverwriteValues.OverwriteName != "" {
				property.AttributeName = propertyOverwriteValues.OverwriteName
			} else if property.PointsToClass != "" {
				if strings.HasSuffix(propertyName, "Name") {
					property.AttributeName = "name"
				} else {
					log.Fatal("Property name for class pointer is not found in meta file and not provided as overwrite.")
				}
			} else {
				property.AttributeName = Underscore(propertyName)
			}

			if propertyOverwriteValues.ValueType != "" {
				property.ValueType = propertyOverwriteValues.ValueType
			} else if valueType, ok := propertyValue["uitype"]; ok {
				property.ValueType = valueType.(string)
			} else {
				log.Fatal("Value type is not found in meta file and not provided as overwrite.")
			}

			if propertyOverwriteValues.Required || propertyValue["isNaming"].(bool) {
				metaDataMap[className].RequiredProperties = append(metaDataMap[className].RequiredProperties, property.AttributeName)
				property.Required = true
			} else if propertyOverwriteValues.ReadOnly {
				metaDataMap[className].ReadOnlyProperties = append(metaDataMap[className].ReadOnlyProperties, property.AttributeName)
				property.Computed = true
			} else {
				metaDataMap[className].OptionalProperties = append(metaDataMap[className].OptionalProperties, property.AttributeName)
				property.Optional = true
				property.Computed = true
			}

			if propertyOverwriteValues.Documentation != "" {
				property.Documentation = propertyOverwriteValues.Documentation
			} else if value, ok := propertyValue["comment"]; ok {
				space := regexp.MustCompile(`\s+`)
				var comment string
				for _, details := range value.([]interface{}) {
					comment = comment + details.(string)
				}
				property.Documentation = space.ReplaceAllString(comment, " ")
			} else if value, ok := propertyValue["label"]; ok {
				property.Documentation = value.(string)
			} else {
				log.Fatal("Documentation is not found in meta file and not provided as overwrite.")
			}
			property.Documentation = cleanDocumentationString(property.Documentation)

			// No need for overwrite yet could be added in the future
			if value, ok := propertyValue["validators"]; ok {
				for _, validatorDetails := range value.([]interface{}) {
					validator := Validator{}
					if min, ok := validatorDetails.(map[string]interface{})["min"]; ok {
						validator.Min = min.(float64)
					}
					if max, ok := validatorDetails.(map[string]interface{})["max"]; ok {
						validator.Max = max.(float64)
					}
					if regexs, ok := validatorDetails.(map[string]interface{})["regexs"]; ok {
						for _, regex := range regexs.([]interface{}) {
							regexStatement := RegexStatement{}
							if regexValue, ok := regex.(map[string]interface{})["regex"]; ok {
								regexStatement.Regex = regexValue.(string)
							}
							if regexType, ok := regex.(map[string]interface{})["type"]; ok {
								regexStatement.Type = regexType.(string)
							}
							validator.RegexList = append(validator.RegexList, regexStatement)
						}
					}
					property.Validators = append(property.Validators, validator)
				}
			}

			if value, ok := propertyValue["validValues"]; ok {
				for _, validValue := range value.([]interface{}) {
					localName := validValue.(map[string]interface{})["localName"].(string)
					if localName != "defaultValue" && !slices.Contains(propertyOverwriteValues.RemoveValidValues, localName) {
						property.ValidValues = append(property.ValidValues, localName)
					}
				}
				for _, validValue := range propertyOverwriteValues.AddValidValues {
					if !slices.Contains(property.ValidValues, validValue) {
						property.ValidValues = append(property.ValidValues, validValue)
					}
				}
			}

			if len(property.ValidValues) > 0 && len(property.Validators) > 0 {
				property.CustomType = true
			}

			if propertyOverwriteValues.DefaultValue != "" {
				property.DefaultValue = propertyOverwriteValues.DefaultValue
			} else if value, ok := propertyValue["default"]; ok {
				if defaultValue, ok := value.(string); ok {
					property.DefaultValue = defaultValue
				} else if defaultValue, ok := value.(float64); ok {
					property.DefaultValue = fmt.Sprintf("%f", defaultValue)
				} else {
					log.Fatal("Default value type is not defined in function, please adjust the code to allow casting.")
				}
			}

			if propertyOverwriteValues.Versions != "" {
				property.Versions = propertyOverwriteValues.Versions
			} else if value, ok := propertyValue["versions"]; ok {
				property.Versions = value.(string)
			} else {
				log.Printf("Versions is not found for property '%s' in meta file and not provided as overwrite.", propertyName)
			}

			metaDataMap[className].Properties[propertyName] = &property
		}
	}
}

// Meta Definition type

type MetaDefinition struct {
	Documentation struct {
		SubCategory string
		UiLocations []string
	}
	Migration []struct {
		Blocks  map[string]map[string]string // TODO make this a struct
		Version int
	}
	Properties map[string]struct {
		Exclude           bool
		DefaultValue      string
		Documentation     string
		OverwriteName     string
		ReadOnly          bool
		RemoveValidValues []string
		AddValidValues    []string
		Required          bool
		TestValues        struct {
			Initial []string
			Changed []string
		}
		MigrationMapping struct {
			Type     string
			Name     string
			Required bool
			Optional bool
			Computed bool
		}
		ValueType string
		Versions  string
	}
	ResourceName string
	Reuired      bool
	DnFormats    []string
	ContainedBy  []string
	RnPrepend    string
}

func getMetaDefinition(fileName string) MetaDefinition {

	var definitionMetaData MetaDefinition

	definition, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("A file '%s' is required to be defined in the definitions folder.", fileName)
		log.Fatal(err)
	}

	err = yaml.Unmarshal(definition, &definitionMetaData)
	if err != nil {
		log.Fatal(err)
	}

	return definitionMetaData
}

// Global Definition type

type GlobalMetaDefinition struct {
	ExcludedClasses         []string
	AlwaysIncludeAsChildren []string
	IgnoreProperties        []string
	PluralResourceNaming    []struct {
		Single string
		Plural string
	}
	Overwrites struct {
		Naming        map[string]string
		Documentation map[string]string
		PropertyName  map[string]string
	}
	Settings struct {
		ExampleAmount  string
		ParentDnAmount string
	}
}

func getGlobalMetaDefinition(fileName string) GlobalMetaDefinition {

	var definitionGlobalMetaData GlobalMetaDefinition

	definition, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("A file '%s' is required to be defined in the definitions folder.", fileName)
		log.Fatal(err)
	}

	err = yaml.Unmarshal(definition, &definitionGlobalMetaData)
	if err != nil {
		log.Fatal(err)
	}

	return definitionGlobalMetaData
}

func getMigrationDefinition(fileName string) map[string]interface{} {

	definitionMigrationData := make(map[string]interface{})

	definition, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("A file '%s' is required to be defined in the definitions folder.", fileName)
		log.Fatal(err)
	}

	err = json.Unmarshal(definition, &definitionMigrationData)
	if err != nil {
		log.Fatal(err)
	}

	return definitionMigrationData
}

func generateMigrationData() {
	log.Printf("Generating migration data for classes.")

	for className, metaData := range metaDataMap {
		if len(metaData.Definition.Migration) != 0 {
			metaData.Migration = true
		}
		if metaData.Migration {
			log.Printf("Generating migration data for class: %s", className)

			migrationMetaData := migrationDefinition["provider_schemas"].(map[string]interface{})["registry.terraform.io/ciscodevnet/aci"].(map[string]interface{})["resource_schemas"].(map[string]interface{})[fmt.Sprintf("aci_%s", metaData.ResourceName)].(map[string]interface{})
			version := int(migrationMetaData["version"].(float64))
			block := migrationMetaData["block"].(map[string]interface{})
			attributes := block["attributes"].(map[string]interface{})

			log.Printf("Migration data for class '%s' version '%d':", className, version)
			log.Printf("migrationMetaData: %v", migrationMetaData)
			log.Printf("attributes: %v", attributes)
		}
	}
	log.Printf("Migration data generation completed.")
}

func generateTestData() {
	log.Printf("Generating test data for all classes.")
	for className, metaData := range metaDataMap {
		log.Printf("Generating test data for class: %s", className)

		for propertyName, property := range metaData.Properties {

			if !property.Required && !property.Optional && property.Computed {
				continue
			}

			initial := metaData.Definition.Properties[propertyName].TestValues.Initial
			if len(initial) > 0 {
				property.TestValues.Initial = initial
			} else if propertyName == "tDn" {
				property.TestValues.Initial = []string{createReferenceToResource(property.PointsToClass, "initial", "id")}
			} else if property.PointsToClass != "" {
				property.TestValues.Initial = []string{createReferenceToResource(property.PointsToClass, "initial", property.AttributeName)}
			} else if len(property.ValidValues) > 0 {
				property.TestValues.Initial = property.ValidValues[0:1]
			} else if property.ValueType == "string" {
				property.TestValues.Initial = []string{fmt.Sprintf("initial_%s", property.AttributeName)}
			} else {
				property.TestValues.Initial = []string{property.AttributeName}
			}

			changed := metaData.Definition.Properties[propertyName].TestValues.Changed
			if len(changed) > 0 {
				property.TestValues.Changed = changed
			} else if propertyName == "tDn" {
				property.TestValues.Changed = []string{createReferenceToResource(property.PointsToClass, "changed", "id")}
			} else if property.PointsToClass != "" {
				property.TestValues.Changed = []string{createReferenceToResource(property.PointsToClass, "changed", property.AttributeName)}
			} else if len(property.ValidValues) > 1 {
				property.TestValues.Changed = property.ValidValues[1:2]
			} else if property.ValueType == "string" {
				property.TestValues.Changed = []string{fmt.Sprintf("changed_%s", property.AttributeName)}
			} else {
				property.TestValues.Changed = []string{property.AttributeName}
			}
			log.Printf("Test values for class '%s' property '%s': %v", className, propertyName, property.TestValues)
		}

		prettyMetaData, _ := json.MarshalIndent(metaDataMap[className], "", "\t")
		log.Printf("MetaDataMap for class %s: %s", className, prettyMetaData)

	}
	log.Printf("Test data generation completed.")

}

func createReferenceToResource(className, configName, attributeName string) string {

	if metaData, ok := metaDataMap[className]; ok {
		return fmt.Sprintf("aci_%s.%s.%s", metaData.ResourceName, configName, attributeName)
	}
	log.Fatal(fmt.Printf("Class '%s' not found in meta data map. Cannot create reference.", className))
	return ""
}
