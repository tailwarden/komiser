import { ReactNode } from 'react';

export type ErrorStateProps = {
  title: string;
  message: string;
  action?: ReactNode;
};

function ErrorState({ title, message, action }: ErrorStateProps) {
  return (
    <div className="flex h-[calc(100vh-156px)] items-center justify-center">
      <div className="flex items-center justify-center text-center">
        <div className="flex max-w-sm flex-col items-center justify-center gap-6 rounded-lg bg-white p-12">
          <picture>
            <img
              src="/assets/img/purplin/fixing.svg"
              className="w-48"
              alt="Purplin"
            />
          </picture>
          <p className="font-medium text-gray-950">{title}</p>
          <p className="text-center text-sm text-gray-500">{message}</p>
          {action && <>{action}</>}
        </div>
      </div>
    </div>
  );
}

export default ErrorState;
