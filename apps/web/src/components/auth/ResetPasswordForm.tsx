'use client';

import { useState, useEffect } from 'react';
import { useResetPasswordMutation } from '@/store/api/authApi';
import { useRouter, useSearchParams } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { Lock, ArrowRight, KeyRound, AlertCircle } from 'lucide-react';
import { motion } from 'framer-motion';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export function ResetPasswordForm() {
  const searchParams = useSearchParams();
  const [email, setEmail] = useState('');
  const [otp, setOtp] = useState('');
  const [password, setPassword] = useState('');
  const [errorMsg, setErrorMsg] = useState('');
  const [resetPassword, { isLoading }] = useResetPasswordMutation();
  const router = useRouter();

  useEffect(() => {
    const emailParam = searchParams.get('email');
    if (emailParam) setEmail(emailParam);
  }, [searchParams]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg('');
    try {
      const res = await resetPassword({ email, otp, password }).unwrap();
      if (res.success) {
        router.push('/auth/login');
      }
    } catch (err: any) {
      setErrorMsg(err?.data?.message || 'Password reset failed. Check OTP and password.');
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
            <h1 className="text-2xl font-extrabold tracking-tight text-[#404042]">Reset Password</h1>
            <p className="text-xs text-slate-500 font-medium">Set your new account password</p>
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
            label="Email Address"
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="name@example.com"
          />

          <Input
            label="6-Digit OTP Code"
            type="text"
            required
            maxLength={6}
            value={otp}
            onChange={(e) => setOtp(e.target.value)}
            placeholder="123456"
            leftIcon={<KeyRound className="w-4 h-4" />}
          />

          <Input
            label="New Password"
            type="password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="NewPassword@123"
            leftIcon={<Lock className="w-4 h-4" />}
          />

          <Button
            type="submit"
            isLoading={isLoading}
            rightIcon={<ArrowRight className="w-4 h-4" />}
            className="w-full py-3.5 mt-2"
          >
            Update Password
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
