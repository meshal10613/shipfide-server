import { Suspense } from 'react';
import { ResetPasswordForm } from '@/components/auth/ResetPasswordForm';

export default function ResetPasswordPage() {
  return (
    <Suspense fallback={<div className="text-center text-xs text-slate-400 py-12">Loading Reset Form...</div>}>
      <ResetPasswordForm />
    </Suspense>
  );
}
