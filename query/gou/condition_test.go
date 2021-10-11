package gou

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestConditionBase(t *testing.T) {
	var conds []Condition
	bytes := ReadFile("conditions/base.json")
	type M = map[string]interface{}
	type F = float64
	type Any = interface{}
	should := func(t assert.TestingT, actual, expected interface{}, msgAndArgs ...interface{}) bool {
		return assert.Equal(t, expected, actual, msgAndArgs...)
	}

	err := jsoniter.Unmarshal(bytes, &conds)
	assert.Nil(t, err)
	assert.Equal(t, 48, len(conds))

	// { "field": "score", "value": 20, "op": "=", "comment": "分数" },
	should(t, conds[0].ToMap(), M{"field": "score", "value": F(20), "op": "=", "comment": "分数"})
	// { "field": "score", "=": 20 },
	should(t, conds[1].ToMap(), M{"field": "score", "value": F(20), "op": "="})
	// { ":score": "分数", "=": 10 },
	should(t, conds[2].ToMap(), M{"field": "score", "value": F(10), "op": "=", "comment": "分数"})

	// { "field": "score", "value": 20, "op": "<", "comment": "分数" },
	should(t, conds[3].ToMap(), M{"field": "score", "value": F(20), "op": "<", "comment": "分数"})
	// { "field": "score", "<": 20 },
	should(t, conds[4].ToMap(), M{"field": "score", "value": F(20), "op": "<"})
	// { ":score": "分数", "<": 20 },
	should(t, conds[5].ToMap(), M{"field": "score", "value": F(20), "op": "<", "comment": "分数"})

	// { "field": "score", "value": 20, "op": "<=", "comment": "分数" },
	should(t, conds[6].ToMap(), M{"field": "score", "value": F(20), "op": "<=", "comment": "分数"})
	// { "field": "score", "<=": 20 },
	should(t, conds[7].ToMap(), M{"field": "score", "value": F(20), "op": "<="})
	// { ":score": "分数", "<=": 20 },
	should(t, conds[8].ToMap(), M{"field": "score", "value": F(20), "op": "<=", "comment": "分数"})

	// { "field": "score", "value": 20, "op": ">", "comment": "分数" },
	should(t, conds[9].ToMap(), M{"field": "score", "value": F(20), "op": ">", "comment": "分数"})
	// { "field": "score", ">": 20 },
	should(t, conds[10].ToMap(), M{"field": "score", "value": F(20), "op": ">"})
	// { ":score": "分数", ">": 20 },
	should(t, conds[11].ToMap(), M{"field": "score", "value": F(20), "op": ">", "comment": "分数"})

	// { "field": "score", "value": 20, "op": ">=", "comment": "分数" },
	should(t, conds[12].ToMap(), M{"field": "score", "value": F(20), "op": ">=", "comment": "分数"})
	// { "field": "score", ">=": 20 },
	should(t, conds[13].ToMap(), M{"field": "score", "value": F(20), "op": ">="})
	// { ":score": "分数", ">=": 20 },
	should(t, conds[14].ToMap(), M{"field": "score", "value": F(20), "op": ">=", "comment": "分数"})

	// { "field": "name", "value": "李", "op": "match", "comment": "姓名" },
	should(t, conds[15].ToMap(), M{"field": "name", "value": "李", "op": "match", "comment": "姓名"})
	// { "field": "name", "match": "李" },
	should(t, conds[16].ToMap(), M{"field": "name", "value": "李", "op": "match"})
	// { ":name": "姓名", "match": "李" },
	should(t, conds[17].ToMap(), M{"field": "name", "value": "李", "op": "match", "comment": "姓名"})

	// { "field": "name", "value": "%明", "op": "like", "comment": "姓名" },
	should(t, conds[18].ToMap(), M{"field": "name", "value": "%明", "op": "like", "comment": "姓名"})
	// { "field": "name", "like": "%明" },
	should(t, conds[19].ToMap(), M{"field": "name", "value": "%明", "op": "like"})
	// { ":name": "姓名", "like": "%明" },
	should(t, conds[20].ToMap(), M{"field": "name", "value": "%明", "op": "like", "comment": "姓名"})

	// { "field": "score", "value": [10, 20], "op": "in", "comment": "分数" },
	should(t, conds[21].ToMap(), M{"field": "score", "value": []Any{F(10), F(20)}, "op": "in", "comment": "分数"})
	// { "field": "score", "in": [10, 20] },
	should(t, conds[22].ToMap(), M{"field": "score", "value": []Any{F(10), F(20)}, "op": "in"})
	// { ":score": "分数", "in": [10, 20] },
	should(t, conds[23].ToMap(), M{"field": "score", "value": []Any{F(10), F(20)}, "op": "in", "comment": "分数"})

	// { "field": "name", "value": ["张三", "李四"], "op": "in", "comment": "姓名" },
	should(t, conds[24].ToMap(), M{"field": "name", "value": []Any{"张三", "李四"}, "op": "in", "comment": "姓名"})
	// { "field": "name", "in": ["张三", "李四"] },
	should(t, conds[25].ToMap(), M{"field": "name", "value": []Any{"张三", "李四"}, "op": "in"})
	// { ":name": "姓名", "in": ["张三", "李四"] },
	should(t, conds[26].ToMap(), M{"field": "name", "value": []Any{"张三", "李四"}, "op": "in", "comment": "姓名"})

	// { "field": "name", "op": "is", "value": "null", "comment": "姓名" },
	should(t, conds[27].ToMap(), M{"field": "name", "value": "null", "op": "is", "comment": "姓名"})
	// { "field": "name", "is": "null" },
	should(t, conds[28].ToMap(), M{"field": "name", "value": "null", "op": "is"})
	// { ":name": "姓名", "is": "null" },
	should(t, conds[29].ToMap(), M{"field": "name", "value": "null", "op": "is", "comment": "姓名"})

	// { "field": "name", "op": "is", "value": "not null", "comment": "姓名" },
	should(t, conds[30].ToMap(), M{"field": "name", "value": "not null", "op": "is", "comment": "姓名"})
	// { "field": "name", "is": "not null" },
	should(t, conds[31].ToMap(), M{"field": "name", "value": "not null", "op": "is"})
	// { ":name": "姓名", "is": "not null" },
	should(t, conds[32].ToMap(), M{"field": "name", "value": "not null", "op": "is", "comment": "姓名"})

	// { "or": true, "field": "name", "op": "match", "value": "李" },
	should(t, conds[33].ToMap(), M{"field": "name", "value": "李", "op": "match", "or": true})
	// { "or": true, "field": "name", "match": "李" },
	should(t, conds[34].ToMap(), M{"field": "name", "value": "李", "op": "match", "or": true})
	// { "or :name": "或姓名", "match": "李" },
	should(t, conds[35].ToMap(), M{"field": "name", "value": "李", "op": "match", "or": true, "comment": "或姓名"})

	// { "or": true, "field": "name", "op": "is", "value": "notnull" },
	should(t, conds[36].ToMap(), M{"field": "name", "value": "notnull", "op": "is", "or": true})
	// { "or": true, "field": "name", "is": "notnull" },
	should(t, conds[37].ToMap(), M{"field": "name", "value": "notnull", "op": "is", "or": true})
	// { "or :name": "或姓名", "is": "notnull" },
	should(t, conds[38].ToMap(), M{"field": "name", "value": "notnull", "op": "is", "or": true, "comment": "或姓名"})

	// { "or": false, "field": "name", "op": "is", "value": "notnull" },
	should(t, conds[39].ToMap(), M{"field": "name", "value": "notnull", "op": "is"})
	// { "or": false, "field": "name", "is": "notnull" },
	should(t, conds[40].ToMap(), M{"field": "name", "value": "notnull", "op": "is"})
	// { ":name": "姓名", "is": "notnull" },
	should(t, conds[41].ToMap(), M{"field": "name", "value": "notnull", "op": "is", "comment": "姓名"})

	// { "field": "id", "value": 20, "op": "=" },
	should(t, conds[42].ToMap(), M{"field": "id", "value": F(20), "op": "="})
	assert.Nil(t, conds[42].ValueExpression)
	// { "field": "name", "value": "张三", "op": "=" },
	should(t, conds[43].ToMap(), M{"field": "name", "value": "张三", "op": "="})
	assert.Nil(t, conds[43].ValueExpression)
	// { "field": "id", "value": "{20}", "op": "=" },
	should(t, conds[44].ToMap(), M{"field": "id", "value": "{20}", "op": "="})
	should(t, conds[44].ValueExpression.ToString(), "20")
	// { "field": "name", "value": "{'张三'}", "op": "=" },
	should(t, conds[45].ToMap(), M{"field": "name", "value": "{'张三'}", "op": "="})
	should(t, conds[45].ValueExpression.ToString(), "'张三'")
	// { "field": "name", "value": "{short_name}", "op": "=" },
	should(t, conds[46].ToMap(), M{"field": "name", "value": "{short_name}", "op": "="})
	should(t, conds[46].ValueExpression.ToString(), "short_name")
	// { "or": true, "field": "name", "value": "{short_name}", "op": "=" }
	should(t, conds[47].ToMap(), M{"field": "name", "value": "{short_name}", "op": "=", "or": true})
	should(t, conds[47].ValueExpression.ToString(), "short_name")
}

func TestConditionValidate(t *testing.T) {
	var errs []Condition
	bytes := ReadFile("conditions/error.json")

	err := jsoniter.Unmarshal(bytes, &errs)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(errs))

	// { ":score": "分数" },
	res := errs[0].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "op")

	// { "=": 10 },
	res = errs[1].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "field")

	// { ":score": "分数", "op": "gt", "value": 20 },
	res = errs[2].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "gt")

	// { ":score": "分数", "op": "=", "value": "{hello world}" },
	res = errs[3].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "value")
	assert.Contains(t, res[0].Error(), "hello world")

	// { "field": "score hello", "op": "=", "value": "hello" },
	res = errs[4].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "field")
	assert.Contains(t, res[0].Error(), "score hello")

	// { "field": "score hello", "op": "=", "value": "{hello world}" }
	res = errs[5].Validate()
	assert.Equal(t, 2, len(res))
	assert.Contains(t, res[0].Error(), "field")
	assert.Contains(t, res[0].Error(), "score hello")
	assert.Contains(t, res[1].Error(), "value")
	assert.Contains(t, res[1].Error(), "hello world")

	// { "field": "score", "op": "=" },
	res = errs[6].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "value")
	assert.Contains(t, res[0].Error(), "query")

	// { "field": "score", "op": "is", "value": "error" }
	res = errs[7].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "null")
	assert.Contains(t, res[0].Error(), "not null")

	// { "field": "score", "op": "is" }
	res = errs[8].Validate()
	assert.Equal(t, 1, len(res))
	assert.Contains(t, res[0].Error(), "value")
	assert.Contains(t, res[0].Error(), "query")

	// { "field": "score", "op": "is", "value": "null" }
	res = errs[9].Validate()
	assert.Equal(t, 0, len(res))

}

func TestConditionUnmarshalJSONError(t *testing.T) {
	var cond Condition
	err := jsoniter.Unmarshal([]byte(`{1}`), &cond)
	assert.NotNil(t, err)
}

func TestConditionMarshalJSON(t *testing.T) {

	cond := Condition{Field: NewExpression("id"), OP: "=", Value: 1}
	bytes, err := jsoniter.Marshal(cond)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), `"field":"id"`)
	assert.Contains(t, string(bytes), `"op":"="`)
	assert.Contains(t, string(bytes), `"value":1`)
	assert.NotContains(t, string(bytes), `"or"`)
	assert.NotContains(t, string(bytes), `"comment"`)

	cond = Condition{Field: NewExpression("id"), OP: "=", Value: 1, OR: true}
	bytes, err = jsoniter.Marshal(cond)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), `"field":"id"`)
	assert.Contains(t, string(bytes), `"op":"="`)
	assert.Contains(t, string(bytes), `"or":true`)
	assert.Contains(t, string(bytes), `"value":1`)
	assert.NotContains(t, string(bytes), `"comment"`)

	cond = Condition{Field: NewExpression("id"), OP: "=", Value: 1, OR: false}
	bytes, err = jsoniter.Marshal(cond)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), `"field":"id"`)
	assert.Contains(t, string(bytes), `"op":"="`)
	assert.Contains(t, string(bytes), `"value":1`)
	assert.NotContains(t, string(bytes), `"or"`)
	assert.NotContains(t, string(bytes), `"comment"`)

	cond = Condition{Field: NewExpression("id"), OP: "=", Value: 1, OR: true, Comment: "Unit-Test"}
	bytes, err = jsoniter.Marshal(cond)
	assert.Nil(t, err)
	assert.Contains(t, string(bytes), `"field":"id"`)
	assert.Contains(t, string(bytes), `"op":"="`)
	assert.Contains(t, string(bytes), `"or":true`)
	assert.Contains(t, string(bytes), `"value":1`)
	assert.Contains(t, string(bytes), `"comment":"Unit-Test"`)

}