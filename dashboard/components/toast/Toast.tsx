import classNames from 'classnames';
import Button from '../button/Button';

import ErrorIcon from '../icons/ErrorIcon';
import CheckIcon from '../icons/CheckIcon';

export type ToastProps = {
  hasError: boolean;
  title: string;
  message: string;
};

type ToastProp = ToastProps & {
  dismissToast: () => void;
};

function Toast({ hasError, title, message, dismissToast }: ToastProp) {
  return (
    <>
      <div
        className={classNames(
          'fixed bottom-4 left-4 right-4 z-40 flex max-w-2xl animate-fade-in-up items-center justify-between overflow-hidden rounded-lg px-6 py-4 text-gray-950 opacity-0 shadow-right sm:left-8',
          {
            'bg-red-50': hasError,
            'bg-green-100': !hasError
          }
        )}
      >
        <div
          className={classNames(
            'absolute bottom-0 left-0 h-1 animate-width-to-fit',
            {
              'bg-red-50': hasError,
              'bg-green-500': !hasError
            }
          )}
        ></div>
        <div className="flex items-start">
          <div
            className={classNames({
              'mt-4 text-red-500 ': hasError,
              'mt-4 text-green-500': !hasError
            })}
          >
            {hasError ? (
              <ErrorIcon width={24} height={24} />
            ) : (
              <CheckIcon width={24} height={24} />
            )}
          </div>
          <div className="p-3">
            <p className="pb-2 text-sm font-medium">{title}</p>
            <p
              className="text-sm text-gray-950 "
              dangerouslySetInnerHTML={{ __html: message }}
            />
          </div>
        </div>
        <div className="w-12"></div>
        <Button style="text" onClick={dismissToast}>
          Dismiss
        </Button>
      </div>
    </>
  );
}

export default Toast;
