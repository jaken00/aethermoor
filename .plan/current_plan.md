GOALS FOR HOME TIME
1. MOVEMENT & AP OVERHAUL
    Implement an Action Point (AP) System
        Instead of "80% chance," give each entity an AP pool.
        Ticks grant AP += Speed.
        Movement costs a flat 10 AP.
        This allows "Slow" entities (Speed 2) to move every 5 ticks and "Fast" entities (Speed 20) to move twice per tick.

    Enforce Activity Locks
        Logic check: if currentActivity != NullActivity { canMove = false }.
        Eating or Resting "channels" for X ticks, locking movement.
    
    Terrain Movement Costs
        Woods cost 15 AP, Mountains cost 25 AP, Plains cost 10 AP.
        This naturally creates "paths of least resistance" for your entities.

2. COMBAT & INJURY SYSTEM

    Implement "Bump" Combat with Stat Scaling
        Damage calculation should account for basic defense:
        Damage=max(1,Attacker.Attack−Defender.Defense)

    The "Injury" Speed Penalty
        If Health < 50%, apply a speed multiplier: Speedcurrent​=Speedbase​×0.5.
        Why: This stops the "infinite loop" chase. Once a wolf bites a rabbit, the rabbit slows down, allowing the wolf to finish the kill.

    Loot & Carcass Generation
        When an entity dies, it should leave behind a Corpse  object at that Vec2 instead of just disappearing.

3. HUMAN GENERATION & ROLES

    Cluster Spawning (Settlement Logic)
        Instead of random spawning, humans spawn in a "Family" group of 3–5 in a TerrainPlains or TerrainWoods tile.
        They share a single Home Vec2 (The Camp).

    Introduction of "Resource Hoarding"
        Give Humans a Backpack (a slice of ResourceEntry).
        Unlike wolves who eat immediately, Humans should "Forage" until their backpack is full, then return to their Home coordinate to "Deposit" resources.
        This creates "Traffic Patterns" that predators (Wolves) can learn to camp.

4. RENOWN & MYTHOLOGY

    The "Kill Credit" System
        When an entity dies, the attacker gains a percentage of the victim's Renown.
        Renowngain​=max(10,Victim.Renown×0.2)

    Title Evolution

        Create a "Title Map" based on Renown thresholds:
            0–100: "Fledgling"
            101–500: "Veteran"
            501–2000: "Elite
            2000+: "Legend"

    The "Legacy" Log
        If an entity with >500 Renown dies, write their entire History []string to a permanent legends.json file.
        This is the "Graveyard" your LLM will eventually use to tell stories about "The fall of Grimclaw."