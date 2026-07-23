'use client';

import { useState } from 'react';
import { useForgotPasswordMutation } from '@/store/api/authApi';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { ArrowRight, Mail, AlertCircle } from 'lucide-react';
import { motion } from 'framer-motion';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export function ForgotPasswordForm() {
  const [email, setEmail] = useState('');
  const [errorMsg, setErrorMsg] = useState('');
  const [forgotPassword, { isLoading }] = useForgotPasswordMutation();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg('');
    try {
      const res = await forgotPassword({ email }).unwrap();
      if (res.success) {
        router.push(`/auth/reset-password?email=${encodeURIComponent(email)}`);
      }
    } catch (err: any) {
      setErrorMsg(err?.data?.message || 'Failed to dispatch password reset OTP.');
    }
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <motion.div
        initial={{ opacity: 0, scale: 0.95 }}
        animate={{ opacity: 1, scale: 1 }}
        className="w-full p-8 rounded-3xl bg-white border border-slate-200 shadow-xl space-y-6"
      >
        <div className="text-center space-y-3">
          <Link href="/" className="inline-block">
            <Image
              src="/logo.png"
              alt="Shipfide Logo"
              width={180}
              height={50}
              className="h-12 w-auto object-contain mx-auto"
              priority
            />
          </Link>
          <div>
            <h1 className="text-2xl font-extrabold tracking-tight text-[#404042]">Forgot Password?</h1>
            <p className="text-xs text-slate-500 font-medium">Enter your account email to receive a reset code</p>
          </div>
        </div>

        {errorMsg && (
          <div className="p-3 rounded-2xl bg-red-50 border border-red-200 text-[#DD0033] text-xs flex items-center gap-2">
            <AlertCircle className="w-4 h-4 shrink-0" />
            <span>{errorMsg}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <Input
            label="Account Email"
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="name@example.com"
            leftIcon={<Mail className="w-4 h-4" />}
          />

          <Button
            type="submit"
            isLoading={isLoading}
            rightIcon={<ArrowRight className="w-4 h-4" />}
            className="w-full py-3.5 mt-2"
          >
            Send Reset Code
          </Button>
        </form>

        <div className="text-center pt-2 border-t border-slate-100">
          <Link href="/auth/login" className="text-xs font-bold text-slate-500 hover:text-[#DD0033]">
            Back to Sign In
          </Link>
        </div>
      </motion.div>

      {/* Manual Bottom Text */}
      <p className="text-center text-[11px] text-slate-500 font-medium">
        Shipfide • Delivering Trust, Every Time
      </p>
    </div>
  );
}
