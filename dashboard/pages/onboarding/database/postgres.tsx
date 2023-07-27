import Head from 'next/head';

import { allDBProviders } from '../../../utils/providerHelper';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../../components/onboarding-wizard/OnboardingWizardLayout';
import LabelledInput from '../../../components/onboarding-wizard/LabelledInput';
import DatabasePurplin from '../../../components/onboarding-wizard/DatabasePurplin';
import CredentialsButton from '../../../components/onboarding-wizard/CredentialsButton';

export default function PostgreSQLCredentials() {
  const database = allDBProviders.POSTGRES;

  const handleNext = () => {
    // TODO: (onboarding-wizard) complete form inputs, validation, submission and navigation
  };

  return (
    <div>
      <Head>
        <title>Configure Postgres - Komiser</title>
        <meta name="description" content="Configure Postgres - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Configure your Postgres database"
          progressBarWidth="81%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              PostgreSQL is a powerful and feature-rich open-source RDBMS known
              for its extensibility and robustness. It offers advanced SQL
              capabilities, support for complex queries, data integrity
              constraints, transactions, and scalability.
            </div>
          </div>

          <div className="flex flex-col space-y-4 py-10">
            <div className="space-y-[0.2]">
              <LabelledInput
                type="text"
                id="hostname"
                label="Hostname"
                subLabel="The server where the Postgres server is hosted"
                placeholder="my-postgres-server"
              />
              <LabelledInput
                type="text"
                id="database"
                label="Database"
                subLabel="The name of the database where Komiser will insert/save the data"
                placeholder="my_database"
              />
              <LabelledInput
                type="text"
                id="username"
                label="Username"
                subLabel="The Postgres username"
                placeholder="user"
              />
              <LabelledInput
                type="text"
                id="password"
                label="Password"
                subLabel="The Postgres password"
                placeholder="Example0000*"
              />
            </div>
          </div>

          <CredentialsButton handleNext={handleNext} nextLabel="Add database" />
        </LeftSideLayout>

        <RightSideLayout>
          <DatabasePurplin database={database} />
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
