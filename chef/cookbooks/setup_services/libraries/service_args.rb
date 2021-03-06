# Copyright 2022 MobiledgeX, Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

class Chef
  class Recipe
    def extract_cmd(service, argsmap, joincmd, skipcmd)
      args = if skipcmd
               []
             else
               [service]
             end
      argsmap.each_key do |x|
        next unless node[service]['args'].key?(x)
        args += ["--#{x}"]
        unless node[service]['args'][x].empty?
          if joincmd
            args[-1] = args[-1] + "=#{node[service]['args'][x]}"
          else
            args += ["'" + "#{node[service]['args'][x]}" + "'"]
          end
        end
      end
      args.join(' ')
    end

    def crmserver_cmd
      # Hash of:
      #   Key = arg name
      #   Value = arg type (false means flag type)
      # NOTE:
      #   Please update chef/cookbooks/setup_services/attributes/default.rb, if argsmap is updated below
      argsmap = {
        'cloudletKey' => true,
        'notifyAddrs' => true,
        'notifySrvAddr' => true,
        'tls' => true,
        'platform' => true,
        'vaultAddr' => true,
        'physicalName' =>  true,
        'region' => true,
        'span' => true,
        'd' => true,
        'cloudletVMImagePath' => true,
        'vmImageVersion' => true,
        'containerVersion' => true,
        'commercialCerts' => true,
        'useVaultPki' => false,
        'chefServerPath' => true,
        'deploymentTag' => true,
        'upgrade' => false,
        'accessKeyFile' => true,
        'accessApiAddr' => true,
        'cacheDir' => true,
        'redisStandaloneAddr' => true,
        'appDNSRoot' => true,
      }
      extract_cmd('crmserver', argsmap, false, false)
    end

    def shepherd_cmd
      # Hash of:
      #   Key = arg name
      #   Value = arg type (false means flag type)
      # NOTE:
      #   Please update chef/cookbooks/setup_services/attributes/default.rb, if argsmap is updated below
      argsmap = {
        'cloudletKey' => true,
        'notifyAddrs' => true,
        'tls' => true,
        'platform' => true,
        'vaultAddr' => true,
        'physicalName' =>  true,
        'region' => true,
        'span' => true,
        'd' => true,
        'useVaultPki' => false,
        'chefServerPath' => true,
        'deploymentTag' => true,
        'accessKeyFile' => true,
        'accessApiAddr' => true,
        'thanosRecvAddr' => true,
        'appDNSRoot' => true,
      }
      extract_cmd('shepherd', argsmap, false, false)
    end

    def cloudlet_prometheus_cmd
      # Hash of:
      #   Key = arg name
      #   Value = arg type (false means flag type)
      argsmap = {
        'config.file' => true,
        'web.listen-address' => true,
        'web.enable-lifecycle' => false,
      }
      extract_cmd('cloudletPrometheus', argsmap, true, true)
    end

    def get_crm_args(harole)
      crmargs = crmserver_cmd.split(' ')
      crmargs.shift()
      if harole != ''
        crmargs.append('--HARole')
        crmargs.append(harole)
      end
      crmargs
    end

    def get_shep_args(harole)
      shepargs = shepherd_cmd.split(' ')
      shepargs.shift()
      if harole != ''
        shepargs.append('--HARole')
        shepargs.append(harole)
      end
      shepargs
    end

    def get_prom_args
      cloudlet_prometheus_cmd.split(' ')
    end

    def get_services_vars(harole)
      {
        crmserver: { cmd: 'crmserver',
                     cmdargs: get_crm_args(harole),
                     env: node['crmserver']['env'],
                     image: node['edgeCloudImage'] + ':' + node['edgeCloudVersion'],
                     volumeMounts: { accesskey_vol: { name: 'accesskey-vol', mountPath: '/root/accesskey' },
                                     cache_vol: { name: 'cache-vol', mountPath: '/root/crm_cache' } },
        },
        shepherd: { cmd: 'shepherd',
                    cmdargs: get_shep_args(harole),
                    env: node['shepherd']['env'],
                    image: node['edgeCloudImage'] + ':' + node['edgeCloudVersion'],
                    volumeMounts: { accesskey_vol: { name: 'accesskey-vol', mountPath: '/root/accesskey' },
                                    prom_cfg_vol: { name: 'prom-cfg-vol', mountPath: '/etc/prometheus' },
                                    prom_tmp_vol: { name: 'prom-tmp-vol', mountPath: '/tmp' } },
        },
        cloudletprometheus: { cmdargs: get_prom_args,
                              env: node['cloudletPrometheus']['env'],
                              image: 'docker.mobiledgex.net/mobiledgex/mobiledgex_public/' + node['prometheusImage'] + ':' + node['prometheusVersion'],
                              volumeMounts: { prom_cfg_vol: { name: 'prom-cfg-vol', mountPath: '/etc/prometheus' },
                                              prom_tmp_vol: { name: 'prom-tmp-vol', mountPath: '/tmp' } },
        },
      }
    end

    def get_hostvols_vars
      { accesskey_vol: { name: 'accesskey-vol', hostPath: '/root/accesskey' },
        cache_vol:     { name: 'cache-vol',     hostPath: '/root/crm_cache' },
        prom_cfg_vol:     { name: 'prom-cfg-vol',     hostPath: '/home/ubuntu/prometheus-vols/cfg' },
        prom_tmp_vol:     { name: 'prom-tmp-vol',     hostPath: '/home/ubuntu/prometheus-vols/tmp' } }
    end

    def get_configmap_vars
      {}
    end

    def get_thanos_remote_write_addr
      region = node.normal['crmserver']['args']['region'].downcase
      deploymentTag = node.normal['crmserver']['args']['deploymentTag'].downcase
      if deploymentTag == "main"
        cmd = region + ".thanos-recv.mobiledgex.net"
      else
        cmd = region + "-" + deploymentTag + ".thanos-recv.mobiledgex.net"
      end
      cmd
    end
  end
end
