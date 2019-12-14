// Test of cognitive complexity.

// Package pkg ...
package pkg

import (
	"fmt"
	ast "go/ast"
	"log"
)

// Test IF and Boolean expr
func f(x int) bool { // MATCH /function f has cognitive complexity 3 (> max enabled 0)/
	if x > 0 && true || false { // +3
		return true
	} else {
		log.Printf("non-positive x: %d", x)
	}
	return false
}

// Test IF
func g(f func() bool) string { // MATCH /function g has cognitive complexity 1 (> max enabled 0)/
	if ok := f(); ok { // +1
		return "it's okay"
	} else {
		return "it's NOT okay!"
	}
}

// Test Boolean expr
func h(a, b, c, d, e, f bool) bool { // MATCH /function h has cognitive complexity 2 (> max enabled 0)/
	return a && b && c || d || e && f //FIXME: complexity should be 3
}

func i(a, b, c, d, e, f bool) bool { // MATCH /function i has cognitive complexity 2 (> max enabled 0)/
	result := a && b && c || d || e // +2
	return result
}

func j(a, b, c, d, e, f bool) bool { // MATCH /function j has cognitive complexity 2 (> max enabled 0)/
	result := z(a && !(b && c)) // +2
	return result
}

// Test Switch expr
func k(a, b, c, d, e, f bool) bool { // MATCH /function k has cognitive complexity 1 (> max enabled 0)/
	switch expr { // +1
	case cond1:
	case cond2:
	default:
	}

	return result
}

// Test nesting FOR expr + nested IF
func l() { // MATCH /function l has cognitive complexity 6 (> max enabled 0)/
	for i := 1; i <= max; i++ { // +1
		for j := 2; j < i; j++ { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				continue
			}
		}

		total += i
	}
	return total
}

// Test nesting IF
func m() { // MATCH /function m has cognitive complexity 6 (> max enabled 0)/
	if i <= max { // +1
		if j < i { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				return 0
			}
		}

		total += i
	}
	return total
}

// Test nesting IF + nested FOR
func n() { // MATCH /function n has cognitive complexity 6 (> max enabled 0)/
	if i > max { // +1
		for j := 2; j < i; j++ { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				continue
			}
		}

		total += i
	}
	return total
}

// Test nesting
func o() { // MATCH /function o has cognitive complexity 12 (> max enabled 0)/
	if i > max { // +1
		if j < i { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				return
			}
		}

		total += i
	}

	if i > max { // +1
		if j < i { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				return
			}
		}

		total += i
	}
}

// Tests TYPE SWITCH
func p() { // MATCH /function p has cognitive complexity 1 (> max enabled 0)/
	switch n := n.(type) { // +1
	case *ast.IfStmt:
		targets := []ast.Node{n.Cond, n.Body, n.Else}
		v.walk(targets...)
		return nil
	case *ast.ForStmt:
		v.walk(n.Body)
		return nil
	case *ast.TypeSwitchStmt:
		v.walk(n.Body)
		return nil
	case *ast.BinaryExpr:
		v.complexity += v.binExpComplexity(n)
		return nil
	}
}

// Test RANGE
func q() { // MATCH /function q has cognitive complexity 1 (> max enabled 0)/
	for _, t := range targets { // +1
		ast.Walk(v, t)
	}
}

// Tests SELECT
func r() { // MATCH /function r has cognitive complexity 1 (> max enabled 0)/
	select { // +1
	case c <- x:
		x, y = y, x+y
	case <-quit:
		fmt.Println("quit")
		return
	}
}

// Test jump to label
func s() { // MATCH /function s has cognitive complexity 3 (> max enabled 0)/
FirstLoop:
	for i := 0; i < 10; i++ { // +1
		break
	}
	for i := 0; i < 10; i++ { // +1
		break FirstLoop // +1
	}
}

func t() { // MATCH /function t has cognitive complexity 2 (> max enabled 0)/
FirstLoop:
	for i := 0; i < 10; i++ { // +1
		goto FirstLoop // +1
	}
}

func u() { // MATCH /function u has cognitive complexity 3 (> max enabled 0)/
FirstLoop:
	for i := 0; i < 10; i++ { // +1
		continue
	}
	for i := 0; i < 10; i++ { // +1
		continue FirstLoop // +1
	}
}

// Tests FUNC LITERAL
func v() { // MATCH /function v has cognitive complexity 2 (> max enabled 0)/
	myFunc := func(b bool) {
		if b { // +1 +1(nesting)

		}
	}
}

func w() { // MATCH /function w has cognitive complexity 3 (> max enabled 0)/
	defer func(b bool) {
		if b { // +1 +1(nesting)

		}
	}(false || true) // +1
}

// Test from Cognitive Complexity white paper
func sumOfPrimes(max int) int { // MATCH /function sumOfPrimes has cognitive complexity 7 (> max enabled 0)/
	total := 0
OUT:
	for i := 1; i <= max; i++ { // +1
		for j := 2; j < i; j++ { // +1 +1(nesting)
			if i%j == 0 { // +1 +2(nesting)
				continue OUT // +1
			}
		}

		total += i
	}
	return total
}
