// ============================================================================
// NOTE: This is a Brigade 1.x script for now.
// ============================================================================

const { events, Job } = require("brigadier");
const { Check } = require("@brigadecore/brigade-utils");

const releaseTagRegex = /^refs\/tags\/(v[0-9]+(?:\.[0-9]+)*(?:\-.+)?)$/;

const goImg = "brigadecore/go-tools:v0.1.0";
const jsImg = "node:12.3.1-stretch";
const kanikoImg = "brigadecore/kaniko:v0.2.0";
const helmImg = "brigadecore/helm-tools:v0.1.0";
const localPath = "/workspaces/brigade";

class MakeTargetJob extends Job {
  constructor(target, img, e, env) {
    super(target, img);
    this.mountPath = localPath;
    this. env = env || {};
    this.env["SKIP_DOCKER"] = "true";
    const matchStr = e.revision.ref.match(releaseTagRegex);
    if (matchStr) {
      this.env["VERSION"] = Array.from(matchStr)[1];
    }
    this.tasks = [
      `cd ${localPath}`,
      `make ${target}`
    ];
  }
}

class PushImageJob extends MakeTargetJob {
  constructor(target, e, p) {
    super(target, kanikoImg, e, {
      "DOCKER_REGISTRY": p.secrets.dockerhubRegistry || "docker.io",
      "DOCKER_ORG": p.secrets.dockerhubOrg,
      "DOCKER_USERNAME": p.secrets.dockerhubUsername,
      "DOCKER_PASSWORD": p.secrets.dockerhubPassword
    });
  }
}

// A map of all jobs
const jobs = {};

// Basic tests:

const testUnitGoJobName = "test-unit-go";
const testUnitGoJob = (e, p) => {
  return new MakeTargetJob(testUnitGoJobName, goImg, e);
}
jobs[testUnitGoJobName] = testUnitGoJob;

const lintGoJobName = "lint-go";
const lintGoJob = (e, p) => {
  return new MakeTargetJob(lintGoJobName, goImg, e);
}
jobs[lintGoJobName] = lintGoJob;

const testUnitJSJobName = "test-unit-js";
const testUnitJSJob = (e, p) => {
  return new MakeTargetJob(testUnitJSJobName, jsImg, e);
}
jobs[testUnitJSJobName] = testUnitJSJob;

const lintJSJobName = "lint-js";
const lintJSJob = (e, p) => {
  return new MakeTargetJob(lintJSJobName, jsImg, e);
}
jobs[lintJSJobName] = lintJSJob;

// Brigadier:

const buildBrigadierJobName = "build-brigadier";
const buildBrigadierJob = (e, p) => {
  return new MakeTargetJob(buildBrigadierJobName, jsImg, e);
}
jobs[buildBrigadierJobName] = buildBrigadierJob;

const publishBrigadierJobName = "publish-brigadier";
const publishBrigadierJob = (e, p) => {
  return new MakeTargetJob(publishBrigadierJobName, jsImg, e, {
    "NPM_ORG": p.secrets.npmOrg,
    "NPM_USERNAME": p.secrets.npmUsername,
    "NPM_EMAIL": p.secrets.npmEmail,
    "NPM_PASSWORD": p.secrets.npmPassword
  });
}
jobs[publishBrigadierJobName] = publishBrigadierJob;

// Docker images:

const buildAPIServerJobName = "build-apiserver";
const buildAPIServerJob = (e, p) => {
  return new MakeTargetJob(buildAPIServerJobName, kanikoImg, e);
}
jobs[buildAPIServerJobName] = buildAPIServerJob;

const pushAPIServerJobName = "push-apiserver";
const pushAPIServerJob = (e, p) => {
  return new PushImageJob(pushAPIServerJobName, e, p);
}
jobs[pushAPIServerJobName] = pushAPIServerJob;

const buildGitInitializerJobName = "build-git-initializer";
const buildGitInitializerJob = (e, p) => {
  return new MakeTargetJob(buildGitInitializerJobName, kanikoImg, e);
}
jobs[buildGitInitializerJobName] = buildGitInitializerJob;

const pushGitInitializerJobName = "push-git-initializer";
const pushGitInitializerJob = (e, p) => {
  return new PushImageJob(pushGitInitializerJobName, e, p);
}
jobs[pushGitInitializerJobName] = pushGitInitializerJob;

const buildLoggerLinuxJobName = "build-logger-linux";
const buildLoggerLinuxJob = (e, p) => {
  return new MakeTargetJob(buildLoggerLinuxJobName, kanikoImg, e);
}
jobs[buildLoggerLinuxJobName] = buildLoggerLinuxJob;

const pushLoggerLinuxJobName = "push-logger-linux";
const pushLoggerLinuxJob = (e, p) => {
  return new PushImageJob(pushLoggerLinuxJobName, e, p);
}
jobs[pushLoggerLinuxJobName] = pushLoggerLinuxJob;

const buildObserverJobName = "build-observer";
const buildObserverJob = (e, p) => {
  return new MakeTargetJob(buildObserverJobName, kanikoImg, e);
}
jobs[buildObserverJobName] = buildObserverJob;

const pushObserverJobName = "push-observer";
const pushObserverJob = (e, p) => {
  return new PushImageJob(pushObserverJobName, e, p);
}
jobs[pushObserverJobName] = pushObserverJob;

const buildSchedulerJobName = "build-scheduler";
const buildSchedulerJob = (e, p) => {
  return new MakeTargetJob(buildSchedulerJobName, kanikoImg, e);
}
jobs[buildSchedulerJobName] = buildSchedulerJob;

const pushSchedulerJobName = "build-scheduler";
const pushSchedulerJob = (e, p) => {
  return new PushImageJob(pushSchedulerJobName, e, p);
}
jobs[pushSchedulerJobName] = pushSchedulerJob;

const buildWorkerJobName = "build-worker";
const buildWorkerJob = (e, p) => {
  return new MakeTargetJob(buildWorkerJobName, kanikoImg, e);
}
jobs[buildWorkerJobName] = buildWorkerJob;

const pushWorkerJobName = "push-worker";
const pushWorkerJob = (e, p) => {
  return new PushImageJob(pushWorkerJobName, e, p);
}
jobs[pushWorkerJobName] = pushWorkerJob;

// Helm chart:

const lintChartJobName = "lint-chart";
const lintChartJob = (e, p) => {
  return new MakeTargetJob(lintChartJobName, helmImg, e);
}
jobs[lintChartJobName] = lintChartJob;

const publishChartJobName = "publish-chart";
const publishChartJob = (e, p) => {
  return new MakeTargetJob(publishChartJobName, helmImg, e, {
    "HELM_REGISTRY": p.secrets.helmRegistry || "ghcr.io",
    "HELM_ORG": p.secrets.helmOrg,
    "HELM_USERNAME": p.secrets.helmUsername,
    "HELM_PASSWORD": p.secrets.helmPassword
  });
}
jobs[publishChartJobName] = publishChartJob;

// CLI:

const buildCLIJobName = "build-cli";
const buildCLIJob = (e, p) => {
  return new MakeTargetJob(buildCLIJobName, goImg, e);
}
jobs[buildCLIJobName] = buildCLIJob;

const publishCLIJobName = "publish-cli";
const publishCLIJob = (e, p) => {
  return new MakeTargetJob(publishCLIJobName, goImg, e, {
    "GITHUB_ORG": p.secrets.githubOrg,
    "GITHUB_REPO": p.secrets.githubRepo,
    "GITHUB_TOKEN": p.secrets.githubToken
  });
}
jobs[publishCLIJobName] = publishCLIJob;

// Run the entire suite of tests, builds, etc. concurrently WITHOUT publishing
// anything initially. If EVERYTHING passes AND this was a push (merge,
// presumably) to the v2 branch, then run concurrent jobs to publish "edge"
// images.
function runSuite(e, p) {
  // Important: To prevent Promise.all() from failing fast, we catch and
  // return all errors. This ensures Promise.all() always resolves. We then
  // iterate over all resolved values looking for errors. If we find one, we
  // throw it so the whole build will fail.
  //
  // Ref: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all#Promise.all_fail-fast_behaviour
  return Promise.all([
    // Basic tests:
    run(e, p, testUnitGoJob(e, p)).catch((err) => { return err }),
    run(e, p, lintGoJob(e, p)).catch((err) => { return err }),
    run(e, p, testUnitJSJob(e, p)).catch((err) => { return err }),
    run(e, p, lintJSJob(e, p)).catch((err) => { return err }),
    // Brigadier:
    run(e, p, buildBrigadierJob(e, p)).catch((err) => { return err }),
    // Docker images:
    run(e, p, buildAPIServerJob(e, p)).catch((err) => { return err }),
    run(e, p, buildGitInitializerJob(e, p)).catch((err) => { return err }),
    run(e, p, buildLoggerLinuxJob(e, p)).catch((err) => { return err }),
    run(e, p, buildObserverJob(e, p)).catch((err) => { return err }),
    run(e, p, buildSchedulerJob(e, p)).catch((err) => { return err }),
    run(e, p, buildWorkerJob(e, p)).catch((err) => { return err }),
    // Helm chart:
    run(e, p, lintChartJob(e, p)).catch((err) => { return err }),
    // CLI:
    run(e, p, buildCLIJob(e, p)).catch((err) => { return err })
  ]).then((values) => {
    values.forEach((value) => {
      if (value instanceof Error) throw value;
    });
  }).then(() => {
    if (e.revision.ref == "v2") {
      // Push "edge" images. For now we're not publishing an "edge" Brigadier or
      // CLI.
      Promise.all([
        run(e, p, pushAPIServerJob(e, p)).catch((err) => { return err }),
        run(e, p, pushGitInitializerJob(e, p)).catch((err) => { return err }),
        run(e, p, pushLoggerLinuxJob(e, p)).catch((err) => { return err }),
        run(e, p, pushObserverJob(e, p)).catch((err) => { return err }),
        run(e, p, pushSchedulerJob(e, p)).catch((err) => { return err }),
        run(e, p, pushWorkerJob(e, p)).catch((err) => { return err })
      ]).then((values) => {
        values.forEach((value) => {
          if (value instanceof Error) throw value;
        }); 
      })
    }
  });
}

// run a specific job identified by the payload.
function runCheck(e, p) {
  const jobName = JSON.parse(e.payload).body.check_run.name;
  const job = jobs[jobName];
  if (job) {
    return run(e, p, job(e, p));
  }
  throw new Error(`No job found with name: ${jobName}`);
}

// run the specified job, sandwiched between two other jobs to report status
// via the GitHub checks API.
function run(e, p, job) {
  console.log("Check requested");
  var check = new Check(e, p, job, `https://brigadecore.github.io/kashti/builds/${e.buildID}`);
  return check.run();
}

// Either of these events should initiate execution of the entire test suite.
events.on("check_suite:requested", runSuite);
events.on("check_suite:rerequested", runSuite);

// This event indicates a specific job is to be re-run.
events.on("check_run:rerequested", runCheck);

// These events MAY indicate that a maintainer has expressed, via a comment,
// that the entire test suite should be run.
events.on("issue_comment:created", (e, p) => Check.handleIssueComment(e, p, runSuite));
events.on("issue_comment:edited", (e, p) => Check.handleIssueComment(e, p, runSuite));

// Pushing new commits to any branch in github triggers a check suite. Such
// events are already handled above. Here we're only concerned with the case
// wherein a new TAG has been pushed-- and even then, we're only concerned with
// tags that look like a semantic version and indicate a formal release should
// be performed.
events.on("push", (e, p) => {
  const matchStr = e.revision.ref.match(releaseTagRegex);
  if (matchStr) {
    // This is an official release with a semantically versioned tag
    return Group.runAll([
      pushAPIServerJob(e, p),
      pushGitInitializerJob(e, p),
      pushLoggerLinuxJob(e, p),
      pushObserverJob(e, p),
      pushSchedulerJob(e, p),
      pushWorkerJob(e, p)
    ])
    .then(() => {
      // All images built and published successfully, so build and publish the
      // rest...
      Group.runAll([
        publishBrigadierJob(e, p),
        publishChartJob(e, p),
        publishCLIJob(e, p)
      ]);
    });
  }
  console.log(`Ref ${e.revision.ref} does not match release tag regex (${releaseTagRegex}); not releasing.`);
});
