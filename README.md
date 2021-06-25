Terraform Runner API [Experimental]
====================

Concept to explore an API Terraform Runner to plan and apply Terraform and Terragrunt infrastructure.

Concept and Motivation
----------------------
There is a gap with Terraform in being able to run it within an automated fashion. Terraform doesn't work well within CI because of the complications of running plans and apply. You often don't want an entire infrastructure to "apply" on every commit. But you may want to run plan, and optionally apply if everything looks good. There are rare cases or only compact infrastructure that you'd want to apply on every commit.
The Terraform Runner explores the concept of a Terraform Runner as an API designed to fill the gap where you can trigger an infrastructure plan or apply on demand as an API, this makes it possible to selectively apply infrastructure from Pull requests, automation, or webhooks

Configuration
-------------

It is ideal to mount these directories to the pod/container. There are no in-built modules.
* `TF_RUNNER_WORKDIR` Workspace containing the directory of terraform modules to run/apply. (Default: `/var/workspace`)
* `TF_PLUGIN_CACHE_DIR` Directory containing all the required providers by the module. (Default: `/var/lib/terraform/providers`)
* `TF_BINARY_PATH` Path to the Terraform binary. (Default: `/usr/local/bin/terraform`)

Goals
-----
1. Multi-tenant
    * How do we run multiple teams/workspaces from a single API?
2. Trigger from Webhook or cURL call
3. Authentication/Authorization built-in

### Stretch Goals:
1. Post to GHE pull-request the result of a plan

How it Works
------------

### Terraform-runner-api can be deployed as a pod:
1. tf-runner-api (container)
    * VolumeMount a workspace or...
    * VolumeMount a temp dir to house workspaces
2. Volume contains TF workspace
    * Could be achieved through git checkout and thus be triggered on commit
3. Post to GHE Pull-request the plan output

### Trigger Methods:
1. cURL call to API
2. webhook trigger (such as from GHE or Gantry)

### API Commands:
1. plan-all (terragrunt)
2. plan/module (terragrunt or terraform)
3. apply-all (terragrunt)
4. apply/module (terragrunt or terraform)
5. update (git pull)
    * ...

Limitations
-----------
* Terraform version installed on the runner-api
* Terragrunt version installed on the runner-api
