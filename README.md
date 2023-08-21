# Project 3: Analysis with Spark

The provided dataset contains records of all legitimate felony, misdemeanor, and violation crimes reported to the New York City Police Department (NYPD) spanning from 2006 to the conclusion of the previous year (2019).

Dataset: https://data.cityofnewyork.us/Public-Safety/NYPD-Complaint-Data-Historic/qgea-i56i

- The dataset's accuracy is a close approximation of current records due to potential revisions and updates. 
- Null values may appear in certain fields due to changes in data collection or unavailable/unknown information.
- Crime complaints with multiple offenses are categorized based on the most serious offense.
- Attempted crimes are recorded regardless of success, except for Attempted Murder, which is recorded as Felony Assault.
- Date and time information is provided for reported crime incidents.
- The Report Date represents when the incident was reported, not necessarily when it occurred.
- The dataset discourages attempts to match incident locations to exact addresses or link to other datasets.
- Offenses occurring in open areas may be geo-coded as happening on nearby streets or intersections.
- Certain offenses may be located at police station houses or precincts for victim identity protection.
- The dataset provides specific coordinate systems for location information.
- The data represents criminal offenses based on New York State Penal Law definitions, not FBI Uniform Crime Report definitions.
- Errors in data transcription may lead to nominal data inconsistencies.
- Only valid complaints are included, while unfounded or voided complaints are excluded.
- Investigation reports that don't imply criminal activity are also excluded.

## Columns and their Decription
- ADDR_PCT_CD: The precinct in which the incident occurred
- BORO: The name of the borough in which the incident occurred
- CMPLNT_FR_DT: Exact date of occurrence for the reported event (or starting date if CMPLNT_TO_DT exists)
- CMPLNT_FR_TM: Exact time of occurrence for the reported event (or starting time if CMPLNT_TO_TM exists)
- CMPLNT_NUM: Randomly generated persistent ID for each complaint
- CMPLNT_TO_DT: Ending date of occurrence for the reported event if exact time is unknown
- CMPLNT_TO_TM: Ending time of occurrence for the reported event if exact time is unknown
- CRM_ATPT_CPTD_CD: Indicator of whether the crime was successfully completed, attempted but failed, or interrupted prematurely
- HADEVELOPT: Name of NYCHA housing development of occurrence, if applicable
- HOUSING_PSA: Development Level Code for housing
- JURIS_DESC: Description of the jurisdiction code
- JURISDICTION_CODE: Jurisdiction responsible for the incident (internal or external)
- KY_CD: Three-digit offense classification code
- LAW_CAT_CD: Level of offense: felony, misdemeanor, violation
- LOC_OF_OCCUR_DESC: Specific location of occurrence in or around the premises
- Longitude: Midblock Longitude coordinate for Global Coordinate System
- Latitude: Midblock Latitude coordinate for Global Coordinate System
- OFNS_DESC: Description of offense corresponding with key code
- PARKS_NM: Name of NYC park, playground, or greenspace of occurrence, if applicable
- PATROL_BORO: Name of the patrol borough in which the incident occurred
- PD_CD: Three-digit internal classification code (more granular than Key Code)
- PD_DESC: Description of internal classification corresponding with PD code
- PREM_TYP_DESC: Specific description of premises (grocery store, residence, street, etc.)
- RPT_DT: Date event was reported to the police
- STATION_NAME: Transit station name
- SUSP_AGE_GROUP: Suspect's Age Group
- SUSP_RACE: Suspect's Race Description
- SUSP_SEX: Suspect's Sex Description
- TRANSIT_DISTRICT: Transit district in which the offense occurred
- VIC_AGE_GROUP: Victim's Age Group
- VIC_RACE: Victim's Race Description
- VIC_SEX: Victim's Sex Description
- X_COORD_CD: X-coordinate for New York State Plane Coordinate System
- Y_COORD_CD: Y-coordinate for New York State Plane Coordinate System

Depict crime rate based on location, time. Type of offense based on location

## Insights to gain using the data
- Types of crimes and their occurences.
  - Hypothesis: More assault crimes are reported than harrassment crimes.
- Variance in crime rate by year, month and hour.
  - Hypothesis: Crime rates has increased since covid-19 pandemic.
- Crime reports grouped by age, gender and ethnicity.
  - Hypothesis: Certain age groups, genders, and ethnicities appear to have a higher number of victims in relation to a specific crime.

## Questions 
- Find type of crimes reported and validate the hypothesis? 
  <img width="928" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/a137a664-3f95-43c1-8819-3c273543134f"> 
  - Our hypothesis is not true. Reported assault crimes are less compared to harrassment. 
- What is the crime rate trend in given dataset? Has crime rate increased or decreased?
  <img width="891" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/c82622ad-7d53-4b17-b20e-0ddb2531a0f7">
  - Our hypothesis is true. Crime rate has increased since pandemic.   
  <img width="604" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/6a7c2df3-d036-4fc2-923b-78d8968e521d">
  
  - There is a higher occurrence of crime events during the afternoon until late evening compared to the early morning.
- How do sex crime rates vary across different age groups, genders, and ethnicities?
  <img width="619" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/c540fe30-7c62-4baf-9e93-1f477d33f63b">
  <img width="595" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/93c7acf1-8d55-4798-a012-7251ceaa508d">
  <img width="805" alt="image" src="https://github.com/usf-cs677-sp23/P3-p3/assets/109916498/3ec3f248-90d2-443a-9ec7-550cbc0278d4">
  - Based on the graphical data, it can be concluded that sexual crimes are more prevalent among individuals under the age of 18, individuals of black race, and females.
 
