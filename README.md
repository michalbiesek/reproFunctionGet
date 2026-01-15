# reproFunctionGet

Tests SDK behavior with function retrieval and schema inspection.

## SDK Version

This repository uses **`github.com/criblio/cribl-control-plane-sdk-go v0.5.0-beta.21`** for testing.

**Note**: On the RC (Release Candidate) version (e.g., `v0.5.0-rc.23`), the schema handling works incorrectly. This repository uses the beta version to demonstrate the correct behavior and compare against the RC version's issues.

## Quick Start

```bash
# Install dependencies
go mod tidy

# Start Cribl server
make start

# Run example
make run
```

## What It Does

The example demonstrates:
1. Authenticating with Cribl server using username/password
2. Retrieving function information (e.g., `aggregate_metrics`)
3. Inspecting function schemas using reflection
4. Debugging SDK response structures

## Environment Variables

- `CRIBL_SERVER_URL` - Server URL (default: `http://localhost:9000`)
- `CRIBL_USERNAME` or `CRIBL_USER` - Username for authentication (default: `admin`)
- `CRIBL_PASSWORD` or `CRIBL_PASS` - Password for authentication (default: `admin`)

The code will automatically authenticate using the `/api/v1/auth/login` endpoint to obtain a bearer token.

## Docker Commands

```bash
make start    # Start server
make stop     # Stop server
make restart  # Restart server
make clean    # Remove containers and volumes
```

## Problem Description

When calling `Functions.Get()` for `aggregate_metrics`, there is a clear difference in how the schema is returned between the beta and RC versions:

### Beta Version (`v0.5.0-beta.21`) - ✅ Correct Behavior

The SDK correctly returns the schema with all properties:

```json
{
  "schema": {
    "cumulative": false,
    "flushOnInputClose": true,
    "passthrough": false,
    "preserveGroupBys": false,
    "shouldTreatDotsAsLiterals": true,
    "sufficientStatsOnly": false,
    "timeWindow": "10s"
  }
}
```

### RC Version (`v0.5.0-rc.23`) - ❌ Incorrect Behavior

The SDK returns an **empty schema object**:

```json
{
  "schema": {}
}
```

**Issue**: The RC version fails to properly deserialize the schema from the API response, resulting in an empty schema object (`{}`). This indicates a regression in:
- Schema deserialization from the API response
- Schema type definitions in the SDK models
- How the SDK parses/structures the schema in the response

This repository demonstrates the correct behavior (beta) and helps identify the regression in the RC version.

Valid Response
```
Fetching function: aggregate_metrics

Response:
{
  "count": 1,
  "items": [
    {
      "__filename": "/opt/cribl/default/cribl/functions/aggregate_metrics/index.js",
      "disabled": false,
      "group": "Metrics",
      "handleSignals": true,
      "id": "aggregate_metrics",
      "loadTime": 1768434545385,
      "modTime": 1763503897000,
      "name": "Aggregate Metrics",
      "schema": {
        "cumulative": false,
        "flushOnInputClose": true,
        "passthrough": false,
        "preserveGroupBys": false,
        "shouldTreatDotsAsLiterals": true,
        "sufficientStatsOnly": false,
        "timeWindow": "10s"
      },
      "sync": true,
      "uischema": {
        "add": {
          "items": {
            "name": {
              "ui:options": {
                "columnWidth": "30%"
              },
              "ui:placeholder": "Enter field name"
            },
            "value": {
              "ui:options": {
                "columnWidth": "70%"
              },
              "ui:placeholder": "Enter value expression",
              "ui:widget": "JsInput"
            }
          },
          "ui:field": "Table"
        },
        "aggregations": {
          "items": {
            "agg": {
              "ui:options": {
                "columnWidth": "80%"
              },
              "ui:placeholder": "Enter an aggregate function",
              "ui:widget": "AggInput"
            },
            "metricType": {
              "ui:options": {
                "allowClear": false,
                "uiEnumOptions": [
                  {
                    "description": "Use the aggregation function to determine the type of metric being output",
                    "label": "Automatic",
                    "value": "automatic"
                  },
                  {
                    "description": "The following functions can be aggregated into a counter metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Counter",
                    "value": "counter"
                  },
                  {
                    "description": "The following functions can be aggregated into a distribution metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`. Distribution is only supported by Datadog.",
                    "label": "Distribution",
                    "value": "distribution"
                  },
                  {
                    "description": "The following functions can be aggregated into a gauge metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Gauge",
                    "value": "gauge"
                  },
                  {
                    "description": "The following functions can be aggregated into a histogram metric: `histogram`",
                    "label": "Histogram",
                    "value": "histogram"
                  },
                  {
                    "description": "The following functions can be aggregated into a summary metric: `summary`",
                    "label": "Summary",
                    "value": "summary"
                  },
                  {
                    "description": "The following functions can be aggregated into a timer metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Timer",
                    "value": "timer"
                  }
                ]
              }
            }
          },
          "ui:field": "Table"
        },
        "flushEventLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to unlimited"
        },
        "flushMemLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to unlimited"
        },
        "groupbys": {
          "ui:field": "Tags",
          "ui:placeholder": "One or more dimensions (or wildcard expression) to group by. Optional."
        },
        "idleTimeLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to the smaller of Time Window and 1 minute"
        },
        "lagTolerance": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to the smaller of Time Window and 1 minute"
        },
        "passthrough": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          }
        },
        "preserveGroupBys": {
          "ui:options": {
            "inline": true,
            "inlineStyle": {
              "padding": "0"
            },
            "width": "50%"
          }
        },
        "sufficientStatsOnly": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          }
        },
        "timeWindow": {
          "ui:options": {
            "labelInline": true,
            "width": "200px"
          }
        },
        "ui:options": {
          "groups": {
            "advanced": {
              "collapsed": true,
              "properties": [
                "flushEventLimit",
                "flushMemLimit",
                "shouldTreatDotsAsLiterals",
                "flushOnInputClose"
              ],
              "title": "Advanced Settings"
            },
            "modeSettings": {
              "collapsed": true,
              "properties": [
                "passthrough",
                "sufficientStatsOnly",
                "preserveGroupBys",
                "prefix"
              ],
              "title": "Output Settings"
            },
            "windowType": {
              "collapsed": true,
              "properties": [
                "cumulative",
                "lagTolerance",
                "idleTimeLimit"
              ],
              "title": "Time Window Settings"
            }
          }
        }
      },
      "version": "0.1"
    }
  ]
}
```

Invalid response
```
Response:
{
  "count": 1,
  "items": [
    {
      "__filename": "/opt/cribl/default/cribl/functions/aggregate_metrics/index.js",
      "disabled": false,
      "group": "Metrics",
      "handleSignals": true,
      "id": "aggregate_metrics",
      "loadTime": 1768434545385,
      "modTime": 1763503897000,
      "name": "Aggregate Metrics",
      "schema": {},
      "sync": true,
      "uischema": {
        "add": {
          "items": {
            "name": {
              "ui:options": {
                "columnWidth": "30%"
              },
              "ui:placeholder": "Enter field name"
            },
            "value": {
              "ui:options": {
                "columnWidth": "70%"
              },
              "ui:placeholder": "Enter value expression",
              "ui:widget": "JsInput"
            }
          },
          "ui:field": "Table"
        },
        "aggregations": {
          "items": {
            "agg": {
              "ui:options": {
                "columnWidth": "80%"
              },
              "ui:placeholder": "Enter an aggregate function",
              "ui:widget": "AggInput"
            },
            "metricType": {
              "ui:options": {
                "allowClear": false,
                "uiEnumOptions": [
                  {
                    "description": "Use the aggregation function to determine the type of metric being output",
                    "label": "Automatic",
                    "value": "automatic"
                  },
                  {
                    "description": "The following functions can be aggregated into a counter metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Counter",
                    "value": "counter"
                  },
                  {
                    "description": "The following functions can be aggregated into a distribution metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`. Distribution is only supported by Datadog.",
                    "label": "Distribution",
                    "value": "distribution"
                  },
                  {
                    "description": "The following functions can be aggregated into a gauge metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Gauge",
                    "value": "gauge"
                  },
                  {
                    "description": "The following functions can be aggregated into a histogram metric: `histogram`",
                    "label": "Histogram",
                    "value": "histogram"
                  },
                  {
                    "description": "The following functions can be aggregated into a summary metric: `summary`",
                    "label": "Summary",
                    "value": "summary"
                  },
                  {
                    "description": "The following functions can be aggregated into a timer metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`",
                    "label": "Timer",
                    "value": "timer"
                  }
                ]
              }
            }
          },
          "ui:field": "Table"
        },
        "flushEventLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to unlimited"
        },
        "flushMemLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to unlimited"
        },
        "groupbys": {
          "ui:field": "Tags",
          "ui:placeholder": "One or more dimensions (or wildcard expression) to group by. Optional."
        },
        "idleTimeLimit": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to the smaller of Time Window and 1 minute"
        },
        "lagTolerance": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          },
          "ui:placeholder": "Defaults to the smaller of Time Window and 1 minute"
        },
        "passthrough": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          }
        },
        "preserveGroupBys": {
          "ui:options": {
            "inline": true,
            "inlineStyle": {
              "padding": "0"
            },
            "width": "50%"
          }
        },
        "sufficientStatsOnly": {
          "ui:options": {
            "inline": true,
            "width": "50%"
          }
        },
        "timeWindow": {
          "ui:options": {
            "labelInline": true,
            "width": "200px"
          }
        },
        "ui:options": {
          "groups": {
            "advanced": {
              "collapsed": true,
              "properties": [
                "flushEventLimit",
                "flushMemLimit",
                "shouldTreatDotsAsLiterals",
                "flushOnInputClose"
              ],
              "title": "Advanced Settings"
            },
            "modeSettings": {
              "collapsed": true,
              "properties": [
                "passthrough",
                "sufficientStatsOnly",
                "preserveGroupBys",
                "prefix"
              ],
              "title": "Output Settings"
            },
            "windowType": {
              "collapsed": true,
              "properties": [
                "cumulative",
                "lagTolerance",
                "idleTimeLimit"
              ],
              "title": "Time Window Settings"
            }
          }
        }
      },
      "version": "0.1"
    }
  ]
}
```

Curl scenario
````
curl 'http://localhost:9000/api/v1/functions/aggregate_metrics' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:146.0) Gecko/20100101 Firefox/146.0' \
  -H 'Accept: application/json' \
  -H 'Accept-Language: en-US,en;q=0.5' \
  -H 'Accept-Encoding: gzip, deflate, br, zstd' \
  -H 'Referer: http://localhost:9000/' \
  -H 'Authorization: Bearer secret otken' \
  -H 'Connection: keep-alive' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'Priority: u=0'
````

```
{"items":[{"name":"Aggregate Metrics","version":"0.1","disabled":false,"handleSignals":true,"group":"Metrics","sync":true,"__filename":"/opt/cribl/default/cribl/functions/aggregate_metrics/index.js","loadTime":1768434545385,"modTime":1763503897000,"id":"aggregate_metrics","schema":{"type":"object","title":"","required":["timeWindow","aggregations"],"properties":{"passthrough":{"type":"boolean","title":"Passthrough mode","description":"Pass through the original events along with the aggregation events","default":false},"preserveGroupBys":{"type":"boolean","title":"Preserve group by fields","description":"Preserve the structure of the original aggregation event's groupby fields","default":false},"sufficientStatsOnly":{"type":"boolean","title":"Sufficient stats mode","description":"Output only statistics that are sufficient for the supplied aggregations","default":false},"prefix":{"type":"string","title":"Output prefix","description":"A prefix that is prepended to all of the fields output by this Aggregations Function"},"timeWindow":{"pattern":"\\d+[sm]$","type":"string","title":"Time window","description":"The time span of the tumbling window for aggregating events. Must be a valid time string (such as 10s).","default":"10s"},"aggregations":{"type":"array","title":"Aggregates","description":"Combination of Aggregation function and output metric type","minItems":1,"items":{"type":"object","required":["agg","metricType"],"additionalProperties":false,"properties":{"metricType":{"title":"Metric type","description":"The output metric type","type":"string","enum":["automatic","counter","distribution","gauge","histogram","summary","timer"],"default":"automatic"},"agg":{"title":"Aggregation","type":"string","description":"Aggregate function to perform on events. Example: sum(bytes).where(action=='REJECT').as(TotalBytes)","aggregationExpression":true}}}},"groupbys":{"type":"array","title":"Group by dimensions","description":"Optional: One or more dimensions to group aggregates by. Supports wildcard expressions. Wrap dimension names in quotes if using literal identifiers, such as 'service.name'. Warning: Using wildcard '*' causes all dimensions in the event to be included, which can result in high cardinality and increased memory usage. Exclude dimensions that can result in high cardinality before using wildcards. Example: !_time, !_numericValue, *","items":{"type":"string"}},"flushEventLimit":{"type":"number","title":"Aggregation event limit","description":"The maximum number of events to include in any given aggregation event","minimum":1},"flushMemLimit":{"type":"string","title":"Aggregation memory limit","description":"The memory usage limit to impose upon aggregations. Defaults to unlimited (all available system memory). Accepts numerals with units like KB and MB (example: 4GB).","pattern":"^\\d+\\s*(?:\\w{2})?$"},"cumulative":{"type":"boolean","title":"Cumulative aggregations","description":"Enable to retain aggregations for cumulative aggregations when flushing out an aggregation table event. When disabled (the default), aggregations are reset to 0 on flush.","default":false},"shouldTreatDotsAsLiterals":{"type":"boolean","title":"Treat dots as literals","description":"Treat dots in dimension names as literals. This is useful for top-level dimensions that contain dots, such as 'service.name'.","default":true},"add":{"title":"Evaluate fields","description":"Set of key-value pairs to evaluate and add/set","type":"array","items":{"type":"object","required":["value"],"properties":{"name":{"type":"string","title":"Name"},"value":{"type":"string","title":"Value expression","description":"JavaScript expression to compute the value (can be constant)","jsExpression":true}}}},"flushOnInputClose":{"type":"boolean","title":"Flush on stream close","description":"Flush aggregations when an input stream is closed. If disabled, Time Window Settings control flush behavior.","default":true}},"dependencies":{"cumulative":{"oneOf":[{"properties":{"cumulative":{"enum":[true]}}},{"properties":{"cumulative":{"enum":[false]},"lagTolerance":{"type":"string","title":"Lag tolerance","description":"The tumbling window tolerance to late events. Must be a valid time string (such as 10s).","pattern":"\\d+[sm]$"},"idleTimeLimit":{"type":"string","title":"Idle bucket time limit","description":"How long to wait before flushing a bucket that has not received events. Must be a valid time string (such as 10s).","pattern":"\\d+[sm]$"}}}]}}},"uischema":{"aggregations":{"ui:field":"Table","items":{"metricType":{"ui:options":{"allowClear":false,"uiEnumOptions":[{"value":"automatic","label":"Automatic","description":"Use the aggregation function to determine the type of metric being output"},{"value":"counter","label":"Counter","description":"The following functions can be aggregated into a counter metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`"},{"value":"distribution","label":"Distribution","description":"The following functions can be aggregated into a distribution metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`. Distribution is only supported by Datadog."},{"value":"gauge","label":"Gauge","description":"The following functions can be aggregated into a gauge metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`"},{"value":"histogram","label":"Histogram","description":"The following functions can be aggregated into a histogram metric: `histogram`"},{"value":"summary","label":"Summary","description":"The following functions can be aggregated into a summary metric: `summary`"},{"value":"timer","label":"Timer","description":"The following functions can be aggregated into a timer metric: `avg`, `count`, `distinct_count`, `dc`, `earliest`, `latest`, `first`, `histogram`, `last`, `max`, `min`, `median`, `mode`, `perc`, `per_second`, `rate`, `stdev`, `stdevp`, `sum`, `sumsq`, `summary`, `variance`, `variancep`"}]}},"agg":{"ui:widget":"AggInput","ui:options":{"columnWidth":"80%"},"ui:placeholder":"Enter an aggregate function"}}},"groupbys":{"ui:field":"Tags","ui:placeholder":"One or more dimensions (or wildcard expression) to group by. Optional."},"timeWindow":{"ui:options":{"labelInline":true,"width":"200px"}},"flushEventLimit":{"ui:options":{"inline":true,"width":"50%"},"ui:placeholder":"Defaults to unlimited"},"flushMemLimit":{"ui:options":{"inline":true,"width":"50%"},"ui:placeholder":"Defaults to unlimited"},"passthrough":{"ui:options":{"inline":true,"width":"50%"}},"sufficientStatsOnly":{"ui:options":{"inline":true,"width":"50%"}},"preserveGroupBys":{"ui:options":{"inline":true,"width":"50%","inlineStyle":{"padding":"0"}}},"lagTolerance":{"ui:options":{"inline":true,"width":"50%"},"ui:placeholder":"Defaults to the smaller of Time Window and 1 minute"},"idleTimeLimit":{"ui:options":{"inline":true,"width":"50%"},"ui:placeholder":"Defaults to the smaller of Time Window and 1 minute"},"add":{"ui:field":"Table","items":{"name":{"ui:options":{"columnWidth":"30%"},"ui:placeholder":"Enter field name"},"value":{"ui:widget":"JsInput","ui:options":{"columnWidth":"70%"},"ui:placeholder":"Enter value expression"}}},"ui:options":{"groups":{"windowType":{"title":"Time Window Settings","collapsed":true,"properties":["cumulative","lagTolerance","idleTimeLimit"]},"modeSettings":{"title":"Output Settings","collapsed":true,"properties":["passthrough","sufficientStatsOnly","preserveGroupBys","prefix"]},"advanced":{"title":"Advanced Settings","collapsed":true,"properties":["flushEventLimit","flushMemLimit","shouldTreatDotsAsLiterals","flushOnInputClose"]}}}}}],"count":1}
```