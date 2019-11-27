// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Okta.Network
{
    /// <summary>
    /// Creates an Okta Network Zone.
    /// 
    /// This resource allows you to create and configure an Okta Network Zone.
    /// 
    /// &gt; This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/network_zone.html.markdown.
    /// </summary>
    public partial class Zone : Pulumi.CustomResource
    {
        /// <summary>
        /// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
        /// </summary>
        [Output("dynamicLocations")]
        public Output<ImmutableArray<string>> DynamicLocations { get; private set; } = null!;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        [Output("gateways")]
        public Output<ImmutableArray<string>> Gateways { get; private set; } = null!;

        /// <summary>
        /// Name of the Network Zone Resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        [Output("proxies")]
        public Output<ImmutableArray<string>> Proxies { get; private set; } = null!;

        /// <summary>
        /// Type of the Network Zone - can either be IP or DYNAMIC only.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("okta:network/zone:Zone", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("okta:network/zone:Zone", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : Pulumi.ResourceArgs
    {
        [Input("dynamicLocations")]
        private InputList<string>? _dynamicLocations;

        /// <summary>
        /// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
        /// </summary>
        public InputList<string> DynamicLocations
        {
            get => _dynamicLocations ?? (_dynamicLocations = new InputList<string>());
            set => _dynamicLocations = value;
        }

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        /// <summary>
        /// Name of the Network Zone Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("proxies")]
        private InputList<string>? _proxies;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        public InputList<string> Proxies
        {
            get => _proxies ?? (_proxies = new InputList<string>());
            set => _proxies = value;
        }

        /// <summary>
        /// Type of the Network Zone - can either be IP or DYNAMIC only.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ZoneArgs()
        {
        }
    }

    public sealed class ZoneState : Pulumi.ResourceArgs
    {
        [Input("dynamicLocations")]
        private InputList<string>? _dynamicLocations;

        /// <summary>
        /// Array of locations ISO-3166-1(2). Format code: countryCode OR countryCode-regionCode.
        /// </summary>
        public InputList<string> DynamicLocations
        {
            get => _dynamicLocations ?? (_dynamicLocations = new InputList<string>());
            set => _dynamicLocations = value;
        }

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        /// <summary>
        /// Name of the Network Zone Resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("proxies")]
        private InputList<string>? _proxies;

        /// <summary>
        /// Array of values in CIDR/range form.
        /// </summary>
        public InputList<string> Proxies
        {
            get => _proxies ?? (_proxies = new InputList<string>());
            set => _proxies = value;
        }

        /// <summary>
        /// Type of the Network Zone - can either be IP or DYNAMIC only.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ZoneState()
        {
        }
    }
}
