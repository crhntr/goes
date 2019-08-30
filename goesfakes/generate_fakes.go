package goesfakes

//go:generate mockgen -source ../interface.go -destination goes_fakes.go -package goesfakes -imports goes=github.com/crhntr/goes -mock_names Typer=Typer,Value=Value,Func=Func,Releaser=Releaser,TypedArray=TypedArray
