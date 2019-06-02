package leetcode_go

import (
	"testing"
)

func TestPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 3,
								Left: &TreeNode{
									Val:   3,
									Left:  nil,
									Right: nil,
								},
								Right: &TreeNode{
									Val:  -2,
									Left: nil,
									Right: &TreeNode{
										Val: 3,
										Left: &TreeNode{
											Val:  3,
											Left: nil,
											Right: &TreeNode{
												Val: 3,
												Left: &TreeNode{
													Val:   3,
													Left:  nil,
													Right: nil,
												},
												Right: &TreeNode{
													Val:   -2,
													Left:  nil,
													Right: nil,
												},
											},
										},
										Right: &TreeNode{
											Val:   -2,
											Left:  nil,
											Right: nil,
										},
									},
								},
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val:  -2,
							Left: nil,
							Right: &TreeNode{
								Val: 3,
								Left: &TreeNode{
									Val: 3,
									Left: &TreeNode{
										Val: 3,
										Left: &TreeNode{
											Val:   3,
											Left:  nil,
											Right: nil,
										},
										Right: &TreeNode{
											Val:  -2,
											Left: nil,
											Right: &TreeNode{
												Val: 3,
												Left: &TreeNode{
													Val:   3,
													Left:  nil,
													Right: nil,
												},
												Right: &TreeNode{
													Val:   -2,
													Left:  nil,
													Right: nil,
												},
											},
										},
									},
									Right: nil,
								},
								Right: &TreeNode{
									Val:   -2,
									Left:  nil,
									Right: nil,
								},
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: -2,
							Left: &TreeNode{
								Val: 3,
								Left: &TreeNode{
									Val:   3,
									Left:  nil,
									Right: nil,
								},
								Right: &TreeNode{
									Val:   -2,
									Left:  nil,
									Right: nil,
								},
							},
							Right: nil,
						},
					},
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:  -2,
						Left: nil,
						Right: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val:  3,
								Left: nil,
								Right: &TreeNode{
									Val: 3,
									Left: &TreeNode{
										Val:   3,
										Left:  nil,
										Right: nil,
									},
									Right: &TreeNode{
										Val: -2,
										Left: &TreeNode{
											Val: 3,
											Left: &TreeNode{
												Val:   3,
												Left:  nil,
												Right: nil,
											},
											Right: &TreeNode{
												Val:  -2,
												Left: nil,
												Right: &TreeNode{
													Val: 3,
													Left: &TreeNode{
														Val:  3,
														Left: nil,
														Right: &TreeNode{
															Val: 3,
															Left: &TreeNode{
																Val:   3,
																Left:  nil,
																Right: nil,
															},
															Right: &TreeNode{
																Val:  -2,
																Left: nil,
																Right: &TreeNode{
																	Val: 3,
																	Left: &TreeNode{
																		Val:   3,
																		Left:  nil,
																		Right: nil,
																	},
																	Right: &TreeNode{
																		Val:  -2,
																		Left: nil,
																		Right: &TreeNode{
																			Val: 3,
																			Left: &TreeNode{
																				Val:  3,
																				Left: nil,
																				Right: &TreeNode{
																					Val: 3,
																					Left: &TreeNode{
																						Val: 3,
																						Left: &TreeNode{
																							Val: 3,
																							Left: &TreeNode{
																								Val:   3,
																								Left:  nil,
																								Right: nil,
																							},
																							Right: &TreeNode{
																								Val:  -2,
																								Left: nil,
																								Right: &TreeNode{
																									Val: 3,
																									Left: &TreeNode{
																										Val:  3,
																										Left: nil,
																										Right: &TreeNode{
																											Val: 3,
																											Left: &TreeNode{
																												Val:   3,
																												Left:  nil,
																												Right: nil,
																											},
																											Right: &TreeNode{
																												Val:   -2,
																												Left:  nil,
																												Right: nil,
																											},
																										},
																									},
																									Right: &TreeNode{
																										Val:   -2,
																										Left:  nil,
																										Right: nil,
																									},
																								},
																							},
																						},
																						Right: nil,
																					},
																					Right: &TreeNode{
																						Val:   -2,
																						Left:  nil,
																						Right: nil,
																					},
																				},
																			},
																			Right: &TreeNode{
																				Val:   -2,
																				Left:  nil,
																				Right: nil,
																			},
																		},
																	},
																},
															},
														},
													},
													Right: &TreeNode{
														Val:   -2,
														Left:  nil,
														Right: nil,
													},
												},
											},
										},
										Right: nil,
									},
								},
							},
							Right: &TreeNode{
								Val:   -2,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:  -2,
							Left: nil,
							Right: &TreeNode{
								Val: 3,
								Left: &TreeNode{
									Val: 3,
									Left: &TreeNode{
										Val: 3,
										Left: &TreeNode{
											Val:   3,
											Left:  nil,
											Right: nil,
										},
										Right: &TreeNode{
											Val: -2,
											Left: &TreeNode{
												Val: 3,
												Left: &TreeNode{
													Val:   3,
													Left:  nil,
													Right: nil,
												},
												Right: &TreeNode{
													Val:   -2,
													Left:  nil,
													Right: nil,
												},
											},
											Right: nil,
										},
									},
									Right: nil,
								},
								Right: &TreeNode{
									Val:   -2,
									Left:  nil,
									Right: nil,
								},
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   -2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
		},
		Right: &TreeNode{
			Val:  -3,
			Left: nil,
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}

	actualResult := pathSum(tree, 8)
	expectedResult := 37
	if expectedResult != actualResult {
		t.Errorf("The expected result should be %d, but we got %d", expectedResult, actualResult)
	}
}

// func makeTree(numbers []int) {
//     for _, number := range numbers {
//
//     }
// }
