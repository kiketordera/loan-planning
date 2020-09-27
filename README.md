# loan-planning
This is a simple loan planning. 

It receives a JSON object with the parameters for the planning and responds with a JSON structure with the plan.

Example of JSON structure allowed in the POST request:
```
 {
  "loanAmount": "5000", "nominalRate": "5.0",
  "duration": 24,
  "startDate": "2018-01-01T00:00:01Z"
}
```
Example of JSON structure answered:

```

{
   "borrowerPayments":[
        {
          "borrowerPaymentAmount":"219.36", 
          "date":"2018-01-01T00:00:00Z", 
          "initialOutstandingPrincipal":"5000.00", 
          "interest":"20.83", 
          "principal":"198.53", 
          "remainingOutstandingPrincipal":"4801.47"
        }, {
          "borrowerPaymentAmount":"219.36", 
          "date":"2018-02-01T00:00:00Z", 
          "initialOutstandingPrincipal":"4801.47", 
          "interest":"20.01", 
          "principal":"199.35", 
          "remainingOutstandingPrincipal":"4602.12"
        }, ... {
          "borrowerPaymentAmount":"219.28", 
          "date":"2019-12-01T00:00:00Z", 
          "initialOutstandingPrincipal":"218.37", 
          "interest":"0.91", 
          "principal":"218.37", 
          "remainingOutstandingPrincipal":"0"
        } 
    ]
}
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

What things you need to install for the software to run
```
golang
git (only Windows)
```


### Installing

With this command you will install the project and all his dependencies

```
go get -u github.com/kiketordera/loan-planning/...
```

Run it in pure Go
```
go run main.go
```

Run it in a Docker container
```
docker  build -t loanplanning .   
docker run -p 8080:8080 loanplanning
```
Then you can make POST request with the structure shown above. For example:

```
curl -H "Content-Type: application/json" -X POST -d '{"loanAmount":"5000","nominalRate":"5.0","duration":"24"}' localhost:8080
```


