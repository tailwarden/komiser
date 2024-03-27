import classNames from 'classnames';
import { OVHCredentials } from '@utils/cloudAccountHelpers';
import LabelledInput from '../onboarding-wizard/LabelledInput';
import { CloudAccountPayload } from '../cloud-account/hooks/useCloudAccounts/useCloudAccount';

interface OVHAccountDetailsProps {
  cloudAccountData?: CloudAccountPayload<OVHCredentials>;
  hasError?: boolean;
}

function OVHAccountDetails({
  cloudAccountData,
  hasError = false
}: OVHAccountDetailsProps) {
  return (
    <div className="flex flex-col space-y-4 py-10">
      <LabelledInput
        type="text"
        id="account-name"
        name="name"
        required={true}
        value={cloudAccountData?.name}
        label="Account name"
        placeholder="my-ovh-account"
      />

      <div
        className={classNames(
          'flex flex-col space-y-1 rounded-md p-5',
          hasError ? 'bg-red-50' : 'bg-gray-50'
        )}
      >
        <LabelledInput
          type="text"
          id="endpoint"
          name="endpoint"
          label="Endpoint"
          required={true}
          value={cloudAccountData?.credentials.endpoint}
          subLabel="The connection endpoint"
          placeholder="ovh-eu"
        />
        <LabelledInput
          type="text"
          id="application_key"
          name="application_key"
          required={true}
          value={cloudAccountData?.credentials.applicationKey}
          label="Application Key"
          subLabel="Your application Key"
          placeholder="my_app_key"
        />
        <LabelledInput
          type="password"
          id="application_secret"
          name="application_secret"
          required={true}
          value={cloudAccountData?.credentials.applicationSecret}
          label="Application Secret"
          subLabel="An Application Secret"
          placeholder="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
        />
        <LabelledInput
          type="text"
          id="my_consumer_key"
          name="my_consumer_key"
          required={true}
          value={cloudAccountData?.credentials.consumerKey}
          label="My Consumer Key"
          subLabel="Your consumer key"
          placeholder="my_consumer_key"
        />
      </div>
      {hasError && (
        <div className="text-sm text-red-500">
          We couldn&apos;t connect to your OVH account. Please check if the
          details are correct.
        </div>
      )}
    </div>
  );
}

export default OVHAccountDetails;
