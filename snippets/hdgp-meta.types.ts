// SPDX-License-Identifier: Apache-2.0
/**
 * HDGP-Core Meta-only guidance shape (development-time weaving).
 * Not an Engine/Judge API — see `schemas/hdgp-core-meta.schema.json` and `docs/ADOPTION_BUNDLE.md`.
 */

export type HdgpcRiskLevel = "low" | "medium" | "high" | "critical";

export interface HdgpcActor {
  type?: string;
  role?: string;
  [key: string]: unknown;
}

export interface HdgpcScene {
  /** INTEGRATION_SPEC §2.4 default: `general` */
  domain?: string;
  /** default: `chat` */
  intent?: string;
  /** default: `medium` */
  risk_level?: HdgpcRiskLevel;
  sensitivity?: string[];
  [key: string]: unknown;
}

export interface HdgpcPolicyRef {
  spec_version?: string;
  strategy_id?: string;
  bundles?: string[];
  override_flags?: string[];
  [key: string]: unknown;
}

/** Root object: `scene` is required by the JSON Schema. */
export interface HdgpcCoreMeta {
  request_id?: string;
  locale?: string;
  channel?: string;
  actor?: HdgpcActor;
  scene: HdgpcScene;
  policy?: HdgpcPolicyRef;
  [key: string]: unknown;
}

/** Defaults from `spec/HDGP_INTEGRATION_SPEC.md` §2.4 */
export function hdgpcDefaultScene(): HdgpcScene {
  return { domain: "general", intent: "chat", risk_level: "medium" };
}
