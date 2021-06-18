Terraform Runner API
====================

An API to run and manage Terraform infrastructure.

Motivation
----------
There is a gap with Terraform in being able to run it within an automated fashion. Terraform doesn't work well within CI because of the complications of running plans and apply. You often don't want an entire infrastructure to "apply" on every commit. But you may want to run plan, and optionally apply if everything looks good. There are rare cases or only compact infrastructure that you'd want to apply on every commit.
The Terraform Runner API is designed to fill the gap where you can trigger an infrastructure plan or apply on demand as an API, this makes it possible to selectively apply infrastructure from Pull requests, automation, or webhooks

