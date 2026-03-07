---
title: "Identity Verification - Input validation rules | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/hybrid-input-validation/"
scraped_at: "2026-03-07T22:04:55+00:00"
---

# Input validation rules

####

Plaid does as much as possible to ensure the user data you collect and verify is clean and standardized. If you are submitting data about the user via the API, you will need to match the formatting and standardization requirements described below.

To help with development, we also provide a [JSON file containing the same validation rules](/schema/identity_verification_api.json).

#### ID numbers

Plaid's Data Source Verification will automatically collect, validate, and verify any of the following ID numbers for countries that you've enabled in the Dashboard Editor.

If you are separately collecting ID numbers outside of Link, or if you have existing customer data you're looking to import into Plaid, you can submit this data programmatically via [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) or [`/link/token/create`](/docs/api/link/#linktokencreate).

The table below lists each ID number that Identity Verification supports. The `API type` is a unique identifier that you should provide via API via the `id_number.type` field.

ID Numbers submitted via API should be stripped of any formatting characters like spaces, periods, and dashes. The table below also includes the exact regular expressions we use to decide whether a given ID Number is well-formed and valid for that region's ID type.

Plaid may perform additional validation to ensure the ID number is valid for the selected type. For example, the SSA will not issue SSNs with specific numbers in certain positions. SSNs that meet any of the following conditions will not be accepted:

- Begin with “9”, “666”, or “000”
- Have the number “00" in positions 4 – 5
- Have the number “0000” in positions 6 – 9

*Need this data for your code? [Download these validation rules in JSON form](/schema/identity_verification_api.json)*

| Country | Country Code | ID Name | API Type | Format |
| --- | --- | --- | --- | --- |
| Argentina | `AR` | Documento Nacional de Identidad (DNI) | `ar_dni` | `\A\d{8,10}\z` |
| Australia | `AU` | Driver’s License | `au_drivers_license` | `\A[A-Z0-9]{6,11}\z` |
| Australia | `AU` | Passport Number | `au_passport` | `\A[A-Z0-9]{7,9}\z` |
| Brazil | `BR` | Cadastro de Pessoas Físicas (CPF) | `br_cpf` | `\A\d{11}\z` |
| Canada | `CA` | Social Insurance Number (SIN) | `ca_sin` | `\A\d{9}\z` |
| Chile | `CL` | Rol Único Nacional (RUN) | `cl_run` | `\A\d{7,8}[0-9Kk]\z` |
| China | `CN` | Resident Identity Card Number | `cn_resident_card` | `\A\d{15} |
| Colombia | `CO` | Número de Identificación Tributaria (NIT) | `co_nit` | `\A\d{8,10}\z` |
| Denmark | `DK` | Civil Personal Registration (CPR) | `dk_cpr` | `\A\d{9,11}\z` |
| Egypt | `EG` | National ID Card Number | `eg_national_id` | `\A\d{14}\z` |
| Hong Kong | `HK` | Hong Kong Identity (HKID) | `hk_hkid` | `\A[A-Z0-9]{8,9}\z` |
| India | `IN` | Permanent Account Number (PAN) | `in_pan` | `\A[A-Z]{5}\d{4}[A-Z]\z` |
| India | `IN` | Voter ID (EPIC) | `in_epic` | `\A[A-Z]{3}\d{7}\z` |
| Italy | `IT` | Codice Fiscale (CF) | `it_cf` | `\A[A-Z]{6}\d{2}[A-Z]\d{2}[A-Z]\d{3}[A-Z]\z` |
| Japan | `JP` | "My Number" Resident Record Code | `jp_my_number` | `\A\d{12}\z` |
| Jordan | `JO` | Civil Identification Number | `jo_civil_id` | `\A\d{10,14}\z` |
| Kenya | `KE` | Huduma Namba | `ke_huduma_namba` | `\A\d{10}\z` |
| Kuwait | `KW` | Civil ID Card Number | `kw_civil_id` | `\A\d{12}\z` |
| Malaysia | `MY` | National Registration Identity Card Number (NRIC) | `my_nric` | `\A\d{12}\z` |
| Mexico | `MX` | Clave Única de Registro de Población (CURP) | `mx_curp` | `\A[A-Z]{4}\d{6}[A-Z]{6}[A-Z\d]\d\z` |
| Mexico | `MX` | Registro Federal de Contribuyentes (RFC) | `mx_rfc` | `\A[A-Z]{4}\d{6}[A-Z\d]{3}\z` |
| New Zealand | `NZ` | Driver’s License | `nz_drivers_license` | `\A[A-Z]{2}\d{6}\z` |
| Nigeria | `NG` | National ID Number | `ng_nin` | `\A\d{11}\z` |
| Oman | `OM` | Identity Card Number | `om_civil_id` | `\A\d{8,9}\z` |
| Philippines | `PH` | PhilSys Number | `ph_psn` | `\A\d{12}\z` |
| Poland | `PL` | PESEL Number | `pl_pesel` | `\A\d{11}\z` |
| Romania | `RO` | Cod Numeric Personal (CNP) | `ro_cnp` | `\A\d{13}\z` |
| Saudi Arabia | `SA` | Biṭāgat Al-ʼaḥwāl | `sa_national_id` | `\A[A-Z0-9]{10}\z` |
| Singapore | `SG` | National Registration Identity Card | `sg_nric` | `\A[A-Z]\d{7}[A-Z]\z` |
| South Africa | `ZA` | Smart ID Card Number | `za_smart_id` | `\A\d{13}\z` |
| Spain | `ES` | Documento Nacional de Identidad (DNI) | `es_dni` | `\A\d{8}[A-Z]\z` |
| Spain | `ES` | Número de Identidad de Extranjero (NIE) | `es_nie` | `\A[A-Z]\d{7}[A-Z]\z` |
| Sweden | `SE` | Personnummer (PIN) | `se_pin` | `\A\d{9,12}\z` |
| Turkey | `TR` | T.C. Kimlik No. | `tr_tc_kimlik` | `\A\d{11}\z` |
| United States | `US` | Social Security Number (SSN or ITIN) | `us_ssn` | `\A\d{9}\z` |
| United States | `US` | Social Security Number (SSN or ITIN) Last 4 | `us_ssn_last_4` | `\A\d{4}\z` |

#### Input validation by country

Address standards vary by country. Some countries do not support postal code or region (the equivalent of "state" or "province" depending on the region). In other countries, postal code and/or region is valid in a sense, but superfluous and not typically collected in native forms.

Separately, variations in identity data coverage mean that collection of "ID Number" is a required part of Data Source Verification in some countries, optional (configured in the Dashboard Editor) for other countries, and not supported in countries that do not support Data Source Verification at all.

This regional variation is reflected in the table below.

The labels for each input field below correspond to the following input validations:

- `address.postal_code` and `address.region`

  - **Required** - If `address` is provided as part of the request, this field must always be present. For a set of accepted `region` codes, see [country subdivision codes in JSON form](https://plaid.com/documents/country_subdivision_codes.json). If a country is absent from this list, it means that `address.region` is not accepted for that country.
  - **Not Accepted** - If `address` is provided as part of the request, this field must be omitted. Specifically, the key should be omitted from the `address` object (as opposed to the key being present with a null value)

For the `region` field, Identity Verification accepts a subset of ISO 3166-2 subdivision codes, excluding subdivision codes that are not typically used in addresses (these exclusions primarily impact Italy and Hong Kong). For a full list of accepted codes, see [country subdivision codes in JSON form](https://plaid.com/documents/country_subdivision_codes.json). Only the portion of the region code after the hyphen should be sent, since the country code can be extrapolated from the `country` field. For example, for Rome, use `country: "IT", region: "RM"`.

- `id_number`

  - **Required** - If `id_number` is not provided for this country, `kyc_check` will treat the Identity Verification Session as having incomplete data and will require user interaction with the Identity Verification modal in order to collect `id_number`
  - **Optional** - If `id_number` is not provided for this country **and** your Identity Verification Template is configured to not collect ID numbers for this country then Identity Verification will automatically run the `kyc_check` step without any user interaction as long as all other required input fields are provided. Similar to the address fields, the API requires you omit the `id_number` key entirely if you are not providing ID number data
  - **Not Accepted** - ID number verification is not supported for this country and the `id_number` field should always be omitted from API requests

*Need this data for your code? [Download these validation rules in JSON form](/schema/identity_verification_api.json)*

| Country | Country Code | `address.postal_code` | `address.region` | `id_number` |
| --- | --- | --- | --- | --- |
| Afghanistan | `AF` | Required | Not Accepted | Not Accepted |
| Albania | `AL` | Required | Not Accepted | Not Accepted |
| Algeria | `DZ` | Required | Not Accepted | Not Accepted |
| Andorra | `AD` | Required | Not Accepted | Not Accepted |
| Angola | `AO` | Not Accepted | Not Accepted | Not Accepted |
| Anguilla | `AI` | Not Accepted | Not Accepted | Not Accepted |
| Antigua & Barbuda | `AG` | Not Accepted | Not Accepted | Not Accepted |
| Argentina | `AR` | Required | Required | Required |
| Armenia | `AM` | Required | Not Accepted | Not Accepted |
| Aruba | `AW` | Not Accepted | Not Accepted | Not Accepted |
| Australia | `AU` | Required | Required | Required |
| Austria | `AT` | Required | Not Accepted | Not Accepted |
| Azerbaijan | `AZ` | Required | Not Accepted | Not Accepted |
| Bahamas | `BS` | Not Accepted | Not Accepted | Not Accepted |
| Bahrain | `BH` | Required | Not Accepted | Not Accepted |
| Bangladesh | `BD` | Required | Not Accepted | Not Accepted |
| Barbados | `BB` | Required | Not Accepted | Not Accepted |
| Belarus | `BY` | Required | Not Accepted | Not Accepted |
| Belgium | `BE` | Required | Not Accepted | Not Accepted |
| Belize | `BZ` | Not Accepted | Not Accepted | Not Accepted |
| Benin | `BJ` | Not Accepted | Not Accepted | Not Accepted |
| Bermuda | `BM` | Required | Not Accepted | Not Accepted |
| Bhutan | `BT` | Required | Not Accepted | Not Accepted |
| Bolivia | `BO` | Not Accepted | Not Accepted | Not Accepted |
| Bosnia & Herzegovina | `BA` | Required | Not Accepted | Not Accepted |
| Botswana | `BW` | Required | Not Accepted | Not Accepted |
| Brazil | `BR` | Required | Required | Required |
| British Virgin Islands | `VG` | Required | Not Accepted | Not Accepted |
| Brunei | `BN` | Required | Not Accepted | Not Accepted |
| Bulgaria | `BG` | Required | Not Accepted | Not Accepted |
| Burkina Faso | `BF` | Not Accepted | Not Accepted | Not Accepted |
| Burundi | `BI` | Required | Not Accepted | Not Accepted |
| Cambodia | `KH` | Required | Not Accepted | Not Accepted |
| Cameroon | `CM` | Required | Not Accepted | Not Accepted |
| Canada | `CA` | Required | Required | Optional |
| Cape Verde | `CV` | Required | Not Accepted | Not Accepted |
| Caribbean Netherlands | `BQ` | Required | Not Accepted | Not Accepted |
| Cayman Islands | `KY` | Required | Not Accepted | Not Accepted |
| Central African Republic | `CF` | Required | Not Accepted | Not Accepted |
| Chad | `TD` | Not Accepted | Not Accepted | Not Accepted |
| Chile | `CL` | Required | Required | Required |
| China | `CN` | Required | Required | Required |
| Colombia | `CO` | Required | Required | Required |
| Comoros | `KM` | Required | Not Accepted | Not Accepted |
| Congo - Brazzaville | `CG` | Required | Not Accepted | Not Accepted |
| Cook Islands | `CK` | Required | Not Accepted | Not Accepted |
| Costa Rica | `CR` | Required | Not Accepted | Not Accepted |
| Croatia | `HR` | Required | Not Accepted | Not Accepted |
| Curaçao | `CW` | Not Accepted | Not Accepted | Not Accepted |
| Cyprus | `CY` | Required | Not Accepted | Not Accepted |
| Czech Republic | `CZ` | Required | Not Accepted | Required |
| Côte d’Ivoire | `CI` | Required | Not Accepted | Not Accepted |
| Denmark | `DK` | Required | Not Accepted | Required |
| Djibouti | `DJ` | Not Accepted | Not Accepted | Not Accepted |
| Dominica | `DM` | Required | Not Accepted | Not Accepted |
| Dominican Republic | `DO` | Required | Not Accepted | Not Accepted |
| Ecuador | `EC` | Required | Not Accepted | Not Accepted |
| Egypt | `EG` | Required | Required | Not Accepted |
| El Salvador | `SV` | Required | Not Accepted | Not Accepted |
| Eritrea | `ER` | Required | Not Accepted | Not Accepted |
| Estonia | `EE` | Required | Not Accepted | Not Accepted |
| Eswatini | `SZ` | Required | Not Accepted | Not Accepted |
| Ethiopia | `ET` | Required | Not Accepted | Not Accepted |
| Faroe Islands | `FO` | Required | Not Accepted | Not Accepted |
| Fiji | `FJ` | Not Accepted | Not Accepted | Not Accepted |
| Finland | `FI` | Required | Not Accepted | Required |
| France | `FR` | Required | Not Accepted | Not Accepted |
| French Polynesia | `PF` | Required | Not Accepted | Not Accepted |
| Gabon | `GA` | Required | Not Accepted | Not Accepted |
| Gambia | `GM` | Required | Not Accepted | Not Accepted |
| Georgia | `GE` | Required | Not Accepted | Not Accepted |
| Germany | `DE` | Required | Not Accepted | Not Accepted |
| Ghana | `GH` | Not Accepted | Not Accepted | Not Accepted |
| Gibraltar | `GI` | Required | Not Accepted | Not Accepted |
| Greece | `GR` | Required | Not Accepted | Not Accepted |
| Greenland | `GL` | Required | Not Accepted | Not Accepted |
| Grenada | `GD` | Required | Not Accepted | Not Accepted |
| Guatemala | `GT` | Required | Required | Not Accepted |
| Guernsey | `GG` | Required | Not Accepted | Not Accepted |
| Guinea | `GN` | Required | Not Accepted | Not Accepted |
| Guinea-Bissau | `GW` | Required | Not Accepted | Not Accepted |
| Guyana | `GY` | Required | Not Accepted | Not Accepted |
| Haiti | `HT` | Required | Not Accepted | Not Accepted |
| Honduras | `HN` | Required | Not Accepted | Not Accepted |
| Hong Kong | `HK` | Not Accepted | Required | Required |
| Hungary | `HU` | Required | Not Accepted | Not Accepted |
| Iceland | `IS` | Required | Not Accepted | Not Accepted |
| India | `IN` | Required | Required | Required |
| Indonesia | `ID` | Required | Required | Not Accepted |
| Iraq | `IQ` | Required | Not Accepted | Not Accepted |
| Ireland | `IE` | Required | Required | Not Accepted |
| Isle of Man | `IM` | Required | Not Accepted | Not Accepted |
| Israel | `IL` | Required | Not Accepted | Not Accepted |
| Italy | `IT` | Required | Required | Optional |
| Jamaica | `JM` | Not Accepted | Not Accepted | Not Accepted |
| Japan | `JP` | Required | Required | Optional |
| Jersey | `JE` | Required | Not Accepted | Not Accepted |
| Jordan | `JO` | Required | Not Accepted | Not Accepted |
| Kazakhstan | `KZ` | Required | Not Accepted | Not Accepted |
| Kenya | `KE` | Required | Not Accepted | Required |
| Kiribati | `KI` | Required | Not Accepted | Not Accepted |
| Kosovo | `XK` | Required | Not Accepted | Not Accepted |
| Kuwait | `KW` | Required | Not Accepted | Not Accepted |
| Kyrgyzstan | `KG` | Required | Not Accepted | Not Accepted |
| Laos | `LA` | Required | Not Accepted | Not Accepted |
| Latvia | `LV` | Required | Not Accepted | Not Accepted |
| Lebanon | `LB` | Required | Not Accepted | Not Accepted |
| Lesotho | `LS` | Required | Not Accepted | Not Accepted |
| Liberia | `LR` | Required | Not Accepted | Not Accepted |
| Liechtenstein | `LI` | Required | Not Accepted | Not Accepted |
| Lithuania | `LT` | Required | Not Accepted | Not Accepted |
| Luxembourg | `LU` | Required | Not Accepted | Not Accepted |
| Macao SAR China | `MO` | Required | Not Accepted | Not Accepted |
| Madagascar | `MG` | Required | Not Accepted | Not Accepted |
| Malawi | `MW` | Not Accepted | Not Accepted | Not Accepted |
| Malaysia | `MY` | Required | Required | Required |
| Maldives | `MV` | Required | Not Accepted | Not Accepted |
| Mali | `ML` | Not Accepted | Not Accepted | Not Accepted |
| Malta | `MT` | Required | Not Accepted | Not Accepted |
| Mauritania | `MR` | Required | Not Accepted | Not Accepted |
| Mauritius | `MU` | Required | Not Accepted | Not Accepted |
| Mexico | `MX` | Required | Required | Required |
| Moldova | `MD` | Required | Not Accepted | Not Accepted |
| Monaco | `MC` | Required | Not Accepted | Not Accepted |
| Mongolia | `MN` | Required | Not Accepted | Not Accepted |
| Montenegro | `ME` | Required | Not Accepted | Not Accepted |
| Montserrat | `MS` | Required | Not Accepted | Not Accepted |
| Morocco | `MA` | Required | Not Accepted | Not Accepted |
| Mozambique | `MZ` | Required | Not Accepted | Not Accepted |
| Myanmar (Burma) | `MM` | Required | Not Accepted | Not Accepted |
| Namibia | `NA` | Required | Not Accepted | Not Accepted |
| Nauru | `NR` | Required | Not Accepted | Not Accepted |
| Nepal | `NP` | Required | Not Accepted | Not Accepted |
| Netherlands | `NL` | Required | Not Accepted | Not Accepted |
| New Zealand | `NZ` | Required | Required | Optional |
| Nicaragua | `NI` | Required | Not Accepted | Not Accepted |
| Niger | `NE` | Required | Not Accepted | Not Accepted |
| Nigeria | `NG` | Required | Required | Required |
| Niue | `NU` | Required | Not Accepted | Not Accepted |
| North Macedonia | `MK` | Required | Not Accepted | Not Accepted |
| Norway | `NO` | Required | Not Accepted | Not Accepted |
| Oman | `OM` | Required | Not Accepted | Not Accepted |
| Pakistan | `PK` | Required | Not Accepted | Not Accepted |
| Palestinian Territories | `PS` | Required | Not Accepted | Not Accepted |
| Panama | `PA` | Not Accepted | Required | Not Accepted |
| Papua New Guinea | `PG` | Required | Not Accepted | Not Accepted |
| Paraguay | `PY` | Required | Not Accepted | Not Accepted |
| Peru | `PE` | Required | Required | Not Accepted |
| Philippines | `PH` | Required | Required | Optional |
| Poland | `PL` | Required | Not Accepted | Optional |
| Portugal | `PT` | Required | Required | Not Accepted |
| Qatar | `QA` | Not Accepted | Not Accepted | Not Accepted |
| Romania | `RO` | Required | Required | Not Accepted |
| Russia | `RU` | Required | Required | Not Accepted |
| Rwanda | `RW` | Required | Not Accepted | Not Accepted |
| Samoa | `WS` | Required | Not Accepted | Not Accepted |
| San Marino | `SM` | Required | Not Accepted | Not Accepted |
| Saudi Arabia | `SA` | Required | Not Accepted | Not Accepted |
| Senegal | `SN` | Required | Not Accepted | Not Accepted |
| Serbia | `RS` | Required | Not Accepted | Not Accepted |
| Seychelles | `SC` | Required | Not Accepted | Not Accepted |
| Sierra Leone | `SL` | Not Accepted | Not Accepted | Not Accepted |
| Singapore | `SG` | Required | Not Accepted | Required |
| Slovakia | `SK` | Required | Not Accepted | Not Accepted |
| Slovenia | `SI` | Required | Not Accepted | Not Accepted |
| Solomon Islands | `SB` | Required | Not Accepted | Not Accepted |
| Somalia | `SO` | Required | Not Accepted | Not Accepted |
| South Africa | `ZA` | Required | Required | Required |
| South Korea | `KR` | Required | Required | Not Accepted |
| South Sudan | `SS` | Not Accepted | Not Accepted | Not Accepted |
| Spain | `ES` | Required | Required | Optional |
| Sri Lanka | `LK` | Required | Not Accepted | Not Accepted |
| St. Kitts & Nevis | `KN` | Required | Not Accepted | Not Accepted |
| St. Lucia | `LC` | Required | Not Accepted | Not Accepted |
| St. Martin | `MF` | Required | Not Accepted | Not Accepted |
| St. Vincent & Grenadines | `VC` | Required | Not Accepted | Not Accepted |
| Sudan | `SD` | Required | Not Accepted | Not Accepted |
| Suriname | `SR` | Required | Not Accepted | Not Accepted |
| Sweden | `SE` | Required | Not Accepted | Required |
| Switzerland | `CH` | Required | Not Accepted | Not Accepted |
| São Tomé & Príncipe | `ST` | Required | Not Accepted | Not Accepted |
| Taiwan | `TW` | Required | Not Accepted | Not Accepted |
| Tajikistan | `TJ` | Required | Not Accepted | Not Accepted |
| Tanzania | `TZ` | Required | Not Accepted | Not Accepted |
| Thailand | `TH` | Required | Required | Not Accepted |
| Timor-Leste | `TL` | Required | Not Accepted | Not Accepted |
| Togo | `TG` | Not Accepted | Not Accepted | Not Accepted |
| Tonga | `TO` | Not Accepted | Not Accepted | Not Accepted |
| Trinidad & Tobago | `TT` | Not Accepted | Not Accepted | Not Accepted |
| Tunisia | `TN` | Required | Not Accepted | Not Accepted |
| Turkey | `TR` | Required | Not Accepted | Required |
| Turkmenistan | `TM` | Required | Not Accepted | Not Accepted |
| Turks & Caicos Islands | `TC` | Required | Not Accepted | Not Accepted |
| Tuvalu | `TV` | Not Accepted | Not Accepted | Not Accepted |
| Uganda | `UG` | Not Accepted | Not Accepted | Not Accepted |
| Ukraine | `UA` | Required | Not Accepted | Not Accepted |
| United Arab Emirates | `AE` | Not Accepted | Required | Not Accepted |
| United Kingdom | `GB` | Required | Not Accepted | Not Accepted |
| United States | `US` | Required | Required | Optional |
| Uruguay | `UY` | Required | Not Accepted | Not Accepted |
| Uzbekistan | `UZ` | Required | Not Accepted | Not Accepted |
| Vanuatu | `VU` | Not Accepted | Not Accepted | Not Accepted |
| Vatican City | `VA` | Required | Not Accepted | Not Accepted |
| Venezuela | `VE` | Required | Not Accepted | Not Accepted |
| Vietnam | `VN` | Required | Not Accepted | Not Accepted |
| Western Sahara | `EH` | Required | Not Accepted | Not Accepted |
| Yemen | `YE` | Not Accepted | Not Accepted | Not Accepted |
| Zambia | `ZM` | Required | Not Accepted | Not Accepted |
| Zimbabwe | `ZW` | Not Accepted | Not Accepted | Not Accepted |

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
