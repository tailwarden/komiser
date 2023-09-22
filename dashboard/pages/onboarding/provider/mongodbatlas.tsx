import MongoDbAtlasAccountDetails from '@components/account-details/MongoDBAtlasAccountDetails';
import { allProviders } from '@utils/providerHelper';
import ProviderContent from '@components/onboarding-wizard/ProviderContent';

export default function MongoDBAtlasCredentials() {
  return (
    <ProviderContent
      provider={allProviders.MONGODB_ATLAS}
      providerName="MongoDB Atlas"
      description="MongoDB Atlas is a fully managed cloud database service provided
  by MongoDB, Inc. It allows users to run their MongoDB databases on
  popular cloud providers such as AWS, Google Cloud, and Azure."
    >
      <MongoDbAtlasAccountDetails />
    </ProviderContent>
  );
}
