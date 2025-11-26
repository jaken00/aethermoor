# Plan for project

## Food and Shelter Movement
Currently running into an issue where we are having these large switch statements and this will eventually become HELL when we want to actually start involving humans and other 
creatures into the world
We need to create a map to determine where these entities need to go for each of these resources
The map will work like: ["RESOURCE"] {SLICE OF TERRAINS THAT HAVE THIS RESOURCE} -> I have no clue how ill do this but I think this will be best and you just look up O(1) to get what terrains 
After we have the list of terrains we want for this resource. We will then get all of the Vec2s that are touching current pos and compare the Resource list to the list that we have. 
If we check the Cell and we can check current entities -> we will do this ONLY to see if there are entities on the aversion list. -> THis means we are looking oh is there any wolves here? no?
okay going in and trying to eat
THere once they are inside the cEll we need seperate logic to eat until threshold? maybe we do 6
Once we hit priority 1 -> go to priority 2..3.. etc
if no priorities go home! 
THen we tick Shelter. We may need to edit shelter beuase it shouldnt havbe a consume rate and is technically not a "NEED" but we need extra logic for reproduction so that rabbits will reproduce
as well as wolves die off if they cant find food
we need to do a death check before we do 