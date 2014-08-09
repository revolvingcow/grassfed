// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tDatabaseController struct {}
var DatabaseController tDatabaseController


func (_ tDatabaseController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DatabaseController.Begin", args).Url
}

func (_ tDatabaseController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DatabaseController.Commit", args).Url
}

func (_ tDatabaseController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("DatabaseController.Rollback", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tApplication struct {}
var Application tApplication



type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).Url
}

func (_ tHome) About(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.About", args).Url
}

func (_ tHome) Overview(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Overview", args).Url
}


type tProfile struct {}
var Profile tProfile


func (_ tProfile) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Profile.Index", args).Url
}

func (_ tProfile) Logon(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Profile.Logon", args).Url
}

func (_ tProfile) History(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Profile.History", args).Url
}

func (_ tProfile) Stats(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Profile.Stats", args).Url
}

func (_ tProfile) Trends(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Profile.Trends", args).Url
}

func (_ tProfile) Add(
		product string,
		calories int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "product", product)
	revel.Unbind(args, "calories", calories)
	return revel.MainRouter.Reverse("Profile.Add", args).Url
}

func (_ tProfile) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Profile.Delete", args).Url
}

func (_ tProfile) Goal(
		calories int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "calories", calories)
	return revel.MainRouter.Reverse("Profile.Goal", args).Url
}

func (_ tProfile) Weight(
		weight float64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "weight", weight)
	return revel.MainRouter.Reverse("Profile.Weight", args).Url
}


