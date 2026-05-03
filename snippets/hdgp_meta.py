# SPDX-License-Identifier: Apache-2.0
"""HDGP-Core Meta-only guidance shape (development-time weaving).

Not an Engine/Judge API — see ``schemas/hdgp-core-meta.schema.json`` and ``docs/ADOPTION_BUNDLE.md``.

Requires Python 3.11+ (``typing.NotRequired``). For 3.10, copy this file and import ``NotRequired`` from ``typing_extensions``.
"""

from __future__ import annotations

from typing import NotRequired, TypedDict


class HdgpcActor(TypedDict, total=False):
    type: str
    role: str


class HdgpcScene(TypedDict, total=False):
    domain: str
    intent: str
    risk_level: str
    sensitivity: list[str]


class HdgpcPolicyRef(TypedDict, total=False):
    spec_version: str
    strategy_id: str
    bundles: list[str]
    override_flags: list[str]


class HdgpcCoreMeta(TypedDict):
    """``scene`` is required by the JSON Schema; other fields are optional."""

    scene: HdgpcScene
    request_id: NotRequired[str]
    locale: NotRequired[str]
    channel: NotRequired[str]
    actor: NotRequired[HdgpcActor]
    policy: NotRequired[HdgpcPolicyRef]


def hdgpc_default_scene() -> HdgpcScene:
    """Defaults from ``spec/HDGP_INTEGRATION_SPEC.md`` §2.4."""
    return {"domain": "general", "intent": "chat", "risk_level": "medium"}
