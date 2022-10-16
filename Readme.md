*** Finance Solver
Backend for a finance solver project

*** Deploy 
```
pack build finance-solver:latest \
 --builder paketobuildpacks/builder:tiny \
 --path . \
&& fly deploy --local-only --image finance-solver:latest
```