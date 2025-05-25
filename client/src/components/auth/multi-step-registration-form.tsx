import type { APIResponse } from '@/api/api'
import { useRegister } from '@/api/auth/auth.query'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { cn } from '@/lib/utils'
import { APP_NAME, APP_VERSION } from '@/utils/config'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link, useNavigate } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import * as z from 'zod'
import Logo from '../logo'

// Add this at the top level of your component
type FieldErrors = {
  email?: string;
  password?: string;
  confirm_password?: string;
  full_name?: string;
  date_of_birth?: string;
  general?: string; // For general errors not tied to a specific field
};

// Define validation schemas for each step
const step1Schema = z.object({
  full_name: z.string().min(2, "Full name must be at least 2 characters").max(255, "Full name must be less than 255 characters"),
})

const step2Schema = z.object({
  email: z.string().email("Please enter a valid email address").max(255, "Email must be less than 255 characters"),
  date_of_birth: z.string().regex(/^\d{4}-\d{2}-\d{2}$/, "Date must be in YYYY-MM-DD format")
})

const step3Schema = z.object({
  password: z.string().min(8, "Password must be at least 8 characters").max(72, "Password must be less than 72 characters"),
  confirm_password: z.string()
})

// Combine all schemas for final submission
const formSchema = step1Schema.merge(step2Schema).merge(step3Schema).superRefine((data, ctx) => {
    if (data.password !== data.confirm_password) {
        ctx.addIssue({
        code: z.ZodIssueCode.custom,
        message: "Passwords don't match",
        path: ["confirm_password"],
      });
    }
  });



type FormData = z.infer<typeof formSchema>

interface StepFormProps {
  formData: Partial<FormData>;
  onSubmit: (data: Partial<FormData>) => Promise<void>;
  handleBack: () => void;
  isSubmitting: boolean;
  step: number;
  fieldErrors: FieldErrors;
  clearFieldError: (field: keyof FieldErrors) => void;
}

// Step 1: Personal Information
const Step1Form = ({
  formData,
  onSubmit,
  isSubmitting,
  fieldErrors,
  clearFieldError
}: StepFormProps) => {
  const form = useForm<z.infer<typeof step1Schema>>({
    resolver: zodResolver(step1Schema),
    mode: 'onBlur',
    defaultValues: formData
  })

  // Set server errors in the form
  useEffect(() => {
    if (fieldErrors.full_name) {
      form.setError('full_name', { message: fieldErrors.full_name });
    }
  }, [fieldErrors, form]);

  // Check if there are any errors for the current step's fields
  const hasStepErrors = Boolean(fieldErrors.full_name);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-6">
        <FormField
          control={form.control}
          name="full_name"
          render={({ field }) => (
            <FormItem className="grid gap-3">
              <FormLabel>Full Name</FormLabel>
              <FormControl>
                <Input
                    
                  {...field}
                  placeholder="Enter your full name"
                  onChange={(e) => {
                    field.onChange(e);
                    clearFieldError('full_name');
                  }}
                />
              </FormControl>
              <FormDescription>
                Your full name will be displayed on your profile
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex justify-end">
          <Button type="submit" disabled={isSubmitting || hasStepErrors}>
            Next
          </Button>
        </div>
      </form>
    </Form>
  )
}

// Step 2: Account Information
const Step2Form = ({ formData, onSubmit, handleBack, isSubmitting, fieldErrors, clearFieldError }: StepFormProps) => {
  const form = useForm<z.infer<typeof step2Schema>>({
    resolver: zodResolver(step2Schema),
    mode: 'onBlur',
    defaultValues: formData
  })

  // Set server errors in the form
  useEffect(() => {
    if (fieldErrors.email) {
      form.setError('email', { message: fieldErrors.email });
    }
    if (fieldErrors.date_of_birth) {
      form.setError('date_of_birth', { message: fieldErrors.date_of_birth });
    }
  }, [fieldErrors, form]);

  // Check if there are any errors for the current step's fields
  const hasStepErrors = Boolean(
    fieldErrors.email ||
    fieldErrors.date_of_birth
  );

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-6">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem className="grid gap-3">
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input 
                  type="email" 
                  {...field} 
                  placeholder="m@example.com" 
                  onChange={(e) => {
                    field.onChange(e);
                    clearFieldError('email');
                  }} 
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="date_of_birth"
          render={({ field }) => (
            <FormItem className="grid gap-3">
              <FormLabel>Date of Birth</FormLabel>
              <FormControl>
                <Input
                  type="date"
                  {...field}
                  onChange={(e) => {
                    // Convert from date input to YYYY-MM-DD
                    field.onChange(e.target.value);
                    clearFieldError('date_of_birth');
                  }}
                />
              </FormControl>
              <FormDescription>
                Your date of birth is required (Format: YYYY-MM-DD)
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex justify-between">
          <Button type="button" variant="outline" onClick={handleBack}>
            Back
          </Button>
          <Button type="submit" disabled={isSubmitting || hasStepErrors}>
            Next
          </Button>
        </div>
      </form>
    </Form>
  )
}

// Step 3: Password
const Step3Form = ({ formData, onSubmit, handleBack, isSubmitting, fieldErrors, clearFieldError }: StepFormProps) => {
  const form = useForm<z.infer<typeof step3Schema>>({
    resolver: zodResolver(step3Schema),
    mode: 'onBlur',
    defaultValues: formData
  })

  // Set server errors in the form
  useEffect(() => {
    if (fieldErrors.password) {
      form.setError('password', { message: fieldErrors.password });
    }
    if (fieldErrors.confirm_password) {
      form.setError('confirm_password', { message: fieldErrors.confirm_password });
    }
  }, [fieldErrors, form]);

  // Check if there are any errors for the current step's fields
  const hasStepErrors = Boolean(
    fieldErrors.password ||
    fieldErrors.confirm_password
  );

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-6">
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem className="grid gap-3">
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input 
                  type="password" 
                  {...field} 
                  placeholder="********" 
                  onChange={(e) => {
                    field.onChange(e);
                    clearFieldError('password');
                  }} 
                />
              </FormControl>
              <FormDescription>
                Password must be at least 8 characters
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="confirm_password"
          render={({ field }) => (
            <FormItem className="grid gap-3">
              <FormLabel>Confirm Password</FormLabel>
              <FormControl>
                <Input 
                  type="password" 
                  {...field} 
                  placeholder="********" 
                  onChange={(e) => {
                    field.onChange(e);
                    clearFieldError('confirm_password');
                  }} 
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex justify-between">
          <Button type="button" variant="outline" onClick={handleBack}>
            Back
          </Button>
          <Button type="submit" disabled={isSubmitting || hasStepErrors}>
            Register
          </Button>
        </div>
      </form>
    </Form>
  )
}

// The multi-step registration form component
export const MultiStepRegistrationForm = ({
  className,
  ...props
}: React.ComponentProps<"div">) => {
  const [step, setStep] = useState(0)
  const totalSteps = 3
  const [formData, setFormData] = useState<Partial<FormData>>({})
  const [isSubmitting, setIsSubmitting] = useState(false)
  // Replace the simple error state with a structured object
  const [fieldErrors, setFieldErrors] = useState<FieldErrors>({})
  const navigate = useNavigate()

  const { mutateAsync } = useRegister();

  // Clear specific field errors when navigating to a different step
  const clearErrorsForStep = (stepNumber: number) => {
    const fieldsToKeep: Record<number, (keyof FieldErrors)[]> = {
      0: ['full_name'],
      1: ['email', 'date_of_birth'],
      2: ['password', 'confirm_password'],
    };

    // Only keep errors relevant to the current step
    const newErrors: FieldErrors = {};
    fieldsToKeep[stepNumber]?.forEach(field => {
      if (fieldErrors[field]) {
        newErrors[field] = fieldErrors[field];
      }
    });

    // Always keep general errors
    if (fieldErrors.general) {
      newErrors.general = fieldErrors.general;
    }

    setFieldErrors(newErrors);
  };

  // Clear a specific field error when user edits that field
  const clearFieldError = (field: keyof FieldErrors) => {
    if (fieldErrors[field]) {
      setFieldErrors(prev => {
        const newErrors = { ...prev };
        delete newErrors[field];
        return newErrors;
      });
    }
  };

  const handleSubmit = async (data: Partial<FormData>) => {
    setIsSubmitting(true)
    try {
      // Clear any existing errors before submission
      setFieldErrors({});

      // Merge form data and update state
      const updatedData = { ...formData, ...data }
      setFormData(updatedData)

      // Proceed to next step or submit
      if (step < totalSteps - 1) {
        setStep(prev => {
          const newStep = prev + 1;
          clearErrorsForStep(newStep);
          return newStep;
        });
      } else {
        // Final submission - validate all required fields
        if (!updatedData.email || !updatedData.password || !updatedData.full_name || !updatedData.date_of_birth) {
          throw new Error("Missing required fields");
        }

        // Call the registration API
        const registrationData = {
          email: updatedData.email,
          password: updatedData.password,
          full_name: updatedData.full_name.trim(),
          date_of_birth: updatedData.date_of_birth.trim(),
        };

        await mutateAsync(registrationData);

        toast.success("Registration successful!", {
          description: "You can now log in to your account."
        });

        navigate({ to: "/login" });
      }
    } catch (error) {
      const err = error as APIResponse<null>;
      console.error("Registration error:", err);

      // Handle and categorize API errors
      if (err.error?.code) {
        // Handle field-specific errors
        if (err.error.code === 'validation_failed' && err.error.details) {
          // If API returns field-specific errors in the details object
          setFieldErrors(prev => ({ ...prev, ...err.error?.details }));

          // Navigate to the appropriate step based on the field with error
          const errorFields = Object.keys(err.error.details) as Array<keyof FieldErrors>;
          if (errorFields.some(field => field === 'email' || field === 'date_of_birth')) {
            setStep(1); // Account details step
          } else if (errorFields.some(field => field === 'password' || field === 'confirm_password')) {
            setStep(2); // Password step
          } else if (errorFields.some(field => field === 'full_name')) {
            setStep(0); // Personal info step
          }
        }
        // Handle specific error codes
        else if (err.error.code === 'duplicate_entry') {
          if (err.error.description?.includes("email") || err.error.description?.includes("Email")) {
            setFieldErrors(prev => ({ ...prev, email: "This email is already registered" }));
            setFormData(prev => ({ ...prev, email: "" }));
            setStep(1);
          } else {
            // Generic duplicate error
            setFieldErrors(prev => ({ ...prev, general: err.error?.description || "Duplicate entry detected" }));
          }
        }
        // General error message from backend
        else {
          const errorMessage = err.error.description || err.message || "An error occurred. Please try again.";
          setFieldErrors(prev => ({ ...prev, general: errorMessage }));
          toast.error(errorMessage);
        }
      }
      // Handle cases where we have a message but no structured error
      else if (err.message) {
        setFieldErrors(prev => ({ ...prev, general: err.message }));
        toast.error(err.message);
      }
      // Handle validation errors from zod
      else if (error instanceof z.ZodError) {
        // Handle validation errors
        const newErrors: FieldErrors = {};
        error.errors.forEach(err => {
          if (err.path) {
            const field = err.path[0] as keyof FieldErrors;
            newErrors[field] = err.message;
          }
        });
        setFieldErrors(newErrors);
        toast.error("Validation failed. Please check your inputs.");
      } else {
        // Generic error
        const errorMessage = error instanceof Error ? error.message : "An unexpected error occurred";
        setFieldErrors(prev => ({ ...prev, general: errorMessage }));
        toast.error("An unexpected error occurred. Please try again.");
      }
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleBack = () => {
    if (step > 0) {
      setStep(prev => {
        const newStep = prev - 1;
        clearErrorsForStep(newStep);
        return newStep;
      });
    }
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <div className="flex flex-col items-center gap-2">
        <Link
          to="/"
          className="flex flex-col items-center gap-2 font-medium"
        >
          <div className="flex size-8 items-center justify-center rounded-md">
            <Logo />
          </div>
          <span className="sr-only">{APP_NAME}.</span>
          <span className="text-sm text-gray-500">{APP_VERSION}</span>
        </Link>
        <h1 className="text-xl font-bold">Create Your Account</h1>
        <div className="text-center text-sm">
          Already have an account?{" "}
          <Link to="/login" className="underline underline-offset-4">
            Sign in
          </Link>
        </div>
      </div>

      {/* Step Indicators */}
      <div className="flex items-center justify-center gap-1">
        {Array.from({ length: totalSteps }).map((_, index) => (
          <div key={index} className="flex items-center">
            <div className={cn(
              "w-2 h-2 rounded-full transition-all",
              index === step ? "w-6 bg-primary" :
                index < step ? "bg-primary" : "bg-muted"
            )} />
          </div>
        ))}
      </div>

      {/* Error Message */}
      {fieldErrors.general && (
        <div className="border border-destructive/50 bg-destructive/10 text-destructive p-3 rounded">
          {fieldErrors.general}
        </div>
      )}

      {/* Render Current Step */}
      {step === 0 && (
        <Step1Form
          formData={formData}
          onSubmit={handleSubmit}
          handleBack={handleBack}
          isSubmitting={isSubmitting}
          step={step}
          fieldErrors={fieldErrors}
          clearFieldError={clearFieldError}
        />
      )}
      {step === 1 && (
        <Step2Form
          formData={formData}
          onSubmit={handleSubmit}
          handleBack={handleBack}
          isSubmitting={isSubmitting}
          step={step}
          fieldErrors={fieldErrors}
          clearFieldError={clearFieldError}
        />
      )}
      {step === 2 && (
        <Step3Form
          formData={formData}
          onSubmit={handleSubmit}
          handleBack={handleBack}
          isSubmitting={isSubmitting}
          step={step}
          fieldErrors={fieldErrors}
          clearFieldError={clearFieldError}
        />
      )}

      <div className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
        By clicking continue, you agree to our <Link to="/terms">Terms of Service</Link>{" "}
        and <Link to="/privacy">Privacy Policy</Link>.
      </div>
    </div>
  );
};
