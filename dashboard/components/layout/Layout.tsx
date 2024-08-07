import * as Sentry from '@sentry/react';
import classNames from 'classnames';
import { useRouter } from 'next/router';
import { ReactNode, useEffect } from 'react';
import settingsService from '@services/settingsService';
import { ToastProvider } from '@components/toast/ToastProvider';
import GithubBanner from './components/github-banner/GithubBanner';
import environment from '../../environments/environment';
import useGithubStarBanner from './hooks/useGithubStarBanner';
import Button from '../button/Button';
import EmptyState from '../empty-state/EmptyState';
import ErrorState from '../error-state/ErrorState';
import Navbar from '../navbar/Navbar';
import GlobalAppContext from './context/GlobalAppContext';
import useGlobalStats from './hooks/useGlobalStats';
import useTelemetry from './hooks/useTelemetry';

type LayoutProps = {
  children: ReactNode;
};

function Layout({ children }: LayoutProps) {
  const { displayBanner, dismissBanner, githubStars } = useGithubStarBanner();
  const { loading, data, error, hasNoAccounts, fetch } = useGlobalStats();
  const { telemetry } = useTelemetry();
  const router = useRouter();
  const canRender = !error && !hasNoAccounts;

  useEffect(() => {
    settingsService.getOnboardingStatus().then(res => {
      if (
        res.onboarded === true &&
        res.status === 'COMPLETE' &&
        router.asPath.includes('/onboarding/')
      ) {
        router.replace('/dashboard/');
      } else if (res.onboarded === false && res.status === 'PENDING_DATABASE') {
        router.replace('/onboarding/choose-database');
      } else if (res.onboarded === false && res.status === 'PENDING_ACCOUNTS') {
        router.replace('/onboarding/choose-cloud');
      }
    });
  }, []);

  useEffect(() => {
    if (telemetry?.telemetry_enabled && environment.production) {
      Sentry.init({
        dsn: environment.SENTRY_URL,
        integrations: [Sentry.browserTracingIntegration()],

        // We recommend adjusting this value in production, or using tracesSampler
        // for finer control
        tracesSampleRate: 1.0
      });
    }
  }, [telemetry]);

  const betaFlagOnboardingWizard = true; // set this to true once wizard gets good support of the backend
  const isOnboarding =
    betaFlagOnboardingWizard && router.pathname.startsWith('/onboarding');

  return (
    <GlobalAppContext.Provider
      value={{
        displayBanner,
        dismissBanner,
        loading,
        data,
        error,
        hasNoAccounts,
        fetch,
        betaFlagOnboardingWizard
      }}
    >
      <ToastProvider>
        {isOnboarding && <>{children}</>}

        {!isOnboarding && (
          <>
            <GithubBanner githubStars={githubStars} />
            <Navbar />
            <main
              className={classNames(
                'relative bg-gray-50 p-6 pb-12 xl:px-8 2xl:px-24',
                displayBanner
                  ? 'mt-[145px] min-h-[calc(100vh-145px)]'
                  : 'mt-[73px] min-h-[calc(100vh-73px)]'
              )}
            >
              {canRender && children}

              {hasNoAccounts && betaFlagOnboardingWizard && !isOnboarding && (
                <EmptyState
                  title="We could not find a cloud account"
                  message="Get Started Onboarding"
                  action={() => {
                    router.push('/onboarding/choose-database');
                  }}
                  actionLabel="Begin Onboarding"
                  secondaryAction={() => {
                    router.push(
                      'https://github.com/tailwarden/komiser/issues/new/choose'
                    );
                  }}
                  secondaryActionLabel="Report an issue"
                  mascotPose="greetings"
                />
              )}

              {/* This block would be removed when onboarding Wizard is stable leaving the block above */}
              {hasNoAccounts && !betaFlagOnboardingWizard && (
                <EmptyState
                  title="We could not find a cloud account"
                  message="It seems you have not connected a cloud account to Komiser. Connect one right now so you can start managing it from your dashboard"
                  action={() => {
                    router.push(
                      'https://docs.komiser.io/docs/introduction/getting-started?utm_source=komiser&utm_medium=referral&utm_campaign=static'
                    );
                  }}
                  actionLabel="Guide to connect account"
                  secondaryAction={() => {
                    router.push(
                      'https://github.com/tailwarden/komiser/issues/new/choose'
                    );
                  }}
                  secondaryActionLabel="Report an issue"
                  mascotPose="thinking"
                />
              )}
              {/* This block would be removed when onboarding Wizard is stable leaving the block above */}

              {error && (
                <ErrorState
                  title="Network request error"
                  message="There was an error fetching the cloud accounts. Please refer to the logs for more info and try again."
                  action={
                    <Button
                      size="lg"
                      style="secondary"
                      onClick={() => router.reload()}
                    >
                      Refresh the page
                    </Button>
                  }
                />
              )}
            </main>
          </>
        )}
      </ToastProvider>
    </GlobalAppContext.Provider>
  );
}

export default Layout;
