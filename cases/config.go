package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage/file"
	"github.com/xackery/xegony/storage/memory"
	"github.com/xeipuuv/gojsonschema"
)

//LoadConfigFromFileToMemory is ran during initialization
func LoadConfigFromFileToMemory() (err error) {

	fr, err := file.New("config", "_config.yml", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new file")
		return
	}

	err = Initialize("config-file", fr, fr, fr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize config-file")
		return
	}

	mr, err := memory.New("", nil, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create new memory")
		return
	}

	err = Initialize("config-memory", mr, mr, mr)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize config-memory")
		return
	}

	fileReader, err := getReader("config-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-file reader")
		return
	}

	memWriter, err := getWriter("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-memory writer")
		return
	}

	page := &model.Page{
		Limit: 100,
	}
	page.Total, err = fileReader.ListConfigTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to get list config count")
		return
	}
	page.Limit = page.Total

	configs, err := fileReader.ListConfig(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list configs")
		return
	}

	for _, config := range configs {
		err = memWriter.CreateConfig(config)
		if err != nil {
			err = errors.Wrap(err, "failed to create config")
			return
		}
	}

	fmt.Printf("%d configs, ", len(configs))
	return
}

//ListConfig lists all configs accessible by provided user
func ListConfig(page *model.Page, user *model.User) (configs []*model.Config, err error) {
	err = validateOrderByConfigField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for config")
		return
	}

	page.Total, err = reader.ListConfigTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list config toal count")
		return
	}

	configs, err = reader.ListConfig(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list config")
		return
	}
	for i, config := range configs {
		err = sanitizeConfig(config, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize config element %d", i)
			return
		}
	}
	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListConfigBySearch will request any config matching the pattern of name
func ListConfigBySearch(page *model.Page, config *model.Config, user *model.User) (configs []*model.Config, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list config by search without guide+")
		return
	}

	err = validateOrderByConfigField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre config")
		return
	}

	err = validateConfig(config, nil, []string{ //optional
		"key",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate config")
		return
	}
	reader, err := getReader("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-memory reader")
		return
	}

	configs, err = reader.ListConfigBySearch(page, config)
	if err != nil {
		err = errors.Wrap(err, "failed to list config by search")
		return
	}

	for _, config := range configs {
		err = sanitizeConfig(config, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize config")
			return
		}
	}

	err = sanitizeConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search config")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateConfig will create an config using provided information
func CreateConfig(config *model.Config, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list config by search without guide+")
		return
	}
	err = prepareConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare config")
		return
	}

	err = validateConfig(config, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate config")
		return
	}
	writer, err := getWriter("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for config")
		return
	}
	err = writer.CreateConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to create config")
		return
	}

	fileWriter, err := getWriter("config-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-file writer")
		return
	}
	err = fileWriter.CreateConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to create config-file")
		return
	}
	err = sanitizeConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize config")
		return
	}
	return
}

//GetConfigValue is a quickhand way to get information
func GetConfigValue(key string) (value string) {
	config := &model.Config{
		Key: key,
	}
	user := &model.User{}
	err := GetConfig(config, user)
	if err != nil {
		fmt.Println("Failed to get config", err.Error())
		//GetConfigValue never errors, it always returns an empty string on failure
		err = nil
		return
	}
	value = config.Value
	return
}

//GetConfigForMySQL returns a connection string for use of establishing a connection for MySQL
func GetConfigForMySQL() (value string) {
	value = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", GetConfigValue("mysqlUsername"), GetConfigValue("mysqlPassword"), GetConfigValue("mysqlHostname"), GetConfigValue("mysqlPort"), GetConfigValue("mysqlDatabase"))
	return
}

//GetConfigForHTTP returns a listen path for http.Listen
func GetConfigForHTTP() (value string) {
	value = fmt.Sprintf("%s:%s", GetConfigValue("httpHostname"), GetConfigValue("httpPort"))
	return
}

//GetConfig gets an config by provided configID
func GetConfig(config *model.Config, user *model.User) (err error) {
	err = prepareConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare config")
		return
	}

	err = validateConfig(config, []string{"key"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate config")
		return
	}

	reader, err := getReader("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-memory reader")
		return
	}

	err = reader.GetConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to get config")
		return
	}

	err = sanitizeConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize config")
		return
	}

	return
}

//EditConfig edits an existing config
func EditConfig(config *model.Config, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list config by search without guide+")
		return
	}
	err = prepareConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare config")
		return
	}

	err = validateConfig(config,
		[]string{"key"}, //required
		[]string{ //optional
			"category",
			"description",
			"value",
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate config")
		return
	}
	writer, err := getWriter("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for config")
		return
	}
	err = writer.EditConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to edit config")
		return
	}

	fileWriter, err := getWriter("config-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-file writer")
		return
	}
	err = fileWriter.EditConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to edit config-file")
		return
	}

	err = sanitizeConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize config")
		return
	}
	return
}

//DeleteConfig deletes an config by provided configID
func DeleteConfig(config *model.Config, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete config without admin+")
		return
	}
	err = prepareConfig(config, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare config")
		return
	}

	err = validateConfig(config, []string{"key"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate config")
		return
	}
	writer, err := getWriter("config-memory")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-memory writer")
		return
	}
	err = writer.DeleteConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to delete config")
		return
	}

	fileWriter, err := getWriter("config-file")
	if err != nil {
		err = errors.Wrap(err, "failed to get config-file writer")
		return
	}
	err = fileWriter.DeleteConfig(config)
	if err != nil {
		err = errors.Wrap(err, "failed to delete config-file")
		return
	}

	return
}

func prepareConfig(config *model.Config, user *model.User) (err error) {
	if config == nil {
		err = fmt.Errorf("empty config")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateConfig(config *model.Config, required []string, optional []string) (err error) {
	schema, err := newSchemaConfig(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(config))
	if err != nil {
		return
	}

	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	return
}

func validateOrderByConfigField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"key",
		"category",
	}

	possibleNames := ""
	for _, name := range validNames {
		if page.OrderBy == name {
			return
		}
		possibleNames += name + ", "
	}
	if len(possibleNames) > 0 {
		possibleNames = possibleNames[0 : len(possibleNames)-2]
	}

	err = &model.ErrValidation{
		Message: "orderBy is invalid. Possible fields: " + possibleNames,
		Reasons: map[string]string{
			"orderBy": "field is not valid",
		},
	}
	return
}

func sanitizeConfig(config *model.Config, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = nil
	}
	return
}

func newSchemaConfig(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyConfig(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyConfig(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func getSchemaPropertyConfig(field string) (prop model.Schema, err error) {
	switch field {
	case "key":
		prop.Type = "string"
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "value":
		prop.Type = "string"
		prop.MaxLength = 30
	case "category":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "description":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 128
		prop.Pattern = "^[a-zA-Z]*$"
	case "icon":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
