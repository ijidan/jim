package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gen"
	"jim/global"
)

// genGormCmd represents the genGorm command
var genGormCmd = &cobra.Command{
	Use:   "genGorm",
	Short: "auto gen gorm",
	Long:  `auto gen gorm`,
	Run: func(cmd *cobra.Command, args []string) {
		// specify the output directory (default: "./query")
		// ### if you want to query without context constrain, set mode gen.WithoutContext ###
		g := gen.NewGenerator(gen.Config{
			OutPath: "../internal/model",
			/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
			//if you want the nullable field generation property to be pointer type, set FieldNullable true
			/* FieldNullable: true,*/
			//if you want to generate index tags from database, set FieldWithIndexTag true
			/* FieldWithIndexTag: true,*/
			//if you want to generate type tags from database, set FieldWithTypeTag true
			/* FieldWithTypeTag: true,*/
			//if you need unit tests for query code, set WithUnitTest true
			/* WithUnitTest: true, */
		})

		// reuse the database connection in Project or create a connection here
		// if you want to use GenerateModel/GenerateModelAs, UseDB is necessray or it will panic
		// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
		g.UseDB(global.Db)

		// apply basic crud api on structs or table models which is specified by table name with function
		// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
		//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))

		// apply diy interfaces on structs or table models
		//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

		g.GenerateAllTable()
		// execute the action of code generation
		g.Execute()

		fmt.Println("genGorm called")
	},
}

func init() {
	rootCmd.AddCommand(genGormCmd)
}
