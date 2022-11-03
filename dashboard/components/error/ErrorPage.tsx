import { ReactNode } from 'react';

type ErrorPageProps = {
  title: string;
  message: string;
  action?: ReactNode;
};

function ErrorPage({ title, message, action }: ErrorPageProps) {
  return (
    <div className="flex h-[calc(100vh-156px)] items-center justify-center">
      <div className="flex items-center justify-center text-center">
        <div className="flex flex-col items-center justify-center max-w-sm bg-white p-12 rounded-lg gap-6">
          <picture>
            <img
              src="/assets/img/branding/purplin/serious.png"
              className="w-48"
              alt="Purplin"
            />
          </picture>
          <p className="font-medium text-black-900">
            {title}
          </p>
          <p className="text-sm text-black-300 text-center">
            {message}
          </p>
          {action && (
            <>
              <div className="mt-8"></div>
              {action}
            </>
          )}
        </div>
      </div>
    </div>
  );
}

export default ErrorPage;
