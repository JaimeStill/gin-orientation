# AI-Assisted Development Workflow: A Pragmatic Starting Point

## Introduction

This workflow applies the scientific method to development projects by treating each advancement as a testable hypothesis. It integrates AI assistance at strategic points while maintaining human judgment for critical decisions. The approach minimizes context switching, focuses on incremental complexity, and builds verified components that compound into larger solutions.

The workflow operates in three distinct phases: Project Development (architectural planning), Planning (experiment design), and Execution (implementation with AI assistance). Each phase has specific objectives, constraints, and AI integration patterns designed to optimize productivity while preventing common pitfalls like scope creep and premature optimization.

## Quick Reference

### Phase Overview

1. **Project Development**: Define vision, create roadmap, establish building blocks
2. **Planning**: Design minimal experiments with clear success criteria  
3. **Execution**: Implement with AI assistance, evaluate results, iterate

### Core Principles

- Start with minimum variables needed for success
- Test one hypothesis at a time with measurable outcomes
- Build complexity through proven components
- Maintain process discipline even when tempted to skip steps
- Use AI for resource generation, humans for judgment

### File Structure

```
project-root/
├── experiments/
│   └── X.X-name/
│── memory/
│   ├── architecture-decisions.md
│   ├── current-state.md
│   ├── development-constraints.md
│   ├── failed-experiments.md
│   ├── learning-insights.md
│   └── verified-patterns.md
│── prompts/
│── src/
│── claude.md
│── readme.md
│── roadmap.md
└── workflow.md
```

## Project Development Phase Prompt

```
I want to develop [PROJECT DESCRIPTION]. Help me establish a science-driven development approach:

OBJECTIVES:
1. Create an analogical framework for understanding this domain
2. Identify fundamental building blocks (minimal viable components)
3. Map dependencies between components  
4. Design a roadmap with 3-5 stages from basic to complete
5. Define the first minimal viable experiment

CONSTRAINTS:
- Each stage should build on previous verified components
- Individual experiments must be able to complete in 15-20 minutes
- Focus on conceptual architecture, not implementation details
- Identify potential failure points and mitigation strategies

CONTEXT:
- I learn through analogical thinking and hands-on experimentation
- I prefer markdown documentation that lives alongside code
- I will use Claude Code for implementation assistance
- Testing and documentation happen in separate roadmap stages

Do not provide code or implementation details. Focus on architectural understanding and experimental design.
```

## Project Layout and Documentation

### roadmap.md

```markdown
# Project Roadmap: [Project Name]

## Vision
[Clear statement of end goal and success criteria]

## Analogical Framework
[Tangible metaphor for understanding the domain - like spelunkers for panic/recover]

## Fundamental Building Blocks
- **Block A**: [Core component] - handles [specific responsibility]
- **Block B**: [Integration component] - coordinates [specific interaction]
- **Block C**: [Interface component] - manages [specific boundary]

## Development Stages

### Stage 1: Foundation ([Component A])
**Objective**: Establish core functionality
**Success Criteria**: [Specific measurable outcome]
**Key Tasks**:
- [ ] Hypothesis 1.1: [Minimal viable implementation]
- [ ] Hypothesis 1.2: [Basic error handling]
- [ ] Hypothesis 1.3: [Core behavior validation]

**Dependencies**: None
**Estimated Experiments**: 3-4
**Failure Risks**: [Known pitfalls and mitigation strategies]

### Stage 2: Integration ([Component A + B])
**Objective**: Coordinate components
**Success Criteria**: [Specific interaction outcome]
**Key Tasks**:
- [ ] Hypothesis 2.1: [Basic coordination]
- [ ] Hypothesis 2.2: [Error propagation]
- [ ] Hypothesis 2.3: [Performance validation]

**Dependencies**: Stage 1 complete
**Estimated Experiments**: 4-5
**Failure Risks**: [Integration complexity, timing issues]

### Stage 3: Interface ([Component A + B + C])
**Objective**: External interaction capability
**Success Criteria**: [User-facing functionality]
**Key Tasks**:
- [ ] Hypothesis 3.1: [Basic interface]
- [ ] Hypothesis 3.2: [Input validation]
- [ ] Hypothesis 3.3: [Output formatting]

**Dependencies**: Stage 2 complete
**Estimated Experiments**: 3-4
**Failure Risks**: [Interface complexity, user experience]

### Stage 4: Critical Infrastructure Testing
**Objective**: Verify critical paths and error conditions
**Success Criteria**: Critical infrastructure proven reliable
**Testing Focus**: Error boundaries, resource limits, failure recovery

### Stage 5: Documentation & Polish
**Objective**: Self-documenting code and clear API
**Success Criteria**: Code intent clear to future developers
**Focus**: Comments, examples, usage patterns
```

### memory/current-state.md

```markdown
# Current Project State

**Last Updated**: [Date]
**Current Stage**: [Stage number and name]
**Active Hypothesis**: [Current experiment]

## Completed Building Blocks
- [Block A]: Verified - handles [responsibility]
- [Block B]: In progress - [current status]

## Architecture Decisions
- **Decision 1**: [Choice made] because [reasoning]
- **Decision 2**: [Choice made] because [reasoning]

## Current Context for AI
**What works**: [Verified patterns and components]
**Current focus**: [Specific problem being solved]
**Known constraints**: [Technical or business limitations]
**Next logical step**: [Planned experiment]

## Key Learnings
- [Insight 1]: [What was learned and why it matters]
- [Insight 2]: [Behavioral understanding gained]

## Context Checkpoint
[Brief summary suitable for initializing AI context in new sessions]
```

### development-constraints.md

```markdown
# Development Constraints

## ACTIVE BOUNDARIES (Always Enforce)
- Package completion takes priority over all other concerns
- No testing implementation until package reaches completion milestone
- No documentation/comments until package reaches completion milestone  
- Testing metrics focus on critical infrastructure, not coverage percentages

## PHASE GATES
- **Development Phase**: Focus solely on functional implementation
- **Testing Phase**: Implement critical infrastructure tests only
- **Documentation Phase**: Add self-documenting comments and API docs

## AI INTERACTION RULES
- If AI suggests testing during development phase, redirect to completion focus
- If AI suggests documentation during development phase, redirect to completion focus
- Testing discussions only occur during dedicated testing roadmap stage
- One hypothesis per experiment, minimal variables only

## CONSTRAINT VIOLATIONS LOG
[Track when constraints are violated and lessons learned]
```

### prompts/planning-phase.md

```markdown
# Planning Phase Prompt Template

Context: [Reference to current-state.md]
Constraints: [Reference to development-constraints.md]
Current Phase: DEVELOPMENT

Design an experiment to test: [SPECIFIC HYPOTHESIS]

REQUIREMENTS:
- Single file or action scope (15-20 minute execution)
- Clear, measurable success criteria
- Minimal variables (isolate what you're testing)
- Builds on verified components: [list from current-state.md]

EXPERIMENTAL DESIGN NEEDED:
1. Hypothesis statement (what you expect to happen)
2. Implementation approach (what you'll build/change)
3. Success criteria (how you'll know it worked)
4. Failure conditions (what would invalidate the hypothesis)
5. Required resources (dependencies, tools, knowledge)

CONSTRAINTS:
- This is a DEVELOPMENT phase experiment
- Focus only on functional implementation
- No testing or documentation suggestions
- Must build on existing verified components

Provide experimental design only, not implementation code.
```

### prompts/execution-phase.md

```markdown
# Execution Phase Prompt Template

Execute this experiment: [Reference to experiment file]
Phase: DEVELOPMENT
Expected outcome: [Specific behavioral result]

CONTEXT:
- Current project state: [brief from current-state.md]
- Verified building blocks: [list available components]
- Active constraints: development-constraints.md applies

EXECUTION GUIDELINES:
- Focus on functional implementation only
- Document deviations from experimental design
- Continue unless core hypothesis becomes invalid
- Capture actual behavior vs. expected behavior

If you suggest testing/documentation during this DEVELOPMENT phase, I will redirect you to functional implementation focus only.
```

### claude.md

```markdown
# Claude Code Configuration

## Project Context
This project follows a science-driven development workflow with three phases: Project Development, Planning, and Execution. Currently in [CURRENT PHASE].

## Current Constraints
Reference: .workflow/development-constraints.md
- Development phase: Functional implementation only
- No testing until completion milestone
- No documentation until completion milestone

## Workflow Integration
- Read .workflow/current-state.md before starting sessions
- Reference .workflow/roadmap.md for overall direction
- Follow experiment designs from experiments/ directory
- Update building-blocks/verified-patterns.md with successful components

## Coding Standards
- Single responsibility focus per experiment
- Build on verified components only
- Document deviations in experiment logs
- Prefer simple, readable solutions

## Communication Style
- Ask for clarification if experiment objectives unclear
- Suggest alternative approaches if current path seems invalid
- Focus on helping execute experiments, not designing them
- Provide specific behavioral descriptions of what code does
```

### experiments/1.1-basic-server/readme.md

```markdown
# Experiment: Basic HTTP Server

**Hypothesis**: A minimal HTTP server can handle basic GET requests and return static responses

**Implementation Approach**: 
- Use Go's net/http package
- Single endpoint responding to GET /health
- Static "OK" response
- Minimal error handling

**Success Criteria**:
- Server starts without errors
- Responds to GET /health with 200 status
- Returns "OK" in response body
- Graceful shutdown capability

**Failure Conditions**:
- Server fails to start
- No response to requests
- Unexpected error messages

**Required Resources**:
- Go net/http package knowledge
- Basic HTTP concepts
- Port availability for testing

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
```

### memory/verified-patterns.md

```markdown
# Verified Patterns

## Basic HTTP Server
**Status**: Verified
**Description**: Minimal server handling GET requests
**Code Location**: experiments/hypothesis-001-basic-server/
**Behavior**: Starts on specified port, responds to health checks, shuts down gracefully
**Usage Pattern**: Import server struct, call Start() method, defer Stop()
**Analogical Model**: Simple receptionist answering basic questions

## [Next Pattern]
**Status**: In Development
**Description**: [What it does]
**Behavior**: [How it actually behaves]
```

### memory/architecture-decisions.md

```markdown
# Architecture Decisions

## Decision Log

### AD-001: HTTP Framework Choice
**Date**: [Date]
**Decision**: Use standard library net/http
**Reasoning**: Minimal dependencies, well understood, sufficient for current scope
**Alternatives Considered**: Gin, Echo, Chi
**Status**: Active

### AD-002: Error Handling Strategy  
**Date**: [Date]
**Decision**: [Choice made]
**Reasoning**: [Why this approach]
**Impact**: [How it affects other components]
**Status**: Active
```

### memory/failed-experiments.md

```markdown
# Failed Experiments Log

## Experiment Failures

### FE-001: Complex Router Implementation
**Date**: [Date]
**Hypothesis**: Custom router would be simpler than standard library
**Failure Point**: Too many variables, context switching between routing logic and business logic
**Root Cause**: Violated minimal variables principle
**Learning**: Use standard library until clear need for custom solution
**Prevention**: Better scope definition in planning phase

### FE-002: [Next failure]
**Date**: [Date]
**Hypothesis**: [What was tested]
**Failure Point**: [Where it broke down]
**Learning**: [What this taught us]
```

### memory/insights.md

```markdown
# Learning Insights

## Analogical Models Developed

### HTTP Server as Receptionist
**Concept**: HTTP request handling
**Analogy**: Office receptionist fielding phone calls
- Requests are phone calls with specific questions
- Routes are different departments to transfer calls
- Middleware is call screening and logging
- Response is the answer provided back to caller

### [Next Analogy]
**Concept**: [Programming concept]
**Analogy**: [Real-world comparison]
**Key Behaviors**: [How the analogy maps to code behavior]

## Behavioral Insights
- [Insight 1]: [Understanding about how something actually works]
- [Insight 2]: [Pattern observed across multiple experiments]
```

## Git Flow Practices

### Commit Message Convention

```
type(scope): description

Types:
- experiment: New experiment or modification
- verify: Successful experiment verification  
- fail: Failed experiment documentation
- roadmap: Changes to project roadmap
- workflow: Workflow or process changes

Examples:
experiment(http-server): test basic GET request handling
verify(http-server): minimal server responds correctly
fail(routing): custom router too complex for current scope
roadmap(stages): add testing stage after core completion
```

### Branch Strategy

```
main
├── experiments/hypothesis-001-basic-server
├── experiments/hypothesis-002-error-handling  
└── stage-1-foundation
```

**Branch Rules**:

- `main`: Verified building blocks only
- `experiments/X.X-name`: Individual experiment branches
- `stage-N-name`: Stage completion branches
- Merge to main only after experiment verification
- Tag major stage completions

## Context Checkpoint Strategy

### Daily Workflow

1. **Session Start**: Review `memory/current-state.md`
2. **Session End**: Update current state with progress/learnings
3. **Experiment Complete**: Update `memory/verified-patterns.md` or `memory/failed-experiments.md`

### Weekly Workflow  

1. **Review**: Analyze completed experiments for patterns
2. **Update**: Refresh roadmap based on learnings
3. **Archive**: Move completed stages to memory banks
4. **Plan**: Identify next week's experimental focus

### Context Rotation

- Keep current-state.md under 200 lines for AI context efficiency
- Archive historical context to memory-banks/ when reaching limits
- Maintain "last 3 experiments" summary in current state
- Reference memory banks by key decision or learning rather than full content

This workflow balances systematic rigor with practical flexibility, allowing natural evolution while maintaining the scientific approach that drives reliable progress.
