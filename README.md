# Sample Size Computation
While performing an AB experiment, the most importing thing to do are
1. Decide the randomisation unit
2. Decide the metrics
3. Compute the sample size needed to power the test

This repo focuses on how to compute sample size needed to power the test.

# Types of metrics

### Additive metric
Metrics which are computed per randonmisation unit are known as additive metric. 
For example: If the randomisation unit of a test is user, then dollar spent/user would be an additive metric

Stats needed for computing sample size
* Metric mean
* Metric standard deviation

**Sample request**

```python
additive_request = {
    "avg": 10,
    "std": 2,
    "mde": 0.05,
    "alpha": 0.05,
    "power": 0.8,
    "two_tailed": True
}
```

### Ratio metric
Metrics which is a ratio of two additive metrics is know as a ratio metrics.
For example: If the randomisation unit of a test is user, then dollar spent/total requests would be a ratio metric

Stats needed for computing sample size
* numerator mean
* denominator mean
* numberator std
* denominator std
* covariance between numberator and denominator

**Sample request**

  ```python
ratio_request = {
    "nummean": 10,
    "denmean": 2,
    "numstd": 5,
    "denstd": 1,
    "covv": 0.7,
    "mde": 0.05,
    "alpha": 0.05,
    "power": 0.8,
    "two_tailed": True
}
```

### Proportion metric
Metrics in which there is a binary success criteria
For example: If randomisation unit of a test is user, then reactivatation rate (reacted users/total users) would be a proportion metric

Stats needed for computing sample size
* Proportion value

```python
prop_request = {
    "prop": 0.2,
    "mde": 0.05,
    "alpha": 0.05,
    "power": 0.8,
    "two_tailed": True
}
```


## API Documentation

There is a separate endpoint for each kind of metrics mentioned above

Here is the sample code


**Function calling API**
```python
import requests
import json

def sampleSizeCompute(metriType, data):
    base_url = 'https://samplesizecalculator-production.up.railway.app/'
    base_url += metriType
    headers = {'content-type': 'application/json'}
    response = requests.post(base_url, data=json.dumps(data), headers=headers)
    if response.status_code == 200:
        samples = response.json()
        return samples['data']
    else:
        return None
```

**Calling API**
```python
# sampleSizeCompute(metric_type, metric_request)

sampleSizeCompute('additive', additive_request)
```
