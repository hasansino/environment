/*
	Package environment implements methods to detect and classify environments
	where services are executed.
	Environment type can be set by variable `ENVIRONMENT`, it accepts values:
		`production`, `avangard`, `staging` and `development`,
	and defaults to `production`.
	The `Avangard` environment is used to describe small portion of production servers,
	usually used to battle-test new features without affecting working product too much.
*/
package environment
