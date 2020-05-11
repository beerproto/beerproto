# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [beer.proto](#beer.proto)
    - [AcidityType](#beerproto.AcidityType)
    - [BitternessRangeType](#beerproto.BitternessRangeType)
    - [BitternessType](#beerproto.BitternessType)
    - [BoilProcedureType](#beerproto.BoilProcedureType)
    - [BoilStepType](#beerproto.BoilStepType)
    - [CarbonationRangeType](#beerproto.CarbonationRangeType)
    - [CarbonationType](#beerproto.CarbonationType)
    - [ColorRangeType](#beerproto.ColorRangeType)
    - [ColorType](#beerproto.ColorType)
    - [ConcentrationType](#beerproto.ConcentrationType)
    - [CultureAdditionType](#beerproto.CultureAdditionType)
    - [CultureInformation](#beerproto.CultureInformation)
    - [CultureInventoryType](#beerproto.CultureInventoryType)
    - [DiastaticPowerType](#beerproto.DiastaticPowerType)
    - [EfficiencyType](#beerproto.EfficiencyType)
    - [EquipmentItemType](#beerproto.EquipmentItemType)
    - [EquipmentType](#beerproto.EquipmentType)
    - [FermentableAdditionType](#beerproto.FermentableAdditionType)
    - [FermentableInventoryType](#beerproto.FermentableInventoryType)
    - [FermentableType](#beerproto.FermentableType)
    - [FermentationProcedureType](#beerproto.FermentationProcedureType)
    - [FermentationStepType](#beerproto.FermentationStepType)
    - [GravityRangeType](#beerproto.GravityRangeType)
    - [GravityType](#beerproto.GravityType)
    - [HopAdditionType](#beerproto.HopAdditionType)
    - [HopInventoryType](#beerproto.HopInventoryType)
    - [IBUEstimateType](#beerproto.IBUEstimateType)
    - [IngredientsType](#beerproto.IngredientsType)
    - [MashProcedureType](#beerproto.MashProcedureType)
    - [MashStepType](#beerproto.MashStepType)
    - [MassType](#beerproto.MassType)
    - [MiscellaneousAdditionType](#beerproto.MiscellaneousAdditionType)
    - [MiscellaneousInventoryType](#beerproto.MiscellaneousInventoryType)
    - [MiscellaneousType](#beerproto.MiscellaneousType)
    - [OilContentType](#beerproto.OilContentType)
    - [PackagingProcedureType](#beerproto.PackagingProcedureType)
    - [PackagingVesselType](#beerproto.PackagingVesselType)
    - [PercentRangeType](#beerproto.PercentRangeType)
    - [PercentType](#beerproto.PercentType)
    - [Recipe](#beerproto.Recipe)
    - [RecipeStyleType](#beerproto.RecipeStyleType)
    - [RecipeType](#beerproto.RecipeType)
    - [SpecificHeatType](#beerproto.SpecificHeatType)
    - [SpecificVolumeType](#beerproto.SpecificVolumeType)
    - [StyleType](#beerproto.StyleType)
    - [TasteType](#beerproto.TasteType)
    - [TemperatureRangeType](#beerproto.TemperatureRangeType)
    - [TemperatureType](#beerproto.TemperatureType)
    - [TimeType](#beerproto.TimeType)
    - [TimingType](#beerproto.TimingType)
    - [UnitType](#beerproto.UnitType)
    - [VarietyInformation](#beerproto.VarietyInformation)
    - [VolumeType](#beerproto.VolumeType)
    - [WaterAdditionType](#beerproto.WaterAdditionType)
    - [WaterBase](#beerproto.WaterBase)
    - [YieldType](#beerproto.YieldType)
    - [Zymocide](#beerproto.Zymocide)
  
    - [AcidityType.AcidityUnitType](#beerproto.AcidityType.AcidityUnitType)
    - [BitternessType.BitternessUnitType](#beerproto.BitternessType.BitternessUnitType)
    - [BoilStepType.BoilStepTypeChillingType](#beerproto.BoilStepType.BoilStepTypeChillingType)
    - [CarbonationType.CarbonationUnitType](#beerproto.CarbonationType.CarbonationUnitType)
    - [ColorType.ColorUnitType](#beerproto.ColorType.ColorUnitType)
    - [ConcentrationType.ConcentrationUnitType](#beerproto.ConcentrationType.ConcentrationUnitType)
    - [CultureBaseForm](#beerproto.CultureBaseForm)
    - [CultureBaseType](#beerproto.CultureBaseType)
    - [CultureInformation.QualitativeRangeType](#beerproto.CultureInformation.QualitativeRangeType)
    - [DiastaticPowerType.DiastaticPowerUnitType](#beerproto.DiastaticPowerType.DiastaticPowerUnitType)
    - [EquipmentItemType.EquipmentBaseForm](#beerproto.EquipmentItemType.EquipmentBaseForm)
    - [FermentableBaseType](#beerproto.FermentableBaseType)
    - [GrainGroup](#beerproto.GrainGroup)
    - [GravityType.GravityUnitType](#beerproto.GravityType.GravityUnitType)
    - [HopVarietyBaseForm](#beerproto.HopVarietyBaseForm)
    - [IBUEstimateType.IBUMethodType](#beerproto.IBUEstimateType.IBUMethodType)
    - [MashStepType.MashStepTypeType](#beerproto.MashStepType.MashStepTypeType)
    - [MassType.MassUnitType](#beerproto.MassType.MassUnitType)
    - [MiscellaneousBaseType](#beerproto.MiscellaneousBaseType)
    - [PackagingVesselType.PackagingVesselTypeType](#beerproto.PackagingVesselType.PackagingVesselTypeType)
    - [PercentType.PercentUnitType](#beerproto.PercentType.PercentUnitType)
    - [RecipeStyleType.StyleCategories](#beerproto.RecipeStyleType.StyleCategories)
    - [RecipeType.RecipeTypeType](#beerproto.RecipeType.RecipeTypeType)
    - [SpecificHeatType.SpecificHeatUnitType](#beerproto.SpecificHeatType.SpecificHeatUnitType)
    - [SpecificVolumeType.SpecificVolumeUnitType](#beerproto.SpecificVolumeType.SpecificVolumeUnitType)
    - [StyleType.StyleCategories](#beerproto.StyleType.StyleCategories)
    - [TemperatureType.TemperatureUnitType](#beerproto.TemperatureType.TemperatureUnitType)
    - [TimeType.TimeUnitType](#beerproto.TimeType.TimeUnitType)
    - [TimingType.UseType](#beerproto.TimingType.UseType)
    - [UnitType.UnitUnitType](#beerproto.UnitType.UnitUnitType)
    - [VarietyInformation.VarietyInformationType](#beerproto.VarietyInformation.VarietyInformationType)
    - [VolumeType.VolumeUnitType](#beerproto.VolumeType.VolumeUnitType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="beer.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## beer.proto
BeerProto

Another beer format, written in protocol buffer.


<a name="beerproto.AcidityType"></a>

### AcidityType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [AcidityType.AcidityUnitType](#beerproto.AcidityType.AcidityUnitType) |  |  |






<a name="beerproto.BitternessRangeType"></a>

### BitternessRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [BitternessType](#beerproto.BitternessType) |  |  |
| maximum | [BitternessType](#beerproto.BitternessType) |  |  |






<a name="beerproto.BitternessType"></a>

### BitternessType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [BitternessType.BitternessUnitType](#beerproto.BitternessType.BitternessUnitType) |  |  |






<a name="beerproto.BoilProcedureType"></a>

### BoilProcedureType
BoilProcedureType defines the procedure for performing a boil. A boil procedure with no steps is the same as a standard single step boil


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pre_boil_size | [VolumeType](#beerproto.VolumeType) |  |  |
| boil_time | [TimeType](#beerproto.TimeType) |  |  |
| boil_steps | [BoilStepType](#beerproto.BoilStepType) | repeated |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| notes | [string](#string) |  |  |






<a name="beerproto.BoilStepType"></a>

### BoilStepType
BoilStepType - a per step representation of a boil process, can be used to support preboil steps, non-boiling pasteurization steps, boiling, whirlpool steps, and chilling


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| end_gravity | [GravityType](#beerproto.GravityType) |  |  |
| chilling_type | [BoilStepType.BoilStepTypeChillingType](#beerproto.BoilStepType.BoilStepTypeChillingType) |  |  |
| description | [string](#string) |  |  |
| end_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| ramp_time | [TimeType](#beerproto.TimeType) |  | The amount of time that passes before this step begins. eg moving from a boiling step (step 1) to a whirlpool step (step 2) may take 5 minutes. Step 2 would have a ramp time of 5 minutes, hop isomerization and bitterness calculations will need to account for this accordingly. |
| step_time | [TimeType](#beerproto.TimeType) |  |  |
| start_gravity | [GravityType](#beerproto.GravityType) |  |  |
| start_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| end_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| name | [string](#string) |  |  |
| start_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |






<a name="beerproto.CarbonationRangeType"></a>

### CarbonationRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [CarbonationType](#beerproto.CarbonationType) |  |  |
| maximum | [CarbonationType](#beerproto.CarbonationType) |  |  |






<a name="beerproto.CarbonationType"></a>

### CarbonationType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [CarbonationType.CarbonationUnitType](#beerproto.CarbonationType.CarbonationUnitType) |  |  |






<a name="beerproto.ColorRangeType"></a>

### ColorRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [ColorType](#beerproto.ColorType) |  |  |
| maximum | [ColorType](#beerproto.ColorType) |  |  |






<a name="beerproto.ColorType"></a>

### ColorType
ColorType supports both grain color properties, such as Lovibond, and wort color properties such as SRM and EBC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [ColorType.ColorUnitType](#beerproto.ColorType.ColorUnitType) |  |  |






<a name="beerproto.ConcentrationType"></a>

### ConcentrationType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [ConcentrationType.ConcentrationUnitType](#beerproto.ConcentrationType.ConcentrationUnitType) |  |  |






<a name="beerproto.CultureAdditionType"></a>

### CultureAdditionType
CultureAdditionType collects the attributes of each culture ingredient for use in a recipe


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| form | [CultureBaseForm](#beerproto.CultureBaseForm) |  |  |
| product_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| cell_count_billions | [int32](#int32) |  |  |
| times_cultured | [int32](#int32) |  |  |
| producer | [string](#string) |  |  |
| type | [CultureBaseType](#beerproto.CultureBaseType) |  |  |
| attenuation | [PercentType](#beerproto.PercentType) |  | The expected, or measured apparent attenuation for a given culture in a given recipe. In comparison to attenuation range, this is a single value. |
| timing | [TimingType](#beerproto.TimingType) |  | The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step. |
| mass | [MassType](#beerproto.MassType) |  |  |
| unit | [UnitType](#beerproto.UnitType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.CultureInformation"></a>

### CultureInformation
CultureInformation collects the attributes of a microbial culture


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| form | [CultureBaseForm](#beerproto.CultureBaseForm) |  |  |
| producer | [string](#string) |  |  |
| temperature_range | [TemperatureRangeType](#beerproto.TemperatureRangeType) |  | The recommended temperature range of fermentation by the culture producer. |
| notes | [string](#string) |  |  |
| best_for | [string](#string) |  | Recommended styles for a particular culture. |
| inventory | [CultureInventoryType](#beerproto.CultureInventoryType) |  |  |
| product_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| alcohol_tolerance | [PercentType](#beerproto.PercentType) |  | The recommended limit of abv by the culture producer before attenuation stops. |
| glucoamylase | [bool](#bool) |  | A glucoamylase positive culture is capable of producing glucoamylase, the enzyme produced through expression of the diastatic gene, which allows yeast to attenuate dextrins and starches leading to a very low FG. This is positive in some saison/brett yeasts as well as the new gulo hybrid by Omega yeast labs. |
| type | [CultureBaseType](#beerproto.CultureBaseType) |  |  |
| flocculation | [CultureInformation.QualitativeRangeType](#beerproto.CultureInformation.QualitativeRangeType) |  | Floculation refers to the ability of yeast to aggregate to form large flocs which drop out of suspension. |
| attenuation_range | [PercentRangeType](#beerproto.PercentRangeType) |  |  |
| max_reuse | [int32](#int32) |  | Maximum number of times to reuse a culture before a new lab source is recommended. |
| pof | [bool](#bool) |  | A POF&#43; culture is capable of producing phenols, which is a common distinctive property of saison, and brett yeasts. |
| zymocide | [Zymocide](#beerproto.Zymocide) |  |  |






<a name="beerproto.CultureInventoryType"></a>

### CultureInventoryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| liquid | [VolumeType](#beerproto.VolumeType) |  |  |
| dry | [MassType](#beerproto.MassType) |  |  |
| slant | [VolumeType](#beerproto.VolumeType) |  |  |
| culture | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.DiastaticPowerType"></a>

### DiastaticPowerType
Diastatic power is a measurement of malted grains enzymatic content. A value of 35 Lintner is needed to self convert, while a value of 100 or more is desirable for base malts


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [DiastaticPowerType.DiastaticPowerUnitType](#beerproto.DiastaticPowerType.DiastaticPowerUnitType) |  |  |






<a name="beerproto.EfficiencyType"></a>

### EfficiencyType
The efficiencyType stores each efficiency component


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| conversion | [PercentType](#beerproto.PercentType) |  | The percentage of sugar from the grain yield that is extracted and converted during the mash |
| lauter | [PercentType](#beerproto.PercentType) |  | The percentage of sugar that makes it from the mash tun to the kettle |
| mash | [PercentType](#beerproto.PercentType) |  | The percentage of sugar that makes it from the grain to the kettle |
| brewhouse | [PercentType](#beerproto.PercentType) |  | The percentage of sugar that makes it from the grain to the fermenter |






<a name="beerproto.EquipmentItemType"></a>

### EquipmentItemType
EquipmentType provides necessary information for individual brewing equipment


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| boil_rate_per_hour | [VolumeType](#beerproto.VolumeType) |  | The volume boiled off during 1 hour, measured before and after at room temperature. |
| type | [string](#string) |  |  |
| form | [EquipmentItemType.EquipmentBaseForm](#beerproto.EquipmentItemType.EquipmentBaseForm) |  |  |
| drain_rate_per_minute | [VolumeType](#beerproto.VolumeType) |  | The volume that leaves the kettle, especially important for non-immersion chillers that cool the wort as it leaves the kettle. |
| specific_heat | [SpecificHeatType](#beerproto.SpecificHeatType) |  | The specific heat of the piece of equipment, expressed in Cal/(g*C), especially important for when the mashtun is not preheated. |
| grain_absorption_rate | [SpecificVolumeType](#beerproto.SpecificVolumeType) |  | The apparent volume absorbed by grain, typical values are 0.125 qt/lb (1.04 L/kg) for a mashtun, 0.08 gal/lb (0.66 L/kg) for BIAB. |
| name | [string](#string) |  |  |
| maximum_volume | [VolumeType](#beerproto.VolumeType) |  |  |
| weight | [MassType](#beerproto.MassType) |  | The weight of the piece of equipment, especially important for when the mashtun is not preheated. |
| loss | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.EquipmentType"></a>

### EquipmentType
Provides necessary information for brewing equipment set


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| equipment_items | [EquipmentItemType](#beerproto.EquipmentItemType) | repeated |  |






<a name="beerproto.FermentableAdditionType"></a>

### FermentableAdditionType
FermentableAdditionType collects the attributes of each fermentable ingredient for use in a recipe fermentable bill


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [FermentableBaseType](#beerproto.FermentableBaseType) |  |  |
| Origin | [string](#string) |  |  |
| grain_group | [GrainGroup](#beerproto.GrainGroup) |  |  |
| yield | [YieldType](#beerproto.YieldType) |  |  |
| color | [ColorType](#beerproto.ColorType) |  |  |
| name | [string](#string) |  |  |
| producer | [string](#string) |  |  |
| product_id | [string](#string) |  |  |
| timing | [TimingType](#beerproto.TimingType) |  | The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step. |
| mass | [MassType](#beerproto.MassType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.FermentableInventoryType"></a>

### FermentableInventoryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mass | [MassType](#beerproto.MassType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.FermentableType"></a>

### FermentableType
FermentableType collects the attributes of a fermentable ingredient to store as record information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| max_in_batch | [PercentType](#beerproto.PercentType) |  | The recommended maximum percentage to use in a grain bill. |
| recommend_mash | [bool](#bool) |  | True if the fermentable must be mashed, false if it can be steeped. |
| protein | [PercentType](#beerproto.PercentType) |  | The percentage of protein. Higher values may indicate a possibility of haze, or lautering issues. |
| product_id | [string](#string) |  |  |
| grain_group | [GrainGroup](#beerproto.GrainGroup) |  |  |
| yield | [YieldType](#beerproto.YieldType) |  |  |
| type | [FermentableBaseType](#beerproto.FermentableBaseType) |  |  |
| producer | [string](#string) |  |  |
| alpha_amylase | [double](#double) |  | Where diastatic power gives the total amount of all enzymes, alpha amylase, also known as dextrinizing units, refers to only the total amount of alpa amylase in the malted grain. A value of 25-50 is desirable for base malt. |
| color | [ColorType](#beerproto.ColorType) |  |  |
| name | [string](#string) |  |  |
| diastatic_power | [DiastaticPowerType](#beerproto.DiastaticPowerType) |  | Diastatic power is a measurement of malted grains enzymatic content. A value of 35 Lintner is needed to self convert, while a value of 100 or more is desirable. |
| moisture | [PercentType](#beerproto.PercentType) |  |  |
| origin | [string](#string) |  |  |
| inventory | [FermentableInventoryType](#beerproto.FermentableInventoryType) |  |  |
| kolbach_index | [double](#double) |  | The Kolbach Index, also known as soluble to total ratio of nitrogen or protein, is used to indcate the degree of malt modification. A value above 35% is desired for simple single infusion mashing, undermodified malt may require multiple step mashes or decoction. |
| notes | [string](#string) |  |  |






<a name="beerproto.FermentationProcedureType"></a>

### FermentationProcedureType
FermentationProcedureType defines the procedure for performing fermentation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| description | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| fermentation_steps | [FermentationStepType](#beerproto.FermentationStepType) | repeated |  |
| name | [string](#string) |  |  |






<a name="beerproto.FermentationStepType"></a>

### FermentationStepType
FermentationStepType - a per step representation of a fermentation action


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |
| end_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| step_time | [TimeType](#beerproto.TimeType) |  |  |
| free_rise | [bool](#bool) |  | Free rise is used to indicate a fermentation step where the exothermic fermentation is allowed to raise the temperature without restriction This is either True or false. |
| start_gravity | [GravityType](#beerproto.GravityType) |  |  |
| end_gravity | [GravityType](#beerproto.GravityType) |  |  |
| start_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| description | [string](#string) |  |  |
| start_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| end_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| vessel | [string](#string) |  |  |






<a name="beerproto.GravityRangeType"></a>

### GravityRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [GravityType](#beerproto.GravityType) |  |  |
| maximum | [GravityType](#beerproto.GravityType) |  |  |






<a name="beerproto.GravityType"></a>

### GravityType
Gravity refers to the both the measurements of percent of sugar content, ie plato and brix, as well as relative density ie specific gravity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [GravityType.GravityUnitType](#beerproto.GravityType.GravityUnitType) |  |  |






<a name="beerproto.HopAdditionType"></a>

### HopAdditionType
HopAdditionType collects the attributes of each hop ingredient for use in a recipe hop bil


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beta_acid | [PercentType](#beerproto.PercentType) |  |  |
| producer | [string](#string) |  |  |
| origin | [string](#string) |  |  |
| year | [string](#string) |  |  |
| form | [HopVarietyBaseForm](#beerproto.HopVarietyBaseForm) |  |  |
| timing | [TimingType](#beerproto.TimingType) |  | The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step |
| name | [string](#string) |  |  |
| product_id | [string](#string) |  |  |
| alpha_acid | [PercentType](#beerproto.PercentType) |  |  |
| mass | [MassType](#beerproto.MassType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.HopInventoryType"></a>

### HopInventoryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mass | [MassType](#beerproto.MassType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.IBUEstimateType"></a>

### IBUEstimateType
Used to differentiate which IBU formula is being used in a recipe. If formula is modified in any way, eg to support whirlpool/flameout additions etc etc, please use `Other` for transparency


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method | [IBUEstimateType.IBUMethodType](#beerproto.IBUEstimateType.IBUMethodType) |  |  |






<a name="beerproto.IngredientsType"></a>

### IngredientsType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| miscellaneous_additions | [MiscellaneousAdditionType](#beerproto.MiscellaneousAdditionType) | repeated | miscellaneous_additions collects all the miscellaneous items for use in a recipe |
| culture_additions | [CultureAdditionType](#beerproto.CultureAdditionType) | repeated | culture_additions collects all the culture items for use in a recipe |
| water_additions | [WaterAdditionType](#beerproto.WaterAdditionType) | repeated | water_additions collects all the water items for use in a recipe |
| fermentable_additions | [FermentableAdditionType](#beerproto.FermentableAdditionType) | repeated | fermentable_additions collects all the fermentable ingredients for use in a recipe |
| hop_additions | [HopAdditionType](#beerproto.HopAdditionType) | repeated | hop_additions collects all the hops for use in a recipe |






<a name="beerproto.MashProcedureType"></a>

### MashProcedureType
This defines the procedure for performing unique mashing processes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| grain_temperature | [TemperatureType](#beerproto.TemperatureType) |  | Initial grain temperature prior to the start of the mash |
| notes | [string](#string) |  |  |
| mash_steps | [MashStepType](#beerproto.MashStepType) | repeated |  |
| name | [string](#string) |  |  |






<a name="beerproto.MashStepType"></a>

### MashStepType
MashStepType - a per step representation occurring during the mash


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step_time | [TimeType](#beerproto.TimeType) |  |  |
| ramp_time | [TimeType](#beerproto.TimeType) |  | The amount of time that passes before this step begins. eg moving from a mash step (step 1) of 148F, to a new temperature step of 156F (step 2) may take 8 minutes to heat the mash. Step 2 would have a ramp time of 8 minutes |
| end_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| description | [string](#string) |  |  |
| infuse_temperature | [TemperatureType](#beerproto.TemperatureType) |  | Temperature of the water for an infusion step |
| start_pH | [AcidityType](#beerproto.AcidityType) |  |  |
| end_pH | [AcidityType](#beerproto.AcidityType) |  |  |
| name | [string](#string) |  |  |
| type | [MashStepType.MashStepTypeType](#beerproto.MashStepType.MashStepTypeType) |  |  |
| amount | [VolumeType](#beerproto.VolumeType) |  |  |
| step_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| water_grain_ratio | [SpecificVolumeType](#beerproto.SpecificVolumeType) |  | Also known as the mash thickness. eg 1.75 qt/lb or 3.65 L/kg |






<a name="beerproto.MassType"></a>

### MassType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [MassType.MassUnitType](#beerproto.MassType.MassUnitType) |  |  |






<a name="beerproto.MiscellaneousAdditionType"></a>

### MiscellaneousAdditionType
MiscellaneousAdditionType collects the attributes of each miscellaneous ingredient for use in a recipe


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| producer | [string](#string) |  |  |
| timing | [TimingType](#beerproto.TimingType) |  | The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step. |
| product_id | [string](#string) |  |  |
| type | [MiscellaneousBaseType](#beerproto.MiscellaneousBaseType) |  |  |
| mass | [MassType](#beerproto.MassType) |  |  |
| unit | [UnitType](#beerproto.UnitType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.MiscellaneousInventoryType"></a>

### MiscellaneousInventoryType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mass | [MassType](#beerproto.MassType) |  |  |
| unit | [UnitType](#beerproto.UnitType) |  |  |
| volume | [VolumeType](#beerproto.VolumeType) |  |  |






<a name="beerproto.MiscellaneousType"></a>

### MiscellaneousType
MiscellaneousType collects the attributes of an ingredient to store as record information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| use_for | [string](#string) |  | Used to describe the purpose of the miscellaneous ingredient, e.g. whirlfloc is used for clarity. |
| notes | [string](#string) |  |  |
| name | [string](#string) |  |  |
| producer | [string](#string) |  |  |
| product_id | [string](#string) |  |  |
| type | [MiscellaneousBaseType](#beerproto.MiscellaneousBaseType) |  |  |
| inventory | [MiscellaneousInventoryType](#beerproto.MiscellaneousInventoryType) |  |  |






<a name="beerproto.OilContentType"></a>

### OilContentType
oil_content collects all information of a hop variety pertaining to oil content, polyphenols, and thiols. Each individual compound is expressed as a percent of the total oil measurement


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| polyphenols | [PercentType](#beerproto.PercentType) |  |  |
| total_oil_ml_per_100g | [double](#double) |  | The total amount of oil, including hydrocarbons, esters, and terpene alcohols in units of ml of oil per 100g of hop mass. |
| farnesene | [PercentType](#beerproto.PercentType) |  |  |
| limonene | [PercentType](#beerproto.PercentType) |  |  |
| nerol | [PercentType](#beerproto.PercentType) |  |  |
| geraniol | [PercentType](#beerproto.PercentType) |  |  |
| b_pinene | [PercentType](#beerproto.PercentType) |  |  |
| linalool | [PercentType](#beerproto.PercentType) |  |  |
| caryophyllene | [PercentType](#beerproto.PercentType) |  |  |
| cohumulone | [PercentType](#beerproto.PercentType) |  |  |
| xanthohumol | [PercentType](#beerproto.PercentType) |  |  |
| humulene | [PercentType](#beerproto.PercentType) |  |  |
| myrcene | [PercentType](#beerproto.PercentType) |  |  |
| pinene | [PercentType](#beerproto.PercentType) |  |  |






<a name="beerproto.PackagingProcedureType"></a>

### PackagingProcedureType
Describes the procedure for packaging your beverage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| packaged_volume | [VolumeType](#beerproto.VolumeType) |  |  |
| description | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| packaging_vessels | [PackagingVesselType](#beerproto.PackagingVesselType) | repeated |  |






<a name="beerproto.PackagingVesselType"></a>

### PackagingVesselType
PackagingVesselType - a per vessel representation of a packaging process


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [PackagingVesselType.PackagingVesselTypeType](#beerproto.PackagingVesselType.PackagingVesselTypeType) |  |  |
| start_gravity | [GravityType](#beerproto.GravityType) |  |  |
| name | [string](#string) |  |  |
| package_date | [string](#string) |  |  |
| step_time | [TimeType](#beerproto.TimeType) |  |  |
| end_gravity | [GravityType](#beerproto.GravityType) |  |  |
| vessel_volume | [VolumeType](#beerproto.VolumeType) |  |  |
| vessel_quantity | [double](#double) |  |  |
| description | [string](#string) |  |  |
| start_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| carbonation | [double](#double) |  |  |
| start_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |
| end_ph | [AcidityType](#beerproto.AcidityType) |  |  |
| end_temperature | [TemperatureType](#beerproto.TemperatureType) |  |  |






<a name="beerproto.PercentRangeType"></a>

### PercentRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [PercentType](#beerproto.PercentType) |  |  |
| maximum | [PercentType](#beerproto.PercentType) |  |  |






<a name="beerproto.PercentType"></a>

### PercentType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [PercentType.PercentUnitType](#beerproto.PercentType.PercentUnitType) |  |  |






<a name="beerproto.Recipe"></a>

### Recipe



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| mashes | [MashProcedureType](#beerproto.MashProcedureType) | repeated | A collection of steps providing process information for common mashing procedures |
| recipes | [RecipeType](#beerproto.RecipeType) | repeated | Records containing a minimal collection of the description of ingredients, procedures and other required parameters necessary to recreate a batch of beer |
| miscellaneous_ingredients | [MiscellaneousType](#beerproto.MiscellaneousType) | repeated | Records for adjuncts which do not contribute to the gravity of the beer |
| styles | [StyleType](#beerproto.StyleType) | repeated | Records detailing the characteristics of the beer styles for which judging guidelines have been established |
| fermentations | [FermentationProcedureType](#beerproto.FermentationProcedureType) | repeated | A collection of steps providing process information for common fermentation procedures |
| boil | [BoilProcedureType](#beerproto.BoilProcedureType) | repeated | A collection of steps providing process information for common boil procedures |
| version | [double](#double) |  | Explicitly encode version within list of records |
| fermentables | [FermentableType](#beerproto.FermentableType) | repeated | Records for any ingredient that contributes to the gravity of the beer |
| timing_object | [TimingType](#beerproto.TimingType) |  | The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step |
| cultures | [CultureInformation](#beerproto.CultureInformation) | repeated | Records detailing the wide array of unique cultures |
| equipments | [EquipmentType](#beerproto.EquipmentType) | repeated | Provides necessary information for brewing equipment |
| packaging | [PackagingProcedureType](#beerproto.PackagingProcedureType) | repeated | A collection of steps providing process information for common packaging procedures |
| hop_varieties | [VarietyInformation](#beerproto.VarietyInformation) | repeated | Records detailing the many properties of unique hop varieties |
| profiles | [WaterBase](#beerproto.WaterBase) | repeated | Records for water profiles used in brewing |






<a name="beerproto.RecipeStyleType"></a>

### RecipeStyleType
RecipeStyleType defines style information stored in a recipe record


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [RecipeStyleType.StyleCategories](#beerproto.RecipeStyleType.StyleCategories) |  |  |
| name | [string](#string) |  |  |
| category | [string](#string) |  |  |
| category_number | [int32](#int32) |  |  |
| style_letter | [string](#string) |  |  |
| style_guide | [string](#string) |  |  |






<a name="beerproto.RecipeType"></a>

### RecipeType
RecipeType composes the information stored in a recipe


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| efficiency | [EfficiencyType](#beerproto.EfficiencyType) |  | Used to store each efficiency component, including conversion, and brewhouse |
| style | [RecipeStyleType](#beerproto.RecipeStyleType) |  |  |
| ibu_estimate | [IBUEstimateType](#beerproto.IBUEstimateType) |  | Used to differentiate which IBU formula is being used in a recipe. If formula is modified in any way, eg to support whirlpool/flameout additions etc etc, please use `Other` for transparency |
| color_estimate | [ColorType](#beerproto.ColorType) |  | The color of the finished beer, using SRM or EBC |
| beer_pH | [AcidityType](#beerproto.AcidityType) |  | The final beer pH at the end of fermentation |
| name | [string](#string) |  |  |
| type | [RecipeType.RecipeTypeType](#beerproto.RecipeType.RecipeTypeType) |  |  |
| coauthor | [string](#string) |  |  |
| original_gravity | [GravityType](#beerproto.GravityType) |  | The gravity of wort when transffered to the fermenter |
| final_gravity | [GravityType](#beerproto.GravityType) |  | The gravity of beer at the end of fermentation |
| carbonation | [double](#double) |  | The final carbonation of the beer when packaged or served |
| fermentation | [FermentationProcedureType](#beerproto.FermentationProcedureType) |  | FermentationProcedureType defines the procedure for performing fermentation |
| author | [string](#string) |  |  |
| ingredients | [IngredientsType](#beerproto.IngredientsType) |  | A collection of all ingredients used for the recipe |
| mash | [MashProcedureType](#beerproto.MashProcedureType) |  | This defines the procedure for performing unique mashing processes |
| packaging | [PackagingProcedureType](#beerproto.PackagingProcedureType) |  | Describes the procedure for packaging your beverage |
| boil | [BoilProcedureType](#beerproto.BoilProcedureType) |  | Defines the procedure for performing a boil. A boil procedure with no steps is the same as a standard single step boil |
| taste | [TasteType](#beerproto.TasteType) |  | Used to store subjective tasting notes, and rating |
| calories_per_pint | [double](#double) |  |  |
| created | [string](#string) |  |  |
| batch_size | [VolumeType](#beerproto.VolumeType) |  | The volume into the fermenter |
| notes | [string](#string) |  |  |
| alcohol_by_volume | [PercentType](#beerproto.PercentType) |  |  |
| apparent_attenuation | [PercentType](#beerproto.PercentType) |  | The total apparent attenuation of the finished beer after fermentation |






<a name="beerproto.SpecificHeatType"></a>

### SpecificHeatType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [SpecificHeatType.SpecificHeatUnitType](#beerproto.SpecificHeatType.SpecificHeatUnitType) |  |  |






<a name="beerproto.SpecificVolumeType"></a>

### SpecificVolumeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [SpecificVolumeType.SpecificVolumeUnitType](#beerproto.SpecificVolumeType.SpecificVolumeUnitType) |  |  |






<a name="beerproto.StyleType"></a>

### StyleType
StyleType provide information for Style categorization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aroma | [string](#string) |  |  |
| ingredients | [string](#string) |  |  |
| category_number | [int32](#int32) |  |  |
| notes | [string](#string) |  |  |
| flavor | [string](#string) |  |  |
| mouthfeel | [string](#string) |  |  |
| final_gravity | [GravityRangeType](#beerproto.GravityRangeType) |  |  |
| style_guide | [string](#string) |  |  |
| color | [ColorRangeType](#beerproto.ColorRangeType) |  |  |
| original_gravity | [GravityRangeType](#beerproto.GravityRangeType) |  |  |
| examples | [string](#string) |  |  |
| name | [string](#string) |  |  |
| carbonation | [CarbonationRangeType](#beerproto.CarbonationRangeType) |  |  |
| alcohol_by_volume | [PercentRangeType](#beerproto.PercentRangeType) |  |  |
| international_bitterness_units | [BitternessRangeType](#beerproto.BitternessRangeType) |  |  |
| appearance | [string](#string) |  |  |
| category | [string](#string) |  |  |
| style_letter | [string](#string) |  |  |
| type | [StyleType.StyleCategories](#beerproto.StyleType.StyleCategories) |  |  |
| overall_impression | [string](#string) |  |  |






<a name="beerproto.TasteType"></a>

### TasteType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notes | [string](#string) |  |  |
| rating | [double](#double) |  |  |






<a name="beerproto.TemperatureRangeType"></a>

### TemperatureRangeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minimum | [TemperatureType](#beerproto.TemperatureType) |  |  |
| maximum | [TemperatureType](#beerproto.TemperatureType) |  |  |






<a name="beerproto.TemperatureType"></a>

### TemperatureType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [TemperatureType.TemperatureUnitType](#beerproto.TemperatureType.TemperatureUnitType) |  |  |






<a name="beerproto.TimeType"></a>

### TimeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [TimeType.TimeUnitType](#beerproto.TimeType.TimeUnitType) |  |  |






<a name="beerproto.TimingType"></a>

### TimingType
The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [TimeType](#beerproto.TimeType) |  | What time during a process step is added, eg a value of 2 days for a dry hop addition would be added 2 days into the fermentation step. |
| duration | [TimeType](#beerproto.TimeType) |  | How long an ingredient addition remains, this was referred to as time in the BeerXML standard. E.G. A 40 minute hop boil additions means to boil for 40 minutes, and a 2 day duration for a dry hop means to remove it after 2 days. |
| continuous | [bool](#bool) |  | A continuous addition is spread out evenly and added during the entire process step, eg 60 minute IPA by dogfish head takes all ofthe hop additions and adds them throughout the entire boil. |
| specific_gravity | [GravityType](#beerproto.GravityType) |  | Used to indicate when an addition is added based on a desired specific gravity. E.G. Add dry hop at when SG is 1.018. |
| ph | [AcidityType](#beerproto.AcidityType) |  | Used to indicate when an addition is added based on a desired specific gravity. eg Add brett when pH is 3.4. |
| step | [int32](#int32) |  | Used to indicate what step this ingredient timing addition is referencing. EG A value of 2 for add_to_fermentation would mean to add during the second fermentation step. |
| use | [TimingType.UseType](#beerproto.TimingType.UseType) |  |  |






<a name="beerproto.UnitType"></a>

### UnitType
UnitType is used where unitless amounts are required, such as 1 apple, or 1 yeast packet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [UnitType.UnitUnitType](#beerproto.UnitType.UnitUnitType) |  |  |






<a name="beerproto.VarietyInformation"></a>

### VarietyInformation
VarietyInformation collects the attributes of a hop variety to store as record information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| inventory | [HopInventoryType](#beerproto.HopInventoryType) |  |  |
| type | [VarietyInformation.VarietyInformationType](#beerproto.VarietyInformation.VarietyInformationType) |  |  |
| oil_content | [OilContentType](#beerproto.OilContentType) |  | Oil Content information object. |
| percent_lost | [PercentType](#beerproto.PercentType) |  | Defined as the percentage of hop alpha lost in 6 months of storage. |
| product_id | [string](#string) |  |  |
| alpha_acid | [PercentType](#beerproto.PercentType) |  |  |
| beta_acid | [PercentType](#beerproto.PercentType) |  |  |
| name | [string](#string) |  |  |
| origin | [string](#string) |  |  |
| substitutes | [string](#string) |  |  |
| year | [string](#string) |  |  |
| form | [HopVarietyBaseForm](#beerproto.HopVarietyBaseForm) |  |  |
| producer | [string](#string) |  |  |
| notes | [string](#string) |  |  |






<a name="beerproto.VolumeType"></a>

### VolumeType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |
| unit | [VolumeType.VolumeUnitType](#beerproto.VolumeType.VolumeUnitType) |  |  |






<a name="beerproto.WaterAdditionType"></a>

### WaterAdditionType
WaterAdditionType collects the attributes of each water addition for use in a recipe


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| flouride | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| sulfate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| producer | [string](#string) |  |  |
| nitrate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| nitrite | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| chloride | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| amount | [VolumeType](#beerproto.VolumeType) |  |  |
| name | [string](#string) |  |  |
| potassium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| magnesium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| iron | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| bicarbonate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| calcium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| carbonate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| sodium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |






<a name="beerproto.WaterBase"></a>

### WaterBase
WaterBase provides unique properties to identify individual records of  brewing water


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| calcium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| nitrite | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| chloride | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| name | [string](#string) |  |  |
| potassium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| carbonate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| iron | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| flouride | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| sulfate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| magnesium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| producer | [string](#string) |  |  |
| bicarbonate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| nitrate | [ConcentrationType](#beerproto.ConcentrationType) |  |  |
| sodium | [ConcentrationType](#beerproto.ConcentrationType) |  |  |






<a name="beerproto.YieldType"></a>

### YieldType
The potential yield of the fermentable ingredient, supporting SG, or percentage. eg 1.037 or 80% are valid yield types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fine_grind | [PercentType](#beerproto.PercentType) |  | Percentage yield compared to succrose of a fine grind. eg 80% |
| coarse_grind | [PercentType](#beerproto.PercentType) |  | Percentage yield compared to succrose of a coarse grind. eg 78% |
| fine_coarse_difference | [PercentType](#beerproto.PercentType) |  | The difference between fine and coarse grind, a difference more than 2 percent can indicate a protein or step mash may be desirable. eg 2% |
| potential | [GravityType](#beerproto.GravityType) |  | The potential yield of the fermentable ingredient for 1 lb of grain mashed in 1 gallon of water. eg 1.037 |






<a name="beerproto.Zymocide"></a>

### Zymocide
Zymocide, also known as killer yeast properties, is common among wine yeast. There are also some ale and brett yeasts that are immune to some zymocidic properties, these are known as killer neutral


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| no1 | [bool](#bool) |  |  |
| no2 | [bool](#bool) |  |  |
| no28 | [bool](#bool) |  |  |
| klus | [bool](#bool) |  |  |
| neutral | [bool](#bool) |  |  |





 


<a name="beerproto.AcidityType.AcidityUnitType"></a>

### AcidityType.AcidityUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| PH | 1 |  |



<a name="beerproto.BitternessType.BitternessUnitType"></a>

### BitternessType.BitternessUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| IBUs | 1 | IBUs |



<a name="beerproto.BoilStepType.BoilStepTypeChillingType"></a>

### BoilStepType.BoilStepTypeChillingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| BATCH | 1 | batch |
| INLINE | 2 | inline |



<a name="beerproto.CarbonationType.CarbonationUnitType"></a>

### CarbonationType.CarbonationUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| VOLS | 1 | vols |



<a name="beerproto.ColorType.ColorUnitType"></a>

### ColorType.ColorUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| EBC | 1 | EBC |
| LOVI | 2 | Lovi |
| SRM | 3 | SRM |



<a name="beerproto.ConcentrationType.ConcentrationUnitType"></a>

### ConcentrationType.ConcentrationUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| PPM | 1 | ppm |
| PPB | 2 | ppb |
| MGL | 3 | mg/l |



<a name="beerproto.CultureBaseForm"></a>

### CultureBaseForm


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_CultureBaseForm | 0 |  |
| LIQUID | 1 | liquid |
| DRY | 2 | dry |
| SLANT | 3 | slant |
| CULTURE | 4 | culture |
| DREGS | 5 | dregs |



<a name="beerproto.CultureBaseType"></a>

### CultureBaseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_CultureBaseType | 0 |  |
| ALE | 1 | ale |
| BACTERIA | 2 | bacteria |
| BRETT | 3 | brett |
| CHAMPAGNE | 4 | champagne |
| KVEIK | 5 | kveik |
| LACTO | 6 | lacto |
| LAGER | 7 | lager |
| MALOLACTIC | 8 | malolactic |
| MIXED_CULTURE | 9 | mixed-culture |
| OTHER_CultureBaseType | 10 | other |
| PEDIO | 11 | pedio |
| SPONTANEOUS | 12 | spontaneous |
| WINE | 13 | wine |



<a name="beerproto.CultureInformation.QualitativeRangeType"></a>

### CultureInformation.QualitativeRangeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_QualitativeRangeType | 0 |  |
| VERY_LOW | 1 | very low |
| LOW | 2 | low |
| MEDIUM_LOW | 3 | medium low |
| MEDIUM | 4 | medium |
| MEDIUM_HIGH | 5 | medium high |
| HIGH | 6 | high |
| VERY_HIGH | 7 | very high |



<a name="beerproto.DiastaticPowerType.DiastaticPowerUnitType"></a>

### DiastaticPowerType.DiastaticPowerUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| lintner | 1 | Lintner |
| WK | 2 | WK |



<a name="beerproto.EquipmentItemType.EquipmentBaseForm"></a>

### EquipmentItemType.EquipmentBaseForm


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| HLT | 1 | HLT |
| MASH_TUN | 2 | Mash Tun |
| LAUTER_TUN | 3 | Lauter Tun |
| BREW_KETTLE | 4 | Brew Kettle |
| FERMENTER | 5 | Fermenter |
| AGING_VESSEL | 6 | Aging Vessel |
| PACKAGING_VESSEL | 7 | Packaging Vessel |



<a name="beerproto.FermentableBaseType"></a>

### FermentableBaseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_FermentableBaseType | 0 |  |
| DRY_EXTRACT | 1 | dry extract |
| EXTRACT | 2 | extract |
| GRAIN | 3 | grain |
| SUGAR | 4 | sugar |
| FRUIT | 5 | fruit |
| JUICE | 6 | juice |
| HONEY | 7 | honey |
| OTHER_FermentableBaseType | 8 | other |



<a name="beerproto.GrainGroup"></a>

### GrainGroup


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_GrainGroup | 0 |  |
| BASE | 1 | base |
| CARAMEL | 2 | caramel |
| FLAKED | 3 | flaked |
| ROASTED | 4 | roasted |
| SPECIALTY | 5 | specialty |
| SMOKED | 6 | smoked |
| ADJUNCT | 7 | adjunct |



<a name="beerproto.GravityType.GravityUnitType"></a>

### GravityType.GravityUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| SG | 1 | sg |
| PLATO | 2 | plato |
| BRIX | 3 | brix |



<a name="beerproto.HopVarietyBaseForm"></a>

### HopVarietyBaseForm


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_HopVarietyBaseForm | 0 |  |
| EXTRACT_HopVarietyBaseForm | 1 | extract |
| LEAF | 2 | leaf |
| LEAFWET | 3 | leaf (wet) |
| PELLET | 4 | pellet |
| POWDER | 5 | powder |
| PLUG | 6 | plug |



<a name="beerproto.IBUEstimateType.IBUMethodType"></a>

### IBUEstimateType.IBUMethodType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| Rager | 1 | Rager |
| Tinseth | 2 | Tinseth |
| Garetz | 3 | Garetz |
| Other | 4 | Other |



<a name="beerproto.MashStepType.MashStepTypeType"></a>

### MashStepType.MashStepTypeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| INFUSION | 1 |  |
| TEMPERATURE | 2 |  |
| DECOCTION | 3 |  |
| SOURING_MASH | 4 |  |
| SOURING_WORT | 5 |  |
| DRAIN_MASH_TUN | 6 |  |
| SPARGE | 7 |  |



<a name="beerproto.MassType.MassUnitType"></a>

### MassType.MassUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| MG | 1 | mg |
| G | 2 | g |
| KG | 3 | kg |
| LB | 4 | lb |
| OZ | 5 | oz |



<a name="beerproto.MiscellaneousBaseType"></a>

### MiscellaneousBaseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| SPICE | 1 | spice |
| FINING | 2 | fining |
| WATER_AGENT | 3 | water agent |
| HERB | 4 | herb |
| FLAVOR | 5 | flavor |
| WOOD | 6 | wood |
| OTHER | 7 | other |



<a name="beerproto.PackagingVesselType.PackagingVesselTypeType"></a>

### PackagingVesselType.PackagingVesselTypeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| KEG | 1 | keg |
| BOTTLE | 2 | bottle |
| CASK | 3 | cask |
| TANK | 4 | tank |
| FIRKIN | 5 | firkin |
| OTHER | 6 | other |



<a name="beerproto.PercentType.PercentUnitType"></a>

### PercentType.PercentUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| PERCENT_SIGN | 1 | % |



<a name="beerproto.RecipeStyleType.StyleCategories"></a>

### RecipeStyleType.StyleCategories


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| BEER | 1 | beer |
| cider | 2 | cider |
| OMBUCHA | 3 | kombucha |
| MEAD | 4 | mead |
| SODA | 5 | soda |
| WINE | 6 | wine |
| OTHER | 7 | other |



<a name="beerproto.RecipeType.RecipeTypeType"></a>

### RecipeType.RecipeTypeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| CIDER | 1 | cider |
| KOMBUCHA | 2 | kombucha |
| SODA | 3 | soda |
| OTHER | 4 | other |
| MEAD | 5 | mead |
| WINE | 6 | wine |
| EXTRACT | 7 | extract |
| PARTIAL_MASH | 8 | partial mash |
| ALL_GRAIN | 9 | all grain |



<a name="beerproto.SpecificHeatType.SpecificHeatUnitType"></a>

### SpecificHeatType.SpecificHeatUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| CALGC | 1 | Cal/(g C) |
| JKGK | 2 | J/(kg K) |
| BTULBF | 3 | BTU/(lb F) |



<a name="beerproto.SpecificVolumeType.SpecificVolumeUnitType"></a>

### SpecificVolumeType.SpecificVolumeUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| QTLB | 1 | qt/lb |
| GALLB | 2 | gal/lb |
| GALOZ | 3 | gal/oz |
| LG | 4 | l/g |
| LKG | 5 | l/kg |
| FLOZOZ | 6 | floz/oz |
| M3KG | 7 | m^3/kg |
| FT3LB | 8 | ft^3/lb |



<a name="beerproto.StyleType.StyleCategories"></a>

### StyleType.StyleCategories


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| BEER | 1 | beer |
| CIDER | 2 | cider |
| KOMBUCHA | 3 | kombucha |
| MEAD | 4 | mead |
| OTHER | 5 | other |
| SODA | 6 | soda |
| WINE | 7 | wine |



<a name="beerproto.TemperatureType.TemperatureUnitType"></a>

### TemperatureType.TemperatureUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| C | 1 |  |
| F | 2 |  |



<a name="beerproto.TimeType.TimeUnitType"></a>

### TimeType.TimeUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| SEC | 1 | sec |
| MIN | 2 | min |
| HR | 3 | hr |
| DAY | 4 | day |
| WEEK | 5 | week |
| MONTH | 6 | month |
| YEAR | 7 | year |



<a name="beerproto.TimingType.UseType"></a>

### TimingType.UseType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| ADD_TO_MASH | 1 | add to mash |
| ADD_TO_BOIL | 2 | add to boil |
| ADD_TO_FERMENTATION | 3 | add to fermentation |
| ADD_TO_PACKAGE | 4 | add to package |



<a name="beerproto.UnitType.UnitUnitType"></a>

### UnitType.UnitUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| ONE | 1 | 1 |
| UNIT | 2 | unit |
| EACH | 3 | each |
| DIMENSIONLESS | 4 | dimensionless |
| PKG | 5 | pkg |



<a name="beerproto.VarietyInformation.VarietyInformationType"></a>

### VarietyInformation.VarietyInformationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_VarietyInformationType | 0 |  |
| AROMA | 1 | aroma |
| BITTERING | 2 | bittering |
| FLAVOR | 3 | flavor |
| AROMA_BITTERING | 4 | aroma/bittering |
| BITTERING_FLAVOR | 5 | bittering/flavor |
| AROMA_FLAVOR | 6 | aroma/flavor |
| AROMA_BITTERING_FLAVOR | 7 | aroma/bittering/flavor |



<a name="beerproto.VolumeType.VolumeUnitType"></a>

### VolumeType.VolumeUnitType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL | 0 |  |
| ML | 1 | ml |
| L | 2 | l |
| TSP | 3 | tsp |
| TBSP | 4 | tbsp |
| FLOZ | 5 | floz |
| CUP | 6 | cup |
| PT | 7 | pt |
| QT | 8 | qt |
| GAL | 9 | gal |
| BBL | 10 | bbl |
| IFOZ | 11 | ifloz |
| IPT | 12 | ipt |
| IQT | 13 | iqt |
| IGAL | 14 | igal |
| IBBL | 15 | ibbl |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

