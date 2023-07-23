import { useRouter } from 'next/router';
import { useContext, useEffect } from 'react';

import GlobalAppContext from '../components/layout/context/GlobalAppContext';

function Home() {
  const { betaFlagOnboardingWizard } = useContext(GlobalAppContext);

  const router = useRouter();

  useEffect(() => {
    if (betaFlagOnboardingWizard) {
      router.push('/onboarding/choose-cloud');
    } else {
      router.push('/dashboard');
    }
  }, []);
}

export default Home;
