package config

type ServerSettingS struct {
	Port string
}

type DatabaseSettingS struct {
	UserName       string
	Password       string
	DriverName     string
	DBName         string
	Charset        string
	Local          string
	Domain         string
	Port           string
	ParseTime      string
	DataSourceName string
}

var sections = make(map[string]interface{})

// 将配置文件中的部分 读取入结构体中
func (s *Setting) ReadSection(k string, v interface{}) error {
	// 读取 k 对应的值到 v 中
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
