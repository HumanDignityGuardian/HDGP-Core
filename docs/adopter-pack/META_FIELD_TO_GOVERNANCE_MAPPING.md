# Meta fields ↔ governance/review questions (HDGP-Core · Meta-only)

Adopters usually ask: “Which Meta fields should we fill, and what governance questions do they answer?”  
This file makes that mapping **explicit**, without implying runtime enforcement.

## Field categories

- **Transparency / statement fields**: help external readers understand scope, assumptions, and disclosures.
- **Routing / review hint fields**: help your team route items to review queues and produce consistent logs.

## Mapping table (minimal)

### `scene.domain`

- **Answers**: “What domain is this output in?” “Which domain-specific review checklist applies?”
- **Typical governance questions**:
  - Which domain policies apply (medical/finance/education/compliance…)?  
  - Does this domain require higher source standards or legal review?
- **Category**: routing / review hint

### `scene.intent`

- **Answers**: “What is the system trying to do here?” “Is this advice/decision support/notification?”
- **Typical governance questions**:
  - Is the system framing output as advice or as a final decision?  
  - Do we need to add “human final decision” reminders in this intent?
- **Category**: routing / review hint (also impacts transparency copy)

### `scene.risk_level`

- **Answers**: “How high-stakes is this output path?” “Do we require human review?”
- **Typical governance questions**:
  - What is the minimum review requirement?  
  - What is the allowed failure mode (e.g. do we stop output when checks are unavailable)?
- **Category**: routing / review hint

### `scene.sensitivity[]`

- **Answers**: “What sensitive aspects exist (privacy/minors/medical/financial)?”
- **Typical governance questions**:
  - Are there special disclosure needs (data boundary, consent, retention)?  
  - Are there special content restrictions or extra reviewer training needs?
- **Category**: routing / review hint

### `request_id`

- **Answers**: “Can we trace this output to a review thread / dispute / incident?”
- **Typical governance questions**:
  - Can we reproduce and explain outcomes?  
  - Can we do incident response without guessing?
- **Category**: routing / traceability

### `policy.*` (semantic references only)

- **Answers**: “Which spec/profile/bundle did we *intend* to align with?”
- **Typical governance questions**:
  - Are we consistent across teams/tenants?  
  - Can we explain “which policy set was referenced” in a governance report?
- **Category**: transparency + future routing hints  
  - **Important**: this is **not** a claim that any Engine evaluated these.

## Disclosure and sources (recommended adopter additions)

HDGP-Core schema intentionally stays small. Many adopters still want two explicit fields:

- **`adopter.<org>.disclosure`**: sponsorship/affiliate/relationship disclosures; “anti-promo” posture.
- **`adopter.<org>.sources`**: source list / provenance / thresholds for factual claims.

These map to governance questions like:

- “Are we promoting or diverting traffic commercially?”  
- “Are we presenting unsupported opinions as facts?”  

If you add these fields, document them in your adopter pack and keep them stable.

