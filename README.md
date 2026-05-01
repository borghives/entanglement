# entanglement

Entangle complexity between different systems knowledge to assure higher probability of intended operation. 

`entanglement` is a Go library designed to manage state, cryptographically link (entangle) properties, and define state correlations for complex systems. It relies on `github.com/borghives/websession` to verify context integrity via token validation.

## Installation

```bash
go get github.com/borghives/entanglement
```

## Core Concepts

### SystemFrame
A `SystemFrame` defines a contextual boundary for a given state. It includes:
- **Name**: The identifier for the frame.
- **Nonce**: A unique value for salt generation.
- **Token**: A verification token to ensure the frame's integrity.
- **Properties**: A key-value map of properties that make up the entangled state.

### State Correlation
The library provides `StateCorrelation` and `TypeStateCorrelation` to map how entities transition between states. This allows you to validate whether a state transition is permitted based on predefined correlations.

### Token Verification
Frames can generate and verify tokens using a `websession.Session`. The token is derived from a salt built using the frame's nonce and name.

## Usage

### Creating and Entangling a Frame

```go
import (
	"github.com/borghives/entanglement"
)

// Initialize a base frame
baseFrame := entanglement.Create("unique-nonce-123", "expected-token-val")

// Create a subframe and entangle properties
subFrame := baseFrame.CreateSubFrame("user-context").
	EntangleProperty("status", "active").
	EntangleProperty("role", "admin")

// Calculate the cryptographic state of all entangled properties
stateHash := subFrame.CalculateEntangledState()
```

### State Correlations

```go
import (
	"github.com/borghives/entanglement"
)

correlations := make(entanglement.TypeStateCorrelation)

// Define allowed transitions for a "Document" frame
// e.g., "draft" can transition to "published"
correlations.AddCorrelation("Document", "draft", "published")

props := entanglement.EntangleProperties{}
props.UpdateCorrelationProperties(correlations)
```

## Dependencies
- [github.com/borghives/websession](https://github.com/borghives/websession)