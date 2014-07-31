// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) About(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.About", args).Url
}

func (_ tApp) Me(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Me", args).Url
}

func (_ tApp) Stats(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Stats", args).Url
}

func (_ tApp) Add(
		id int,
		product string,
		calories int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "product", product)
	revel.Unbind(args, "calories", calories)
	return revel.MainRouter.Reverse("App.Add", args).Url
}

func (_ tApp) Goal(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Goal", args).Url
}

func (_ tApp) SetGoal(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.SetGoal", args).Url
}

func (_ tApp) Streak(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Streak", args).Url
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


