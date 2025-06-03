# Experiment 1.1: Basic Gin Router Handles Multiple Endpoints

## Hypothesis

Gin router handles GET/POST to different endpoints

## Implementation Approach

- Initialize Gin router with `gin.Default()`
- Register GET endpoint at `/health` returning simple status
- Register POST endpoint at `/echo` returning request confirmation
- Start server on localhost:8080

## Success Criteria

- Server starts without errors
- GET /health returns 200 with "OK" message  
- POST /echo returns 200 with confirmation message
- Server responds within reasonable time (<100ms)

## Failure Conditions

- Server fails to start
- Endpoints return 404 or wrong status codes
- Server crashes on valid requests

## Required Resources

- Go environment with Gin framework (`go get github.com/gin-gonic/gin`)
- Basic HTTP testing capability (curl/Postman/browser)
- Understanding of HTTP methods and status codes

## Measurement Method

Manual testing with curl commands, observe server logs for request handling

## Execution Log

**Date**: [Date]
**Duration**: [Time spent]
**Result**: [SUCCESS/FAILURE]

### What Happened

[Actual behavior observed]

### Deviations from Design

[Any changes made during implementation]

### Key Learnings

[Behavioral insights gained]

### Next Steps

[What this enables for next experiment]
