import { Suspense } from 'react';
import { VerifyEmailForm } from '@/components/auth/VerifyEmailForm';

export default function VerifyEmailPage() {
  return (
    <Suspense fallback={<div className="text-center text-xs text-slate-400 py-12">Loading Verification...</div>}>
      <VerifyEmailForm />
    </Suspense>
  );
}
