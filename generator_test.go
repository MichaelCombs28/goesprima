package goesprima

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LiteralTest struct {
	Literal
	Expect string
}

func TestLiteralString(t *testing.T) {
	tests := []LiteralTest{
		{
			Literal: StringLiteral("hello"),
			Expect:  `"hello"`,
		},
		{
			Literal: BoolLiteral(false),
			Expect:  "false",
		},
		{
			Literal: LiteralValueNull,
			Expect:  "null",
		},
		{
			Literal: LiteralValueUndefined,
			Expect:  "undefined",
		},
		{
			Literal: NumberLiteral(10.123123),
			Expect:  "10.123123",
		},
		{
			Literal: NumberLiteral(big.NewFloat(10.123123)),
			Expect:  "10.123123",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expect, test.Literal.String())
	}
}

func TestGenerator(t *testing.T) {

	expectation := `import Amplify from "@aws-amplify/core";
import { Auth } from "@aws-amplify/auth";
export const USER_POOL_CLIENT_ID_TODOUSERS = process.env.REACT_APP_USER_POOL_CLIENT_ID_TODOUSERS;
export const config = {
  Auth: {
    region: USER_POOL_REGION_TODOUSERS,
    userPoolId: USER_POOL_ID_TODOUSERS,
    userPoolWebClientId: USER_POOL_CLIENT_ID_TODOUSERS,
    storage: localStorage,
  },
  SecondaryAuths: null,
  API: {
    endpoints: [
      {
        name: REST_NAME_TODORESTAPI,
        endpoint: REST_URL_TODORESTAPI,
        custom_header: async () => {
          return {
            Authorization: "Bearer " + (await Auth.currentSession()).getAccessToken().getJwtToken(),
          };
        },
      },
    ],
  },
};
Amplify.configure(config);`
	g := NewGenerator()
	g.AddStatements(
		&ImportDeclaration{
			Source: "@aws-amplify/core",
			Specifiers: []ImportDeclarationSpecifier{
				&ImportDefaultSpecifier{
					Local: &Identifier{
						Name: "Amplify",
					},
				},
			},
		},
		&ImportDeclaration{
			Source: "@aws-amplify/auth",
			Specifiers: []ImportDeclarationSpecifier{
				&ImportSpecifier{
					NamedImports: []NamedImport{
						{
							Imported: &Identifier{
								Name: "Auth",
							},
						},
					},
				},
			},
		},
		&ExportNamedDeclaration{
			Declaration: &VariableDeclaration{
				Kind: VariableDeclarationTypeConst,
				Declarations: []VariableDeclarator{
					{
						ID: &Identifier{
							Name: "USER_POOL_CLIENT_ID_TODOUSERS",
						},
						Init: &StaticMemberExpression{
							Object: &StaticMemberExpression{
								Object: &Identifier{
									Name: "process",
								},
								Property: &Identifier{
									Name: "env",
								},
							},
							Property: &Identifier{
								Name: "REACT_APP_USER_POOL_CLIENT_ID_TODOUSERS",
							},
						},
					},
				},
			},
		},
		&ExportNamedDeclaration{
			Declaration: &VariableDeclaration{
				Kind: VariableDeclarationTypeConst,
				Declarations: []VariableDeclarator{
					{
						ID: &Identifier{
							Name: "config",
						},
						Init: &ObjectExpression{
							Properties: []ObjectExpressionProperty{
								&Property{
									Key: &Identifier{
										Name: "Auth",
									},
									Value: &ObjectExpression{
										Properties: []ObjectExpressionProperty{
											&Property{
												Key: &Identifier{
													Name: "region",
												},
												Value: &Identifier{
													Name: "USER_POOL_REGION_TODOUSERS",
												},
											},
											&Property{
												Key: &Identifier{
													Name: "userPoolId",
												},
												Value: &Identifier{
													Name: "USER_POOL_ID_TODOUSERS",
												},
											},
											&Property{
												Key: &Identifier{
													Name: "userPoolWebClientId",
												},
												Value: &Identifier{
													Name: "USER_POOL_CLIENT_ID_TODOUSERS",
												},
											},
											&Property{
												Key: &Identifier{
													Name: "storage",
												},
												Value: &Identifier{
													Name: "localStorage",
												},
											},
										},
									},
								},
								&Property{
									Key: &Identifier{
										Name: "SecondaryAuths",
									},
									Value: LiteralValueNull,
								},
								&Property{
									Key: &Identifier{
										Name: "API",
									},
									Value: &ObjectExpression{
										Properties: []ObjectExpressionProperty{
											&Property{
												Key: &Identifier{
													Name: "endpoints",
												},
												Value: &ArrayExpression{
													Elements: []ArrayExpressionElement{
														&ObjectExpression{
															Properties: []ObjectExpressionProperty{
																&Property{
																	Key: &Identifier{
																		Name: "name",
																	},
																	Value: &Identifier{
																		Name: "REST_NAME_TODORESTAPI",
																	},
																},
																&Property{
																	Key: &Identifier{
																		Name: "endpoint",
																	},
																	Value: &Identifier{
																		Name: "REST_URL_TODORESTAPI",
																	},
																},
																&Property{
																	Key: &Identifier{
																		Name: "custom_header",
																	},
																	Value: &ArrowFunctionExpression{
																		Async: true,
																		Body: BlockStatement{
																			Items: []Statement{
																				&ReturnStatement{
																					Argument: &ObjectExpression{
																						Properties: []ObjectExpressionProperty{
																							&Property{
																								Key: &Identifier{
																									Name: "Authorization",
																								},
																								Value: &BinaryExpression{
																									Left:     StringLiteral("Bearer "),
																									Operator: BinaryOperatorADD,
																									Right: &StaticMemberExpression{
																										Object: &AwaitExpression{
																											Arguement: &StaticMemberExpression{
																												Object: &Identifier{
																													Name: "Auth",
																												},
																												Property: &CallExpression{
																													Callee: &Identifier{
																														Name: "currentSession",
																													},
																												},
																											},
																										},
																										Property: &StaticMemberExpression{
																											Object: &CallExpression{
																												Callee: &Identifier{
																													Name: "getAccessToken",
																												},
																											},
																											Property: &CallExpression{
																												Callee: &Identifier{
																													Name: "getJwtToken",
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		&ExpressionStatement{
			Expression: &StaticMemberExpression{
				Object: &Identifier{
					Name: "Amplify",
				},
				Property: &CallExpression{
					Callee: &Identifier{
						Name: "configure",
					},
					Arguments: []ArgumentListElement{
						&Identifier{
							Name: "config",
						},
					},
				},
			},
		},
	)
	s := g.String()
	assert.Equal(t, expectation, s)
}
