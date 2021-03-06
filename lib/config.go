package lib

import (
	"fmt"
    "os"
	"strings"
	configparser "github.com/alyu/configparser"
)

var specChineseConfig = SpecText{

	synopsisText: "创建配置文件用以存储配置项",

	paramText: "[options]",

	syntaxText: ` 
    ossutil config [-e endpoint] [-I id] [-K key] [-T token] [-L language] [-c file] 
`,

	detailHelpText: ` 
    该命令创建配置文件，将用户设置的配置项信息存储进该配置文件，配置项用
    以访问OSS时提供访问信息（某命令是否需要配置项，参见其是否支持
    --config_file选项，具体可见该命令的帮助）。

    配置文件路径可由用户指定，默认为` + DefaultConfigFile + `。如果配置文件存在，假
    设其为:a，ossutil会将文件a另存为：a.bak，然后重新创建文件a并写入配置，
    此时，如果a.bak存在，其会被文件a覆盖。

用法:

    该命令有两种用法，交互式1)和非交互式2)，推荐用法为交互式，因为交互
    式用法拥有更好的安全性。

    1) ossutil config [-c file]
        该用法提供一种交互式的方法来配置信息，ossutil交互式地询问用户如下
    信息：
        (1) config file
            配置文件路径，如果用户键入回车，ossutil会使用默认的配置文件：
        ` + DefaultConfigFile + `。
        (2) language
            当首次配置（配置文件不存在）时，ossutil会向用户询问语言设置，可选
        为中文或者英文，如果不配置ossutil将根据用户输入的language选项配置，如
        果此时用户也未输入language选项，将配置成默认语言中文。
            如果配置文件已存在，ossutil会综合用户输入的language选项和配置文件
        中的语言信息，配置该项，而不会询问。
            ossutil在运行时会从配置文件中读取该language选项，如果该选项不存在
        或者非法，将采用默认语言：` + DefaultLanguage + `。
        (3) endpoint, accessKeyID, accessKeySecret
            回车代表着跳过相应配置项的设置。注意：endpoint应该为一个二级域
        名(SLD)。  
        (4) bucket-endpoint
            ossutil询问用户是否有bucket-endpoint配对，请输入'y'或者'n'来进行
        配置或者跳过配置。如果用户在输入bucket信息时键入回车，则代表着结束
        bucket-endpoint的配置。注意：此处的endpoint应该为一个二级域名。
            如果配置了bucket-endpoint选项，当对某bucket进行操作时，ossutil会
        在该选项中寻找该bucket对应的endpoint，如果找到，该endpoint会覆盖基本
        配置中endpoint。
        (5) bucket-cname
            与bucket-endpoint配置类似。
            如果配置了bucket-endpoint选项，当对某bucket进行操作时，ossutil会
        在该选项中寻找该bucket对应的endpoint，如果找到，则找到的endpoint会覆
        盖bucket-endpoint选项和基本配置中的endpoint。
        
        即优先级：bucket-cname > bucket-endpoint > endpoint > 默认endpoint

    2) ossutil config options
        如果用户使用命令时输入了除--config_file之外的任何选项，则该命令进入非
    交互式模式。所有的配置项应当使用选项指定。


配置文件格式：

    [Credentials]
        endpoint = ` + DefaultEndpoint + `
        accessKeyID = your_key_id
        accessKeySecret = your_key_secret
        stsToken = your_sts_token
        language = 中文
    [Bucket-Endpoint]
        bucket1 = endpoint1
        bucket2 = endpoint2
        ...
    [Bucket-Cname]
        bucket3 = cname1
        bucket4 = cname2
        ...
`,

	sampleText: ` 
    ossutil config
    ossutil config -e oss-cn-hangzhou.aliyuncs.com -c ~/.myconfig
`,
}

var specEnglishConfig = SpecText{

	synopsisText: "Create configuration file to store credentials",

	paramText: "[options]",

	syntaxText: ` 
    ossutil config [-e endpoint] [-I id] [-K key] [-T token] [-L language] [-c file] 
`,

	detailHelpText: ` 
    The command create a configuration file and stores credentials
    information user specified. Credentials information is used when
    access OSS(if a command supports --config_file option, then the 
    information is useful to the command).

    The configuration file can be specified by user, which in default
    is ` + DefaultConfigFile + `. If the configuration file exist, suppose the file 
    is: a, ossutil will save a as a.bak, and rewrite file a, at this time, 
    if file a.bak exists, a.bak will be rewrited.

Useage:

    There are two useages for the command, one is interactive(shows
    in 1) ), which is recommended because of safety problem. another is
    non interactive(shows in 2) ).

    1) ossutil config [-c file]
        The useage provides an interactive way to configure credentials.
    Interactively ossutil asks you for:
        (1) config file
            If user enter carriage return, ossutil use the default file.
        (2) language
            When configure for the first time(config file not exit), ossutil 
        will ask user to set the language(support Chinese or English at this 
        time), if user did not input the language, ossutil will set language 
        to the value of --language option, if user did not specify --language 
        option, ossutil will set it to Chinese.
            If config file exists, ossutil will set the language according to 
        --language option and configuration in config file, instead of asking 
        user for it.
            ossutil will read the language configuration when run command, if 
        the configuration does not exist or is invalid, ossutil will show in 
        default language: ` + DefaultLanguage + `.
        (3) endpoint, accessKeyID, accessKeySecret
            Carriage return means skip the configuration of these options.
        Notice that endpoint means a second-level domain(SLD).
        (4) bucket-endpoint
            ossutil ask you if there are any bucket-endpoint pairs, please
        enter 'y' or 'n' to configure the pairs or skip. If you enter carriage
        return when configure bucket, it means the pairs' configuration is
        ended. Notice that endpoint means a second-level domain(SLD).
            When access a bucket, ossutil will search for endpoint corresponding 
        to the bucket in this section, if found, the endpoint has priority over 
        the endpoint in the base section.
        (5) bucket-cname
            Similar to bucket-endpoint configuration.
            When access a bucket, ossutil will search for endpoint corresponding 
        tothe bucket in this section, if found, the endpoint has priority over 
        the endpoint in bucket-endpoint and the endpoint in the base section.

        PRI: bucket-cname > bucket-endpoint > endpoint > default endpoint

    2) ossutil config options
        If any options except --config_file is specified, the command enter
    the non interactive mode. All the configurations should be specified by
    options.


Credential File Format:

    [Credentials]
        endpoint = ` + DefaultEndpoint + ` 
        accessKeyID = your_key_id
        accessKeySecret = your_key_secret
        stsToken = your_sts_token
    [Bucket-Endpoint]
        bucket1 = endpoint1
        bucket2 = endpoint2
        ...
    [Bucket-Cname]
        bucket3 = cname1
        bucket4 = cname2
        ...
`,

	sampleText: ` 
    ossutil config
    ossutil config -e oss-cn-hangzhou.aliyuncs.com -c ~/.myconfig
`,
}

// ConfigCommand is the command config user's credentials information
type ConfigCommand struct {
	command Command
}

var configCommand = ConfigCommand{
	command: Command{
		name:        "config",
		nameAlias:   []string{"cfg", "config"},
		minArgc:     0,
		maxArgc:     0,
		specChinese: specChineseConfig,
		specEnglish: specEnglishConfig,
		group:       GroupTypeAdditionalCommand,
		validOptionNames: []string{
			OptionConfigFile,
			OptionEndpoint,
			OptionAccessKeyID,
			OptionAccessKeySecret,
			OptionSTSToken,
            OptionLanguage,
		},
	},
}

// function for RewriteLoadConfiger interface
func (cc *ConfigCommand) rewriteLoadConfig(configFile string) error {
    // read config file, if error exist, do not print error
    var err error
    if cc.command.configOptions, err = LoadConfig(configFile); err != nil {
        cc.command.configOptions = OptionMapType{}
    }
	return nil
}

// function for AssembleOptioner interface
func (cc *ConfigCommand) rewriteAssembleOptions() {
    // only assemble language option
	if val, _ := GetString(OptionLanguage, cc.command.options); val == "" {
        if val, ok := cc.command.configOptions[OptionLanguage]; ok {
            opval := val.(string)
            cc.command.options[OptionLanguage] = &opval
            delete(cc.command.configOptions, OptionLanguage)
        }
    }

    if val, _ := GetString(OptionLanguage, cc.command.options); val == "" {
        def := OptionMap[OptionLanguage].def
        cc.command.options[OptionLanguage] = &def
    }
}


// function for FormatHelper interface
func (cc *ConfigCommand) formatHelpForWhole() string {
	return cc.command.formatHelpForWhole()
}

func (cc *ConfigCommand) formatIndependHelp() string {
	return cc.command.formatIndependHelp()
}

// Init simulate inheritance, and polymorphism 
func (cc *ConfigCommand) Init(args []string, options OptionMapType) error {
	return cc.command.Init(args, options, cc)
}


// RunCommand simulate inheritance, and polymorphism 
func (cc *ConfigCommand) RunCommand() error {
	configFile, _ := GetString(OptionConfigFile, cc.command.options)
	delete(cc.command.options, OptionConfigFile)
    language, _ := GetString(OptionLanguage, cc.command.options)
    delete(cc.command.options, OptionLanguage)

	// filter user input options
	cc.filterNonInputOptions()

	var err error
	if len(cc.command.options) == 0 {
		err = cc.runCommandInteractive(configFile, language)
	} else {
		err = cc.runCommandNonInteractive(configFile, language)
	}
	return err
}

func (cc *ConfigCommand) filterNonInputOptions() {
	for name := range cc.command.options {
		if val, err := GetString(name, cc.command.options); err != nil || val == "" {
			delete(cc.command.options, name)
		}
	}
}

func (cc *ConfigCommand) runCommandInteractive(configFile, language string) error {
    llanguage := strings.ToLower(language)
    if llanguage == LEnglishLanguage {
	    fmt.Println("The command creates a configuration file and stores credentials.")
    } else {
        fmt.Println("该命令创建将一个配置文件，在其中存储配置信息。")
    }

	if configFile == "" {
        if llanguage == LEnglishLanguage {
		    fmt.Printf("\nPlease enter the config file path(default " + DefaultConfigFile + ", carriage return will use the default path):")
        } else {
		    fmt.Printf("\n请输入配置文件路径（默认为：" + DefaultConfigFile + "，回车将使用默认路径）：")
        }

		if _, err := fmt.Scanln(&configFile); err != nil {
            if llanguage == LEnglishLanguage {
			    fmt.Println("No config file entered, will use the default config file " + DefaultConfigFile + "\n")
            } else {
		        fmt.Println("未输入配置文件路径，将使用默认配置文件：" + DefaultConfigFile + "。\n")
            }
		}
	}

	configFile = DecideConfigFile(configFile)
    if llanguage == LEnglishLanguage {
	    fmt.Println("For the following settings, carriage return means skip the configuration. Please try \"help config\" to see the meaning of the settings.\n")
    } else {
        fmt.Println("对于下述配置，回车将跳过相关配置项的设置，配置项的具体含义，请使用\"help config\"命令查看。\n")
    }

	if err := cc.configInteractive(configFile, language); err != nil {
		return err
	}
	return nil
}

func (cc *ConfigCommand) configInteractive(configFile, language string) error {
	var val string
	config := configparser.NewConfiguration()
	section := config.NewSection(CREDSection)

    // if config file not exist, config Language
    llanguage := strings.ToLower(language)
	section.Add(OptionLanguage, language)
    if _, err := os.Stat(configFile); err != nil {
        if llanguage == LEnglishLanguage {
		    fmt.Printf("Please enter language(%s):", OptionMap[OptionLanguage].minVal)
        } else {
            fmt.Printf("请输入语言(%s)：", OptionMap[OptionLanguage].minVal)
        }
		if _, err := fmt.Scanln(&val); err == nil {
            vals := strings.Split(OptionMap[OptionLanguage].minVal, "|")
            if FindPosCaseInsen(val, vals) == -1 {
                return fmt.Errorf("invalid option value of %s, the value: %s is not anyone of %s", OptionLanguage, val, OptionMap[OptionLanguage].minVal)
            }
			section.Add(OptionLanguage, val)
		}
    }

    credOptionList := CredOptionList[1:]
	for _, name := range credOptionList {
        if llanguage == LEnglishLanguage {
		    fmt.Printf("Please enter %s:", name)
        } else {
            fmt.Printf("请输入%s：", name)
        }
		if _, err := fmt.Scanln(&val); err == nil {
			section.Add(name, val)
		}
	}

	for _, sec := range []string{BucketEndpointSection, BucketCnameSection} {
        if llanguage == LEnglishLanguage {
		    fmt.Printf("\nIs there any %s configurations(yes or no)?", sec)
        } else {
		    fmt.Printf("\n是否需要配置：%s(y or n)?", sec)
        }
		if _, err := fmt.Scanln(&val); err == nil && (val == "yes" || val == "y") {
			section = config.NewSection(sec)
			nameList := strings.SplitN(sec, "-", 2)
			for {
				bucket := ""
				host := ""
                if llanguage == LEnglishLanguage {
				    fmt.Printf("Please enter the %s:", nameList[0])
                } else {
				    fmt.Printf("请输入%s：", nameList[0])
                }
				if _, err := fmt.Scanln(&bucket); err != nil || "" == strings.TrimSpace(bucket) {
                    if llanguage == LEnglishLanguage {
					    fmt.Printf("No %s entered, the configuration of %s ended.\n", nameList[0], sec)
                    } else {
                        fmt.Printf("未输入%s，%s项的配置结束。\n", nameList[0], sec)
                    }
					break
				}
				fmt.Printf("Please enter the %s:", nameList[1])
				_, _ = fmt.Scanln(&host)
				section.Add(bucket, host)
			}
		}
	}

	if err := configparser.Save(config, configFile); err != nil {
		return err
	}
	return nil
}

func (cc *ConfigCommand) runCommandNonInteractive(configFile, language string) error {
	configFile = DecideConfigFile(configFile)
	config := configparser.NewConfiguration()
	section := config.NewSection(CREDSection)
	section.Add(OptionLanguage, language)
	for name := range CredOptionMap {
		if val, _ := GetString(name, cc.command.options); val != "" {
			section.Add(name, val)
		}
	}
	if err := configparser.Save(config, configFile); err != nil {
		return err
	}
	return nil
}
