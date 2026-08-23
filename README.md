# BeerProto

This repository contains the beer.proto file.

[Module documentation](https://buf.build/rossmerr/beerprotoapis/docs/main)

To generate a copy of BeerProto in your language use the following command changing the template to match your language.

```bash
buf generate beerprotoapis --template=buf.gen.go.yaml
```

## Canonical units

Every measurement in BeerProto is a `{value, unit}` pair. The two fields answer
different questions, and this is the single most important thing to understand
before reading or writing BeerProto:

> **`value` is always stored in the type's canonical unit.**
> **`unit` carries the unit the measurement is intended to be presented in.**

`unit` is a presentation instruction, **not** a description of `value`.

### Worked example

A brewer enters a hop charge of **12 oz**. One ounce avoirdupois is exactly
28.349523125 g, so:

```
MassType {
  value: 340.1942775        // 12 oz expressed in grams, the canonical unit
  unit:  MASS_UNIT_OZ       // present this to the user in ounces
}
```

Reading that message back:

- Its magnitude is **340.1942775 grams**. Anything doing arithmetic — summing a
  grist, checking stock, computing IBU — uses `value` directly, as grams.
- Its magnitude is **not** 340.1942775 ounces, and `unit` must never be used to
  convert `value`.
- A UI showing this to the brewer converts 340.1942775 g into ounces at the
  moment of display and renders **12 oz**.

Change `unit` to `MASS_UNIT_KG` and nothing about the quantity changes — the same
340.1942775 g now renders as 0.34 kg.

### Rules

**Producers MUST** convert into the canonical unit before setting `value`, and
set `unit` to whatever the measurement should be displayed in. If you have no
presentation preference, set `unit` to the canonical unit.

**Consumers MUST NOT** convert `value` using `unit`. Doing so double-converts and
silently corrupts the quantity. Treat `value` as canonical unconditionally, and
read `unit` only when rendering.

**At an interop boundary — BeerJSON, BeerXML, or any other format whose
`{value, unit}` pairs are self-describing — convert.** Those formats mean "12" and
"oz" together; BeerProto does not. Write `value` converted into `unit` on export,
and normalise into the canonical unit on import.

**`UNSPECIFIED` means "no presentation preference stated", not "no unit".**
`value` is canonical regardless. Use presence of the enclosing message, not
`unit == UNSPECIFIED` and not `value == 0`, to test whether a measurement was
supplied at all.

### The canonical unit for each type

`value` is always in the unit named here.

| Message | `value` is in | Canonical enum constant |
|---|---|---|
| `VolumeType` | millilitres | `VOLUME_UNIT_ML` |
| `MassType` | grams | `MASS_UNIT_G` |
| `TemperatureType` | degrees Celsius | `TEMPERATURE_UNIT_C` |
| `TimeType` | seconds | `TIME_UNIT_SEC` |
| `ConcentrationType` | milligrams per litre | `CONCENTRATION_UNIT_MGL` |
| `PartsPerType` | parts per million | `PARTS_PER_UNIT_MILLION` |
| `ColorType` | EBC | `COLOR_UNIT_EBC` |
| `GravityType` | specific gravity | `GRAVITY_UNIT_SG` |
| `CarbonationType` | volumes CO₂ | `CARBONATION_UNIT_VOLS` |
| `DiastaticPowerType` | degrees Windisch–Kolbach | `DIASTATIC_POWER_UNIT_WK` |
| `PressureType` | pascals | `PRESSURE_UNIT_PASCAL` |
| `SpecificHeatType` | joules per kilogram-kelvin | `SPECIFIC_HEAT_UNIT_JKGK` |
| `SpecificVolumeType` | litres per kilogram | `SPECIFIC_VOLUME_UNIT_LKG` |
| `ViscosityType` | millipascal-seconds | `VISCOSITY_UNIT_MPAS` |
| `CellCountType` | billions of cells | `CELL_COUNT_UNIT_BILLION` |

These six have exactly one real unit in their enum, so canonical and presentation
are necessarily the same:

| Message | `value` is in | Canonical enum constant |
|---|---|---|
| `BitternessType` | IBU | `BITTERNESS_UNIT_IBUS` |
| `PercentType` | percent | `PERCENT_UNIT_PERCENT_SIGN` |
| `AcidityType` | pH | `ACIDITY_UNIT_PH` |
| `MolarType` | grams per mole | `MOLAR_UNIT_GMOL` |
| `PitchRateType` | million cells / mL / °P | `PITCH_RATE_UNIT_MILLION_CELLS_PER_ML_PER_PLATO` |
| `BufferingCapacityType` | mEq per kilogram per pH | `BUFFERING_CAPACITY_UNIT_MEQ_KG_PH` |

The twelve `*RangeType` messages carry no unit of their own — each holds its
element type as `minimum` and `maximum`, so both members follow the element's
rule. Convert both, not just one.

### Excluded types — where `unit` *does* describe `value`

Three types have no canonical unit, so the rule above cannot apply to them. **For
these three, and only these three, `unit` describes `value` in the ordinary
self-describing way.** Read them as a conventional `{value, unit}` pair.

| Message | Why there is no canonical unit |
|---|---|
| `UnitType` | Its units are discrete counts — `ONE`, `UNIT`, `EACH`, `PKG`, `DIMENSIONLESS`. No conversion between them exists or can exist: a package is not a number of "each". Compare values only when the units already match. |
| `EnzymeActivityType` | `DU`, `WK` and `SKB` are three different laboratory assays, not three scalings of one quantity. There is no defensible conversion factor between them. |
| `RateType` | `RateUnit` spans four dimensions in one enum — volume flow (`L_PER_HOUR`, `BBL_PER_HOUR`), percent rate (`PERCENT_PER_HOUR`), specific volume (`L_PER_KG`) and temperature rate (`C_PER_MINUTE`). A single canonical unit is not meaningful. |

### Why these units

The canonical units are **metric, at the smallest unit a brewing quantity is
actually measured in** — not strict SI. Strict SI would put volume in m³ and
temperature in kelvin, which no brewing formula uses and no brewer reads.

Preferring the smaller metric unit is not about floating-point range: IEEE-754
doubles are scale-invariant, so `20.0` and `20000.0` carry the same precision. It
is about keeping typical values **integral**, because integer-valued doubles add
and subtract exactly below 2⁵³. In litres `0.1 + 0.2 != 0.3`; in millilitres
`100 + 200 == 300`. Anything accumulating many small additions — water salt
doses, step infusions, hop charges — benefits directly.

Four choices need their reasons recorded, because they are not obvious:

- **`TimeType` → seconds.** `TimeType.value` is an `int64`, the only measurement
  in BeerProto that is not a `double`. Minutes would make any sub-minute duration
  — a whirlpool rest, a short hop stand — unrepresentable, and would silently
  round a 30 second step up to one minute.
- **`ConcentrationType` → mg/L, `PartsPerType` → ppm.** For dilute aqueous
  solutions 1 mg/L is numerically 1 ppm, because a litre of brewing liquor weighs
  a kilogram, so the two interchange without arithmetic. The types stay distinct
  all the same: a concentration is a mass per volume, a parts-per is a pure ratio,
  and the equivalence stops holding for anything not around 1 kg/L.
- **`ColorType` → EBC, never Lovibond.** EBC and SRM are linearly related and
  convert losslessly. Lovibond is not a third scale to convert through: the
  Daniels relation `SRM = 1.3546 × °L − 0.76` is a curve fit of MCU→SRM data, not
  a unit conversion, and must never be applied to an individual malt colour.
- **`GravityType` → SG.** SG ↔ °Plato runs through a cubic and is not
  algebraically lossless. The error sits well below hydrometer and refractometer
  resolution, so SG is safe as the stored form — but extract-basis calculations
  (attenuation, real extract) should convert to °Plato explicitly and do their
  arithmetic there.

### Implementation notes

- **Convert once, at the edges.** Into canonical on ingest, out of canonical at
  the point of display. Every layer in between works in canonical units and needs
  no conversion at all — which is the entire point of the convention.
- **Round-trips accumulate error.** Metric → imperial → metric through
  floating-point conversions is lossy; use exact rational arithmetic where your
  implementation offers it, and prefer fewer conversions over better ones.
- **`DiastaticPowerType` is affine, not a scaling**: `°WK = (3.5 × °L) − 16`, so
  `°L = (°WK + 16) / 3.5`. Zero does not survive that conversion — 0 °L is
  −16 °WK — with two consequences. Canonical °WK values are **negative** for any
  malt below 4.57 °L, so `value > 0` must never be used as a presence or validity
  test. And a *missing* diastatic power is not 0 °WK; test presence on the
  enclosing message.
- **Beware code that predates this convention.** Anything that converts using
  `unit` — a server applying an inventory delta, an exporter writing BeerJSON —
  will double-convert under these semantics and must be updated before canonical
  data reaches it.
