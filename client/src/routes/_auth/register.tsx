import { MultiStepRegistrationForm } from '@/components/auth/multi-step-registration-form'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_auth/register')({
  component: RouteComponent,
})

function RouteComponent() {
  return <>
    <MultiStepRegistrationForm className="mx-auto max-w-sm" />
  </>
}
