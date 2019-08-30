package goesfakes

//go:generate mockgen -source ../interface.go -destination interface_fakes.go -package goesfakes -imports goes=github.com/crhntr/goes -mock_names Runtimer=Runtime,Booler=Booler,Caller=Caller,Floater=Floater,Getter=Getter,Indexer=Indexer,InstanceOfer=InstanceOfer,Inter=Inter,Invoker=Invoker,Lengther=Lengther,Newer=Newer,Setter=Setter,SetIndexer=SetIndexer,Stringer=Stringer,Truther=Truther,Typer=Typer,Value=Value,Func=Func,Releaser=Releaser,TypedArray=TypedArray
