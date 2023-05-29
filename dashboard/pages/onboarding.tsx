import Head from 'next/head';

import OnboardingWizardLayout from '../components/onboarding-wizard/OnboardingWizardLayout';
import DashboardTopStats from '../components/dashboard/components/top-stats/DashboardTopStats';
import DashboardCostExplorer from '../components/dashboard/components/cost-explorer/DashboardCostExplorer';

export default function Onboarding() {
  return (
    <div>
      <Head>
        <title>Onboarding - Komiser</title>
        <meta name="description" content="Onboarding - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <DashboardTopStats />
        <DashboardCostExplorer />
      </OnboardingWizardLayout>
    </div>
  );
}
