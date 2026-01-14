# reproFunctionGet

Tests SDK behavior with function retrieval and schema inspection.

## SDK Version

This repository uses **`github.com/criblio/cribl-control-plane-sdk-go v0.5.0-beta.21`** for testing.

This version corresponds to the TypeScript SDK release [v0.5.0-beta.21](https://github.com/criblio/cribl-control-plane-sdk-typescript/releases/tag/v0.5.0-beta.21).

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