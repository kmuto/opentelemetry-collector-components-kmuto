[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# samsa

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### samsa.battery.remaining

Battery remaining percent value

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| 1 | Gauge | Double |

#### Attributes

| Name | Description | Values |
| ---- | ----------- | ------ |
| samsa.device.name | Robot name | Any Str |