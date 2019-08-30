//+build !js

package greeting_test

import (
	"github.com/crhntr/goes/examples/greeting"
	"github.com/crhntr/goes/goesfakes"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewHelloBox(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	div := goesfakes.NewValue(ctrl)
	div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box"))
	div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("hello"))

	document := goesfakes.NewValue(ctrl)
	document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

	box := greeting.NewHelloBox("hello")

	box.Create(document)
}

func TestHelloBox_SetMessage(t *testing.T) {
	{
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		div := goesfakes.NewValue(ctrl)
		gomock.InOrder(
			div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box")),

			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("hello")),
			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("Hello, world!")),
		)

		document := goesfakes.NewValue(ctrl)
		document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

		box := greeting.NewHelloBox("hello")

		box.Create(document)
		box.SetMessage("Hello, world!")

    if box.Message() != "Hello, world!" {
      t.Fail()
    }
	}

	t.Run("when initialized with an empty string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msgTxtFromDom := goesfakes.NewValue(ctrl)
		msgTxtFromDom.EXPECT().String().Return("Hello, world!").AnyTimes()

		div := goesfakes.NewValue(ctrl)
		gomock.InOrder(
			div.EXPECT().Call(gomock.Eq("setAttribute"), gomock.Eq("id"), gomock.Eq("hello-box")),
			div.EXPECT().Set(gomock.Eq("innerText"), gomock.Eq("")),

			div.EXPECT().Get(gomock.Eq("innerText")).Return(msgTxtFromDom), // <- what changed
		)

		document := goesfakes.NewValue(ctrl)
		document.EXPECT().Call(gomock.Eq("createElement"), gomock.Eq("div")).Return(div).Times(1)

		box := greeting.NewHelloBox("")

		box.Create(document)

    if box.Message() != "Hello, world!" {
      t.Fail()
    }
	})
}
