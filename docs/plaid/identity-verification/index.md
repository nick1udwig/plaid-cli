---
title: "Identity Verification - Introduction to Identity Verification | Plaid Docs"
source_url: "https://plaid.com/docs/identity-verification/"
scraped_at: "2026-03-07T22:04:54+00:00"
---

# Introduction to Identity Verification

#### Verify the identity of your users for global KYC and anti-fraud.

Get started with Identity Verification

[API Reference](/docs/api/products/identity-verification/)[Request Access](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access)[Quickstart](https://github.com/plaid/idv-quickstart)[Demo](https://plaid.coastdemo.com/share/68a50b5fc54642a282e49c0e?zoom=100)

#### Overview

Plaid Identity Verification (IDV) lets you verify the identity of your customers and seamlessly stitches together verification methods. Using Identity Verification, you can verify identification documents, phone numbers, name, date of birth, ID numbers, addresses, and more. Identity Verification also integrates directly with [Monitor](/docs/monitor/) for an end-to-end verification and KYC solution, allowing you to check for the presence of a user on watchlists and PEP lists, in addition to verifying identity.

The user starts the Identity Verification process...

![The user starts the Identity Verification process...](/assets/img/docs/idv/tour/idv-link-1.png)

![...selects their country from a configurable list...](/assets/img/docs/idv/tour/idv-link-2.png)

![...enters their phone number...](/assets/img/docs/idv/tour/idv-link-3.png)

![...and, if you've enabled SMS-verification, verifies it.](/assets/img/docs/idv/tour/idv-link-4.png)

![They provide their full name...](/assets/img/docs/idv/tour/idv-link-5.png)

![...address...](/assets/img/docs/idv/tour/idv-link-6.png)

![...date of birth...](/assets/img/docs/idv/tour/idv-link-7.png)

![...and, if you've enabled Data Source Verification, an ID number.](/assets/img/docs/idv/tour/idv-link-8.png)

![With Document Verification enabled, they pick a document from a configurable list...](/assets/img/docs/idv/tour/idv-link-9.png)

![...and photograph the front...](/assets/img/docs/idv/tour/idv-link-10.png)

![...and the back.](/assets/img/docs/idv/tour/idv-link-11.png)

![With Selfie Check enabled, the user grants camera access...](/assets/img/docs/idv/tour/idv-link-12.png)

![...takes a selfie...](/assets/img/docs/idv/tour/idv-link-13.png)

![...decides whether to save their profile with Plaid...](/assets/img/docs/idv/tour/idv-link-14.png)

![...and the flow is done.](/assets/img/docs/idv/tour/idv-link-15.png)

Plaid Identity Verification can also be used together with [Plaid Identity](/docs/identity/) in a single workflow to provide full identity verification and fraud reduction. Identity Verification is used to verify that the user of your app is the person they claim to be, while [Identity](/docs/identity/) is used to confirm that the ownership information on their linked bank or credit card account matches this verified identity.

Identity Verification can be used to verify end users in nearly 190 countries. To integrate with Identity Verification, your company must be based in the US, Canada, or UK.

#### Identity Verification checks

Plaid offers a variety of verification checks to choose from. You can combine multiple verification checks within a single Identity Verification template via the Workflow Management editor, and even run certain checks conditionally, such as only running a particular verification if another verification fails or indicates a high-risk session, or running different checks based on your user's country of origin.

After running your selected checks, Identity Verification will also display a set of granular risk assessments (email risk, phone risk, etc.) in the Dashboard for each session, along with an overall risk score, helping you to understand which factors contributed to the session passing or failing and which sessions have been identified as the highest risk.

When configuring your checks, it is typically recommended to start with the default templates and risk rules, then fine-tune them as needed based on your results.

##### SMS Verification

SMS Verification verifies a user's phone number by asking them to enter a code sent via SMS.

##### Data Source Verification

Data Source Verification (formerly known as Lightning Verification) verifies a user's name, address, date of birth, phone number, and ID number (such as SSN) against available records. These records are sourced from high-quality, trusted databases such as voter and driver registration records, property records, and credit bureau records. You can configure the level of matching required for each field in your template’s Identity Rules. The results of this check will be summarized in the `kyc_check` object in the API.

Note that certain end user segments, such as recent immigrants and end users aged 18-21, are less likely to pass data source verification, due to having a "thin file" (i.e., fewer data source records available to verify against). If these groups constitute a substantial portion of the audience you will be verifying, you may wish to enable Document Verification as a backup verification method.

##### Document Verification

Document Verification prompts the user to upload an image of an identity document. Plaid will use anti-fraud checks to verify that the document appears to be legitimate and is unexpired. Plaid will also verify the date of birth and name on the document against the data provided by the user (other fields that may be present on the document, such as address or national ID number, will not be verified). Plaid supports a wide range of documents; you can review and configure which document types are supported for each country. You can also require certain classes of document types, e.g. only photo IDs.

If the user is on a non-mobile device, Plaid will display a QR code that they can use to perform Document Verification on mobile. They will resume the computer-based flow once the Document Verification upload is complete. This handoff process is automatic and does not require any integration work on your part.

Document Verification can be used as a potential fallback for Data Source Verification (e.g. for verifying thin file customers) or as a step-up method of verification for customers with a lower [Trust Index Score](/docs/identity-verification/#trust-index-scores-us-only). It can also be used as a primary method of verification, especially in countries where Data Source Verification is not supported.

##### AAMVA Check (beta, US only)

AAMVA (American Association of Motor Vehicle Administrators) checks can optionally be enabled as part of Document Verification. When enabled, if the end user uploads an image of a drivers license from a [participating US state](https://www.aamva.org/it-systems-participation-map?id=594), Identity Verification will compare the text on the uploaded drivers license with the AAMVA's records. Fields compared include name, gender, height, eye color, address, date of birth, and drivers license number.

##### Selfie Check

When Selfie Check is enabled, the user will be asked to take a selfie on their mobile device as part of the verification process. Plaid will verify that the selfie submitted is a genuine, live video of a real human. If Document Verification is enabled, Selfie Check will also verify that the selfie matches the photo in the document.

If the user is on a non-mobile device, Plaid will display a QR code that they can use to perform the Selfie Check on mobile. They will resume the computer-based flow once the Selfie Check is complete. This handoff process is automatic and does not require any integration work on your part.

After the end user has passed a Selfie Check, you can [issue additional on-demand liveness checks](/docs/identity-verification/#retries) at any time as a form of step-up verification. These checks will check both for liveness and for a match to the previous Selfie Check.

##### Age verification

If both Selfie Checks and Document Verification are enabled, an age verification check will automatically be run. Identity Verification will check that the user's photos appear consistent with the date of birth the user has provided.

##### AML Screening

If you're also using [Monitor](/docs/monitor/), AML Screening allows you to screen a user against government watchlists during an IDV session by incorporating one of your Monitor screening programs into an IDV template.

##### Address requirements

Within the workflow configuration, you can set a number of address requirements. You can optionally require that the user provide a residential address and/or a physical address (i.e., block PO box addresses). You can also require that any ID documents presented in the Document Verification check are issued by the same country as the country in the user's address.

##### Risk Check

Plaid automatically runs anti-fraud checks in the background of a user’s verification session and summarizes the results into several categories. The data collected in each category is analyzed and assigned a risk level. Whether the overall risk check passes or fails is determined by the template’s Risk Rules, which you can configure. For more details on the risk check process and what each check represents, see [Risk checks](/docs/identity-verification/risk-checks/).

##### Trust Index Scores (US only)

For verifications of US users, Identity Verification will calculate a Trust Index Score, which provides an overall holistic risk assessment based on the Risk Check results. Via the Dashboard, you can view the primary drivers of the Trust Index Score for any given verification. You can also configure a minimum Trust Index Threshold in addition to your template's Risk Rules. The Dashboard will display several recommended thresholds, as well as what percentage of your traffic would be blocked at each score threshold, once you have enough traffic assessed.

You can also configure conditional step-up verifications based on the Trust Index Score. For example, you can allow users with a high Trust Index Score to bypass documentary verification, while requiring for users with a lower score.

##### Checks not performed by Identity Verification

Identity Verification is not a background check tool; it will not report whether an end user has a criminal record.

Identity Verification does not check citizenship, immigration status, or residency status, although you may be able to build your own business logic to determine this supported by Identity Verification results. For example, a valid passport from a country will typically indicate that the end user is a citizen of that country.

Identity Verification does not check for duplicate user submissions based on identity data such as phone number or national ID number (although multiple submissions from the same device, or multiple submissions using different names but the same ostensibly unique PII such as phone number, are tracked and will increase the risk score). Instead, it checks for duplicates using the `client_user_id` provided when calling [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see [Client user ID](/docs/identity-verification/#client-user-id).

Per Plaid's [terms and policies](https://plaid.com/legal/), Identity Verification is not intended to be used by end users under the age of 18.

#### Identity Verification integration process

Identity Verification is not available by default in the Sandbox environment. To obtain Sandbox access, [request full Production access for Identity Verification](https://dashboard.plaid.com/settings/team/products), which will automatically grant Sandbox access, or, if you are not ready for Production access, [contact Sales](https://plaid.com/products/identity-verification/#contact-form). To obtain Sandbox access for Identity Verification if you are already using another Plaid product, you can also contact your account manager or submit an [access request ticket](https://dashboard.plaid.com/support/new/admin/account-administration/request-product-access).

The steps below explain how to set up the standard Identity Verification flow. For details on other flow options, including the backend-only flow and hosted UI flow, see [Integration options for Identity Verification](/docs/identity-verification/#integration-options-for-identity-verification).

Note that unlike other Plaid products, which are configured primarily via the API, most options for Identity Verification are configured on the web, via the [Template Editor](https://dashboard.plaid.com/identity_verification/templates) within the Identity Verification Dashboard. Detailed documentation on configuring Identity Verification templates is available via in-app help within the Template Editor. To request access to the Template Editor and Identity Verification Dashboard, contact your Plaid Account Manager, or file a [Product Access request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).

1. Visit the [Plaid Dashboard](https://dashboard.plaid.com/) and select **Identity Verification** from the left sidebar. If you don't see this option, you may need to talk to your Plaid Account Manager and ask them to enable this product for you, or file a [Product access request](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access).
2. Select or create the template you want to use, click the **Integration** button, and copy the `template_id`.
   - (Optional) If you plan to integrate with [Monitor](/docs/monitor/), make sure to enable the **AML Screening** option and select the appropriate Monitor program within the IDV template editor.
3. (Optional) Call [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) if you already have information about your user that you can fill out on their behalf. See [pre-populating user information](/docs/identity-verification/#streamlining-link-by-pre-populating-user-information-hybrid-link-flow) for more details.
4. Call [`/link/token/create`](/docs/api/link/#linktokencreate). In addition to the required parameters, you will need to provide the following:
   - For `identity_verification.template_id`, use the `template_id` you copied in step 2.
   - For `products`, use `["identity_verification"]`. Identity Verification is mutually exclusive with other products.
   - Provide `user.client_user_id` and optionally `user.email_address`. See below for details on these fields.
5. On your web or mobile client, create an instance of Link using the `link_token` returned by [`/link/token/create`](/docs/api/link/#linktokencreate); for more details, see the [Link documentation](/docs/link/).
6. When the user has successfully completed the client-side Link flow, you will receive an [`onSuccess`](/docs/link/web/#onsuccess) callback and be able to review the session in the Dashboard. The `onSuccess` callback will return a Link session ID that can be used as the `identity_verification_id` in the next step. Note that `onSuccess` does *not* indicate that the user has passed the Identity Verification checks, but simply that they have submitted all the information requested in the Link flow. Unlike some other Plaid products, Identity Verification will not return a public token in `onSuccess`.
7. To retrieve the status of your user's verification attempt, make a call to [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget), passing in the verification session ID. You can retrieve this session ID from the metadata in the Link callbacks, from the [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) response, or from the [Identity Verification webhooks](/docs/identity-verification/webhooks/) that Plaid sends during the process.
8. You can retrieve a list of all of your user's verification attempts by making a call to [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist), passing in the `client_user_id` and the `template_id`.

##### Events and callbacks

You can optionally track the progress of users through the Link flow via client-side callbacks and events. For more information, see [Link callbacks](/docs/identity-verification/link/).

##### Mobile support

Identity Verification is fully supported by Plaid's mobile SDKs. Because the end user may need to use the camera during the Identity Verification flow, your app must request certain permissions in order to fully support Identity Verification flows on mobile. For more details, see the documentation for the [Android](/docs/link/android/#enable-camera-support-identity-verification-only) and [iOS](/docs/link/ios/#camera-support-identity-verification-only) SDKs.

##### Client User Id

The mandatory `client_user_id` field should be a unique and persistent identifier for your customer, such as the `id` field on your users table.

Identity Verification intelligently handles sessions being started with the same `client_user_id` multiple times. If your customer starts a Link session, closes it, reopens it, and reopens your Link integration, their session will resume from where they left off. Likewise, if your customer has completed their Link session in the past (by either failing verification or passing), Plaid will not serve them another session unless you've manually authorized another attempt from the Dashboard or made a call to [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry).

Plaid indexes the `client_user_id` you provide, allowing you to look up users using your internal id.

If you do not want to expose your internal id to Plaid directly, you can symmetrically encrypt the identifier with a secret key that only your servers know.

##### Supplementing with email

During [`/link/token/create`](/docs/api/link/#linktokencreate) or [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate), Identity Verification accepts `email_address` as an optional additional parameter. While the field is optional, we highly recommend providing it. Identity Verification will include the user's email in the Link session and perform a number of anti-fraud related checks, including analyzing social networks, abusive emails, and email address trustworthiness.

#### Integration options for Identity Verification

Identity Verification offers several options for streamlined flows.

The Plaid Link flow is designed to collect high quality user data, through an optimized UI and input validation. If you are bypassing the Link data entry UI and are instead passing user data via API, be sure you are providing data in the correct format expected by the API and have taken steps to avoid issues such as typos or incorrect data entry (e.g. an end user entering a date of birth using a different date format than your app expects). Providing unvalidated or malformatted data via the API will result in degraded success rates for Data Source Verification.

For more details on data formatting requirements and validation rules you can use for input you collect, see [input validation](/docs/identity-verification/hybrid-input-validation/).

##### Streamlining Link by pre-populating user information (hybrid Link flow)

If you already know some information about your user (such as their name, phone number, or address), you can simplify the verification process by making a call to [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate), passing in any information you know about your user, including their `client_user_id`. When you open up a Link session against this same `client_user_id`, Identity Verification can make use of the information that you passed in. Any fields or screens requesting information you already provided will be skipped (not shown) in the UI. This means that an end user will not be able to override or change pre-populated user information during the Link flow. For data format specifications, see [Input Validation](/docs/identity-verification/hybrid-input-validation/).

##### Data Source Checks without UI (backend flow)

If you have already obtained a user's information and their consent to share it with Plaid, and do not need interactive capabilities such as Document Check, Selfie Check, or SMS verification, you can bypass the Identity Verification UI entirely and run an entirely programmatic verification:

1. Create a template that requires only Data Source (KYC) check and does not use SMS verification.
2. Call [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) with all the user data required by this template and `gave_consent: true`.
3. Wait for the `STATUS_UPDATED` webhook to fire for the `id` returned by [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate). Alternatively, you can poll [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) with this id until the `status` is no longer `active`.
4. Call [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget) with the id returned by [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) to retrieve the status of the verification.

When using the backend-only flow, the Network Risk, Device and IP Address, and Behavioral Analytics checks will be unavailable, as they require client-side input.

##### Verification links (hosted flow)

Identity Verification offers the option to generate Plaid-hosted verification links. These links can be used in scenarios where your user signed up for your service outside of an app or website -- for example if a user is opening a bank account or applying for a loan in person at a brick-and-mortar retail location. To generate a verification link, click **Create verification** from the Identity Verification template view in the Dashboard. Alternatively, you can call [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) with `is_shareable: true`; a verification link will be returned in the `shareable_url` field of the response. Once you have shared the verification link, the user can open the link and complete the verification on their phone or other device. Unlike Plaid [Hosted Links](/docs/link/hosted-link/), Identity Verification links do not expire.

Note that verification links are intended to be sent directly to your users; hosting verification links via iframe will be blocked.

##### Auto-fill

When auto-fill is enabled, Plaid will use the user's phone number and date of birth to attempt to auto-fill other fields. The user will be asked to confirm their auto-filled information is accurate before submitting. Auto-fill can only be enabled when using SMS Verification and Data Source Verification and is only available for US verifications.

Enabling auto-fill when possible is strongly recommended, as it has a substantial positive impact on Link flow conversion.

When auto-fill is enabled, full SSN collection cannot be guaranteed; in some cases, 4 digits of the SSN may be returned, even if you have enabled collecting the 9 digit SSN. If it is critical that you collect the end user's full SSN within the Identity Verification flow, you should not enable auto-fill.

Note that auto-fill is not the same as the Plaid [Returning User Experience](https://plaid.com/docs/link/returning-user/). The Returning User Experience is an automatic experience within Link, based on the end user's saved profile that they created during a previous Link session, while auto-fill is a configurable setting and uses Data Source Verification records for auto-filling data.

##### Hiding verification outcome

You can customize many aspects of the Identity Verification flow, including the final pane shown to the end user. If you have a more complex user risk evaluation model that uses Identity Verification as one of multiple inputs, you should select the option in the Workflow pane to "Remove Final Status Screen", preventing the user from seeing a result that may not reflect the true outcome of your verification process.

##### Overrides

If you believe the result of an Identity Verification step was incorrect, you can manually override the result via the Dashboard by clicking the **Override result** button next to the specific step on the verification detail page. It is not possible to override a result via the API. Only the result of failed steps can be overridden; you cannot override the result of a step that has passed verification. You cannot override the result of the session as a whole, but overriding a step's result will cause the verification logic to be re-evaluated for that session, which may cause the session to pass.

##### Retries

If you want to give a user another attempt to verify their identity -- for example, if they failed verification due to a typo -- you can issue them a retry. This can be done from the Dashboard when reviewing a verification by clicking **Actions > Request retry from this user** and selecting the steps to include, or via API by calling [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) with their `client_user_id`, `template_id`, and the `strategy` you want to use.

Retries can be issued from any verification check. Any steps included in the retry will be reset and cleared of any PII entered by the user or passed via API in previous verification attempts. Data Source data can be passed programmatically for a retry by calling [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) and passing the user’s data in the `user` object. [`/identity_verification/create`](/docs/api/products/identity-verification/#identity_verificationcreate) should only be used to create an initial session for users who have never been verified before.

Each retry will generate a new `session_id`, with the same `client_user_id` and `template_id` used in the original session. Retries will use the most recent version of the template, which may be different from the version used in the original verification session. Retries are billed at the same rate as regular sessions.

While allowing retries is useful to help users who have legitimately made an error, allowing too many retry attempts can attract fraud. We recommend starting with no more than two retry attempts allowed per user.

You can also issue retries for successful verifications. In addition to retrying steps, successful verifications allow you to issue a liveness retry, which prompts the user to take a new selfie and both tests it for liveness and compares it against their previous selfie. To use this approach, select **Actions > Request retry from this user > Authenticate User > Start Re-authentication** or call [`/identity_verification/retry`](/docs/api/products/identity-verification/#identity_verificationretry) with `steps.liveness: true`.

###### Step-up verification

Retries can also be used as part of a "step-up" verification process. To use retries in this way, send the user through Identity Verification, and examine the results via the Dashboard or [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget). If their risk results meet your criteria for requiring a step-up, you can issue a retry using a different `template_id`. For example, you might send users through a template that doesn't require documentary verification, but if their risk check failed, you might step them up to a new template that does require documentary verification.

##### Financial Account Matching

If you are a customer of [Identity](/docs/identity/) or [Assets](/docs/assets/), you can seamlessly compare a user's financial accounts with their Identity Verification profile. When a user verifies their identity and links their accounts with Plaid, we detect if they are linking a bank account that belongs to them. In the Plaid Dashboard, you’ll see "match" scores against the account owner’s name, email, phone number, and address.

To turn on this feature, you will need to enable it in the template editor and call [`/identity/get`](/docs/api/products/identity/#identityget), [`/identity/match`](/docs/api/products/identity/#identitymatch), or [`/asset_report/create`](/docs/api/products/assets/#asset_reportcreate) with the same `client_user_id` that was used for Identity Verification. Financial Account Matching will only run if the `client_user_id` is associated with a successful (passed) Identity Verification session.

##### Fraud labeling

Identity Verification offers the ability to flag a session as fraudulent via the Dashboard, using the **Mark as fraud** or **Mark as not fraud** buttons. Labeling fraudulent sessions will allow Identity Verification to refine its fraud model, resulting in higher fraud detection accuracy for your integration. If necessary, you can also perform a one-time bulk upload of fraud labels; contact your Account Manager for more details. Fraud labeling a session does not change its pass/fail outcome.

#### Supported languages

Identity Verification currently supports flows in English, French, Spanish, Japanese, and Portuguese. Link will attempt to auto-detect the correct language based on the user's browser settings. Users can also manually change the language of their session through a dropdown in the Link UI. Unlike other Plaid products, Identity Verification does not use the language setting in [`/link/token/create`](/docs/api/link/#linktokencreate).

#### Supported countries

Identity Verification supports end users in all countries, with the exception of North Korea, Libya, DR Congo, Iran, Cuba, and Syria.

Data Source checks are available in the following countries: Argentina, Australia, Austria, Belgium, Brazil, Canada, Chile, China, Colombia, Czechia, Denmark, Finland, France, Germany, Gibraltar, Hong Kong, India, Ireland, Italy, Japan, Kenya, Luxembourg, Malaysia, Mexico, Netherlands, New Zealand, Nigeria, Norway, Philippines, Poland, Portugal, Singapore, Slovakia, Spain, Sweden, Switzerland, Turkey, United Kingdom, United States.

Note that Data Source Verification success rates vary by country. For more details on a specific country, contact your account manager. If you are running Data Source Verification checks on users outside the US, you may wish to use Documentary Verification as a fallback method in order to increase verification success rates.

Selfie Checks and Documentary Verification are available in all supported countries.

#### Supported documents

Document Verification supports over 16,000 distinct government-issued identification document types. You can check which document types are supported in a specific country from the Dashboard. From the Template Editor, under the **Workflow** tab, click **Assign Countries** and either add a new country or select one of the countries you've already configured for the template. Click **Physical Document Collection Options** to see which document types are supported for this country.

#### Data retention and redaction

Identity Verification will retain data about user sessions as long as you maintain an active Identity Verification contract with Plaid. To request that user personally identifiable information (PII) be redacted on a rolling basis, such as after 90 days, contact support or your Plaid Account Manager to request a custom retention period. Redaction is permanent, and once you have set up a custom retention period, you cannot regain access to redacted data. You can also redact data on an ad-hoc basis via the Dashboard.

If you end your contract with Plaid, you will lose access to the Identity Verification Dashboard and API, including access to information about previous sessions. If your use case requires that you retain user session data after your contract ends, use the API (via [`/identity_verification/list`](/docs/api/products/identity-verification/#identity_verificationlist) and [`/identity_verification/get`](/docs/api/products/identity-verification/#identity_verificationget)) to export the data prior to the end of your contract.

#### Protect Dashboard

The [Protect Dashboard](https://dashboard.plaid.com/sandbox/protect) is an alternative UI for managing Identity Verification rules and viewing activity.

![](/assets/img/docs/protect/protect-dashboard.png)

Example Protect Dashboard main page

The Protect Dashboard provides capabilities not available in the Identity Verification Dashboard, including configurable workflows and approval logic based on custom rule groups. Rule groups can reference over 100 user attributes, including Trust Index score, location, age, and email domain.

You can query this data directly using custom filters. For example, creating a filter such as "Trust Index < 30" or "more than 5 unique IP addresses" immediately returns matching live and historical sessions.

These same filter criteria and user attributes can also be used to create rule groups, which you can then use in your approval logic via the "Configure" tab. For example, a rule group can be used to create a workflow that rejects users under a specified age or requires a higher Trust Index score for sessions originating from a specific country. Running a search shows which historical sessions would have matched the rule group, effectively backtesting the logic against production traffic.

From any aggregated view, you can drill into an individual session to see the full event timeline and signal values used in decisioning. Data can be exported for audits, compliance, or downstream analysis.

To access the Protect Dashboard, click the "Protect" icon on the left hand navigation pane of the Dashboard. The Protect Dashboard is available for all Identity Verification customers based in the US or UK.

Identity Verification customers will have the following limitations when using the Protect Dashboard:

- Only two weeks of history is available to search or visualize
- Only two custom rule groups can be created

To upgrade to two years of history and unlimited ruleset groups, [Contact Sales](https://plaid.com/contact/) or your Account Manager to learn more about Protect.

#### Testing Identity Verification

For more information about testing Identity Verification in Sandbox, see [Testing Identity Verification](/docs/identity-verification/testing/).

#### Sample app and additional resources

For a sample integration, see the [Identity Verification Quickstart](https://github.com/plaid/idv-quickstart) on GitHub.

The PDF resources below provide an in-depth view of IDV integration and configuration options. Note that they are not updated as frequently as the docs and may not reflect the most recent updates to Identity Verification capabilities.

- [Identity Verification and Monitor solution guide](https://plaid.com/documents/plaid-idv-monitor-solution-guide.pdf): A detailed, step-by-step walkthrough for integrating and configuring Identity Verification and Monitor.
- [Identity Verification and Monitor no-code guide](https://plaid.com/documents/idv-monitor-no-code-testing-guide.pdf): A walkthrough of setting up a Identity Verification and Monitor integration with no code, using the [verification links](/docs/identity-verification/#verification-links-hosted-flow) option. This demo flow can help you quickly work with Identity Verification and Monitor and try out the end-user and admin experiences.

#### Identity Verification pricing

Identity Verification has a base fee for every verification attempted, as well as separate fees for Data Source verification, Selfie checks, and Document checks. For more details, see the [Identity Verification fee model](/docs/account/billing/#identity-verification-fee-model). To view the exact pricing you may be eligible for, [apply for Production access](https://dashboard.plaid.com/overview/production) or [contact Sales](https://plaid.com/contact/).

#### Next steps

To learn more about building with Identity Verification, see the [Identity Verification API Reference](/docs/api/products/identity-verification/).

If you're ready to launch to Production, see the Launch Center.

[#### Launch Center

See next steps to launch in Production

Launch](https://dashboard.plaid.com/developers/launch-center)

#### Launch Center

See next steps to launch in Production

[Launch](https://dashboard.plaid.com/developers/launch-center)

Developer community

[![GitHub](/assets/img/icons/icon-social-github.svg)](https://github.com/plaid)

[![Stack Overflow](/assets/img/icons/icon-social-stack-overflow.svg)](https://stackoverflow.com/questions/tagged/plaid?tab=Newest)

[![YouTube](/assets/img/icons/icon-social-youtube-black.svg)](https://www.youtube.com/c/PlaidInc)

[![Discord](/assets/img/icons/discord-footer.svg)](https://discord.gg/sf57M8DW3y)
