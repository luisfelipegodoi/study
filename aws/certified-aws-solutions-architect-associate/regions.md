# AWS Regions

- AWS has Regions all around the world
- AWS Regions are clusters of data centers
- Example of regions: us-east-1, eu-west-1, sa-east-1, etc.
- Most services provided by AWS are regions scoped. This means a data for a service used in one region is not replicated in another region
- How to choice AWS Region? Compliance, Proximity, Available Services, Pricing.

## AWS Availability Zones (AZ)

- AWS Regions are divided into availability zones
- Each region has between 2 and 6 AZs. Usually they have 3
- Each AZ is one or more discrete data center with redundant power, networking an connectivity
- All AZs from a region are geographically separated from each other
- Even if they are separated, they have high bandwidth connection between them

## AWS Edge Locations / Points of Presence

- PoP: Points of Presence. Located in many peace of world, used to delivery content to final users.
- Amazon has 400+ points of presence (400+ Edge Locations & 10+ Regional Caches) in 90+ cities across 40+ countries