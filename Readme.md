*** Finance Solver
Backend for a finance solver project

*** Deploy 
```
pack build finance-solver:latest \
 --builder paketobuildpacks/builder:tiny \
 --path . \
&& fly deploy --local-only --image finance-solver:latest
```


*** Create table
```
CREATE TABLE expenses (  
  id SERIAL PRIMARY KEY,
  category TEXT NOT NULL,
  price REAL
)
```