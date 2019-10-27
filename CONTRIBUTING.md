## How Can I Contribute?

* Ensure that your go version supports `go.mod` (>= 1.11)

* Clone the project
* Install Go dependencies:

```
go mod download
```

* Switch to public folder and install npm dependencies and angular cli:

```
cd public
npm install -g @angular/cli
npm install
```

* Deploy to a local server:

```
ng serve
```

* Once you implemented your frontend changes, build the artifact, build it as golang assets:

```
go-bindata-assetfs -o template.go dist/ dist/assets/images/
```
