import { FormEvent, useState } from 'react';
import router from 'next/router';
import Head from 'next/head';

import { allDBProviders } from '@utils/providerHelper';

import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '@components/onboarding-wizard/OnboardingWizardLayout';
import LabelledInput from '@components/onboarding-wizard/LabelledInput';
import DatabasePurplin from '@components/onboarding-wizard/DatabasePurplin';
import CredentialsButton from '@components/onboarding-wizard/CredentialsButton';
import settingsService from '@services/settingsService';

import Toast from '@components/toast/Toast';
import DatabaseErrorMessage from '@components/onboarding-wizard/DatabaseErrorMessage';
import { useToast } from '@components/toast/ToastProvider';

export default function PostgreSQLCredentials() {
  const databaseProvider = allDBProviders.POSTGRES;

  const { toast, showToast, dismissToast } = useToast();

  const [isError, setIsError] = useState<boolean>(false);
  const [hostname, setHostname] = useState<string>('');
  const [database, setDatabase] = useState<string>('');
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  const handleNext = (e: FormEvent) => {
    e.preventDefault();

    const payload = JSON.stringify({
      type: 'POSTGRES',
      hostname,
      database,
      username,
      password
    });

    settingsService.saveDatabaseConfig(payload).then(res => {
      setIsError(false);

      if (res === Error) {
        setIsError(true);
      } else {
        showToast({
          hasError: false,
          title: 'Database connected',
          message:
            'Your Postgres database has been successfully connected to Komiser.'
        });
        router.push('/onboarding/choose-cloud/');
      }
    });
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

          {isError && <DatabaseErrorMessage />}

          <form onSubmit={handleNext}>
            <div className="flex flex-col space-y-4 py-10">
              <div className="space-y-[0.2]">
                <LabelledInput
                  type="text"
                  id="hostname"
                  label="Hostname"
                  required
                  value={hostname}
                  onChange={e => setHostname(e.target.value)}
                  subLabel="The server where the Postgres server is hosted"
                  placeholder="my-postgres-server"
                />
                <LabelledInput
                  type="text"
                  id="database"
                  label="Database"
                  required
                  value={database}
                  onChange={e => setDatabase(e.target.value)}
                  subLabel="The name of the database where Komiser will insert/save the data"
                  placeholder="my_database"
                />
                <LabelledInput
                  type="text"
                  id="username"
                  label="Username"
                  required
                  value={username}
                  onChange={e => setUsername(e.target.value)}
                  subLabel="The Postgres username"
                  placeholder="user"
                />
                <LabelledInput
                  type="password"
                  id="password"
                  label="Password"
                  required
                  value={password}
                  onChange={e => setPassword(e.target.value)}
                  subLabel="The Postgres password"
                  placeholder="Example0000*"
                />
              </div>
            </div>
            <CredentialsButton nextLabel="Add database" />
          </form>
        </LeftSideLayout>

        <RightSideLayout>
          <DatabasePurplin database={databaseProvider} />
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
