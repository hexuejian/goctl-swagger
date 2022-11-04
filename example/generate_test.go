/**
  @author: xuejian.he
  @since: 2022/11/4
  @desc: //TODO
**/
package example

import (
	"github.com/hexuejian/goctl-swagger/action"
	"testing"
)

func TestGeneratorTest(t *testing.T) {
	if err := action.GeneratorTest("user.json"); err != nil {
		t.Errorf("GeneratorTest() error = %v", err)
	}
}
