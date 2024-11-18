package cf

var CustomNodeList = `{
    "custom_nodes": [
        {
            "author": "Dr.Lt.Data",
            "title": "ComfyUI-Manager",
            "id": "manager",
            "reference": "https://github.com/ltdrdata/ComfyUI-Manager",
            "files": [
                "https://github.com/ltdrdata/ComfyUI-Manager"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Manager itself is also a custom node."
        },
        {
            "author": "Dr.Lt.Data",
            "title": "ComfyUI Impact Pack",
            "id": "impact",
            "reference": "https://github.com/ltdrdata/ComfyUI-Impact-Pack",
            "files": [
                "https://github.com/ltdrdata/ComfyUI-Impact-Pack"
            ],
            "pip": ["ultralytics"],
            "install_type": "git-clone",
            "description": "This extension offers various detector nodes and detailer nodes that allow you to configure a workflow that automatically enhances facial details. And provide iterative upscaler.",
            "preemptions":["SAMLoader"]
        },
        {
            "author": "Dr.Lt.Data",
            "title": "ComfyUI Inspire Pack",
            "id": "inspire",
            "reference": "https://github.com/ltdrdata/ComfyUI-Inspire-Pack",
            "nodename_pattern": "Inspire$",
            "files": [
                "https://github.com/ltdrdata/ComfyUI-Inspire-Pack"
            ],
            "install_type": "git-clone",
            "description": "This extension provides various nodes to support Lora Block Weight and the Impact Pack. Provides many easily applicable regional features and applications for Variation Seed."
        },
        {
            "author": "comfyanonymous",
            "title": "ComfyUI_experiments",
            "id": "comfy-exp",
            "reference": "https://github.com/comfyanonymous/ComfyUI_experiments",
            "files": [
                "https://github.com/comfyanonymous/ComfyUI_experiments"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ModelSamplerTonemapNoiseTest, TonemapNoiseWithRescaleCFG, ReferenceOnlySimple, RescaleClassifierFreeGuidanceTest, ModelMergeBlockNumber, ModelMergeSDXL, ModelMergeSDXLTransformers, ModelMergeSDXLDetailedTransformers."
        },
        {
            "author": "comfyanonymous",
            "title": "TensorRT Node for ComfyUI",
            "id": "tensorrt",
            "reference": "https://github.com/comfyanonymous/ComfyUI_TensorRT",
            "files": [
                "https://github.com/comfyanonymous/ComfyUI_TensorRT"
            ],
            "install_type": "git-clone",
            "description": "This node enables the best performance on NVIDIA RTX™ Graphics Cards  (GPUs) for Stable Diffusion by leveraging NVIDIA TensorRT."
        },
        {
            "author": "Stability-AI",
            "title": "Stability API nodes for ComfyUI",
            "id": "sai-api",
            "reference": "https://github.com/Stability-AI/ComfyUI-SAI_API",
            "files": [
                "https://github.com/Stability-AI/ComfyUI-SAI_API"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Stability SD3, Stability Outpainting, Stability Search and Replace, Stability Image Core, Stability Inpainting, Stability Remove Background, Stability Creative Upscale.\nAdd API key to environment variable 'SAI_API_KEY'\nAlternatively you can write your API key to file 'sai_platform_key.txt'\nYou can also use and/or override the above by entering your API key in the 'api_key_override' field of each node."
        },
        {
            "author": "Stability-AI",
            "title": "stability-ComfyUI-nodes",
            "id": "sai-nodes",
            "reference": "https://github.com/Stability-AI/stability-ComfyUI-nodes",
            "files": [
                "https://github.com/Stability-AI/stability-ComfyUI-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ColorBlend, ControlLoraSave, GetImageSize. NOTE: Control-LoRA recolor example uses these nodes."
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI's ControlNet Auxiliary Preprocessors",
            "id": "comfyui_controlnet_aux",
            "reference": "https://github.com/Fannovel16/comfyui_controlnet_aux",
            "files": [
                "https://github.com/Fannovel16/comfyui_controlnet_aux"
            ],
            "preemptions":[
                "AIO_Preprocessor",
                "AnimalPosePreprocessor",
                "AnimeFace_SemSegPreprocessor",
                "AnimeLineArtPreprocessor",
                "BAE-NormalMapPreprocessor",
                "BinaryPreprocessor",
                "CannyEdgePreprocessor",
                "ColorPreprocessor",
                "DSINE-NormalMapPreprocessor",
                "DWPreprocessor",
                "DensePosePreprocessor",
                "DepthAnythingPreprocessor",
                "DiffusionEdge_Preprocessor",
                "FacialPartColoringFromPoseKps",
                "FakeScribblePreprocessor",
                "HEDPreprocessor",
                "HintImageEnchance",
                "ImageGenResolutionFromImage",
                "ImageGenResolutionFromLatent",
                "ImageIntensityDetector",
                "ImageLuminanceDetector",
                "InpaintPreprocessor",
                "LeReS-DepthMapPreprocessor",
                "LineArtPreprocessor",
                "LineartStandardPreprocessor",
                "M-LSDPreprocessor",
                "Manga2Anime_LineArt_Preprocessor",
                "MaskOptFlow",
                "MediaPipe-FaceMeshPreprocessor",
                "MeshGraphormer-DepthMapPreprocessor",
                "MiDaS-DepthMapPreprocessor",
                "MiDaS-NormalMapPreprocessor",
                "OneFormer-ADE20K-SemSegPreprocessor",
                "OneFormer-COCO-SemSegPreprocessor",
                "OpenposePreprocessor",
                "PiDiNetPreprocessor",
                "PixelPerfectResolution",
                "SAMPreprocessor",
                "SavePoseKpsAsJsonFile",
                "ScribblePreprocessor",
                "Scribble_XDoG_Preprocessor",
                "SemSegPreprocessor",
                "ShufflePreprocessor",
                "TEEDPreprocessor",
                "TilePreprocessor",
                "UniFormer-SemSegPreprocessor",
                "Unimatch_OptFlowPreprocessor",
                "Zoe-DepthMapPreprocessor",
                "Zoe_DepthAnythingPreprocessor"],
            "install_type": "git-clone",
            "description": "Plug-and-play ComfyUI node sets for making ControlNet hint images."
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI Frame Interpolation",
            "id": "frame-interpolation",
            "reference": "https://github.com/Fannovel16/ComfyUI-Frame-Interpolation",
            "files": [
                "https://github.com/Fannovel16/ComfyUI-Frame-Interpolation"
            ],
            "install_type": "git-clone",
            "description": "A custom node suite for Video Frame Interpolation in ComfyUI"
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI Loopchain",
            "id": "loopchain",
            "reference": "https://github.com/Fannovel16/ComfyUI-Loopchain",
            "files": [
                "https://github.com/Fannovel16/ComfyUI-Loopchain"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes which can be useful for animation in ComfyUI. The main focus of this extension is implementing a mechanism called loopchain. A loopchain in this case is the chain of nodes only executed repeatly in the workflow. If a node chain contains a loop node from this extension, it will become a loop chain."
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI MotionDiff",
            "id": "motiondiff",
            "reference": "https://github.com/Fannovel16/ComfyUI-MotionDiff",
            "files": [
                "https://github.com/Fannovel16/ComfyUI-MotionDiff"
            ],
            "install_type": "git-clone",
            "description": "Implementation of MDM, MotionDiffuse and ReMoDiffuse into ComfyUI."
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI-Video-Matting",
            "id": "video-matting",
            "reference": "https://github.com/Fannovel16/ComfyUI-Video-Matting",
            "files": [
                "https://github.com/Fannovel16/ComfyUI-Video-Matting"
            ],
            "install_type": "git-clone",
            "description": "A minimalistic implementation of [a/Robust Video Matting (RVM)](https://mirror.ghproxy.com/https://github.com/PeterL1n/RobustVideoMatting/) in ComfyUI"
        },
        {
            "author": "Fannovel16",
            "title": "ComfyUI-MagickWand",
            "id": "magicwand",
            "reference": "https://github.com/Fannovel16/ComfyUI-MagickWand",
            "files": [
                "https://github.com/Fannovel16/ComfyUI-MagickWand"
            ],
            "install_type": "git-clone",
            "description": "Proper implementation of ImageMagick - the famous software suite for editing and manipulating digital images to ComfyUI using [a/wandpy](https://mirror.ghproxy.com/https://github.com/emcconville/wand).\nNOTE: You need to install ImageMagick, manually."
        },
        {
            "author": "time-river",
            "title": "CLIPSeg",
            "id": "clipseg",
            "reference": "https://github.com/time-river/ComfyUI-CLIPSeg",
            "files": [
                "https://github.com/time-river/ComfyUI-CLIPSeg/raw/main/custom_nodes/clipseg.py"
            ],
            "install_type": "copy",
            "description": "The CLIPSeg node generates a binary mask for a given input image and text prompt.\nNOTE:This custom node is a forked custom node with hotfixes applied from the [a/original repository](https://mirror.ghproxy.com/https://github.com/biegert/ComfyUI-CLIPSeg), which is no longer maintained."
        },
        {
            "author": "BlenderNeko",
            "title": "ComfyUI Cutoff",
            "id": "cutoff",
            "reference": "https://github.com/BlenderNeko/ComfyUI_Cutoff",
            "files": [
                "https://github.com/BlenderNeko/ComfyUI_Cutoff"
            ],
            "install_type": "git-clone",
            "description": "These custom nodes provides features that allow for better control over the effects of the text prompt."
        },
        {
            "author": "BlenderNeko",
            "title": "Advanced CLIP Text Encode",
            "id": "adv-encode",
            "reference": "https://github.com/BlenderNeko/ComfyUI_ADV_CLIP_emb",
            "files": [
                "https://github.com/BlenderNeko/ComfyUI_ADV_CLIP_emb"
            ],
            "install_type": "git-clone",
            "description": "Advanced CLIP Text Encode (if you need A1111 like prompt. you need this. But Cutoff node includes this feature, already.)"
        },
        {
            "author": "BlenderNeko",
            "title": "ComfyUI Noise",
            "id": "comfy-noise",
            "reference": "https://github.com/BlenderNeko/ComfyUI_Noise",
            "files": [
                "https://github.com/BlenderNeko/ComfyUI_Noise"
            ],
            "install_type": "git-clone",
            "description": "This extension contains 6 nodes for ComfyUI that allows for more control and flexibility over the noise."
        },
        {
            "author": "BlenderNeko",
            "title": "Tiled sampling for ComfyUI",
            "id": "tiled-sampling",
            "reference": "https://github.com/BlenderNeko/ComfyUI_TiledKSampler",
            "files": [
                "https://github.com/BlenderNeko/ComfyUI_TiledKSampler"
            ],
            "install_type": "git-clone",
            "description": "This extension contains a tiled sampler for ComfyUI. It allows for denoising larger images by splitting it up into smaller tiles and denoising these. It tries to minimize any seams for showing up in the end result by gradually denoising all tiles one step at the time and randomizing tile positions for every step."
        },
        {
            "author": "BlenderNeko",
            "title": "SeeCoder [WIP]",
            "id": "seecoder",
            "reference": "https://github.com/BlenderNeko/ComfyUI_SeeCoder",
            "files": [
                "https://github.com/BlenderNeko/ComfyUI_SeeCoder"
            ],
            "install_type": "git-clone",
            "description": "It provides the capability to generate CLIP from an image input, unlike unCLIP, which works in all models. (To use this extension, you need to download the required model file from **Install Models**)"
        },
        {
            "author": "jags111",
            "title": "Efficiency Nodes for ComfyUI Version 2.0+",
            "id": "eff-nodes",
            "reference": "https://github.com/jags111/efficiency-nodes-comfyui",
            "files": [
                "https://github.com/jags111/efficiency-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "A collection of ComfyUI custom nodes to help streamline workflows and reduce total node count.[w/NOTE: This node is originally created by LucianoCirino, but the [a/original repository](https://mirror.ghproxy.com/https://github.com/LucianoCirino/efficiency-nodes-comfyui) is no longer maintained and has been forked by a new maintainer. To use the forked version, you should uninstall the original version and **REINSTALL** this one.]"
        },
        {
            "author": "jags111",
            "title": "Jags_VectorMagic",
            "id": "vectormagic",
            "reference": "https://github.com/jags111/ComfyUI_Jags_VectorMagic",
            "files": [
                "https://github.com/jags111/ComfyUI_Jags_VectorMagic"
            ],
            "install_type": "git-clone",
            "description": "a collection of nodes to explore Vector and image manipulation"
        },
        {
            "author": "jags111",
            "title": "Jags_Audiotools",
            "id": "audiotools",
            "reference": "https://github.com/jags111/ComfyUI_Jags_Audiotools",
            "files": [
                "https://github.com/jags111/ComfyUI_Jags_Audiotools"
            ],
            "install_type": "git-clone",
            "description": "This extension offers various audio generation tools"
        },
        {
            "author": "Derfuu",
            "title": "Derfuu_ComfyUI_ModdedNodes",
            "id": "derfuu",
            "reference": "https://github.com/Derfuu/Derfuu_ComfyUI_ModdedNodes",
            "nodename_pattern": "^DF_",
            "files": [
                "https://github.com/Derfuu/Derfuu_ComfyUI_ModdedNodes"
            ],
            "install_type": "git-clone",
            "description": "Automate calculation depending on image sizes or something you want."
        },
        {
            "author": "paulo-coronado",
            "title": "comfy_clip_blip_node",
            "id": "blip",
            "reference": "https://github.com/paulo-coronado/comfy_clip_blip_node",
            "files": [
                "https://github.com/paulo-coronado/comfy_clip_blip_node"
            ],
            "install_type": "git-clone",
            "apt_dependency": [
                "rustc",
                "cargo"
            ],
            "description": "CLIPTextEncodeBLIP: This custom node provides a CLIP Encoder that is capable of receiving images as input."
        },
        {
            "author": "WASasquatch",
            "title": "WAS Node Suite",
            "id": "was",
            "reference": "https://github.com/WASasquatch/was-node-suite-comfyui",
            "pip": ["numba"],
            "files": [
                "https://github.com/WASasquatch/was-node-suite-comfyui"
            ],
            "install_type": "git-clone",
            "description": "A node suite for ComfyUI with many new nodes, such as image processing, text processing, and more."
        },
        {
            "author": "WASasquatch",
            "title": "ComfyUI Preset Merger",
            "id": "preset-merger",
            "reference": "https://github.com/WASasquatch/ComfyUI_Preset_Merger",
            "files": [
                "https://github.com/WASasquatch/ComfyUI_Preset_Merger"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ModelMergeByPreset. Merge checkpoint models by preset"
        },
        {
            "author": "WASasquatch",
            "title": "PPF_Noise_ComfyUI",
            "id": "ppf",
            "reference": "https://github.com/WASasquatch/PPF_Noise_ComfyUI",
            "files": [
                "https://github.com/WASasquatch/PPF_Noise_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes: WAS_PFN_Latent. Perlin Power Fractal Noisey Latents"
        },
        {
            "author": "WASasquatch",
            "title": "Power Noise Suite for ComfyUI",
            "id": "power-noise",
            "reference": "https://github.com/WASasquatch/PowerNoiseSuite",
            "files": [
                "https://github.com/WASasquatch/PowerNoiseSuite"
            ],
            "install_type": "git-clone",
            "description": "Power Noise Suite contains nodes centered around latent noise input, and diffusion, as well as latent adjustments."
        },
        {
            "author": "WASasquatch",
            "title": "FreeU_Advanced",
            "id": "freeu-adv",
            "reference": "https://github.com/WASasquatch/FreeU_Advanced",
            "files": [
                "https://github.com/WASasquatch/FreeU_Advanced"
            ],
            "install_type": "git-clone",
            "description": "This custom node provides advanced settings for FreeU."
        },
        {
            "author": "WASasquatch",
            "title": "ASTERR",
            "id": "asterr",
            "reference": "https://github.com/WASasquatch/ASTERR",
            "files": [
                "https://github.com/WASasquatch/ASTERR"
            ],
            "install_type": "git-clone",
            "description": "Abstract Syntax Trees Evaluated Restricted Run (ASTERR) is a Python Script executor for ComfyUI. [w/Warning:ASTERR runs Python Code from a Web Interface! It is highly recommended to run this in a closed-off environment, as it could have potential security risks.]"
        },
        {
            "author": "WASasquatch",
            "title": "WAS_Extras",
            "id": "was-extras",
            "reference": "https://github.com/WASasquatch/WAS_Extras",
            "files": [
                "https://github.com/WASasquatch/WAS_Extras"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Conditioning (Blend), Inpainting VAE Encode (WAS), VividSharpen. Experimental nodes, or other random extra helper nodes."
        },
        {
            "author": "SaltAI",
            "title": "SaltAI-Open-Resources",
            "id": "saltai-open-resource",
            "reference": "https://github.com/get-salt-AI/SaltAI",
            "pip": ["numba"],
            "files": [
                "https://github.com/get-salt-AI/SaltAI"
            ],
            "install_type": "git-clone",
            "description": "This repository is a collection of open-source nodes and workflows for ComfyUI, a dev tool that allows users to create node-based workflows often powered by various AI models to do pretty much anything.\nOur mission is to seamlessly connect people and organizations with the world’s foremost AI innovations, anywhere, anytime. Our vision is to foster a flourishing AI ecosystem where the world’s best developers can build and share their work, thereby redefining how software is made, pushing innovation forward, and ensuring as many people as possible can benefit from the positive promise of AI technologies.\nWe believe that ComfyUI is a powerful tool that can help us achieve our mission and vision, by enabling anyone to explore the possibilities and limitations of AI models in a visual and interactive way, without coding if desired.\nWe hope that by sharing our nodes and workflows, we can inspire and empower more people to create amazing AI-powered content with ComfyUI."
        },
        {
            "author": "SaltAI",
            "title": "SaltAI_Language_Toolkit",
            "id": "saltai_language_toolkit",
            "reference": "https://github.com/get-salt-AI/SaltAI_Language_Toolkit",
            "files": [
                "https://github.com/get-salt-AI/SaltAI_Language_Toolkit"
            ],
            "install_type": "git-clone",
            "description": "The project integrates the Retrieval Augmented Generation (RAG) tool [a/Llama-Index](https://www.llamaindex.ai/), [a/Microsoft's AutoGen](https://microsoft.github.io/autogen/), and [a/LlaVA-Next](https://mirror.ghproxy.com/https://github.com/LLaVA-VL/LLaVA-NeXT) with ComfyUI's adaptable node interface, enhancing the functionality and user experience of the platform."
        },
        {
            "author": "omar92",
            "title": "Quality of life Suit:V2",
            "id": "qol",
            "reference": "https://github.com/omar92/ComfyUI-QualityOfLifeSuit_Omar92",
            "files": [
                "https://github.com/omar92/ComfyUI-QualityOfLifeSuit_Omar92"
            ],
            "install_type": "git-clone",
            "description": "openAI suite, String suite, Latent Tools, Image Tools: These custom nodes provide expanded functionality for image and string processing, latent processing, as well as the ability to interface with models such as ChatGPT/DallE-2.\nNOTE: Currently, this extension does not support the new OpenAI API, leading to compatibility issues."
        },
        {
            "author": "lilly1987",
            "title": "simple wildcard for ComfyUI",
            "id": "simle-wildcard",
            "reference": "https://github.com/lilly1987/ComfyUI_node_Lilly",
            "files": [
                "https://github.com/lilly1987/ComfyUI_node_Lilly"
            ],
            "install_type": "git-clone",
            "description": "These custom nodes provides a feature to insert arbitrary inputs through wildcards in the prompt. Additionally, this tool provides features that help simplify workflows, such as VAELoaderDecoder and SimplerSample."
        },
        {
            "author": "sylym",
            "title": "Vid2vid",
            "id": "vid2vid",
            "reference": "https://github.com/sylym/comfy_vid2vid",
            "files": [
                "https://github.com/sylym/comfy_vid2vid"
            ],
            "install_type": "git-clone",
            "description": "A node suite for ComfyUI that allows you to load image sequence and generate new image sequence with different styles or content."
        },
        {
            "author": "EllangoK",
            "title": "ComfyUI-post-processing-nodes",
            "id": "post-processing",
            "reference": "https://github.com/EllangoK/ComfyUI-post-processing-nodes",
            "files": [
                "https://github.com/EllangoK/ComfyUI-post-processing-nodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of post processing nodes for ComfyUI, simply download this repo and drag."
        },
        {
            "author": "LEv145",
            "title": "ImagesGrid",
            "id": "imagesgrid",
            "reference": "https://github.com/LEv145/images-grid-comfy-plugin",
            "files": [
                "https://github.com/LEv145/images-grid-comfy-plugin"
            ],
            "install_type": "git-clone",
            "description": "This tool provides a viewer node that allows for checking multiple outputs in a grid, similar to the X/Y Plot extension."
        },
        {
            "author": "diontimmer",
            "title": "ComfyUI-Vextra-Nodes",
            "id": "vextra",
            "reference": "https://github.com/diontimmer/ComfyUI-Vextra-Nodes",
            "files": [
                "https://github.com/diontimmer/ComfyUI-Vextra-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Pixel Sort, Swap Color Mode, Solid Color, Glitch This, Add Text To Image, Play Sound, Prettify Prompt, Generate Noise, Flatten Colors"
        },
        {
            "author": "CYBERLOOM-INC",
            "title": "ComfyUI-nodes-hnmr",
            "id": "hnmr",
            "reference": "https://github.com/CYBERLOOM-INC/ComfyUI-nodes-hnmr",
            "files": [
                "https://github.com/CYBERLOOM-INC/ComfyUI-nodes-hnmr"
            ],
            "install_type": "git-clone",
            "description": "Provide various custom nodes for Latent, Sampling, Model, Loader, Image, Text. This is the fixed version of the original [a/ComfyUI-nodes-hnmr](https://mirror.ghproxy.com/https://github.com/hnmr293/ComfyUI-nodes-hnmr) by hnmr293."
        },
        {
            "author": "BadCafeCode",
            "title": "Masquerade Nodes",
            "id": "masquerade",
            "reference": "https://github.com/BadCafeCode/masquerade-nodes-comfyui",
            "files": [
                "https://github.com/BadCafeCode/masquerade-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "This is a low-dependency node pack primarily dealing with masks. The author recommends using Impact-Pack instead (unless you specifically have trouble installing dependencies)."
        },
        {
            "author": "Jcd1230",
            "title": "Rembg Background Removal Node for ComfyUI",
            "id": "rembg",
            "reference": "https://github.com/Jcd1230/rembg-comfyui-node",
            "files": [
                "https://github.com/Jcd1230/rembg-comfyui-node"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Image Remove Background (rembg)"
        },
        {
            "author": "YinBailiang",
            "title": "MergeBlockWeighted_fo_ComfyUI",
            "id": "mbw",
            "reference": "https://github.com/YinBailiang/MergeBlockWeighted_fo_ComfyUI",
            "files": [
                "https://github.com/YinBailiang/MergeBlockWeighted_fo_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes: MergeBlockWeighted"
        },
        {
            "author": "trojblue",
            "title": "trNodes",
            "id": "trnodes",
            "reference": "https://github.com/trojblue/trNodes",
            "files": [
                "https://github.com/trojblue/trNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: image_layering, color_correction, model_router"
        },
        {
            "author": "szhublox",
            "title": "Auto-MBW",
            "id": "auto-mbw",
            "reference": "https://github.com/szhublox/ambw_comfyui",
            "files": [
                "https://github.com/szhublox/ambw_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Auto-MBW for ComfyUI loosely based on sdweb-auto-MBW. Nodes: auto merge block weighted"
        },
        {
            "author": "city96",
            "title": "ComfyUI_NetDist",
            "id": "netdist",
            "reference": "https://github.com/city96/ComfyUI_NetDist",
            "files": [
                "https://github.com/city96/ComfyUI_NetDist"
            ],
            "install_type": "git-clone",
            "description": "Run ComfyUI workflows on multiple local GPUs/networked machines. Nodes: Remote images, Local Remote control"
        },
        {
            "author": "city96",
            "title": "Latent-Interposer",
            "id": "latent-interposer",
            "reference": "https://github.com/city96/SD-Latent-Interposer",
            "files": [
                "https://github.com/city96/SD-Latent-Interposer"
            ],
            "install_type": "git-clone",
            "description": "Custom node to convert the lantents between SDXL and SD v1.5 directly without the VAE decoding/encoding step."
        },
        {
            "author": "city96",
            "title": "SD-Latent-Upscaler",
            "id": "latent-upscaler",
            "reference": "https://github.com/city96/SD-Latent-Upscaler",
            "files": [
                "https://github.com/city96/SD-Latent-Upscaler"
            ],
            "pip": ["huggingface-hub"],
            "install_type": "git-clone",
            "description": "Upscaling stable diffusion latents using a small neural network."
        },
        {
            "author": "city96",
            "title": "ComfyUI_DiT [WIP]",
            "id": "dit",
            "reference": "https://github.com/city96/ComfyUI_DiT",
            "files": [
                "https://github.com/city96/ComfyUI_DiT"
            ],
            "pip": ["huggingface-hub"],
            "install_type": "git-clone",
            "description": "Testbed for [a/DiT(Scalable Diffusion Models with Transformers)](https://mirror.ghproxy.com/https://github.com/facebookresearch/DiT). [w/None of this code is stable, expect breaking changes if for some reason you want to use this.]"
        },
        {
            "author": "city96",
            "title": "ComfyUI_ColorMod",
            "id": "colormod",
            "reference": "https://github.com/city96/ComfyUI_ColorMod",
            "files": [
                "https://github.com/city96/ComfyUI_ColorMod"
            ],
            "install_type": "git-clone",
            "description": "This extension currently has two sets of nodes - one set for editing the contrast/color of images and another set for saving images as 16 bit PNG files."
        },
        {
            "author": "city96",
            "title": "Extra Models for ComfyUI",
            "id": "extramodels",
            "reference": "https://github.com/city96/ComfyUI_ExtraModels",
            "files": [
                "https://github.com/city96/ComfyUI_ExtraModels"
            ],
            "install_type": "git-clone",
            "description": "This extension aims to add support for various random image diffusion models to ComfyUI."
        },
        {
            "author": "city96",
            "title": "ComfyUI-GGUF",
            "id": "gguf",
            "reference": "https://github.com/city96/ComfyUI-GGUF",
            "files": [
                "https://github.com/city96/ComfyUI-GGUF"
            ],
            "install_type": "git-clone",
            "description": "GGUF Quantization support for native ComfyUI models\nThis is currently very much WIP. These custom nodes provide support for model files stored in the GGUF format popularized by llama.cpp.\nWhile quantization wasn't feasible for regular UNET models (conv2d), transformer/DiT models such as flux seem less affected by quantization. This allows running it in much lower bits per weight variable bitrate quants on low-end GPUs."
        },
        {
            "author": "SLAPaper",
            "title": "ComfyUI-Image-Selector",
            "id": "image-selector",
            "reference": "https://github.com/SLAPaper/ComfyUI-Image-Selector",
            "files": [
                "https://github.com/SLAPaper/ComfyUI-Image-Selector"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI, which can select one or some of images from a batch."
        },
        {
            "author": "SLAPaper",
            "title": "StableDiffusion-dpmpp_2m_alt-Sampler",
            "id": "dpmpp2m-alt",
            "reference": "https://github.com/SLAPaper/StableDiffusion-dpmpp_2m_alt-Sampler",
            "files": [
                "https://github.com/SLAPaper/StableDiffusion-dpmpp_2m_alt-Sampler"
            ],
            "install_type": "git-clone",
            "description": "the sampler introduced by [a/hallatore](https://mirror.ghproxy.com/https://github.com/AUTOMATIC1111/stable-diffusion-webui/discussions/8457)\ncode extracted from [a/smZNodes](https://mirror.ghproxy.com/https://github.com/shiimizu/ComfyUI_smZNodes).[w/NOTE:ComfyUI-dpmpp_2m_alt-Sampler is renamed to StableDiffusion-dpmpp_2m_alt-Sampler. Please reinstall.]"
        },
        {
            "author": "flyingshutter",
            "title": "As_ComfyUI_CustomNodes",
            "reference": "https://github.com/flyingshutter/As_ComfyUI_CustomNodes",
            "files": [
                "https://github.com/flyingshutter/As_ComfyUI_CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Manipulation nodes for Image, Latent"
        },
        {
            "author": "Zuellni",
            "title": "Zuellni/ComfyUI-Custom-Nodes",
            "reference": "https://github.com/Zuellni/ComfyUI-Custom-Nodes",
            "files": [
                "https://github.com/Zuellni/ComfyUI-Custom-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: DeepFloyd, Filter, Select, Save, Decode, Encode, Repeat, Noise, Noise"
        },
        {
            "author": "Zuellni",
            "title": "ComfyUI ExLlamaV2 Nodes",
            "id": "exllamav2",
            "reference": "https://github.com/Zuellni/ComfyUI-ExLlama-Nodes",
            "files": [
                "https://github.com/Zuellni/ComfyUI-ExLlama-Nodes"
            ],
            "install_type": "git-clone",
            "description": "A simple local text generator for ComfyUI utilizing [a/ExLlamaV2](https://mirror.ghproxy.com/https://github.com/turboderp/exllamav2).\n[w/NOTE:Manual package installation is required.]"
        },
        {
            "author": "Zuellni",
            "title": "ComfyUI PickScore Nodes",
            "id": "pickscore",
            "reference": "https://github.com/Zuellni/ComfyUI-PickScore-Nodes",
            "files": [
                "https://github.com/Zuellni/ComfyUI-PickScore-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Image scoring nodes for ComfyUI using PickScore with a batch of images to predict which ones fit a given prompt the best."
        },
        {
            "author": "AlekPet",
            "title": "AlekPet/ComfyUI_Custom_Nodes_AlekPet",
            "id": "alekpet",
            "reference": "https://github.com/AlekPet/ComfyUI_Custom_Nodes_AlekPet",
            "files": [
                "https://github.com/AlekPet/ComfyUI_Custom_Nodes_AlekPet"
            ],
            "install_type": "git-clone",
            "description": "Nodes: PoseNode, PainterNode, TranslateTextNode, TranslateCLIPTextEncodeNode, DeepTranslatorTextNode, DeepTranslatorCLIPTextEncodeNode, ArgosTranslateTextNode, ArgosTranslateCLIPTextEncodeNode, PreviewTextNode, HexToHueNode, ColorsCorrectNode, IDENode."
        },
        {
            "author": "pythongosssss",
            "title": "ComfyUI WD 1.4 Tagger",
            "id": "wd14",
            "reference": "https://github.com/pythongosssss/ComfyUI-WD14-Tagger",
            "files": [
                "https://github.com/pythongosssss/ComfyUI-WD14-Tagger"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI extension allowing the interrogation of booru tags from images."
        },
        {
            "author": "pythongosssss",
            "title": "pythongosssss/ComfyUI-Custom-Scripts",
            "id": "pygos-script",
            "reference": "https://github.com/pythongosssss/ComfyUI-Custom-Scripts",
            "files": [
                "https://github.com/pythongosssss/ComfyUI-Custom-Scripts"
            ],
            "install_type": "git-clone",
            "description": "This extension provides: Auto Arrange Graph, Workflow SVG, Favicon Status, Image Feed, Latent Upscale By, Lock Nodes & Groups, Lora Subfolders, Preset Text, Show Text, Touch Support, Link Render Mode, Locking, Node Finder, Quick Nodes, Show Image On Menu, Show Text, Workflow Managements, Custom Widget Default Values"
        },
        {
            "author": "strimmlarn",
            "title": "ComfyUI_Strimmlarns_aesthetic_score",
            "id": "aesthetic-score",
            "reference": "https://github.com/strimmlarn/ComfyUI-Strimmlarns-Aesthetic-Score",
            "js_path": "strimmlarn",
            "files": [
                "https://github.com/strimmlarn/ComfyUI-Strimmlarns-Aesthetic-Score"
            ],
            "install_type": "git-clone",
            "description": "Nodes: CalculateAestheticScore, LoadAesteticModel, AesthetlcScoreSorter, ScoreToNumber.\nAesthetic score for ComfyUI"
        },
        {
            "author": "TinyTerra",
            "title": "ComfyUI_tinyterraNodes",
            "id": "ttn",
            "reference": "https://github.com/TinyTerra/ComfyUI_tinyterraNodes",
            "files": [
                "https://github.com/TinyTerra/ComfyUI_tinyterraNodes"
            ],
            "install_type": "git-clone",
            "nodename_pattern": "^ttN ",
            "description": "This extension offers various pipe nodes, extensive XYZ plotting, fullscreen image viewer based on node history, dynamic widgets, interface customization, and more."
        },
        {
            "author": "Jordach",
            "title": "comfy-plasma",
            "id": "plasma",
            "reference": "https://github.com/Jordach/comfy-plasma",
            "files": [
                "https://github.com/Jordach/comfy-plasma"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Plasma Noise, Random Noise, Greyscale Noise, Pink Noise, Brown Noise, Plasma KSampler"
        },
        {
            "author": "bvhari",
            "title": "ImageProcessing",
            "id": "imageprocessing",
            "reference": "https://github.com/bvhari/ComfyUI_ImageProcessing",
            "files": [
                "https://github.com/bvhari/ComfyUI_ImageProcessing"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom nodes to apply various image processing techniques."
        },
        {
            "author": "bvhari",
            "title": "ComfyUI_PerpWeight",
            "id": "perpweight",
            "reference": "https://github.com/bvhari/ComfyUI_PerpWeight",
            "files": [
                "https://github.com/bvhari/ComfyUI_PerpWeight"
            ],
            "install_type": "git-clone",
            "description": "A novel weighting scheme for token vectors from CLIP. Allows a wider range of values for the weight. Inspired by Perp-Neg."
        },
        {
            "author": "bvhari",
            "title": "ComfyUI_SUNoise",
            "id": "sunoise",
            "reference": "https://github.com/bvhari/ComfyUI_SUNoise",
            "files": [
                "https://github.com/bvhari/ComfyUI_SUNoise"
            ],
            "install_type": "git-clone",
            "description": "Scaled Uniform Noise for Ancestral and Stochastic samplers"
        },
        {
            "author": "bvhari",
            "title": "ComfyUI_PerpCFG",
            "reference": "https://github.com/bvhari/ComfyUI_PerpCFG",
            "files": [
                "https://github.com/bvhari/ComfyUI_PerpCFG"
            ],
            "install_type": "git-clone",
            "description": "Perpendicular CFG for reducing oversaturation issues with high guidance scale values."
        },
        {
            "author": "ssitu",
            "title": "UltimateSDUpscale",
            "id": "usdu",
            "reference": "https://github.com/ssitu/ComfyUI_UltimateSDUpscale",
            "files": [
                "https://github.com/ssitu/ComfyUI_UltimateSDUpscale"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for the Ultimate Stable Diffusion Upscale script by Coyote-A."
        },
        {
            "author": "ssitu",
            "title": "Restart Sampling",
            "id": "restart-sampling",
            "reference": "https://github.com/ssitu/ComfyUI_restart_sampling",
            "files": [
                "https://github.com/ssitu/ComfyUI_restart_sampling"
            ],
            "install_type": "git-clone",
            "description": "Unofficial ComfyUI nodes for restart sampling based on the paper 'Restart Sampling for Improving Generative Processes' ([a/paper](https://arxiv.org/abs/2306.14878), [a/repo](https://mirror.ghproxy.com/https://github.com/Newbeeer/diffusion_restart_sampling))"
        },
        {
            "author": "ssitu",
            "title": "ComfyUI roop",
            "id": "roop",
            "reference": "https://github.com/ssitu/ComfyUI_roop",
            "files": [
                "https://github.com/ssitu/ComfyUI_roop"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for the roop A1111 webui script."
        },
        {
            "author": "ssitu",
            "title": "ComfyUI fabric",
            "id": "fabric",
            "reference": "https://github.com/ssitu/ComfyUI_fabric",
            "files": [
                "https://github.com/ssitu/ComfyUI_fabric"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes based on the paper [a/FABRIC: Personalizing Diffusion Models with Iterative Feedback](https://arxiv.org/abs/2307.10159) (Feedback via Attention-Based Reference Image Conditioning)"
        },
        {
            "author": "space-nuko",
            "title": "Disco Diffusion",
            "id": "disco",
            "reference": "https://github.com/space-nuko/ComfyUI-Disco-Diffusion",
            "files": [
                "https://github.com/space-nuko/ComfyUI-Disco-Diffusion"
            ],
            "install_type": "git-clone",
            "description": "Modularized version of Disco Diffusion for use with ComfyUI."
        },
        {
            "author": "space-nuko",
            "title": "OpenPose Editor",
            "id": "openpose-editor",
            "reference": "https://github.com/space-nuko/ComfyUI-OpenPose-Editor",
            "files": [
                "https://github.com/space-nuko/ComfyUI-OpenPose-Editor"
            ],
            "install_type": "git-clone",
            "description": "A port of the openpose-editor extension for stable-diffusion-webui. NOTE: Requires [a/this ComfyUI patch](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI/pull/711) to work correctly"
        },
        {
            "author": "space-nuko",
            "title": "nui suite",
            "id": "nui",
            "reference": "https://github.com/space-nuko/nui-suite",
            "files": [
                "https://github.com/space-nuko/nui-suite"
            ],
            "install_type": "git-clone",
            "description": "NODES: Dynamic Prompts Text Encode, Feeling Lucky Text Encode, Output String"
        },
        {
            "author": "Nourepide",
            "title": "Allor Plugin",
            "id": "allor",
            "reference": "https://github.com/Nourepide/ComfyUI-Allor",
            "files": [
                "https://github.com/Nourepide/ComfyUI-Allor"
            ],
            "install_type": "git-clone",
            "description": "Allor is a plugin for ComfyUI with an emphasis on transparency and performance."
        },
        {
            "author": "melMass",
            "title": "MTB Nodes",
            "id": "mtb",
            "reference": "https://github.com/melMass/comfy_mtb",
            "files": [
                "https://github.com/melMass/comfy_mtb"
            ],
            "nodename_pattern": "\\(mtb\\)$",
            "install_type": "git-clone",
            "description": "NODES: Face Swap, Film Interpolation, Latent Lerp, Int To Number, Bounding Box, Crop, Uncrop, ImageBlur, Denoise, ImageCompare, RGV to HSV, HSV to RGB, Color Correct, Modulo, Deglaze Image, Smart Step, ..."
        },
        {
            "author": "xXAdonesXx",
            "title": "NodeGPT",
            "id": "nodegpt",
            "reference": "https://github.com/xXAdonesXx/NodeGPT",
            "files": [
                "https://github.com/xXAdonesXx/NodeGPT"
            ],
            "install_type": "git-clone",
            "description": "Implementation of AutoGen inside ComfyUI. This repository is under development, and not everything is functioning correctly yet."
        },
        {
            "author": "ciri",
            "title": "ComfyUI Model Downloader",
            "id": "model-downloader",
            "reference": "https://github.com/ciri/comfyui-model-downloader",
            "files": [
                "https://github.com/ciri/comfyui-model-downloader"
            ],
            "install_type": "git-clone",
            "description": "This node allows downloading models directly within ComfyUI for easier use and integration."
        },
        {
            "author": "Suzie1",
            "title": "Comfyroll Studio",
            "id": "comfyroll",
            "reference": "https://github.com/Suzie1/ComfyUI_Comfyroll_CustomNodes",
            "files": [
                "https://github.com/Suzie1/ComfyUI_Comfyroll_CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for SDXL and SD1.5 including Multi-ControlNet, LoRA, Aspect Ratio, Process Switches, and many more nodes. NOTE: Maintainer is changed to Suzie1 from RockOfFire. [w/Using an outdated version has resulted in reported issues with updates not being applied. Trying to reinstall the software is advised.]"
        },
        {
            "author": "bmad4ever",
            "title": "ComfyUI-Bmad-DirtyUndoRedo",
            "reference": "https://github.com/bmad4ever/ComfyUI-Bmad-DirtyUndoRedo",
            "files": [
                "https://github.com/bmad4ever/ComfyUI-Bmad-DirtyUndoRedo"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI extension that adds undo (and redo) functionality."
        },
        {
            "author": "bmad4ever",
            "title": "Bmad Nodes",
            "id": "bmad",
            "reference": "https://github.com/bmad4ever/comfyui_bmad_nodes",
            "files": [
                "https://github.com/bmad4ever/comfyui_bmad_nodes"
            ],
            "install_type": "git-clone",
            "description": "This custom node offers the following functionalities: API support for setting up API requests, computer vision primarily for masking or collages, and general utility to streamline workflow setup or implement essential missing features."
        },
        {
            "author": "bmad4ever",
            "title": "comfyui_ab_sampler",
            "id": "ab-sampler",
            "reference": "https://github.com/bmad4ever/comfyui_ab_samplercustom",
            "files": [
                "https://github.com/bmad4ever/comfyui_ab_samplercustom"
            ],
            "install_type": "git-clone",
            "description": "Experimental sampler node. Sampling alternates between A and B inputs until only one remains, starting with A. B steps run over a 2x2 grid, where 3/4's of the grid are copies of the original input latent. When the optional mask is used, the region outside the defined roi is copied from the original latent at the end of every step."
        },
        {
            "author": "bmad4ever",
            "title": "Lists Cartesian Product",
            "reference": "https://github.com/bmad4ever/comfyui_lists_cartesian_product",
            "files": [
                "https://github.com/bmad4ever/comfyui_lists_cartesian_product"
            ],
            "install_type": "git-clone",
            "description": "Given a set of lists, the node adjusts them so that when used as input to another node all the possible argument permutations are computed."
        },
        {
            "author": "bmad4ever",
            "title": "comfyui_wfc_like",
            "id": "wfc",
            "reference": "https://github.com/bmad4ever/comfyui_wfc_like",
            "files": [
                "https://github.com/bmad4ever/comfyui_wfc_like"
            ],
            "install_type": "git-clone",
            "description": "An 'opinionated' Wave Function Collapse implementation with a set of nodes for comfyui"
        },
        {
            "author": "bmad4ever",
            "title": "comfyui_quilting",
            "id": "quilting",
            "reference": "https://github.com/bmad4ever/comfyui_quilting",
            "files": [
                "https://github.com/bmad4ever/comfyui_quilting"
            ],
            "install_type": "git-clone",
            "description": "image and latent quilting nodes for comfyui"
        },
        {
            "author": "FizzleDorf",
            "title": "FizzNodes",
            "id": "fizz",
            "reference": "https://github.com/FizzleDorf/ComfyUI_FizzNodes",
            "files": [
                "https://github.com/FizzleDorf/ComfyUI_FizzNodes"
            ],
            "install_type": "git-clone",
            "description": "Scheduled prompts, scheduled float/int values and wave function nodes for animations and utility. compatable with [a/framesync](https://www.framesync.xyz/) and [a/keyframe-string-generator](https://www.chigozie.co.uk/keyframe-string-generator/) for audio synced animations in Comfyui."
        },
        {
            "author": "FizzleDorf",
            "title": "ComfyUI-AIT",
            "id": "ait",
            "reference": "https://github.com/FizzleDorf/ComfyUI-AIT",
            "files": [
                "https://github.com/FizzleDorf/ComfyUI-AIT"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI implementation of Facebook Meta's [a/AITemplate](https://mirror.ghproxy.com/https://github.com/facebookincubator/AITemplate) repo for faster inference using cpp/cuda. This new repo is behind the old version but is a much more stable foundation to keep AIT online. Please be patient as the repo will eventually include the same features as before.\nNOTE: You can find the old AIT extension in the legacy channel."
        },
        {
            "author": "filipemeneses",
            "title": "Pixelization",
            "id": "pixelization",
            "reference": "https://github.com/filipemeneses/comfy_pixelization",
            "files": [
                "https://github.com/filipemeneses/comfy_pixelization"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node that pixelizes images."
        },
        {
            "author": "shiimizu",
            "title": "smZNodes",
            "id": "smz",
            "reference": "https://github.com/shiimizu/ComfyUI_smZNodes",
            "files": [
                "https://github.com/shiimizu/ComfyUI_smZNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes such as CLIP Text Encode++ to achieve identical embeddings from stable-diffusion-webui for ComfyUI."
        },
        {
            "author": "shiimizu",
            "title": "Tiled Diffusion & VAE for ComfyUI",
            "id": "tiled-diffusion",
            "reference": "https://github.com/shiimizu/ComfyUI-TiledDiffusion",
            "files": [
                "https://github.com/shiimizu/ComfyUI-TiledDiffusion"
            ],
            "install_type": "git-clone",
            "description": "The extension enables large image drawing & upscaling with limited VRAM via the following techniques:\n1.Two SOTA diffusion tiling algorithms: [a/Mixture of Diffusers](https://mirror.ghproxy.com/https://github.com/albarji/mixture-of-diffusers) and [a/MultiDiffusion](https://mirror.ghproxy.com/https://github.com/omerbt/MultiDiffusion)\n2.pkuliyi2015's Tiled VAE algorithm."
        },
        {
            "author": "shiimizu",
            "title": "ComfyUI PhotoMaker Plus",
            "id": "photomaker-plus",
            "reference": "https://github.com/shiimizu/ComfyUI-PhotoMaker-Plus",
            "files": [
                "https://github.com/shiimizu/ComfyUI-PhotoMaker-Plus"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI reference implementation for [a/PhotoMaker](https://mirror.ghproxy.com/https://github.com/TencentARC/PhotoMaker) models.\nNOTE: PhotoMaker V2 is supported."
        },
        {
            "author": "shiimizu",
            "title": "Semantic-aware Guidance (S-CFG)",
            "id": "s-cfg",
            "reference": "https://github.com/shiimizu/ComfyUI-semantic-aware-guidance",
            "files": [
                "https://github.com/shiimizu/ComfyUI-semantic-aware-guidance"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node for Semantic-aware Guidance based on the [a/paper](https://arxiv.org/abs/2404.05384) 'Rethinking the Spatial Inconsistency in Classifier-Free Diffusion Guidance'"
        },
        {
            "author": "ZaneA",
            "title": "ImageReward",
            "reference": "https://github.com/ZaneA/ComfyUI-ImageReward",
            "files": [
                "https://github.com/ZaneA/ComfyUI-ImageReward"
            ],
            "install_type": "git-clone",
            "description": "NODES: ImageRewardLoader, ImageRewardScore"
        },
        {
            "author": "SeargeDP",
            "title": "SeargeSDXL",
            "id": "searge",
            "reference": "https://github.com/SeargeDP/SeargeSDXL",
            "files": [
                "https://github.com/SeargeDP/SeargeSDXL"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for easier use of SDXL in ComfyUI including an img2img workflow that utilizes both the base and refiner checkpoints."
        },
        {
            "author": "SeargeDP",
            "title": "Searge-LLM for ComfyUI v1.0",
            "reference": "https://github.com/SeargeDP/ComfyUI_Searge_LLM",
            "files": [
                "https://github.com/SeargeDP/ComfyUI_Searge_LLM"
            ],
            "install_type": "git-clone",
            "description": "A prompt-generator or prompt-improvement node for ComfyUI, utilizing the power of a language model to turn a provided text-to-image prompt into a more detailed and improved prompt."
        },
        {
            "author": "cubiq",
            "title": "Simple Math",
            "id": "simplemath",
            "reference": "https://github.com/cubiq/ComfyUI_SimpleMath",
            "files": [
                "https://github.com/cubiq/ComfyUI_SimpleMath"
            ],
            "install_type": "git-clone",
            "description": "custom node for ComfyUI to perform simple math operations"
        },
        {
            "author": "cubiq",
            "title": "ComfyUI_IPAdapter_plus",
            "id": "ipadapter",
            "reference": "https://github.com/cubiq/ComfyUI_IPAdapter_plus",
            "files": [
                "https://github.com/cubiq/ComfyUI_IPAdapter_plus"
            ],
            "preemptions": [
                "IPAAdapterFaceIDBatch",
                "IPAdapter",
                "IPAdapterAdvanced",
                "IPAdapterBatch",
                "IPAdapterClipVisionEnhancer",
                "IPAdapterClipVisionEnhancerBatch",
                "IPAdapterCombineEmbeds",
                "IPAdapterCombineParams",
                "IPAdapterCombineWeights",
                "IPAdapterEmbeds",
                "IPAdapterEmbedsBatch",
                "IPAdapterEncoder",
                "IPAdapterFaceID",
                "IPAdapterFromParams",
                "IPAdapterInsightFaceLoader",
                "IPAdapterLoadEmbeds",
                "IPAdapterMS",
                "IPAdapterModelLoader",
                "IPAdapterNoise",
                "IPAdapterPreciseComposition",
                "IPAdapterPreciseCompositionBatch",
                "IPAdapterPreciseStyleTransfer",
                "IPAdapterPreciseStyleTransferBatch",
                "IPAdapterPromptScheduleFromWeightsStrategy",
                "IPAdapterRegionalConditioning",
                "IPAdapterSaveEmbeds",
                "IPAdapterStyleComposition",
                "IPAdapterStyleCompositionBatch",
                "IPAdapterTiled",
                "IPAdapterTiledBatch",
                "IPAdapterUnifiedLoader",
                "IPAdapterUnifiedLoaderCommunity",
                "IPAdapterUnifiedLoaderFaceID",
                "IPAdapterWeights",
                "IPAdapterWeightsFromStrategy",
                "PrepImageForClipVision"
            ],
            "pip": ["insightface"],
            "install_type": "git-clone",
            "description": "ComfyUI reference implementation for IPAdapter models. The code is mostly taken from the original IPAdapter repository and laksjdjf's implementation, all credit goes to them. I just made the extension closer to ComfyUI philosophy."
        },
        {
            "author": "cubiq",
            "title": "ComfyUI InstantID (Native Support)",
            "id": "instantid",
            "reference": "https://github.com/cubiq/ComfyUI_InstantID",
            "files": [
                "https://github.com/cubiq/ComfyUI_InstantID"
            ],
            "install_type": "git-clone",
            "description": "Native [a/InstantID](https://mirror.ghproxy.com/https://github.com/InstantID/InstantID) support for ComfyUI.\nThis extension differs from the many already available as it doesn't use diffusers but instead implements InstantID natively and it fully integrates with ComfyUI.\nPlease note this still could be considered beta stage, looking forward to your feedback."
        },
        {
            "author": "cubiq",
            "title": "Face Analysis for ComfyUI",
            "id": "faceanalysis",
            "reference": "https://github.com/cubiq/ComfyUI_FaceAnalysis",
            "files": [
                "https://github.com/cubiq/ComfyUI_FaceAnalysis"
            ],
            "install_type": "git-clone",
            "description": "This extension uses [a/DLib](http://dlib.net/) to calculate the Euclidean and Cosine distance between two faces.\nNOTE: Install the Shape Predictor, Face Recognition model from the Install models menu."
        },
        {
            "author": "cubiq",
            "title": "PuLID_ComfyUI",
            "id": "pulid",
            "reference": "https://github.com/cubiq/PuLID_ComfyUI",
            "files": [
                "https://github.com/cubiq/PuLID_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "[a/PuLID](https://mirror.ghproxy.com/https://github.com/ToTheBeginning/PuLID) ComfyUI native implementation."
        },
        {
            "author": "cubiq",
            "title": "Flux blocks patcher sampler",
            "reference": "https://github.com/cubiq/Block_Patcher_ComfyUI",
            "files": [
                "https://github.com/cubiq/Block_Patcher_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This is an (very) advanced and (very) experimental custom node for the ComfyUI. It allows you to iteratively change the blocks weights of Flux models and check the difference each value makes."
        },
        {
            "author": "shockz0rz",
            "title": "comfy-easy-grids",
            "id": "easy-grids",
            "reference": "https://github.com/shockz0rz/comfy-easy-grids",
            "files": [
                "https://github.com/shockz0rz/comfy-easy-grids"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes for creating image grids, sequences, and batches in ComfyUI."
        },
        {
            "author": "yolanother",
            "title": "Comfy UI Prompt Agent",
            "id": "prompt-agent",
            "reference": "https://github.com/yolanother/DTAIComfyPromptAgent",
            "files": [
                "https://github.com/yolanother/DTAIComfyPromptAgent"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Prompt Agent, Prompt Agent (String). This script provides a prompt agent node for the Comfy UI stable diffusion client."
        },
        {
            "author": "yolanother",
            "title": "Image to Text Node",
            "id": "dta-img2txt",
            "reference": "https://github.com/yolanother/DTAIImageToTextNode",
            "files": [
                "https://github.com/yolanother/DTAIImageToTextNode"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Image URL to Text, Image to Text."
        },
        {
            "author": "yolanother",
            "title": "Comfy UI Online Loaders",
            "id": "dta-loader",
            "reference": "https://github.com/yolanother/DTAIComfyLoaders",
            "files": [
                "https://github.com/yolanother/DTAIComfyLoaders"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Submit Image (Parameters), Submit Image. A collection of loaders that use a shared common online data source rather than relying on the files to be present locally."
        },
        {
            "author": "yolanother",
            "title": "Comfy AI DoubTech.ai Image Sumission Node",
            "id": "dta-submit",
            "reference": "https://github.com/yolanother/DTAIComfyImageSubmit",
            "files": [
                "https://github.com/yolanother/DTAIComfyImageSubmit"
            ],
            "install_type": "git-clone",
            "description": "A ComfyAI submit node to upload images to DoubTech.ai"
        },
        {
            "author": "yolanother",
            "title": "Comfy UI QR Codes",
            "id": "dta-qr",
            "reference": "https://github.com/yolanother/DTAIComfyQRCodes",
            "files": [
                "https://github.com/yolanother/DTAIComfyQRCodes"
            ],
            "install_type": "git-clone",
            "description": "This extension introduces QR code nodes for the Comfy UI stable diffusion client. NOTE: ComfyUI qrcode extension required."
        },
        {
            "author": "yolanother",
            "title": "Variables for Comfy UI",
            "id": "dta-var",
            "reference": "https://github.com/yolanother/DTAIComfyVariables",
            "files": [
                "https://github.com/yolanother/DTAIComfyVariables"
            ],
            "install_type": "git-clone",
            "description": "Nodes: String, Int, Float, Short String, CLIP Text Encode (With Variables), String Format, Short String Format. This extension introduces quality of life improvements by providing variable nodes and shared global variables."
        },
        {
            "author": "sipherxyz",
            "title": "comfyui-art-venture",
            "id": "artventure",
            "reference": "https://github.com/sipherxyz/comfyui-art-venture",
            "files": [
                "https://github.com/sipherxyz/comfyui-art-venture"
            ],
            "install_type": "git-clone",
            "description": "A comprehensive set of custom nodes for ComfyUI, focusing on utilities for image processing, JSON manipulation, model operations and working with object via URLs"
        },
        {
            "author": "SOELexicon",
            "title": "LexMSDBNodes",
            "id": "lexmsdb",
            "reference": "https://github.com/SOELexicon/ComfyUI-LexMSDBNodes",
            "files": [
                "https://github.com/SOELexicon/ComfyUI-LexMSDBNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: MSSqlTableNode, MSSqlSelectNode. This extension provides custom nodes to interact with MSSQL."
        },
        {
            "author": "pants007",
            "title": "pants",
            "reference": "https://github.com/pants007/comfy-pants",
            "files": [
                "https://github.com/pants007/comfy-pants"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Make Square Node, Interrogate Node, TextEncodeAIO"
        },
        {
            "author": "evanspearman",
            "title": "ComfyMath",
            "id": "comfymath",
            "reference": "https://github.com/evanspearman/ComfyMath",
            "files": [
                "https://github.com/evanspearman/ComfyMath"
            ],
            "install_type": "git-clone",
            "description": "Provides Math Nodes for ComfyUI. Boolean Logic, Integer Arithmetic, Floating Point Arithmetic and Functions, Vec2, Vec3, and Vec4 Arithmetic and Functions"
        },
        {
            "author": "civitai",
            "title": "Civitai Comfy Nodes",
            "id": "civitai",
            "reference": "https://github.com/civitai/civitai_comfy_nodes",
            "files": [
                "https://github.com/civitai/civitai_comfy_nodes"
            ],
            "install_type": "git-clone",
            "description": "Tired of manually downloading and moving models, LoRAs, and more to the right places?\nSick of scouring Civitai for that one mystical LoRA someone was using to make that cool image?\nWant to be share a fully reproducable workflow?"
        },
        {
            "author": "andersxa",
            "title": "CLIP Directional Prompt Attention",
            "id": "prompt-attention",
            "reference": "https://github.com/andersxa/comfyui-PromptAttention",
            "files": [
                "https://github.com/andersxa/comfyui-PromptAttention"
            ],
            "pip": ["scikit-learn", "matplotlib"],
            "install_type": "git-clone",
            "description": "Nodes: CLIP Directional Prompt Attention Encode. Direction prompt attention tries to solve the problem of contextual words (or parts of the prompt) having an effect on much later or irrelevant parts of the prompt."
        },
        {
            "author": "ArtVentureX",
            "title": "AnimateDiff",
            "reference": "https://github.com/ArtVentureX/comfyui-animatediff",
            "pip": ["flash_attn"],
            "files": [
                "https://github.com/ArtVentureX/comfyui-animatediff"
            ],
            "install_type": "git-clone",
            "description": "AnimateDiff integration for ComfyUI, adapts from sd-webui-animatediff.\n[w/You only need to download one of [a/mm_sd_v14.ckpt](https://huggingface.co/guoyww/animatediff/resolve/main/mm_sd_v14.ckpt) | [a/mm_sd_v15.ckpt](https://huggingface.co/guoyww/animatediff/resolve/main/mm_sd_v15.ckpt). Put the model weights under %%ComfyUI/custom_nodes/comfyui-animatediff/models%%. DO NOT change model filename.]"
        },
        {
            "author": "twri",
            "title": "SDXL Prompt Styler",
            "id": "twri-styler",
            "reference": "https://github.com/twri/sdxl_prompt_styler",
            "files": [
                "https://github.com/twri/sdxl_prompt_styler"
            ],
            "install_type": "git-clone",
            "description": "SDXL Prompt Styler is a node that enables you to style prompts based on predefined templates stored in a JSON file."
        },
        {
            "author": "wolfden",
            "title": "SDXL Prompt Styler (customized version by wolfden)",
            "id": "wolfden-styler",
            "reference": "https://github.com/wolfden/ComfyUi_PromptStylers",
            "files": [
                "https://github.com/wolfden/ComfyUi_PromptStylers"
            ],
            "install_type": "git-clone",
            "description": "These custom nodes provide a variety of customized prompt stylers based on [a/twri/SDXL Prompt Styler](https://mirror.ghproxy.com/https://github.com/twri/sdxl_prompt_styler)."
        },
        {
            "author": "wolfden",
            "title": "ComfyUi_String_Function_Tree",
            "id": "str-func-tree",
            "reference": "https://github.com/wolfden/ComfyUi_String_Function_Tree",
            "files": [
                "https://github.com/wolfden/ComfyUi_String_Function_Tree"
            ],
            "install_type": "git-clone",
            "description": "This custom node provides the capability to manipulate multiple string inputs."
        },
        {
            "author": "daxthin",
            "title": "DZ-FaceDetailer",
            "id": "dz-facedetailer",
            "reference": "https://github.com/nicofdga/DZ-FaceDetailer",
            "files": [
                "https://github.com/nicofdga/DZ-FaceDetailer"
            ],
            "install_type": "git-clone",
            "description": "Face Detailer is a custom node for the 'ComfyUI' framework inspired by !After Detailer extension from auto1111, it allows you to detect faces using Mediapipe and YOLOv8n to create masks for the detected faces."
        },
        {
            "author": "asagi4",
            "title": "ComfyUI prompt control",
            "id": "prompt-control",
            "reference": "https://github.com/asagi4/comfyui-prompt-control",
            "files": [
                "https://github.com/asagi4/comfyui-prompt-control"
            ],
            "install_type": "git-clone",
            "description": "Nodes for convenient prompt editing. The aim is to make basic generations in ComfyUI completely prompt-controllable."
        },
        {
            "author": "asagi4",
            "title": "ComfyUI-CADS",
            "id": "cads",
            "reference": "https://github.com/asagi4/ComfyUI-CADS",
            "files": [
                "https://github.com/asagi4/ComfyUI-CADS"
            ],
            "install_type": "git-clone",
            "description": "Attempts to implement [a/CADS](https://arxiv.org/abs/2310.17347) for ComfyUI. Credit also to the [a/A1111 implementation](https://mirror.ghproxy.com/https://github.com/v0xie/sd-webui-cads/tree/main) that I used as a reference."
        },
        {
            "author": "asagi4",
            "title": "asagi4/comfyui-utility-nodes",
            "id": "asagi-nodes",
            "reference": "https://github.com/asagi4/comfyui-utility-nodes",
            "files": [
                "https://github.com/asagi4/comfyui-utility-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MUJinjaRender, MUSimpleWildcard"
        },
        {
            "author": "asagi4",
            "title": "Adaptive Guidance for ComfyUI",
            "id": "comfyui-adaptive-guidance",
            "reference": "https://github.com/asagi4/ComfyUI-Adaptive-Guidance",
            "files": [
                "https://github.com/asagi4/ComfyUI-Adaptive-Guidance"
            ],
            "install_type": "git-clone",
            "description": "An implementation of adaptive guidance for ComfyUI\nSee [a/https://bcv-uniandes.github.io/adaptiveguidance-wp](https://bcv-uniandes.github.io/adaptiveguidance-wp)"
        },
        {
            "author": "jamesWalker55",
            "title": "ComfyUI - P2LDGAN Node",
            "id": "p2ldgan",
            "reference": "https://github.com/jamesWalker55/comfyui-p2ldgan",
            "files": [
                "https://github.com/jamesWalker55/comfyui-p2ldgan"
            ],
            "install_type": "git-clone",
            "description": "Nodes: P2LDGAN. This integrates P2LDGAN into ComfyUI. P2LDGAN extracts lineart from input images.\n[w/To use this extension, you need to download the [a/p2ldgan model](https://drive.google.com/file/d/1To4V_Btc3QhCLBWZ0PdSNgC1cbm3isHP) and save it in the %%ComfyUI/custom_nodes/comfyui-p2ldgan/checkpoints%% directory.]"
        },
        {
            "author": "jamesWalker55",
            "title": "Various ComfyUI Nodes by Type",
            "id": "jameswalker-nodes",
            "reference": "https://github.com/jamesWalker55/comfyui-various",
            "files": [
                "https://github.com/jamesWalker55/comfyui-various"
            ],
            "nodename_pattern": "^JW",
            "install_type": "git-clone",
            "description": "Nodes: JWInteger, JWFloat, JWString, JWImageLoadRGB, JWImageResize, ..."
        },
        {
            "author": "adieyal",
            "title": "DynamicPrompts Custom Nodes",
            "id": "dynamicprompt",
            "reference": "https://github.com/adieyal/comfyui-dynamicprompts",
            "files": [
                "https://github.com/adieyal/comfyui-dynamicprompts"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Random Prompts, Combinatorial Prompts, I'm Feeling Lucky, Magic Prompt, Jinja2 Templates. ComfyUI-DynamicPrompts is a custom nodes library that integrates into your existing ComfyUI Library. It provides nodes that enable the use of Dynamic Prompts in your ComfyUI."
        },
        {
            "author": "mihaiiancu",
            "title": "mihaiiancu/Inpaint",
            "id": "inpaint",
            "reference": "https://github.com/mihaiiancu/ComfyUI_Inpaint",
            "files": [
                "https://github.com/mihaiiancu/ComfyUI_Inpaint"
            ],
            "install_type": "git-clone",
            "description": "Nodes: InpaintMediapipe. This node provides a simple interface to inpaint."
        },
        {
            "author": "kwaroran",
            "title": "abg-comfyui",
            "id": "abg",
            "reference": "https://github.com/kwaroran/abg-comfyui",
            "files": [
                "https://github.com/kwaroran/abg-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Remove Image Background (abg). A Anime Background Remover node for comfyui, based on this hf space, works same as AGB extention in automatic1111."
        },
        {
            "author": "bash-j",
            "title": "Mikey Nodes",
            "id": "mikey",
            "reference": "https://github.com/bash-j/mikey_nodes",
            "files": [
                "https://github.com/bash-j/mikey_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Prompt With Style, Prompt With SDXL, Resize Image for SDXL, Save Image With Prompt Data, HaldCLUT, Empty Latent Ratio Select/Custom SDXL"
        },
        {
            "author": "blib-la",
            "title": "blibla-comfyui-extensions",
            "id": "blibla-comfyui-extensions",
            "reference": "https://github.com/blib-la/blibla-comfyui-extensions",
            "files": [
                "https://github.com/blib-la/blibla-comfyui-extensions"
            ],
            "install_type": "git-clone",
            "description": "node color customization, custom colors, dot reroutes, link rendering options, straight lines, group freezing, node pinning, automated arrangement of nodes, copy image\n[w/failfast-comfyui-extensions is renamed to blibla-comfyui-extensions. Please resintall to this.]"
        },
        {
            "author": "Pfaeff",
            "title": "pfaeff-comfyui",
            "id": "pfaeff",
            "reference": "https://github.com/Pfaeff/pfaeff-comfyui",
            "files": [
                "https://github.com/Pfaeff/pfaeff-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes: AstropulsePixelDetector, BackgroundRemover, ImagePadForBetterOutpaint, InpaintingPipelineLoader, Inpainting, ..."
        },
        {
            "author": "wallish77",
            "title": "wlsh_nodes",
            "id": "wlsh",
            "reference": "https://github.com/wallish77/wlsh_nodes",
            "files": [
                "https://github.com/wallish77/wlsh_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Checkpoint Loader with Name, Save Prompt Info, Outpaint to Image, CLIP Positive-Negative, SDXL Quick Empty Latent, Empty Latent by Ratio, Time String, SDXL Steps, SDXL Resolutions ..."
        },
        {
            "author": "Kosinkadink",
            "title": "ComfyUI-Advanced-ControlNet",
            "id": "adv-cnet",
            "reference": "https://github.com/Kosinkadink/ComfyUI-Advanced-ControlNet",
            "files": [
                "https://github.com/Kosinkadink/ComfyUI-Advanced-ControlNet"
            ],
            "install_type": "git-clone",
            "description": "Nodes for scheduling ControlNet strength across timesteps and batched latents, as well as applying custom weights and attention masks."
        },
        {
            "author": "Kosinkadink",
            "title": "AnimateDiff Evolved",
            "id": "ad-evolved",
            "reference": "https://github.com/Kosinkadink/ComfyUI-AnimateDiff-Evolved",
            "files": [
                "https://github.com/Kosinkadink/ComfyUI-AnimateDiff-Evolved"
            ],
            "install_type": "git-clone",
            "description": "A forked repository that actively maintains [a/AnimateDiff](https://mirror.ghproxy.com/https://github.com/ArtVentureX/comfyui-animatediff), created by ArtVentureX.\n\nImproved AnimateDiff integration for ComfyUI, adapts from sd-webui-animatediff.\n[w/Download one or more motion models from [a/Original Models](https://huggingface.co/guoyww/animatediff/tree/main) | [a/Finetuned Models](https://huggingface.co/manshoety/AD_Stabilized_Motion/tree/main). See README for additional model links and usage. Put the model weights under %%ComfyUI/custom_nodes/ComfyUI-AnimateDiff-Evolved/models%%. You are free to rename the models, but keeping original names will ease use when sharing your workflow.]"
        },
        {
            "author": "Kosinkadink",
            "title": "ComfyUI-VideoHelperSuite",
            "id": "vhs",
            "reference": "https://github.com/Kosinkadink/ComfyUI-VideoHelperSuite",
            "files": [
                "https://github.com/Kosinkadink/ComfyUI-VideoHelperSuite"
            ],
            "install_type": "git-clone",
            "description": "Nodes related to video workflows"
        },
        {
            "author": "Gourieff",
            "title": "ReActor Node for ComfyUI",
            "id": "reactor",
            "reference": "https://github.com/Gourieff/comfyui-reactor-node",
            "files": [
                "https://github.com/Gourieff/comfyui-reactor-node"
            ],
            "install_type": "git-clone",
            "description": "The Fast and Simple 'roop-like' Face Swap Extension Node for ComfyUI, based on ReActor (ex Roop-GE) SD-WebUI Face Swap Extension"
        },
        {
            "author": "Gourieff",
            "title": "ComfyUI-FutureWarningIgnore",
            "id": "futureignore",
            "reference": "https://github.com/Gourieff/ComfyUI-FutureWarningIgnore",
            "files": [
                "https://github.com/Gourieff/ComfyUI-FutureWarningIgnore/raw/main/0_FutureWarningIgnore.py"
            ],
            "install_type": "copy",
            "description": "This extension collapses 'future warning' messages in your Console"
        },
        {
            "author": "imb101",
            "title": "FaceSwap",
            "id": "faceswap",
            "reference": "https://github.com/imb101/ComfyUI-FaceSwap",
            "files": [
                "https://github.com/imb101/ComfyUI-FaceSwap"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FaceSwapNode. Very basic custom node to enable face swapping in ComfyUI. (roop)"
        },
        {
            "author": "Chaoses-Ib",
            "title": "ComfyUI_Ib_CustomNodes",
            "id": "ib-nodes",
            "reference": "https://github.com/Chaoses-Ib/ComfyUI_Ib_CustomNodes",
            "files": [
                "https://github.com/Chaoses-Ib/ComfyUI_Ib_CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LoadImageFromPath. Load Image From Path loads the image from the source path and does not have such problems."
        },
        {
            "author": "AIrjen",
            "title": "One Button Prompt",
            "id": "1button",
            "reference": "https://github.com/AIrjen/OneButtonPrompt",
            "files": [
                "https://github.com/AIrjen/OneButtonPrompt"
            ],
            "install_type": "git-clone",
            "description": "One Button Prompt has a prompt generation node for beginners who have problems writing a good prompt, or advanced users who want to get inspired. It generates an entire prompt from scratch. It is random, but controlled. You simply load up the script and press generate, and let it surprise you."
        },
        {
            "author": "coreyryanhanson",
            "title": "ComfyQR",
            "id": "comfyqr",
            "reference": "https://github.com/coreyryanhanson/ComfyQR",
            "files": [
                "https://github.com/coreyryanhanson/ComfyQR"
            ],
            "install_type": "git-clone",
            "description": "QR generation within ComfyUI. Contains nodes suitable for workflows from generating basic QR images to techniques with advanced QR masking."
        },
        {
            "author": "coreyryanhanson",
            "title": "ComfyQR-scanning-nodes",
            "id": "comfyqr-scanning",
            "reference": "https://github.com/coreyryanhanson/ComfyQR-scanning-nodes",
            "files": [
                "https://github.com/coreyryanhanson/ComfyQR-scanning-nodes"
            ],
            "install_type": "git-clone",
            "description": "A set of ComfyUI nodes to quickly test generated QR codes for scannability. A companion project to ComfyQR."
        },
        {
            "author": "dimtoneff",
            "title": "ComfyUI PixelArt Detector",
            "id": "pixelart-detector",
            "reference": "https://github.com/dimtoneff/ComfyUI-PixelArt-Detector",
            "files": [
                "https://github.com/dimtoneff/ComfyUI-PixelArt-Detector"
            ],
            "install_type": "git-clone",
            "description": "This node manipulates the pixel art image in ways that it should look pixel perfect (downscales, changes palette, upscales etc.)."
        },
        {
            "author": "hylarucoder",
            "title": "comfyui-copilot",
            "reference": "https://github.com/hylarucoder/comfyui-copilot",
            "files": [
                "https://github.com/hylarucoder/comfyui-copilot"
            ],
            "install_type": "git-clone",
            "description": "NODES:Eagle Image Node for PNGInfo, SDXL Resolution Presets (ws), SDXL Prompt Styler, SDXL Prompt Styler Advanced"
        },
        {
            "author": "theUpsider",
            "title": "Styles CSV Loader Extension for ComfyUI",
            "id": "styles-csv-loader",
            "reference": "https://github.com/theUpsider/ComfyUI-Styles_CSV_Loader",
            "files": [
                "https://github.com/theUpsider/ComfyUI-Styles_CSV_Loader"
            ],
            "install_type": "git-clone",
            "description": "This extension allows users to load styles from a CSV file, primarily for migration purposes from the automatic1111 Stable Diffusion web UI."
        },
        {
            "author": "theUpsider",
            "title": "ComfyUI-Logic",
            "id": "comfy-logic",
            "reference": "https://github.com/theUpsider/ComfyUI-Logic",
            "files": [
                "https://github.com/theUpsider/ComfyUI-Logic"
            ],
            "install_type": "git-clone",
            "description": "An extension to ComfyUI that introduces logic nodes and conditional rendering capabilities."
        },
        {
            "author": "M1kep",
            "title": "Comfy_KepListStuff",
            "id": "keplist",
            "reference": "https://github.com/M1kep/Comfy_KepListStuff",
            "files": [
                "https://github.com/M1kep/Comfy_KepListStuff"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Range(Step), Range(Num Steps), List Length, Image Overlay, Stack Images, Empty Images, Join Image Lists, Join Float Lists. This extension provides various list manipulation nodes"
        },
        {
            "author": "M1kep",
            "title": "ComfyLiterals",
            "id": "comfyliterals",
            "reference": "https://github.com/M1kep/ComfyLiterals",
            "files": [
                "https://github.com/M1kep/ComfyLiterals"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Int, Float, String, Operation, Checkpoint"
        },
        {
            "author": "M1kep",
            "title": "KepPromptLang",
            "id": "kepprompt",
            "reference": "https://github.com/M1kep/KepPromptLang",
            "files": [
                "https://github.com/M1kep/KepPromptLang"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Build Gif, Special CLIP Loader. It offers various manipulation capabilities for the internal operations of the prompt."
        },
        {
            "author": "M1kep",
            "title": "Comfy_KepMatteAnything",
            "id": "kepmatte",
            "reference": "https://github.com/M1kep/Comfy_KepMatteAnything",
            "files": [
                "https://github.com/M1kep/Comfy_KepMatteAnything"
            ],
            "install_type": "git-clone",
            "description": "This extension provides a custom node that allows the use of [a/Matte Anything](https://mirror.ghproxy.com/https://github.com/hustvl/Matte-Anything) in ComfyUI."
        },
        {
            "author": "M1kep",
            "title": "Comfy_KepKitchenSink",
            "id": "kepkitchen",
            "reference": "https://github.com/M1kep/Comfy_KepKitchenSink",
            "files": [
                "https://github.com/M1kep/Comfy_KepKitchenSink"
            ],
            "install_type": "git-clone",
            "description": "Nodes: KepRotateImage"
        },
        {
            "author": "M1kep",
            "title": "ComfyUI-OtherVAEs",
            "id": "kep-othervae",
            "reference": "https://github.com/M1kep/ComfyUI-OtherVAEs",
            "files": [
                "https://github.com/M1kep/ComfyUI-OtherVAEs"
            ],
            "install_type": "git-clone",
            "description": "Nodes: TAESD VAE Decode"
        },
        {
            "author": "M1kep",
            "title": "ComfyUI-KepOpenAI",
            "id": "kep-openai",
            "reference": "https://github.com/M1kep/ComfyUI-KepOpenAI",
            "files": [
                "https://github.com/M1kep/ComfyUI-KepOpenAI"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-KepOpenAI is a user-friendly node that serves as an interface to the GPT-4 with Vision (GPT-4V) API. This integration facilitates the processing of images coupled with text prompts, leveraging the capabilities of the OpenAI API to generate text completions that are contextually relevant to the provided inputs."
        },
        {
            "author": "uarefans",
            "title": "ComfyUI-Fans",
            "id": "fans",
            "reference": "https://github.com/uarefans/ComfyUI-Fans",
            "files": [
                "https://github.com/uarefans/ComfyUI-Fans"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Fans Styler (Max 10 Style), Fans Text Concat (Until 10 text), Fans Prompt Styler Postive (Can replace {prompt} word in your csv files), Fans Prompt Styler Negative (With sentence structure)."
        },
        {
            "author": "NicholasMcCarthy",
            "title": "ComfyUI_TravelSuite",
            "id": "travel",
            "reference": "https://github.com/NicholasMcCarthy/ComfyUI_TravelSuite",
            "files": [
                "https://github.com/NicholasMcCarthy/ComfyUI_TravelSuite"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom nodes to apply various latent travel techniques."
        },
        {
            "author": "ManglerFTW",
            "title": "ComfyI2I",
            "id": "comfyi2i",
            "reference": "https://github.com/ManglerFTW/ComfyI2I",
            "files": [
                "https://github.com/ManglerFTW/ComfyI2I"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes to perform image 2 image functions in ComfyUI."
        },
        {
            "author": "m-sokes",
            "title": "ComfyUI Sokes Nodes",
            "id": "sokes",
            "reference": "https://github.com/m-sokes/ComfyUI-Sokes-Nodes",
            "files": [
                "https://github.com/m-sokes/ComfyUI-Sokes-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Empty Latent Randomizer (9 Inputs)"
        },
        {
            "author": "Extraltodeus",
            "title": "noise latent perlinpinpin",
            "id": "perlipinpin",
            "reference": "https://github.com/Extraltodeus/noise_latent_perlinpinpin",
            "files": [
                "https://github.com/Extraltodeus/noise_latent_perlinpinpin"
            ],
            "install_type": "git-clone",
            "description": "Nodes: NoisyLatentPerlin. This allows to create latent spaces filled with perlin-based noise that can actually be used by the samplers."
        },
        {
            "author": "Extraltodeus",
            "title": "LoadLoraWithTags",
            "reference": "https://github.com/Extraltodeus/LoadLoraWithTags",
            "files": [
                "https://github.com/Extraltodeus/LoadLoraWithTags"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LoadLoraWithTags. Save/Load trigger words for loras from a json and auto fetch them on civitai if they are missing."
        },
        {
            "author": "Extraltodeus",
            "title": "sigmas_tools_and_the_golden_scheduler",
            "id": "sigmas-tools",
            "reference": "https://github.com/Extraltodeus/sigmas_tools_and_the_golden_scheduler",
            "files": [
                "https://github.com/Extraltodeus/sigmas_tools_and_the_golden_scheduler"
            ],
            "install_type": "git-clone",
            "description": "A few nodes to mix sigmas and a custom scheduler that uses phi, then one using eval() to be able to schedule with custom formulas."
        },
        {
            "author": "Extraltodeus",
            "title": "ComfyUI-AutomaticCFG",
            "id": "autocfg",
            "reference": "https://github.com/Extraltodeus/ComfyUI-AutomaticCFG",
            "files": [
                "https://github.com/Extraltodeus/ComfyUI-AutomaticCFG"
            ],
            "install_type": "git-clone",
            "description": "My own version 'from scratch' of a self-rescaling CFG. It isn't much but it's honest work.\nTLDR: set your CFG at 8 to try it. No burned images and artifacts anymore. CFG is also a bit more sensitive because it's a proportion around 8. Low scale like 4 also gives really nice results since your CFG is not the CFG anymore. Also in general even with relatively low settings it seems to improve the quality."
        },
        {
            "author": "Extraltodeus",
            "title": "Vector_Sculptor_ComfyUI",
            "id": "vector-sculptor",
            "reference": "https://github.com/Extraltodeus/Vector_Sculptor_ComfyUI",
            "files": [
                "https://github.com/Extraltodeus/Vector_Sculptor_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "The main node makes your conditioning go towards similar concepts so to enrich your composition or further away so to make it more precise. It gathers similar pre-cond vectors for as long as the cosine similarity score diminishes. If it climbs back it stops. This allows to set a relative direction to similar concepts.\nThere are examples at the end but [a/you can also check this imgur album](https://imgur.com/a/WvPd81Y) which demonstrates the capability of improving variety."
        },
        {
            "author": "Extraltodeus",
            "title": "Stable-Diffusion-temperature-settings",
            "id": "sd-temperature",
            "reference": "https://github.com/Extraltodeus/Stable-Diffusion-temperature-settings",
            "files": [
                "https://github.com/Extraltodeus/Stable-Diffusion-temperature-settings"
            ],
            "install_type": "git-clone",
            "description": "Provides the ability to set the temperature for both UNET and CLIP. For ComfyUI."
        },
        {
            "author": "Extraltodeus",
            "title": "Uncond-Zero-for-ComfyUI",
            "id": "uncond-zero",
            "reference": "https://github.com/Extraltodeus/Uncond-Zero-for-ComfyUI",
            "files": [
                "https://github.com/Extraltodeus/Uncond-Zero-for-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Allows to sample without generating any uncond with Stable Diffusion!"
        },
        {
            "author": "Extraltodeus",
            "title": "pre_cfg_comfy_nodes_for_ComfyUI",
            "id": "precfg",
            "reference": "https://github.com/Extraltodeus/pre_cfg_comfy_nodes_for_ComfyUI",
            "files": [
                "https://github.com/Extraltodeus/pre_cfg_comfy_nodes_for_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes to prepare the noise predictions before the CFG function"
        },
        {
            "author": "Extraltodeus",
            "title": "Skimmed_CFG",
            "id": "skimmed-cfg",
            "reference": "https://github.com/Extraltodeus/Skimmed_CFG",
            "files": [
                "https://github.com/Extraltodeus/Skimmed_CFG"
            ],
            "install_type": "git-clone",
            "description": "A powerful anti-burn allowing much higher CFG scales for latent diffusion models (for ComfyUI)"
        },
        {
            "author": "Extraltodeus",
            "title": "DistanceSampler",
            "id": "distancesampler",
            "reference": "https://github.com/Extraltodeus/DistanceSampler",
            "files": [
                "https://github.com/Extraltodeus/Skimmed_CFG"
            ],
            "install_type": "git-clone",
            "description": "Heuristic modification of the Heun sampler using a custom function based on normalized distances. For ComfyUI."
        },
        {
            "author": "JPS",
            "title": "JPS Custom Nodes for ComfyUI",
            "id": "jps-nodes",
            "reference": "https://github.com/JPS-GER/ComfyUI_JPS-Nodes",
            "files": [
                "https://github.com/JPS-GER/ComfyUI_JPS-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Various nodes to handle SDXL Resolutions, SDXL Basic Settings, IP Adapter Settings, Revision Settings, SDXL Prompt Styler, Crop Image to Square, Crop Image to Target Size, Get Date-Time String, Resolution Multiply, Largest Integer, 5-to-1 Switches for Integer, Images, Latents, Conditioning, Model, VAE, ControlNet"
        },
        {
            "author": "hustille",
            "title": "hus' utils for ComfyUI",
            "id": "husutil",
            "reference": "https://github.com/hustille/ComfyUI_hus_utils",
            "files": [
                "https://github.com/hustille/ComfyUI_hus_utils"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes primarily for seed and filename generation"
        },
        {
            "author": "hustille",
            "title": "ComfyUI_Fooocus_KSampler",
            "id": "fooocus-ksampler",
            "reference": "https://github.com/hustille/ComfyUI_Fooocus_KSampler",
            "files": [
                "https://github.com/hustille/ComfyUI_Fooocus_KSampler"
            ],
            "install_type": "git-clone",
            "description": "Nodes: KSampler With Refiner (Fooocus). The KSampler from [a/Fooocus](https://mirror.ghproxy.com/https://github.com/lllyasviel/Fooocus) as a ComfyUI node [w/NOTE: This patches basic ComfyUI behaviour - don't use together with other samplers. Or perhaps do? Other samplers might profit from those changes ... ymmv.]"
        },
        {
            "author": "badjeff",
            "title": "LoRA Tag Loader for ComfyUI",
            "id": "lora-tag-loader",
            "reference": "https://github.com/badjeff/comfyui_lora_tag_loader",
            "files": [
                "https://github.com/badjeff/comfyui_lora_tag_loader"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node to read LoRA tag(s) from text and load it into checkpoint model."
        },
        {
            "author": "rgthree",
            "title": "rgthree's ComfyUI Nodes",
            "id": "rgthree",
            "reference": "https://github.com/rgthree/rgthree-comfy",
            "files": [
                "https://github.com/rgthree/rgthree-comfy"
            ],
            "nodename_pattern": " \\(rgthree\\)$",
            "install_type": "git-clone",
            "description": "Nodes: Seed, Reroute, Context, Lora Loader Stack, Context Switch, Fast Muter. These custom nodes helps organize the building of complex workflows."
        },
        {
            "author": "AIGODLIKE",
            "title": "AIGODLIKE-COMFYUI-TRANSLATION",
            "id": "translation",
            "reference": "https://github.com/AIGODLIKE/AIGODLIKE-COMFYUI-TRANSLATION",
            "files": [
                "https://github.com/AIGODLIKE/AIGODLIKE-COMFYUI-TRANSLATION"
            ],
            "install_type": "git-clone",
            "description": "It provides language settings. (Contribution from users of various languages is needed due to the support for each language.)"
        },
        {
            "author": "AIGODLIKE",
            "title": "AIGODLIKE-ComfyUI-Studio",
            "id": "comfy-studio",
            "reference": "https://github.com/AIGODLIKE/AIGODLIKE-ComfyUI-Studio",
            "files": [
                "https://github.com/AIGODLIKE/AIGODLIKE-ComfyUI-Studio"
            ],
            "install_type": "git-clone",
            "description": "Improve the interactive experience of using ComfyUI, such as making the loading of ComfyUI models more intuitive and making it easier to create model thumbnails"
        },
        {
            "author": "AIGODLIKE",
            "title": "ComfyUI-CUP",
            "id": "comfycup",
            "reference": "https://github.com/AIGODLIKE/ComfyUI-CUP",
            "files": [
                "https://github.com/AIGODLIKE/ComfyUI-CUP"
            ],
            "install_type": "git-clone",
            "description": "Bridge between ComfyUI and blender's ComfyUI-BlenderAI-node addon."
        },
        {
            "author": "AIGODLIKE",
            "title": "ComfyUI-ToonCrafter",
            "id": "tooncrafter",
            "reference": "https://github.com/AIGODLIKE/ComfyUI-ToonCrafter",
            "files": [
                "https://github.com/AIGODLIKE/ComfyUI-ToonCrafter"
            ],
            "install_type": "git-clone",
            "description": "This project is used to enable [a/ToonCrafter](https://mirror.ghproxy.com/https://github.com/ToonCrafter/ToonCrafter) to be used in ComfyUI.\nYou can use it to achieve generative keyframe animation\nAnd use it in Blender for animation rendering and prediction"
        },
        {
            "author": "syllebra",
            "title": "BilboX's ComfyUI Custom Nodes",
            "id": "bilbox",
            "reference": "https://github.com/syllebra/bilbox-comfyui",
            "files": [
                "https://github.com/syllebra/bilbox-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes: BilboX's PromptGeek Photo Prompt. This provides a convenient way to compose photorealistic prompts into ComfyUI. Post-Processing: adds various post processing effects. Bonus: Option to show a distant server shutdown menu."
        },
        {
            "author": "Girish Gopaul",
            "title": "Save Image with Generation Metadata",
            "id": "image-saver",
            "reference": "https://github.com/giriss/comfy-image-saver",
            "files": [
                "https://github.com/giriss/comfy-image-saver"
            ],
            "install_type": "git-clone",
            "description": "All the tools you need to save images with their generation metadata on ComfyUI. Compatible with Civitai & Prompthero geninfo auto-detection. Works with png, jpeg and webp."
        },
        {
            "author": "shingo1228",
            "title": "ComfyUI-send-Eagle(slim)",
            "id": "send-eagle",
            "reference": "https://github.com/shingo1228/ComfyUI-send-eagle-slim",
            "files": [
                "https://github.com/shingo1228/ComfyUI-send-eagle-slim"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Send Webp Image to Eagle. This is an extension node for ComfyUI that allows you to send generated images in webp format to Eagle. This extension node is a re-implementation of the Eagle linkage functions of the previous ComfyUI-send-Eagle node, focusing on the functions required for this node."
        },
        {
            "author": "shingo1228",
            "title": "ComfyUI-SDXL-EmptyLatentImage",
            "id": "sdxl-emptylatent",
            "reference": "https://github.com/shingo1228/ComfyUI-SDXL-EmptyLatentImage",
            "files": [
                "https://github.com/shingo1228/ComfyUI-SDXL-EmptyLatentImage"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SDXL Empty Latent Image. An extension node for ComfyUI that allows you to select a resolution from the pre-defined json files and output a Latent Image."
        },
        {
            "author": "laksjdjf",
            "title": "pfg-ComfyUI",
            "id": "pfg",
            "reference": "https://github.com/laksjdjf/pfg-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/pfg-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI version of https://mirror.ghproxy.com/https://github.com/laksjdjf/pfg-webui. (To use this extension, you need to download the required model file from **Install Models**)"
        },
        {
            "author": "laksjdjf",
            "title": "cgem156-ComfyUI🍌",
            "id": "cgem156",
            "reference": "https://github.com/laksjdjf/cgem156-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/cgem156-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "The custom nodes of laksjdjf have been integrated into the node pack of cgem156🍌.\nNOTE:This includes the attention couple feature."
        },
        {
            "author": "laksjdjf",
            "title": "cd-tuner_negpip-ComfyUI",
            "id": "cdtuner",
            "reference": "https://github.com/laksjdjf/cd-tuner_negpip-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/cd-tuner_negpip-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Apply CDTuner, Apply Negapip. This extension provides the [a/CD(Color/Detail) Tuner](https://mirror.ghproxy.com/https://github.com/hako-mikan/sd-webui-cd-tuner) and the [a/Negative Prompt in the Prompt](https://mirror.ghproxy.com/https://github.com/hako-mikan/sd-webui-negpip) features."
        },
        {
            "author": "laksjdjf",
            "title": "LCMSampler-ComfyUI",
            "id": "lcm-sampler",
            "reference": "https://github.com/laksjdjf/LCMSampler-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/LCMSampler-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This extension node is intended for the use of LCM conversion for SSD-1B-anime. It does not guarantee operation with the original LCM (as it cannot load weights in the current version). To take advantage of fast generation with LCM, a node for using TAESD as a decoder is also provided. This is inspired by ComfyUI-OtherVAEs."
        },
        {
            "author": "laksjdjf",
            "title": "LoRTnoC-ComfyUI",
            "id": "lortnoc",
            "reference": "https://github.com/laksjdjf/LoRTnoC-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/LoRTnoC-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This is a repository for using LoRTnoC (LoRA with hint block of ControlNet) on ComfyUI.\nNOTE:Please place the model file in the same location as controlnet. (Is this too arbitrary?)"
        },
        {
            "author": "laksjdjf",
            "title": "Batch-Condition-ComfyUI",
            "id": "batch-condition",
            "reference": "https://github.com/laksjdjf/Batch-Condition-ComfyUI",
            "files": [
                "https://github.com/laksjdjf/Batch-Condition-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes:CLIP Text Encode (Batch), String Input, Batch String"
        },
        {
            "author": "alsritter",
            "title": "asymmetric-tiling-comfyui",
            "id": "asymmetric",
            "reference": "https://github.com/alsritter/asymmetric-tiling-comfyui",
            "files": [
                "https://github.com/alsritter/asymmetric-tiling-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Asymmetric_Tiling_KSampler. "
        },
        {
            "author": "meap158",
            "title": "GPU temperature protection",
            "id": "gputemp",
            "reference": "https://github.com/meap158/ComfyUI-GPU-temperature-protection",
            "files": [
                "https://github.com/meap158/ComfyUI-GPU-temperature-protection"
            ],
            "install_type": "git-clone",
            "description": "Pause image generation when GPU temperature exceeds threshold."
        },
        {
            "author": "meap158",
            "title": "ComfyUI-Prompt-Expansion",
            "id": "promtp-expansion",
            "reference": "https://github.com/meap158/ComfyUI-Prompt-Expansion",
            "files": [
                "https://github.com/meap158/ComfyUI-Prompt-Expansion"
            ],
            "install_type": "git-clone",
            "description": "Dynamic prompt expansion, powered by GPT-2 locally on your device."
        },
        {
            "author": "meap158",
            "title": "ComfyUI-Background-Replacement",
            "id": "bg-replacement",
            "reference": "https://github.com/meap158/ComfyUI-Background-Replacement",
            "files": [
                "https://github.com/meap158/ComfyUI-Background-Replacement"
            ],
            "install_type": "git-clone",
            "description": "Instantly replace your image's background."
        },
        {
            "author": "TeaCrab",
            "title": "ComfyUI-TeaNodes",
            "id": "teanodes",
            "reference": "https://github.com/TeaCrab/ComfyUI-TeaNodes",
            "files": [
                "https://github.com/TeaCrab/ComfyUI-TeaNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:TC_EqualizeCLAHE, TC_SizeApproximation, TC_ImageResize, TC_ImageScale, TC_ColorFill."
        },
        {
            "author": "nagolinc",
            "title": "ComfyUI_FastVAEDecorder_SDXL",
            "reference": "https://github.com/nagolinc/ComfyUI_FastVAEDecorder_SDXL",
            "files": [
                "https://github.com/nagolinc/ComfyUI_FastVAEDecorder_SDXL"
            ],
            "install_type": "git-clone",
            "description": "Based off of: [a/Birch-san/diffusers-play/approx_vae](https://mirror.ghproxy.com/https://github.com/Birch-san/diffusers-play/tree/main/approx_vae). This ComfyUI node allows you to quickly preview SDXL 1.0 latents."
        },
        {
            "author": "nagolinc",
            "title": "comfyui_openai_node",
            "reference": "https://github.com/nagolinc/comfyui_openai_node",
            "files": [
                "https://github.com/nagolinc/comfyui_openai_node"
            ],
            "install_type": "git-clone",
            "description": "This provides a single node openai > Open AI query node\nthat takes a system prompt and user message and sends them to chatGPT 3.5\nNote, you MUST have an OPEN AI API key stored in the environment variable OPENAI_API_KEY in order for this to work."
        },
        {
            "author": "bradsec",
            "title": "ResolutionSelector for ComfyUI",
            "id": "resolution-selector",
            "reference": "https://github.com/bradsec/ComfyUI_ResolutionSelector",
            "files": [
                "https://github.com/bradsec/ComfyUI_ResolutionSelector"
            ],
            "install_type": "git-clone",
            "description": "A custom node for Stable Diffusion ComfyUI to enable easy selection of image resolutions for SDXL SD15 SD21"
        },
        {
            "author": "kohya-ss",
            "title": "ControlNet-LLLite-ComfyUI",
            "id": "lllite",
            "reference": "https://github.com/kohya-ss/ControlNet-LLLite-ComfyUI",
            "files": [
                "https://github.com/kohya-ss/ControlNet-LLLite-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LLLiteLoader"
        },
        {
            "author": "jjkramhoeft",
            "title": "ComfyUI-Jjk-Nodes",
            "id": "jjk",
            "reference": "https://github.com/jjkramhoeft/ComfyUI-Jjk-Nodes",
            "files": [
                "https://github.com/jjkramhoeft/ComfyUI-Jjk-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: SDXLRecommendedImageSize, JjkText, JjkShowText, JjkConcat. A set of custom nodes for ComfyUI - focused on text and parameter utility"
        },
        {
            "author": "dagthomas",
            "title": "SDXL Auto Prompter",
            "id": "autoprompt",
            "reference": "https://github.com/dagthomas/comfyui_dagthomas",
            "files": [
                "https://github.com/dagthomas/comfyui_dagthomas"
            ],
            "install_type": "git-clone",
            "description": "Easy prompting for generation of endless random art pieces and photographs!"
        },
        {
            "author": "marhensa",
            "title": "Recommended Resolution Calculator",
            "id": "resoultion-calc",
            "reference": "https://github.com/marhensa/sdxl-recommended-res-calc",
            "files": [
                "https://github.com/marhensa/sdxl-recommended-res-calc"
            ],
            "install_type": "git-clone",
            "description": "Input your desired output final resolution, it will automaticaly set the initial recommended SDXL ratio/size and its Upscale Factor to reach that output final resolution, also there's an option for 2x/4x reverse Upscale Factor. These all to avoid using bad/arbitary initial ratio/resolution."
        },
        {
            "author": "Nuked",
            "title": "ComfyUI-N-Nodes",
            "id": "nnodes",
            "reference": "https://github.com/Nuked88/ComfyUI-N-Nodes",
            "files": [
                "https://github.com/Nuked88/ComfyUI-N-Nodes"
            ],
            "install_type": "git-clone",
            "description": "A suite of custom nodes for ConfyUI that includes GPT text-prompt generation, LoadVideo,SaveVideo,LoadFramesFromFolder and FrameInterpolator"
        },
        {
            "author": "Nuked",
            "title": "ComfyUI-N-Sidebar",
            "id": "nsidebar",
            "reference": "https://github.com/Nuked88/ComfyUI-N-Sidebar",
            "files": [
                "https://github.com/Nuked88/ComfyUI-N-Sidebar"
            ],
            "install_type": "git-clone",
            "description": "A simple sidebar for ComfyUI."
        },
        {
            "author": "richinsley",
            "title": "Comfy-LFO",
            "id": "lfo",
            "reference": "https://github.com/richinsley/Comfy-LFO",
            "files": [
                "https://github.com/richinsley/Comfy-LFO"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LFO_Triangle, LFO_Sine, SawtoothNode, SquareNode, PulseNode. ComfyUI custom nodes to create Low Frequency Oscillators."
        },
        {
            "author": "Beinsezii",
            "title": "bsz-cui-extras",
            "id": "bsz",
            "reference": "https://github.com/Beinsezii/bsz-cui-extras",
            "files": [
                "https://github.com/Beinsezii/bsz-cui-extras"
            ],
            "install_type": "git-clone",
            "description": "This contains all-in-one 'principled' nodes for T2I, I2I, refining, and scaling. Additionally it has many tools for directly manipulating the color of latents, high res fix math, and scripted image post-processing."
        },
        {
            "author": "youyegit",
            "title": "tdxh_node_comfyui",
            "id": "tdxh",
            "reference": "https://github.com/youyegit/tdxh_node_comfyui",
            "files": [
                "https://github.com/youyegit/tdxh_node_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Add Switch on nodes, Make nodes amount small! It helps conveniently to use less nodes for doing the same things."
        },
        {
            "author": "Sxela",
            "title": "ComfyWarp",
            "id": "warp",
            "reference": "https://github.com/Sxela/ComfyWarp",
            "files": [
                "https://github.com/Sxela/ComfyWarp"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LoadFrameSequence, LoadFrame"
        },
        {
            "author": "skfoo",
            "title": "ComfyUI-Coziness",
            "id": "coziness",
            "reference": "https://github.com/skfoo/ComfyUI-Coziness",
            "files": [
                "https://github.com/skfoo/ComfyUI-Coziness"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MultiLora Loader, Lora Text Extractor. Provides a node for assisting in loading loras through text."
        },
        {
            "author": "YOUR-WORST-TACO",
            "title": "ComfyUI-TacoNodes",
            "id": "taco",
            "reference": "https://github.com/YOUR-WORST-TACO/ComfyUI-TacoNodes",
            "files": [
                "https://github.com/YOUR-WORST-TACO/ComfyUI-TacoNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:TacoLatent, TacoAnimatedLoader, TacoImg2ImgAnimatedLoader, TacoGifMaker."
        },
        {
            "author": "Lerc",
            "title": "Canvas Tab",
            "id": "canvastab",
            "reference": "https://github.com/Lerc/canvas_tab",
            "files": [
                "https://github.com/Lerc/canvas_tab"
            ],
            "install_type": "git-clone",
            "description": "This extension provides a full page image editor with mask support. There are two nodes, one to receive images from the editor and one to send images to the editor."
        },
        {
            "author": "Ttl",
            "title": "ComfyUI Neural Network Latent Upscale",
            "id": "nnlatent",
            "reference": "https://github.com/Ttl/ComfyUi_NNLatentUpscale",
            "files": [
                "https://github.com/Ttl/ComfyUi_NNLatentUpscale"
            ],
            "install_type": "git-clone",
            "preemptions": ["NNLatentUpscale"],
            "description": "Nodes:NNLatentUpscale, A custom ComfyUI node designed for rapid latent upscaling using a compact neural network, eliminating the need for VAE-based decoding and encoding."
        },
        {
            "author": "spro",
            "title": "Latent Mirror node for ComfyUI",
            "id": "latentmirror",
            "reference": "https://github.com/spro/comfyui-mirror",
            "files": [
                "https://github.com/spro/comfyui-mirror"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Latent Mirror. Node to mirror a latent along the Y (vertical / left to right) or X (horizontal / top to bottom) axis."
        },
        {
            "author": "Tropfchen",
            "title": "Embedding Picker",
            "id": "embedding-picker",
            "reference": "https://github.com/Tropfchen/ComfyUI-Embedding_Picker",
            "files": [
                "https://github.com/Tropfchen/ComfyUI-Embedding_Picker"
            ],
            "install_type": "git-clone",
            "description": "Tired of forgetting and misspelling often weird names of embeddings you use? Or perhaps you use only one, cause you forgot you have tens of them installed?"
        },
        {
            "author": "Acly",
            "title": "ComfyUI Nodes for External Tooling",
            "id": "external-tooling",
            "reference": "https://github.com/Acly/comfyui-tooling-nodes",
            "files": [
                "https://github.com/Acly/comfyui-tooling-nodes"
            ],
            "install_type": "git-clone",
            "description": "Provides nodes and server API extensions geared towards using ComfyUI as a backend for external tools."
        },
        {
            "author": "Acly",
            "title": "ComfyUI Inpaint Nodes",
            "id": "inpaint-nodes",
            "reference": "https://github.com/Acly/comfyui-inpaint-nodes",
            "files": [
                "https://github.com/Acly/comfyui-inpaint-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes for better inpainting with ComfyUI. Adds various ways to pre-process inpaint areas. Supports the Fooocus inpaint model, a small and flexible patch which can be applied to any SDXL checkpoint and will improve consistency when generating masked areas."
        },
        {
            "author": "picturesonpictures",
            "title": "comfy_PoP",
            "id": "pop",
            "reference": "https://github.com/picturesonpictures/comfy_PoP",
            "files": ["https://github.com/picturesonpictures/comfy_PoP"],
            "install_type": "git-clone",
            "description": "A collection of custom nodes for ComfyUI. Includes a quick canny edge detection node with unconventional settings, simple LoRA stack nodes for workflow efficiency, and a customizable aspect ratio node."
        },
        {
            "author": "Dream Project",
            "title": "Dream Project Animation Nodes",
            "id": "dream-anime",
            "reference": "https://github.com/alt-key-project/comfyui-dream-project",
            "files": [
                "https://github.com/alt-key-project/comfyui-dream-project"
            ],
            "install_type": "git-clone",
            "description": "This extension offers various nodes that are useful for Deforum-like animations in ComfyUI."
        },
        {
            "author": "Dream Project",
            "title": "Dream Video Batches",
            "id": "dream-video",
            "reference": "https://github.com/alt-key-project/comfyui-dream-video-batches",
            "files": [
                "https://github.com/alt-key-project/comfyui-dream-video-batches"
            ],
            "install_type": "git-clone",
            "description": "Provide utilities for batch based video generation workflows (s.a. AnimateDiff and Stable Video Diffusion)."
        },
        {
            "author": "seanlynch",
            "title": "ComfyUI Optical Flow",
            "id": "optical-flow",
            "reference": "https://github.com/seanlynch/comfyui-optical-flow",
            "files": [
                "https://github.com/seanlynch/comfyui-optical-flow"
            ],
            "install_type": "git-clone",
            "description": "This package contains three nodes to help you compute optical flow between pairs of images, usually adjacent frames in a video, visualize the flow, and apply the flow to another image of the same dimensions. Most of the code is from Deforum, so this is released under the same license (MIT)."
        },
        {
            "author": "ealkanat",
            "title": "ComfyUI Easy Padding",
            "id": "easy-padding",
            "reference": "https://github.com/ealkanat/comfyui-easy-padding",
            "files": [
                "https://github.com/ealkanat/comfyui-easy-padding"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Easy Padding is a simple custom ComfyUI node that helps you to add padding to images on ComfyUI."
        },
        {
            "author": "ArtBot2023",
            "title": "Character Face Swap",
            "id": "char-faceswap",
            "reference": "https://github.com/ArtBot2023/CharacterFaceSwap",
            "files": [
                "https://github.com/ArtBot2023/CharacterFaceSwap"
            ],
            "install_type": "git-clone",
            "description": "Character face swap with LoRA and embeddings."
        },
        {
            "author": "mav-rik",
            "title": "Facerestore CF (Code Former)",
            "id": "face-cf",
            "reference": "https://github.com/mav-rik/facerestore_cf",
            "files": [
                "https://github.com/mav-rik/facerestore_cf"
            ],
            "install_type": "git-clone",
            "description": "This is a copy of [a/facerestore custom node](https://civitai.com/models/24690/comfyui-facerestore-node) with a bit of a change to support CodeFormer Fidelity parameter. These ComfyUI nodes can be used to restore faces in images similar to the face restore option in AUTOMATIC1111 webui.\nNOTE: To use this node, you need to download the face restoration model and face detection model from the 'Install models' menu."
        },
        {
            "author": "braintacles",
            "title": "braintacles-nodes",
            "id": "braintacles",
            "reference": "https://github.com/braintacles/braintacles-comfyui-nodes",
            "files": [
                "https://github.com/braintacles/braintacles-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: CLIPTextEncodeSDXL-Multi-IO, CLIPTextEncodeSDXL-Pipe, Empty Latent Image from Aspect-Ratio, Random Find and Replace."
        },
        {
            "author": "hayden-fr",
            "title": "ComfyUI-Model-Manager",
            "id": "modelmanager",
            "reference": "https://github.com/hayden-fr/ComfyUI-Model-Manager",
            "files": [
                "https://github.com/hayden-fr/ComfyUI-Model-Manager"
            ],
            "install_type": "git-clone",
            "description": "Manage models: browsing, download and delete."
        },
        {
            "author": "hayden-fr",
            "title": "ComfyUI-Image-Browsing",
            "id": "image-browsing",
            "reference": "https://github.com/hayden-fr/ComfyUI-Image-Browsing",
            "files": [
                "https://github.com/hayden-fr/ComfyUI-Image-Browsing"
            ],
            "install_type": "git-clone",
            "description": "Image Browsing: browsing, download and delete."
        },
        {
            "author": "ali1234",
            "title": "comfyui-job-iterator",
            "id": "job-iterator",
            "reference": "https://github.com/ali1234/comfyui-job-iterator",
            "files": [
                "https://github.com/ali1234/comfyui-job-iterator"
            ],
            "install_type": "git-clone",
            "description": "Implements iteration over sequences within a single workflow run. [w/NOTE: This node replaces the execution of ComfyUI for iterative processing functionality.]"
        },
        {
            "author": "jmkl",
            "title": "ComfyUI Ricing",
            "id": "ricing",
            "reference": "https://github.com/jmkl/ComfyUI-ricing",
            "files": [
                "https://github.com/jmkl/ComfyUI-ricing"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom user.css and some script stuff. mainly for web interface."
        },
        {
            "author": "budihartono",
            "title": "Otonx's Custom Nodes",
            "id": "otonx",
            "reference": "https://github.com/budihartono/comfyui_otonx_nodes",
            "files": [
                "https://github.com/budihartono/comfyui_otonx_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: OTX Multiple Values, OTX KSampler Feeder. This extension provides custom nodes for ComfyUI created for personal projects. Made available for reference. Nodes may be updated or changed intermittently or not at all. Review & test before use."
        },
        {
            "author": "ramyma",
            "title": "A8R8 ComfyUI Nodes",
            "id": "a8r8",
            "reference": "https://github.com/ramyma/A8R8_ComfyUI_nodes",
            "files": [
                "https://github.com/ramyma/A8R8_ComfyUI_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Base64Image Input Node, Base64Image Output Node. [a/A8R8](https://mirror.ghproxy.com/https://github.com/ramyma/a8r8) supporting nodes to integrate with ComfyUI"
        },
        {
            "author": "spinagon",
            "title": "Seamless tiling Node for ComfyUI",
            "id": "seamless",
            "reference": "https://github.com/spinagon/ComfyUI-seamless-tiling",
            "files": [
                "https://github.com/spinagon/ComfyUI-seamless-tiling"
            ],
            "install_type": "git-clone",
            "description": "Node for generating almost seamless textures, based on similar setting from A1111."
        },
        {
            "author": "BiffMunky",
            "title": "Endless ️🌊✨ Nodes",
            "id": "endless",
            "reference": "https://github.com/tusharbhutt/Endless-Nodes",
            "files": [
                "https://github.com/tusharbhutt/Endless-Nodes"
            ],
            "install_type": "git-clone",
            "description": "A small set of nodes I created for various numerical and text inputs.  Features image saver with ability to have JSON saved to separate folder, parameter collection nodes, two aesthetic scoring models, switches for text and numbers, and conversion of string to numeric and vice versa."
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-HQ-Image-Save",
            "id": "hq-image-save",
            "reference": "https://github.com/spacepxl/ComfyUI-HQ-Image-Save",
            "files": [
                "https://github.com/spacepxl/ComfyUI-HQ-Image-Save"
            ],
            "install_type": "git-clone",
            "description": "Add Image Save nodes for TIFF 16 bit and EXR 32 bit formats. Probably only useful if you're applying a LUT or other color corrections, and care about preserving as much color accuracy as possible."
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-Image-Filters",
            "id": "image-fitlers",
            "reference": "https://github.com/spacepxl/ComfyUI-Image-Filters",
            "files": [
                "https://github.com/spacepxl/ComfyUI-Image-Filters"
            ],
            "install_type": "git-clone",
            "description": "Image and matte filtering nodes for ComfyUI image/filters/*"
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-RAVE",
            "id": "rave",
            "reference": "https://github.com/spacepxl/ComfyUI-RAVE",
            "files": [
                "https://github.com/spacepxl/ComfyUI-RAVE"
            ],
            "install_type": "git-clone",
            "description": "Unofficial ComfyUI implementation of [a/RAVE](https://rave-video.github.io/)"
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-StyleGan",
            "id": "stylegan",
            "reference": "https://github.com/spacepxl/ComfyUI-StyleGan",
            "files": [
                "https://github.com/spacepxl/ComfyUI-StyleGan"
            ],
            "install_type": "git-clone",
            "description": "Basic support for StyleGAN2 and StyleGAN3 models."
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-Florence-2",
            "id": "florence2-spacepxl",
            "reference": "https://github.com/spacepxl/ComfyUI-Florence-2",
            "files": [
                "https://github.com/spacepxl/ComfyUI-Florence-2"
            ],
            "install_type": "git-clone",
            "description": "[a/https://huggingface.co/microsoft/Florence-2-large-ft](https://huggingface.co/microsoft/Florence-2-large-ft)\nLarge or base model, support for captioning and bbox task modes, more coming soon."
        },
        {
            "author": "spacepxl",
            "title": "ComfyUI-Depth-Pro",
            "reference": "https://github.com/spacepxl/ComfyUI-Depth-Pro",
            "files": [
                "https://github.com/spacepxl/ComfyUI-Depth-Pro"
            ],
            "install_type": "git-clone",
            "description": "Based on [a/https://mirror.ghproxy.com/https://github.com/apple/ml-depth-pro](https://mirror.ghproxy.com/https://github.com/apple/ml-depth-pro)"
        },
        {
            "author": "PTA",
            "title": "auto nodes layout",
            "id": "autolayout",
            "reference": "https://github.com/phineas-pta/comfyui-auto-nodes-layout",
            "files": [
                "https://github.com/phineas-pta/comfyui-auto-nodes-layout"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI extension to apply better nodes layout algorithm to ComfyUI workflow (mostly for visualization purpose)"
        },
        {
            "author": "receyuki",
            "title": "SD Prompt Reader",
            "id": "sdpromptreader",
            "reference": "https://github.com/receyuki/comfyui-prompt-reader-node",
            "files": [
                "https://github.com/receyuki/comfyui-prompt-reader-node"
            ],
            "install_type": "git-clone",
            "description": "The ultimate solution for managing image metadata and multi-tool compatibility. ComfyUI node version of the SD Prompt Reader."
        },
        {
            "author": "cubiq",
            "title": "ComfyUI Essentials",
            "id": "essentials",
            "reference": "https://github.com/cubiq/ComfyUI_essentials",
            "files": [
                "https://github.com/cubiq/ComfyUI_essentials"
            ],
            "install_type": "git-clone",
            "description": "Essential nodes that are weirdly missing from ComfyUI core. With few exceptions they are new features and not commodities. I hope this will be just a temporary repository until the nodes get included into ComfyUI."
        },
        {
            "author": "Clybius",
            "title": "ComfyUI-Latent-Modifiers",
            "id": "latent-modifier",
            "reference": "https://github.com/Clybius/ComfyUI-Latent-Modifiers",
            "files": [
                "https://github.com/Clybius/ComfyUI-Latent-Modifiers"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Latent Diffusion Mega Modifier. ComfyUI nodes which modify the latent during the diffusion process. (Sharpness, Tonemap, Rescale, Extra Noise)"
        },
        {
            "author": "Clybius",
            "title": "ComfyUI Extra Samplers",
            "id": "extra-samplers",
            "reference": "https://github.com/Clybius/ComfyUI-Extra-Samplers",
            "files": [
                "https://github.com/Clybius/ComfyUI-Extra-Samplers"
            ],
            "install_type": "git-clone",
            "description": "Nodes: SamplerCustomNoise, SamplerCustomNoiseDuo, SamplerCustomModelMixtureDuo, SamplerRES_Momentumized, SamplerDPMPP_DualSDE_Momentumized, SamplerCLYB_4M_SDE_Momentumized, SamplerTTM, SamplerLCMCustom\nThis extension provides various custom samplers not offered by the default nodes in ComfyUI."
        },
        {
            "author": "mcmonkeyprojects",
            "title": "Dynamic Thresholding",
            "id": "dynamic-thresholding",
            "reference": "https://github.com/mcmonkeyprojects/sd-dynamic-thresholding",
            "files": [
                "https://github.com/mcmonkeyprojects/sd-dynamic-thresholding"
            ],
            "install_type": "git-clone",
            "description": "Adds nodes for Dynamic Thresholding, CFG scheduling, and related techniques."
        },
        {
            "author": "Tropfchen",
            "title": "YARS: Yet Another Resolution Selector",
            "id": "yars",
            "reference": "https://github.com/Tropfchen/ComfyUI-yaResolutionSelector",
            "files": [
                "https://github.com/Tropfchen/ComfyUI-yaResolutionSelector"
            ],
            "install_type": "git-clone",
            "description": "A slightly different Resolution Selector node, allowing to freely change base resolution and aspect ratio, with options to maintain the pixel count or use the base resolution as the highest or lowest dimension."
        },
        {
            "author": "chrisgoringe",
            "title": "Noise variation and batch noise tools",
            "id": "cg-noisetools",
            "reference": "https://github.com/chrisgoringe/cg-noisetools",
            "files": [
                "https://github.com/chrisgoringe/cg-noisetools"
            ],
            "install_type": "git-clone",
            "description": "Nodes to create small variations on noise, to shape noise, and to control noise in batches. Replaces the old 'variation-seed' nodes."
        },
        {
            "author": "chrisgoringe",
            "title": "Image chooser",
            "id": "image-chooser",
            "reference": "https://github.com/chrisgoringe/cg-image-picker",
            "files": [
                "https://github.com/chrisgoringe/cg-image-picker"
            ],
            "install_type": "git-clone",
            "description": "A custom node that pauses the flow while you choose which image (or latent) to pass on to the rest of the workflow."
        },
        {
            "author": "chrisgoringe",
            "title": "Use Everywhere (UE Nodes)",
            "id": "ue",
            "reference": "https://github.com/chrisgoringe/cg-use-everywhere",
            "files": [
                "https://github.com/chrisgoringe/cg-use-everywhere"
            ],
            "install_type": "git-clone",
            "nodename_pattern": "(^(Prompts|Anything) Everywhere|Simple String)",
            "description": "A set of nodes that allow data to be 'broadcast' to some or all unconnected inputs. Greatly reduces link spaghetti."
        },
        {
            "author": "chrisgoringe",
            "title": "Prompt Info",
            "id": "promptinfo",
            "reference": "https://github.com/chrisgoringe/cg-prompt-info",
            "files": [
                "https://github.com/chrisgoringe/cg-prompt-info"
            ],
            "install_type": "git-clone",
            "description": "Prompt Info"
        },
        {
            "author": "chrisgoringe",
            "title": "Comfy Controller",
            "id": "cg-comfycontroller",
            "reference": "https://github.com/chrisgoringe/cg-controller",
            "files": [
                "https://github.com/chrisgoringe/cg-controller"
            ],
            "install_type": "git-clone",
            "description": "A simple controller panel that gathers all widgets from red nodes"
        },
        {
            "author": "TGu-97",
            "title": "TGu Utilities",
            "id": "tgu",
            "reference": "https://github.com/TGu-97/ComfyUI-TGu-utils",
            "files": [
                "https://github.com/TGu-97/ComfyUI-TGu-utils"
            ],
            "install_type": "git-clone",
            "description": "Nodes: MPN Switch, MPN Reroute, PN Switch. This is a set of custom nodes for ComfyUI. Mainly focus on control switches."
        },
        {
            "author": "seanlynch",
            "title": "SRL's nodes",
            "id": "srl",
            "reference": "https://github.com/seanlynch/srl-nodes",
            "files": [
                "https://github.com/seanlynch/srl-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: SRL Conditional Interrupt, SRL Format String, SRL Eval, SRL Filter Image List. This is a collection of nodes I find useful. Note that at least one module allows execution of arbitrary code. Do not use any of these nodes on a system that allow untrusted users to control workflows or inputs.[w/WARNING: The custom nodes in this extension are vulnerable to **security risks** because they allow the execution of arbitrary code through the workflow]"
        },
        {
            "author": "alpertunga-bile",
            "title": "prompt-generator",
            "reference": "https://github.com/alpertunga-bile/prompt-generator-comfyui",
            "files": [
                "https://github.com/alpertunga-bile/prompt-generator-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Custom AI prompt generator node for ComfyUI."
        },
        {
            "author": "kijai",
            "title": "KJNodes for ComfyUI",
            "id": "kjnodes",
            "reference": "https://github.com/kijai/ComfyUI-KJNodes",
            "files": [
                "https://github.com/kijai/ComfyUI-KJNodes"
            ],
            "install_type": "git-clone",
            "description": "Various quality of life -nodes for ComfyUI, mostly just visual stuff to improve usability."
        },
        {
            "author": "kijai",
            "title": "ComfyUI-CCSR",
            "id": "ccsr",
            "reference": "https://github.com/kijai/ComfyUI-CCSR",
            "files": [
                "https://github.com/kijai/ComfyUI-CCSR"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI- CCSR upscaler node"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-SVD",
            "id": "kijai-svd",
            "reference": "https://github.com/kijai/ComfyUI-SVD",
            "files": [
                "https://github.com/kijai/ComfyUI-SVD"
            ],
            "install_type": "git-clone",
            "description": "Preliminary use of SVD in ComfyUI.\nNOTE: Quick Implementation, Unstable. See details on repositories."
        },
        {
            "author": "kijai",
            "title": "Marigold depth estimation in ComfyUI",
            "id": "marigold",
            "reference": "https://github.com/kijai/ComfyUI-Marigold",
            "files": [
                "https://github.com/kijai/ComfyUI-Marigold"
            ],
            "install_type": "git-clone",
            "description": "This is a wrapper node for Marigold depth estimation: [https://mirror.ghproxy.com/https://github.com/prs-eth/Marigold](https://mirror.ghproxy.com/https://github.com/kijai/ComfyUI-Marigold). Currently using the same diffusers pipeline as in the original implementation, so in addition to the custom node, you need the model in diffusers format.\nNOTE: See details in repo to install."
        },
        {
            "author": "kijai",
            "title": "Geowizard depth and normal estimation in ComfyUI",
            "id": "geowizard",
            "reference": "https://github.com/kijai/ComfyUI-Geowizard",
            "files": [
                "https://github.com/kijai/ComfyUI-Geowizard"
            ],
            "install_type": "git-clone",
            "description": "This is a diffusers (0.27.2) wrapper node for Geowizard: [https://mirror.ghproxy.com/https://github.com/fuxiao0719/GeoWizard]. The model is autodownloaded from Hugginface to ComfyUI/models/diffusers/geowizard"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-depth-fm",
            "id": "depth-fm",
            "reference": "https://github.com/kijai/ComfyUI-depth-fm",
            "files": [
                "https://github.com/kijai/ComfyUI-depth-fm"
            ],
            "install_type": "git-clone",
            "description": "Fast and accurate monocular depth estimation."
        },
        {
            "author": "kijai",
            "title": "ComfyUI-DDColor",
            "id": "ddcolor-kijai",
            "reference": "https://github.com/kijai/ComfyUI-DDColor",
            "files": [
                "https://github.com/kijai/ComfyUI-DDColor"
            ],
            "install_type": "git-clone",
            "description": "Node to use [a/DDColor](https://mirror.ghproxy.com/https://github.com/piddnad/DDColor) in ComfyUI."
        },
        {
            "author": "kijai",
            "title": "Animatediff MotionLoRA Trainer",
            "id": "motionlora-trainer",
            "reference": "https://github.com/kijai/ComfyUI-ADMotionDirector",
            "files": [
                "https://github.com/kijai/ComfyUI-ADMotionDirector"
            ],
            "install_type": "git-clone",
            "description": "This is a trainer for AnimateDiff MotionLoRAs, based on the implementation of MotionDirector by ExponentialML.\nNOTE:[a/ADMotionDirector](https://mirror.ghproxy.com/https://github.com/ExponentialML/AnimateDiff-MotionDirector)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-moondream",
            "id": "moondream",
            "reference": "https://github.com/kijai/ComfyUI-moondream",
            "files": [
                "https://github.com/kijai/ComfyUI-moondream"
            ],
            "install_type": "git-clone",
            "description": "Moondream image to text query node with batch support"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-SUPIR",
            "id": "supir",
            "reference": "https://github.com/kijai/ComfyUI-SUPIR",
            "files": [
                "https://github.com/kijai/ComfyUI-SUPIR"
            ],
            "install_type": "git-clone",
            "description": "Wrapper nodes to use SUPIR upscaling process in ComfyUI"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-DynamiCrafterWrapper",
            "id": "dynamicrafter-kijai",
            "reference": "https://github.com/kijai/ComfyUI-DynamiCrafterWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-DynamiCrafterWrapper"
            ],
            "install_type": "git-clone",
            "description": "Wrapper nodes to use DynamiCrafter image2video and frame interpolation models in ComfyUI\nAnd this extension supports ToonCrafter as well"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-APISR",
            "id": "apisr",
            "reference": "https://github.com/kijai/ComfyUI-APISR-KJ",
            "files": [
                "https://github.com/kijai/ComfyUI-APISR-KJ"
            ],
            "install_type": "git-clone",
            "description": "Node to use [a/APISR](https://mirror.ghproxy.com/https://github.com/Kiteretsu77/APISR) upscale models in ComfyUI.[w/NOTE: repo name is changed from ComfyUI-APISR -> ComfyUI-APISR-KJ]"
        },
        {
            "author": "kijai",
            "title": "DiffusionLight implementation for ComfyUI",
            "id": "diffusionlight",
            "reference": "https://github.com/kijai/ComfyUI-DiffusionLight",
            "files": [
                "https://github.com/kijai/ComfyUI-DiffusionLight"
            ],
            "install_type": "git-clone",
            "description": "This is simplified implementation of the [a/DiffusionLight](https://mirror.ghproxy.com/https://github.com/DiffusionLight/DiffusionLight) method of creating light probes. You will need the included LoRA, place it in ComfyUI/loras folder like usual, it's converted from the original diffusers one."
        },
        {
            "author": "kijai",
            "title": "ComfyUI-ELLA-wrapper",
            "reference": "https://github.com/kijai/ComfyUI-ELLA-wrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-ELLA-wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI wrapper nodes to use the Diffusers implementation of ELLA"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-LaVi-Bridge-Wrapper",
            "reference": "https://github.com/kijai/ComfyUI-LaVi-Bridge-Wrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-LaVi-Bridge-Wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI wrapper node to test LaVi-Bridge using Diffusers"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-BrushNet-Wrapper",
            "reference": "https://github.com/kijai/ComfyUI-BrushNet-Wrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-BrushNet-Wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI wrapper nodes to use the Diffusers implementation of BrushNet"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-IC-Light",
            "id": "ic-light-kijai",
            "reference": "https://github.com/kijai/ComfyUI-IC-Light",
            "files": [
                "https://github.com/kijai/ComfyUI-IC-Light"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI native nodes for IC-Light"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-DepthAnythingV2",
            "id": "depth-anything-v2",
            "reference": "https://github.com/kijai/ComfyUI-DepthAnythingV2",
            "files": [
                "https://github.com/kijai/ComfyUI-DepthAnythingV2"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/DepthAnythingV2](https://depth-anything-v2.github.io/)\nNOTE:Models autodownload to ComfyUI/models/depthanything from [a/https://huggingface.co/Kijai/DepthAnythingV2-safetensors/tree/main](https://huggingface.co/Kijai/DepthAnythingV2-safetensors/tree/main)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-Florence2",
            "id": "florence2-kijai",
            "reference": "https://github.com/kijai/ComfyUI-Florence2",
            "files": [
                "https://github.com/kijai/ComfyUI-Florence2"
            ],
            "preemptions":[
                "DownloadAndLoadFlorence2Lora",
                "DownloadAndLoadFlorence2Model",
                "Florence2ModelLoader",
                "Florence2Run"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use Florence2 VLM for image vision tasks: object detection, captioning, segmentation and ocr"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-LuminaWrapper",
            "id": "lumina",
            "reference": "https://github.com/kijai/ComfyUI-LuminaWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-LuminaWrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI wrapper nodes for Lumina models"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-MimicMotionWrapper",
            "id": "mimicmotion-kijai",
            "reference": "https://github.com/kijai/ComfyUI-MimicMotionWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-MimicMotionWrapper"
            ],
            "install_type": "git-clone",
            "description": "Optimized wrapper nodes for MimicMotion: [a/https://mirror.ghproxy.com/https://github.com/tencent/MimicMotion](https://mirror.ghproxy.com/https://github.com/tencent/MimicMotion)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-OpenDiTWrapper",
            "id": "opendit-kijai",
            "reference": "https://github.com/kijai/ComfyUI-OpenDiTWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-OpenDiTWrapper"
            ],
            "install_type": "git-clone",
            "description": "Wrapper nodes for OpenDiT: [a/OpenDiT](https://mirror.ghproxy.com/https://github.com/NUS-HPC-AI-Lab/OpenDiT/), supports Open-Sora t2i and i2i"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-LivePortraitKJ",
            "id": "liveportrait-kijai",
            "reference": "https://github.com/kijai/ComfyUI-LivePortraitKJ",
            "files": [
                "https://github.com/kijai/ComfyUI-LivePortraitKJ"
            ],
            "install_type": "git-clone",
            "description": "Nodes for [a/LivePortrait](https://mirror.ghproxy.com/https://github.com/KwaiVGI/LivePortrait)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-KwaiKolorsWrapper",
            "id": "kwaikolors",
            "reference": "https://github.com/kijai/ComfyUI-KwaiKolorsWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-KwaiKolorsWrapper"
            ],
            "install_type": "git-clone",
            "description": "Rudimentary wrapper that runs [a/Kwai-Kolors](https://huggingface.co/Kwai-Kolors/Kolors) text2image pipeline using diffusers."
        },
        {
            "author": "kijai",
            "title": "ComfyUI-segment-anything-2",
            "id": "segment-anything-2",
            "reference": "https://github.com/kijai/ComfyUI-segment-anything-2",
            "files": [
                "https://github.com/kijai/ComfyUI-segment-anything-2"
            ],
            "preemptions":[
                "DownloadAndLoadSAM2Model",
                "Florence2toCoordinates",
                "Sam2AutoSegmentation",
                "Sam2Segmentation",
                "Sam2VideoSegmentation",
                "Sam2VideoSegmentationAddPoints"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use [a/segment-anything-2](https://mirror.ghproxy.com/https://github.com/facebookresearch/segment-anything-2) for image or video segmentation."
        },
        {
            "author": "kijai",
            "title": "ComfyUI nodes for ControlNext-SVD v2",
            "reference": "https://github.com/kijai/ComfyUI-ControlNeXt-SVD",
            "files": [
                "https://github.com/kijai/ComfyUI-ControlNeXt-SVD"
            ],
            "install_type": "git-clone",
            "description": "These nodes include my wrapper for the original diffusers pipeline, as well as work in progress native ComfyUI implementation.\nFor the diffusers wrapper models should be downloaded automatically, for the native version you can get the unet [a/here](https://huggingface.co/Kijai/ControlNeXt-SVD-V2-Comfy/blob/main/controlnext-svd_v2-unet-fp16_converted.safetensors)."
        },
        {
            "author": "kijai",
            "title": "ComfyUI Flux Trainer",
            "reference": "https://github.com/kijai/ComfyUI-FluxTrainer",
            "files": [
                "https://github.com/kijai/ComfyUI-FluxTrainer"
            ],
            "install_type": "git-clone",
            "description": "Currently supports LoRA training, and untested full finetune with code from kohya's scripts: [a/https://mirror.ghproxy.com/https://github.com/kohya-ss/sd-scripts](https://mirror.ghproxy.com/https://github.com/kohya-ss/sd-scripts)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI CogVideoX Wrapper",
            "reference": "https://github.com/kijai/ComfyUI-CogVideoXWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-CogVideoXWrapper"
            ],
            "install_type": "git-clone",
            "description": "Diffusers wrapper for CogVideoX -models: [a/https://mirror.ghproxy.com/https://github.com/THUDM/CogVideo](https://mirror.ghproxy.com/https://github.com/THUDM/CogVideo)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI PyramidFlow Wrapper",
            "reference": "https://github.com/kijai/ComfyUI-PyramidFlowWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-PyramidFlowWrapper"
            ],
            "install_type": "git-clone",
            "description": "Wrapper for PyramidFlow -models: [a/https://mirror.ghproxy.com/https://github.com/jy0205/Pyramid-Flow](https://mirror.ghproxy.com/https://github.com/jy0205/Pyramid-Flow)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI Llava-OneVision",
            "reference": "https://github.com/kijai/ComfyUI-LLaVA-OneVision",
            "files": [
                "https://github.com/kijai/ComfyUI-LLaVA-OneVision"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use the OneVision LLaVA models: [a/https://mirror.ghproxy.com/https://github.com/LLaVA-VL/LLaVA-NeXT](https://mirror.ghproxy.com/https://github.com/LLaVA-VL/LLaVA-NeXT)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI wrapper nodes for LVCD",
            "reference": "https://github.com/kijai/ComfyUI-LVCDWrapper",
            "files": [
                "https://github.com/kijai/ComfyUI-LVCDWrapper"
            ],
            "install_type": "git-clone",
            "description": "Original repo: [a/https://mirror.ghproxy.com/https://github.com/luckyhzt/LVCD](https://mirror.ghproxy.com/https://github.com/luckyhzt/LVCD)"
        },
        {
            "author": "kijai",
            "title": "ComfyUI-Lotus",
            "reference": "https://github.com/kijai/ComfyUI-Lotus",
            "files": [
                "https://github.com/kijai/ComfyUI-Lotus"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use Lotus depth/normal prediction.\nNOTE:The necessary models can be downloaded from ComfyUI-Manager."
        },
        {
            "author": "kijai",
            "title": "ComfyUI-MoGe",
            "reference": "https://github.com/kijai/ComfyUI-MoGe",
            "files": [
                "https://github.com/kijai/ComfyUI-MoGe"
            ],
            "install_type": "git-clone",
            "description": "NODES:(Down)load MoGe Model, MoGe Process"
        },
        {
            "author": "hhhzzyang",
            "title": "Comfyui-Lama",
            "id": "lama",
            "reference": "https://github.com/hhhzzyang/Comfyui_Lama",
            "files": [
                "https://github.com/hhhzzyang/Comfyui_Lama"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LamaaModelLoad, LamaApply, YamlConfigLoader. a costumer node is realized to remove anything/inpainting anything from a picture by mask inpainting.[w/WARN:This extension includes the entire model, which can result in a very long initial installation time, and there may be some compatibility issues with older dependencies and ComfyUI.]"
        },
        {
            "author": "audioscavenger",
            "title": "Save Image Extended for ComfyUI",
            "id": "save-image-extended",
            "reference": "https://github.com/audioscavenger/save-image-extended-comfyui",
            "files": [
                "https://github.com/audioscavenger/save-image-extended-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Upgrade the Save File node: customize subfolders, file names with checkpoint names, or any sampler attribute your want! [w/NOTE: This node is a fork from @thedyze, since the [a/original repository](https://mirror.ghproxy.com/https://github.com/thedyze/save-image-extended-comfyui) is no longer maintained. Simply *uninstall* the original version and **REINSTALL** this one.]"
        },
        {
            "author": "SOELexicon",
            "title": "ComfyUI-LexTools",
            "id": "lextools",
            "reference": "https://github.com/SOELexicon/ComfyUI-LexTools",
            "files": [
                "https://github.com/SOELexicon/ComfyUI-LexTools"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-LexTools is a Python-based image processing and analysis toolkit that uses machine learning models for semantic image segmentation, image scoring, and image captioning."
        },
        {
            "author": "mikkel",
            "title": "ComfyUI - Text Overlay Plugin",
            "id": "textoverlay-mikkel",
            "reference": "https://github.com/mikkel/ComfyUI-text-overlay",
            "files": [
                "https://github.com/mikkel/ComfyUI-text-overlay"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI Text Overlay Plugin provides functionalities for superimposing text on images. Users can select different font types, set text size, choose color, and adjust the text's position on the image."
        },
        {
            "author": "avatechai",
            "title": "Avatar Graph",
            "id": "avatar-graph",
            "reference": "https://github.com/avatechai/avatar-graph-comfyui",
            "files": [
                "https://github.com/avatechai/avatar-graph-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Include nodes for sam + bpy operation, that allows workflow creations for generative 2d character rig."
        },
        {
            "author": "TRI3D-LC",
            "title": "tri3d-comfyui-nodes",
            "id": "tri3d",
            "reference": "https://github.com/TRI3D-LC/tri3d-comfyui-nodes",
            "files": [
                "https://github.com/TRI3D-LC/tri3d-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: tri3d-extract-hand, tri3d-fuzzification, tri3d-position-hands, tri3d-atr-parse."
        },
        {
            "author": "TRI3D-LC",
            "title": "ComfyUI-MiroBoard",
            "id": "miroboard",
            "reference": "https://github.com/TRI3D-LC/ComfyUI-MiroBoard",
            "files": [
                "https://github.com/TRI3D-LC/ComfyUI-MiroBoard"
            ],
            "install_type": "git-clone",
            "description": "Nodes: add-image-miro-board."
        },
        {
            "author": "storyicon",
            "title": "segment anything",
            "id": "sam",
            "reference": "https://github.com/storyicon/comfyui_segment_anything",
            "files": [
                "https://github.com/storyicon/comfyui_segment_anything"
            ],
            "install_type": "git-clone",
            "description": "Based on GroundingDino and SAM, use semantic strings to segment any element in an image. The comfyui version of sd-webui-segment-anything."
        },
        {
            "author": "storyicon",
            "title": "ComfyUI MuseV Evolved",
            "id": "musev-evolved",
            "reference": "https://github.com/storyicon/comfyui_musev_evolved",
            "files": [
                "https://github.com/storyicon/comfyui_musev_evolved"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MuseVImg2Vid (comfyui_musev_evolved)\nNOTE: Download [a/MuseV](https://huggingface.co/TMElyralab/MuseV) to ComfyUI/models/diffusers"
        },
        {
            "author": "a1lazydog",
            "title": "ComfyUI-AudioScheduler",
            "id": "audioscheduler",
            "reference": "https://github.com/a1lazydog/ComfyUI-AudioScheduler",
            "files": [
                "https://github.com/a1lazydog/ComfyUI-AudioScheduler"
            ],
            "install_type": "git-clone",
            "description": "Load mp3 files and use the audio nodes to power animations and prompt scheduling. Use with FizzNodes."
        },
        {
            "author": "whatbirdisthat",
            "title": "cyberdolphin",
            "reference": "https://github.com/whatbirdisthat/cyberdolphin",
            "files": [
                "https://github.com/whatbirdisthat/cyberdolphin"
            ],
            "install_type": "git-clone",
            "description": "Cyberdolphin Suite of ComfyUI nodes for wiring up things."
        },
        {
            "author": "chrish-slingshot",
            "title": "CrasH Utils",
            "id": "crash",
            "reference": "https://github.com/chrish-slingshot/CrasHUtils",
            "files": [
                "https://github.com/chrish-slingshot/CrasHUtils"
            ],
            "install_type": "git-clone",
            "description": "A mixture of effects and quality of life nodes. Nodes: ImageGlitcher (gives an image a cool glitchy effect), ColorStylizer (highlights a single color in an image), QueryLocalLLM (queries a local LLM API though oobabooga), SDXLReslution (resolution picker for the standard SDXL resolutions, the complete list), SDXLResolutionSplit (splits the SDXL resolution into width and height). "
        },
        {
            "author": "spinagon",
            "title": "ComfyUI-seam-carving",
            "id": "seamcarving",
            "reference": "https://github.com/spinagon/ComfyUI-seam-carving",
            "files": [
                "https://github.com/spinagon/ComfyUI-seam-carving"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Image Resize (seam carving). Seam carving (image resize) for ComfyUI. Based on [a/https://mirror.ghproxy.com/https://github.com/li-plus/seam-carving](https://mirror.ghproxy.com/https://github.com/li-plus/seam-carving). With seam carving algorithm, the image could be intelligently resized while keeping the important contents undistorted. The carving process could be further guided, so that an object could be removed from the image without apparent artifacts."
        },
        {
            "author": "YMC",
            "title": "ymc-node-suite-comfyui",
            "id": "ymc-suite",
            "reference": "https://github.com/YMC-GitHub/ymc-node-suite-comfyui",
            "files": [
                "https://github.com/YMC-GitHub/ymc-node-suite-comfyui"
            ],
            "install_type": "git-clone",
            "description": "ymc 's nodes for comfyui. This extension is composed of nodes that provide various utility features such as text, region, and I/O."
        },
        {
            "author": "YMC",
            "title": "ymc-node-as-x-type",
            "reference": "https://github.com/YMC-GitHub/ymc-node-as-x-type",
            "files": [
                "https://github.com/YMC-GitHub/ymc-node-as-x-type"
            ],
            "install_type": "git-clone",
            "description": "some comfyui custom nodes to set it as known type"
        },
        {
            "author": "chibiace",
            "title": "ComfyUI-Chibi-Nodes",
            "id": "chibi",
            "reference": "https://github.com/chibiace/ComfyUI-Chibi-Nodes",
            "files": [
                "https://github.com/chibiace/ComfyUI-Chibi-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Loader, Prompts, ImageTool, Wildcards, LoadEmbedding, ConditionText, SaveImages, ..."
        },
        {
            "author": "DigitalIO",
            "title": "ComfyUI-stable-wildcards",
            "id": "stable-wildcards",
            "reference": "https://github.com/DigitalIO/ComfyUI-stable-wildcards",
            "files": [
                "https://github.com/DigitalIO/ComfyUI-stable-wildcards"
            ],
            "install_type": "git-clone",
            "description": "Wildcard implementation that can be reproduced with workflows."
        },
        {
            "author": "THtianhao",
            "title": "ComfyUI-Portrait-Maker",
            "id": "portrait-maker",
            "reference": "https://github.com/THtianhao/ComfyUI-Portrait-Maker",
            "files": [
                "https://github.com/THtianhao/ComfyUI-Portrait-Maker"
            ],
            "install_type": "git-clone",
            "description": "Nodes:RetainFace, FaceFusion, RatioMerge2Image, MaskMerge2Image, ReplaceBoxImg, ExpandMaskBox, FaceSkin, SkinRetouching, PortraitEnhancement, ..."
        },
        {
            "author": "THtianhao",
            "title": "ComfyUI-FaceChain",
            "id": "facechain",
            "reference": "https://github.com/THtianhao/ComfyUI-FaceChain",
            "files": [
                "https://github.com/THtianhao/ComfyUI-FaceChain"
            ],
            "install_type": "git-clone",
            "description": "The official ComfyUI version of facechain greatly improves the speed of reasoning and has great custom process controls."
        },
        {
            "author": "zer0TF",
            "title": "Cute Comfy",
            "id": "cutecomfy",
            "reference": "https://github.com/zer0TF/cute-comfy",
            "files": [
                "https://github.com/zer0TF/cute-comfy"
            ],
            "install_type": "git-clone",
            "description": "Adds a configurable folder watcher that auto-converts Comfy metadata into a Civitai-friendly format for automatic resource tagging when you upload images. Oh, and it makes your UI awesome, too. 💜"
        },
        {
            "author": "chflame163",
            "title": "ComfyUI_MSSpeech_TTS",
            "id": "msspeech",
            "reference": "https://github.com/chflame163/ComfyUI_MSSpeech_TTS",
            "files": [
                "https://github.com/chflame163/ComfyUI_MSSpeech_TTS"
            ],
            "install_type": "git-clone",
            "description": "A text-to-speech plugin used under ComfyUI. It utilizes the Microsoft Speech TTS interface to convert text content into MP3 format audio files."
        },
        {
            "author": "chflame163",
            "title": "ComfyUI_WordCloud",
            "id": "wordcloud",
            "reference": "https://github.com/chflame163/ComfyUI_WordCloud",
            "files": [
                "https://github.com/chflame163/ComfyUI_WordCloud"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Word Cloud, Load Text File"
        },
        {
            "author": "chflame163",
            "title": "ComfyUI Layer Style",
            "id": "layerstyle",
            "reference": "https://github.com/chflame163/ComfyUI_LayerStyle",
            "files": [
                "https://github.com/chflame163/ComfyUI_LayerStyle"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes for ComfyUI it generate image like Adobe Photoshop's Layer Style. the Drop Shadow is first completed node, and follow-up work is in progress."
        },
        {
            "author": "chflame163",
            "title": "ComfyUI Face Similarity",
            "id": "face-similarity",
            "reference": "https://github.com/chflame163/ComfyUI_FaceSimilarity",
            "files": [
                "https://github.com/chflame163/ComfyUI_FaceSimilarity"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI. It compare two images to rate facial similarity."
        },
        {
            "author": "chflame163",
            "title": "ComfyUI_CatVTON_Wrapper",
            "id": "catvton-wrapper",
            "reference": "https://github.com/chflame163/ComfyUI_CatVTON_Wrapper",
            "files": [
                "https://github.com/chflame163/ComfyUI_CatVTON_Wrapper"
            ],
            "install_type": "git-clone",
            "description": "[a/CatVTON](https://mirror.ghproxy.com/https://github.com/Zheng-Chong/CatVTON) warpper for ComfyUI"
        },
        {
            "author": "drustan-hawk",
            "title": "primitive-types",
            "reference": "https://github.com/drustan-hawk/primitive-types",
            "files": [
                "https://github.com/drustan-hawk/primitive-types"
            ],
            "install_type": "git-clone",
            "description": "Small collection of typed primitive nodes."
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-mixlab-nodes",
            "id": "mixlab",
            "reference": "https://github.com/shadowcz007/comfyui-mixlab-nodes",
            "files": [
                "https://github.com/shadowcz007/comfyui-mixlab-nodes"
            ],
            "install_type": "git-clone",
            "description": "3D, ScreenShareNode & FloatingVideoNode, SpeechRecognition & SpeechSynthesis, GPT, LoadImagesFromLocal, Layers, Other Nodes, ..."
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-ultralytics-yolo",
            "id": "yolo",
            "reference": "https://github.com/shadowcz007/comfyui-ultralytics-yolo",
            "files": [
                "https://github.com/shadowcz007/comfyui-ultralytics-yolo"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Detect By Label."
        },
        {
            "author": "shadowcz007",
            "title": "Consistency Decoder",
            "id": "consistency-decoder",
            "reference": "https://github.com/shadowcz007/comfyui-consistency-decoder",
            "files": [
                "https://github.com/shadowcz007/comfyui-consistency-decoder"
            ],
            "install_type": "git-clone",
            "description": "[a/openai Consistency Decoder](https://mirror.ghproxy.com/https://github.com/openai/consistencydecoder). After downloading the [a/OpenAI VAE model](https://openaipublic.azureedge.net/diff-vae/c9cebd3132dd9c42936d803e33424145a748843c8f716c0814838bdc8a2fe7cb/decoder.pt), place it in the model/vae directory for use."
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-Image-reward",
            "reference": "https://github.com/shadowcz007/comfyui-Image-reward",
            "files": [
                "https://github.com/shadowcz007/comfyui-Image-reward"
            ],
            "install_type": "git-clone",
            "description": "[a/ImageReward](https://mirror.ghproxy.com/https://github.com/THUDM/ImageReward): Human preference learning in text-to-image generation. This is a [a/paper](https://arxiv.org/abs/2304.05977) from NeurIPS 2023"
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-sound-lab",
            "id": "soundlab",
            "reference": "https://github.com/shadowcz007/comfyui-sound-lab",
            "files": [
                "https://github.com/shadowcz007/comfyui-sound-lab"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Music Gen, Audio Play, Stable Audio"
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-edit-mask",
            "id": "edit-mask",
            "reference": "https://github.com/shadowcz007/comfyui-edit-mask",
            "files": [
                "https://github.com/shadowcz007/comfyui-edit-mask"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Edit Mask"
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-liveportrait",
            "id": "liveportrait",
            "reference": "https://github.com/shadowcz007/comfyui-liveportrait",
            "files": [
                "https://github.com/shadowcz007/comfyui-liveportrait"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI version of [a/LivePortrait](https://mirror.ghproxy.com/https://github.com/KwaiVGI/LivePortrait)."
        },
        {
            "author": "shadowcz007",
            "title": "comfyui-try-on",
            "reference": "https://github.com/shadowcz007/comfyui-try-on",
            "files": [
                "https://github.com/shadowcz007/comfyui-try-on"
            ],
            "install_type": "git-clone",
            "description": "Virtual try-on for creating a personal brand wardrobe collection."
        },
        {
            "author": "ostris",
            "title": "Ostris Nodes ComfyUI",
            "id": "ostris",
            "reference": "https://github.com/ostris/ostris_nodes_comfyui",
            "files": [
                "https://github.com/ostris/ostris_nodes_comfyui"
            ],
            "install_type": "git-clone",
            "nodename_pattern": "- Ostris$",
            "description": "This is a collection of custom nodes for ComfyUI that I made for some QOL. I will be adding much more advanced ones in the future once I get more familiar with the API."
        },
        {
            "author": "0xbitches",
            "title": "Latent Consistency Model for ComfyUI",
            "id": "lcm",
            "reference": "https://github.com/0xbitches/ComfyUI-LCM",
            "files": [
                "https://github.com/0xbitches/ComfyUI-LCM"
            ],
            "install_type": "git-clone",
            "description": "This custom node implements a Latent Consistency Model sampler in ComfyUI. (LCM)"
        },
        {
            "author": "aszc-dev",
            "title": "Core ML Suite for ComfyUI",
            "id": "coreml",
            "reference": "https://github.com/aszc-dev/ComfyUI-CoreMLSuite",
            "files": [
                "https://github.com/aszc-dev/ComfyUI-CoreMLSuite"
            ],
            "install_type": "git-clone",
            "description": "This extension contains a set of custom nodes for ComfyUI that allow you to use Core ML models in your ComfyUI workflows. The models can be obtained here, or you can convert your own models using coremltools. The main motivation behind using Core ML models in ComfyUI is to allow you to utilize the ANE (Apple Neural Engine) on Apple Silicon (M1/M2) machines to improve performance."
        },
        {
            "author": "taabata",
            "title": "Syrian Falcon Nodes",
            "id": "syrian",
            "reference": "https://github.com/taabata/Comfy_Syrian_Falcon_Nodes",
            "files": [
                "https://github.com/taabata/Comfy_Syrian_Falcon_Nodes/raw/main/SyrianFalconNodes.py"
            ],
            "install_type": "copy",
            "description": "Nodes:Prompt editing, Word as Image"
        },
        {
            "author": "taabata",
            "title": "LCM_Inpaint-Outpaint_Comfy",
            "id": "lcm-inpaint-outpaint",
            "reference": "https://github.com/taabata/LCM_Inpaint_Outpaint_Comfy",
            "files": [
                "https://github.com/taabata/LCM_Inpaint_Outpaint_Comfy"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom nodes for inpainting/outpainting using the new latent consistency model (LCM)"
        },
        {
            "author": "noxinias",
            "title": "ComfyUI_NoxinNodes",
            "id": "noxin",
            "reference": "https://github.com/noxinias/ComfyUI_NoxinNodes",
            "files": [
                "https://github.com/noxinias/ComfyUI_NoxinNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Noxin Complete Chime, Noxin Scaled Resolutions, Load from Noxin Prompt Library, Save to Noxin Prompt Library"
        },
        {
            "author": "apesplat",
            "title": "ezXY scripts and nodes",
            "id": "ezxy",
            "reference": "https://github.com/GMapeSplat/ComfyUI_ezXY",
            "files": [
                "https://github.com/GMapeSplat/ComfyUI_ezXY"
            ],
            "install_type": "git-clone",
            "description": "Extensions/Patches: Enables linking float and integer inputs and ouputs. Values are automatically cast to the correct type and clamped to the correct range. Works with both builtin and custom nodes.[w/NOTE: This repo patches ComfyUI's validate_inputs and map_node_over_list functions while running. May break depending on your version of ComfyUI. Can be deactivated in config.yaml.]Nodes: A collection of nodes for facilitating the generation of XY plots. Capable of plotting changes over most primitive values.[w/Does not work with current version of Comfyui]"
        },
        {
            "author": "kinfolk0117",
            "title": "SimpleTiles",
            "id": "simpletiles",
            "reference": "https://github.com/kinfolk0117/ComfyUI_SimpleTiles",
            "files": [
                "https://github.com/kinfolk0117/ComfyUI_SimpleTiles"
            ],
            "install_type": "git-clone",
            "description": "Nodes:TileSplit, TileMerge."
        },
        {
            "author": "kinfolk0117",
            "title": "ComfyUI_GradientDeepShrink",
            "id": "deepshrink",
            "reference": "https://github.com/kinfolk0117/ComfyUI_GradientDeepShrink",
            "files": [
                "https://github.com/kinfolk0117/ComfyUI_GradientDeepShrink"
            ],
            "install_type": "git-clone",
            "description": "Nodes:GradientPatchModelAddDownscale (Kohya Deep Shrink)."
        },
        {
            "author": "kinfolk0117",
            "title": "ComfyUI_Pilgram",
            "id": "pilgram",
            "reference": "https://github.com/kinfolk0117/ComfyUI_Pilgram",
            "files": [
                "https://github.com/kinfolk0117/ComfyUI_Pilgram"
            ],
            "install_type": "git-clone",
            "description": "Use [a/Pilgram2](https://mirror.ghproxy.com/https://github.com/mgineer85/pilgram2) filters in ComfyUI"
        },
        {
            "author": "kinfolk0117",
            "title": "Gridswapper",
            "reference": "https://github.com/kinfolk0117/ComfyUI_GridSwapper",
            "files": [
                "https://github.com/kinfolk0117/ComfyUI_GridSwapper"
            ],
            "install_type": "git-clone",
            "description": "Gridswapper takes a batch of latents and spreads them over the necessary amount of grids. It then automatically shuffles the images in the grids for each step. So, a batch of 12 latents for a 2x2 grid will generate 3 grid images in each step. It will then shuffle around the images for the next step. This makes it possible for all images to influence the others during the denoising process. This approach works well for generating 2-4 grids."
        },
        {
            "author": "Fictiverse",
            "title": "ComfyUI Fictiverse Nodes",
            "id": "fictverse",
            "reference": "https://github.com/Fictiverse/ComfyUI_Fictiverse",
            "files": [
                "https://github.com/Fictiverse/ComfyUI_Fictiverse"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Color correction."
        },
        {
            "author": "idrirap",
            "title": "ComfyUI-Lora-Auto-Trigger-Words",
            "id": "lora-auto-trigger",
            "reference": "https://github.com/idrirap/ComfyUI-Lora-Auto-Trigger-Words",
            "files": [
                "https://github.com/idrirap/ComfyUI-Lora-Auto-Trigger-Words"
            ],
            "install_type": "git-clone",
            "description": "This project is a fork of [a/https://mirror.ghproxy.com/https://github.com/Extraltodeus/LoadLoraWithTags](https://mirror.ghproxy.com/https://github.com/Extraltodeus/LoadLoraWithTags) The aim of these custom nodes is to get an easy access to the tags used to trigger a lora."
        },
        {
            "author": "aianimation55",
            "title": "Comfy UI FatLabels",
            "id": "fatlab",
            "reference": "https://github.com/aianimation55/ComfyUI-FatLabels",
            "files": [
                "https://github.com/aianimation55/ComfyUI-FatLabels"
            ],
            "install_type": "git-clone",
            "description": "It's a super simple custom node for Comfy UI, to generate text, with a font size option. Useful for bigger labelling of nodes, helpful for wider screen captures or tutorials. Plus you can of course use the text within your generations."
        },
        {
            "author": "noEmbryo",
            "title": "noEmbryo nodes",
            "id": "noembryo",
            "reference": "https://github.com/noembryo/ComfyUI-noEmbryo",
            "files": [
                "https://github.com/noembryo/ComfyUI-noEmbryo"
            ],
            "install_type": "git-clone",
            "description": "PromptTermList (1-6): are some nodes that help with the creation of Prompts inside ComfyUI. Resolution Scale outputs image dimensions using a scale factor. Regex Text Chopper outputs the chopped parts of a text using RegEx."
        },
        {
            "author": "mikkel",
            "title": "ComfyUI - Mask Bounding Box",
            "id": "mask-bbox",
            "reference": "https://github.com/mikkel/comfyui-mask-boundingbox",
            "files": [
                "https://github.com/mikkel/comfyui-mask-boundingbox"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI Mask Bounding Box Plugin provides functionalities for selecting a specific size mask from an image. Can be combined with ClipSEG to replace any aspect of an SDXL image with an SD1.5 output."
        },
        {
            "author": "ParmanBabra",
            "title": "ComfyUI-Malefish-Custom-Scripts",
            "id": "malefish",
            "reference": "https://github.com/ParmanBabra/ComfyUI-Malefish-Custom-Scripts",
            "files": [
                "https://github.com/ParmanBabra/ComfyUI-Malefish-Custom-Scripts"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Multi Lora Loader, Random (Prompt), Combine (Prompt), CSV Prompts Loader"
        },
        {
            "author": "IAmMatan.com",
            "title": "ComfyUI Serving toolkit",
            "id": "serving-toolkit",
            "reference": "https://github.com/matan1905/ComfyUI-Serving-Toolkit",
            "files": [
                "https://github.com/matan1905/ComfyUI-Serving-Toolkit"
            ],
            "install_type": "git-clone",
            "description": "This extension adds nodes that allow you to easily serve your workflow (for example using a discord bot) "
        },
        {
            "author": "PCMonsterx",
            "title": "ComfyUI-CSV-Loader",
            "id": "csv-loader",
            "reference": "https://github.com/PCMonsterx/ComfyUI-CSV-Loader",
            "files": [
                "https://github.com/PCMonsterx/ComfyUI-CSV-Loader"
            ],
            "install_type": "git-clone",
            "description": "CSV Loader for prompt building within ComfyUI interface. Allows access to positive/negative prompts associated with a name. Selections are being pulled from CSV files."
        },
        {
            "author": "Trung0246",
            "title": "ComfyUI-0246",
            "id": "0246",
            "reference": "https://github.com/Trung0246/ComfyUI-0246",
            "files": [
                "https://github.com/Trung0246/ComfyUI-0246"
            ],
            "install_type": "git-clone",
            "description": "Random nodes for ComfyUI I made to solve my struggle with ComfyUI (ex: pipe, process). Have varying quality."
        },
        {
            "author": "fexli",
            "title": "fexli-util-node-comfyui",
            "id": "fexli",
            "reference": "https://github.com/fexli/fexli-util-node-comfyui",
            "files": [
                "https://github.com/fexli/fexli-util-node-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FEImagePadForOutpaint, FEColorOut, FEColor2Image, FERandomizedColor2Image"
        },
        {
            "author": "AbyssBadger0",
            "title": "ComfyUI_BadgerTools",
            "id": "badger",
            "reference": "https://github.com/AbyssBadger0/ComfyUI_BadgerTools",
            "files": [
                "https://github.com/AbyssBadger0/ComfyUI_BadgerTools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageOverlap-badger, FloatToInt-badger, IntToString-badger, FloatToString-badger, ImageNormalization-badger, ImageScaleToSide-badger, NovelToFizz-badger."
        },
        {
            "author": "palant",
            "title": "Image Resize for ComfyUI",
            "id": "image-resize",
            "reference": "https://github.com/palant/image-resize-comfyui",
            "files": [
                "https://github.com/palant/image-resize-comfyui"
            ],
            "install_type": "git-clone",
            "description": "This custom node provides various tools for resizing images. The goal is resizing without distorting proportions, yet without having to perform any calculations with the size of the original image. If a mask is present, it is resized and modified along with the image."
        },
        {
            "author": "palant",
            "title": "Integrated Nodes for ComfyUI",
            "reference": "https://github.com/palant/integrated-nodes-comfyui",
            "files": [
                "https://github.com/palant/integrated-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "This tool will turn entire workflows or parts of them into single integrated nodes. In a way, it is similar to the Node Templates functionality but hides the inner structure. This is useful if all you want is to reuse and quickly configure a bunch of nodes without caring how they are interconnected."
        },
        {
            "author": "palant",
            "title": "Extended Save Image for ComfyUI",
            "reference": "https://github.com/palant/extended-saveimage-comfyui",
            "files": [
                "https://github.com/palant/extended-saveimage-comfyui"
            ],
            "install_type": "git-clone",
            "description": "This custom node is largely identical to the usual Save Image but allows saving images also in JPEG and WEBP formats, the latter with both lossless and lossy compression. Metadata is embedded in the images as usual, and the resulting images can be used to load a workflow."
        },
        {
            "author": "whmc76",
            "title": "ComfyUI-Openpose-Editor-Plus",
            "id": "openpose-editor-plus",
            "reference": "https://github.com/whmc76/ComfyUI-Openpose-Editor-Plus",
            "files": [
                "https://github.com/whmc76/ComfyUI-Openpose-Editor-Plus"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Openpose Editor Plus"
        },
        {
            "author": "martijnat",
            "title": "comfyui-previewlatent",
            "reference": "https://github.com/martijnat/comfyui-previewlatent",
            "files": [
                "https://github.com/martijnat/comfyui-previewlatent"
            ],
            "install_type": "git-clone",
            "description": "a ComfyUI plugin for previewing latents without vae decoding. Useful for showing intermediate results and can be used a faster 'preview image' if you don't wan't to use vae decode."
        },
        {
            "author": "banodoco",
            "title": "Steerable Motion",
            "id": "steerable-motion",
            "reference": "https://github.com/banodoco/steerable-motion",
            "files": [
                "https://github.com/banodoco/steerable-motion"
            ],
            "install_type": "git-clone",
            "description": "Steerable Motion is a ComfyUI node for batch creative interpolation. Our goal is to feature the best methods for steering motion with images as video models evolve."
        },
        {
            "author": "gemell1",
            "title": "ComfyUI_GMIC",
            "id": "gmic",
            "reference": "https://github.com/gemell1/ComfyUI_GMIC",
            "files": [
                "https://github.com/gemell1/ComfyUI_GMIC"
            ],
            "install_type": "git-clone",
            "description": "Nodes:GMIC Image Processing."
        },
        {
            "author": "LonicaMewinsky",
            "title": "ComfyBreakAnim",
            "id": "breakanim",
            "reference": "https://github.com/LonicaMewinsky/ComfyUI-MakeFrame",
            "files": [
                "https://github.com/LonicaMewinsky/ComfyUI-MakeFrame"
            ],
            "install_type": "git-clone",
            "description": "Nodes:BreakFrames, GetKeyFrames, MakeGrid."
        },
        {
            "author": "TheBarret",
            "title": "ZSuite",
            "id": "zsuite",
            "reference": "https://github.com/TheBarret/ZSuite",
            "files": [
                "https://github.com/TheBarret/ZSuite"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Prompter, RF Noise, SeedMod."
        },
        {
            "author": "romeobuilderotti",
            "title": "ComfyUI PNG Metadata",
            "id": "pngmeta",
            "reference": "https://github.com/romeobuilderotti/ComfyUI-PNG-Metadata",
            "files": [
                "https://github.com/romeobuilderotti/ComfyUI-PNG-Metadata"
            ],
            "install_type": "git-clone",
            "description": "Add custom Metadata fields to your saved PNG files."
        },
        {
            "author": "ka-puna",
            "title": "comfyui-yanc",
            "id": "yanc",
            "reference": "https://github.com/ka-puna/comfyui-yanc",
            "files": [
                "https://github.com/ka-puna/comfyui-yanc"
            ],
            "install_type": "git-clone",
            "description": "NOTE: Concatenate Strings, Format Datetime String, Integer Caster, Multiline String, Truncate String. Yet Another Node Collection, a repository of simple nodes for ComfyUI. This repository eases the addition or removal of custom nodes to itself."
        },
        {
            "author": "amorano",
            "title": "Jovimetrix Composition Nodes",
            "id": "jovimetrix",
            "reference": "https://github.com/Amorano/Jovimetrix",
            "files": [
                "https://github.com/Amorano/Jovimetrix"
            ],
            "nodename_pattern": " \\(JOV\\)$",
            "install_type": "git-clone",
            "description": "Webcam, MIDI, Spout and GLSL shader support. Animation via tick. Parameter manipulation with wave generator. Math operations, universal value converstion, shape mask generation, image channel ops, batch splitting/merging/randomizing, load image/video from URL, Dynamic bus routing, support for GIPHY, save output anywhere! flatten, crop, transform; check color blindness, make stereograms or stereoscopic images, and much more."
        },
        {
            "author": "Umikaze-job",
            "title": "select_folder_path_easy",
            "reference": "https://github.com/Umikaze-job/select_folder_path_easy",
            "files": [
                "https://github.com/Umikaze-job/select_folder_path_easy"
            ],
            "install_type": "git-clone",
            "description": "This extension simply connects the nodes and specifies the output path of the generated images to a manageable path."
        },
        {
            "author": "Niutonian",
            "title": "ComfyUi-NoodleWebcam",
            "id": "noodle-webcam",
            "reference": "https://github.com/Niutonian/ComfyUi-NoodleWebcam",
            "files": [
                "https://github.com/Niutonian/ComfyUi-NoodleWebcam"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Noodle webcam is a node that records frames and send them to your favourite node."
        },
        {
            "author": "Feidorian",
            "title": "feidorian-ComfyNodes",
            "id": "feidorian",
            "reference": "https://github.com/Feidorian/feidorian-ComfyNodes",
            "nodename_pattern": "^Feidorian_",
            "files": [
                "https://github.com/Feidorian/feidorian-ComfyNodes"
            ],
            "install_type": "git-clone",
            "description": "This extension provides various custom nodes. literals, loaders, logic, output, switches"
        },
        {
            "author": "wutipong",
            "title": "ComfyUI-TextUtils",
            "reference": "https://github.com/wutipong/ComfyUI-TextUtils",
            "files": [
                "https://github.com/wutipong/ComfyUI-TextUtils"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Create N-Token String"
        },
        {
            "author": "natto-maki",
            "title": "ComfyUI-NegiTools",
            "id": "negitools",
            "reference": "https://github.com/natto-maki/ComfyUI-NegiTools",
            "files": [
                "https://github.com/natto-maki/ComfyUI-NegiTools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:OpenAI DALLe3, OpenAI Translate to English, String Function, Seed Generator"
        },
        {
            "author": "LonicaMewinsky",
            "title": "ComfyUI-RawSaver",
            "id": "rawsaver",
            "reference": "https://github.com/LonicaMewinsky/ComfyUI-RawSaver",
            "files": [
                "https://github.com/LonicaMewinsky/ComfyUI-RawSaver"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SaveTifImage. ComfyUI custom node for purpose of saving image as uint16 tif file."
        },
        {
            "author": "jojkaart",
            "title": "ComfyUI-sampler-lcm-alternative",
            "id": "lmc-alt",
            "reference": "https://github.com/jojkaart/ComfyUI-sampler-lcm-alternative",
            "files": [
                "https://github.com/jojkaart/ComfyUI-sampler-lcm-alternative"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LCMScheduler, SamplerLCMAlternative, SamplerLCMCycle. ComfyUI Custom Sampler nodes that add a new improved LCM sampler functions"
        },
        {
            "author": "GTSuya-Studio",
            "title": "ComfyUI-GTSuya-Nodes",
            "id": "gtsuya",
            "reference": "https://github.com/GTSuya-Studio/ComfyUI-Gtsuya-Nodes",
            "files": [
                "https://github.com/GTSuya-Studio/ComfyUI-Gtsuya-Nodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-GTSuya-Nodes is a ComfyUI extension designed to add several wildcards supports into ComfyUI. Wildcards allow you to use __name__ syntax in your prompt to get a random line from a file named name.txt in a wildcards directory."
        },
        {
            "author": "oyvindg",
            "title": "ComfyUI-TrollSuite",
            "id": "troll",
            "reference": "https://github.com/oyvindg/ComfyUI-TrollSuite",
            "files": [
                "https://github.com/oyvindg/ComfyUI-TrollSuite"
            ],
            "install_type": "git-clone",
            "description": "Nodes: BinaryImageMask, ImagePadding, LoadLastCreatedImage, RandomMask, TransparentImage."
        },
        {
            "author": "drago87",
            "title": "ComfyUI_Dragos_Nodes",
            "id": "dragos",
            "reference": "https://github.com/drago87/ComfyUI_Dragos_Nodes",
            "files": [
                "https://github.com/drago87/ComfyUI_Dragos_Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:File Padding, Image Info, VAE Loader With Name"
        },
        {
            "author": "bronkula",
            "title": "comfyui-fitsize",
            "id": "fitsize",
            "reference": "https://github.com/bronkula/comfyui-fitsize",
            "files": [
                "https://github.com/bronkula/comfyui-fitsize"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Fit Size From Int/Image/Resize, Load Image And Resize To Fit, Pick Image From Batch/List, Crop Image Into Even Pieces, Image Region To Mask... A simple set of nodes for making an image fit within a bounding box"
        },
        {
            "author": "toyxyz",
            "title": "ComfyUI_toyxyz_test_nodes",
            "id": "toyxyz",
            "reference": "https://github.com/toyxyz/ComfyUI_toyxyz_test_nodes",
            "files": [
                "https://github.com/toyxyz/ComfyUI_toyxyz_test_nodes"
            ],
            "install_type": "git-clone",
            "description": "This node was created to send a webcam to ComfyUI in real time. This node is recommended for use with LCM."
        },
        {
            "author": "thecooltechguy",
            "title": "ComfyUI Stable Video Diffusion",
            "reference": "https://github.com/thecooltechguy/ComfyUI-Stable-Video-Diffusion",
            "files": [
                "https://github.com/thecooltechguy/ComfyUI-Stable-Video-Diffusion"
            ],
            "install_type": "git-clone",
            "description": "Easily use Stable Video Diffusion inside ComfyUI!"
        },
        {
            "author": "thecooltechguy",
            "title": "ComfyUI-ComfyRun",
            "reference": "https://github.com/thecooltechguy/ComfyUI-ComfyRun",
            "files": [
                "https://github.com/thecooltechguy/ComfyUI-ComfyRun"
            ],
            "install_type": "git-clone",
            "description": "The easiest way to run & share any ComfyUI workflow [a/https://comfyrun.com](https://comfyrun.com)"
        },
        {
            "author": "thecooltechguy",
            "title": "ComfyUI-MagicAnimate",
            "reference": "https://github.com/thecooltechguy/ComfyUI-MagicAnimate",
            "files": [
                "https://github.com/thecooltechguy/ComfyUI-MagicAnimate"
            ],
            "install_type": "git-clone",
            "description": "Easily use Magic Animate within ComfyUI!\n[w/WARN: This extension requires 15GB disk space.]"
        },
        {
            "author": "thecooltechguy",
            "title": "ComfyUI-ComfyWorkflows",
            "reference": "https://github.com/thecooltechguy/ComfyUI-ComfyWorkflows",
            "files": [
                "https://github.com/thecooltechguy/ComfyUI-ComfyWorkflows"
            ],
            "install_type": "git-clone",
            "description": "The best way to run, share, & discover thousands of ComfyUI workflows."
        },
        {
            "author": "Danand",
            "title": "Comfy Couple",
            "reference": "https://github.com/Danand/ComfyUI-ComfyCouple",
            "files": [
                "https://github.com/Danand/ComfyUI-ComfyCouple"
            ],
            "install_type": "git-clone",
            "description": " If you want to draw two different characters together without blending their features, so you could try to check out this custom node."
        },
        {
            "author": "42lux",
            "title": "ComfyUI-safety-checker",
            "reference": "https://github.com/42lux/ComfyUI-safety-checker",
            "files": [
                "https://github.com/42lux/ComfyUI-safety-checker"
            ],
            "install_type": "git-clone",
            "description": "A NSFW/Safety Checker Node for ComfyUI."
        },
        {
            "author": "sergekatzmann",
            "title": "ComfyUI_Nimbus-Pack",
            "reference": "https://github.com/sergekatzmann/ComfyUI_Nimbus-Pack",
            "files": [
                "https://github.com/sergekatzmann/ComfyUI_Nimbus-Pack"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Square Adapter Node, Image Resize And Crop Node"
        },
        {
            "author": "komojini",
            "title": "ComfyUI_SDXL_DreamBooth_LoRA_CustomNodes",
            "reference": "https://github.com/komojini/ComfyUI_SDXL_DreamBooth_LoRA_CustomNodes",
            "files": [
                "https://github.com/komojini/ComfyUI_SDXL_DreamBooth_LoRA_CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:XL DreamBooth LoRA, S3 Bucket LoRA"
        },
        {
            "author": "komojini",
            "title": "komojini-comfyui-nodes",
            "id": "komojini-nodes",
            "reference": "https://github.com/komojini/komojini-comfyui-nodes",
            "files": [
                "https://github.com/komojini/komojini-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:YouTube Video Loader. Custom ComfyUI Nodes for video generation"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "APISR IN COMFYUI",
            "id": "apisr-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-APISR",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-APISR"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of APISR for ComfyUI, both image and video"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-Text_Image-Composite [WIP]",
            "id": "txtimg-composite",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Text_Image-Composite",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Text_Image-Composite"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Text_Image_Zho, Text_Image_Multiline_Zho, RGB_Image_Zho, AlphaChanelAddByMask, ImageComposite_Zho, ..."
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-Gemini",
            "id": "gemini",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Gemini",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Gemini"
            ],
            "install_type": "git-clone",
            "description": "Using Gemini-pro & Gemini-pro-vision in ComfyUI."
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "comfyui-portrait-master-zh-cn",
            "id": "portrait-master-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/comfyui-portrait-master-zh-cn",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/comfyui-portrait-master-zh-cn"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Portrait Master 简体中文版."
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-Q-Align",
            "id": "qalign-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Q-Align",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Q-Align"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Q-Align Scoring. Implementation of [a/Q-Align](https://arxiv.org/abs/2312.17090) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-InstantID",
            "id": "instantid-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-InstantID",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-InstantID"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/InstantID](https://mirror.ghproxy.com/https://github.com/InstantID/InstantID) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI PhotoMaker (ZHO)",
            "id": "photomaker-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-PhotoMaker-ZHO",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-PhotoMaker-ZHO"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/PhotoMaker](https://mirror.ghproxy.com/https://github.com/TencentARC/PhotoMaker) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-Qwen-VL-API",
            "id": "qwen-vl-api",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Qwen-VL-API",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Qwen-VL-API"
            ],
            "install_type": "git-clone",
            "description": "QWen-VL-Plus & QWen-VL-Max in ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-SVD-ZHO (WIP)",
            "id": "svd-zho",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-SVD-ZHO",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-SVD-ZHO"
            ],
            "install_type": "git-clone",
            "description": "My Workflows + Auxiliary nodes for Stable Video Diffusion (SVD)"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI SegMoE",
            "id": "segmoe",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-SegMoE",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-SegMoE"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/SegMoE: Segmind Mixture of Diffusion Experts](https://mirror.ghproxy.com/https://github.com/segmind/segmoe) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI YoloWorld-EfficientSAM",
            "id": "yoloworld",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-YoloWorld-EfficientSAM",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-YoloWorld-EfficientSAM"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/YOLO-World + EfficientSAM](https://huggingface.co/spaces/SkalskiP/YOLO-World) & [a/YOLO-World](https://mirror.ghproxy.com/https://github.com/AILab-CVC/YOLO-World) for ComfyUI\nNOTE: Install the efficient_sam model from the Install models menu.\n[w/When installing or updating this custom node, many installation packages may be downgraded due to the installation of requirements.\n!! python3.12 is incompatible.]"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-PixArt-alpha-Diffusers",
            "id": "pixart-alpha",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-PixArt-alpha-Diffusers",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-PixArt-alpha-Diffusers"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/PixArt-alpha-Diffusers](https://mirror.ghproxy.com/https://github.com/PixArt-alpha/PixArt-alpha) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-BRIA_AI-RMBG",
            "id": "bria-ai-rmbg",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-BRIA_AI-RMBG",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-BRIA_AI-RMBG"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of BRIA RMBG Model for ComfyUI."
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "DepthFM IN COMFYUI",
            "id": "depthfm",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-DepthFM",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-DepthFM"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/DepthFM](https://mirror.ghproxy.com/https://github.com/CompVis/depth-fm) for ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-BiRefNet-ZHO",
            "id": "birefnet",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-BiRefNet-ZHO",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-BiRefNet-ZHO"
            ],
            "install_type": "git-clone",
            "description": "Better version for [a/BiRefNet](https://mirror.ghproxy.com/https://github.com/zhengpeng7/birefnet) in ComfyUI | Both img and video"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "Phi-3-mini in ComfyUI",
            "id": "phi3mini",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Phi-3-mini",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Phi-3-mini"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Phi3mini_4k_ModelLoader_Zho, Phi3mini_4k_Zho, Phi3mini_4k_Chat_Zho"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-ArtGallery",
            "id": "artgallery",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-ArtGallery",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-ArtGallery"
            ],
            "install_type": "git-clone",
            "description": "Prompt Visualization | Art Gallery\n[w/WARN: Installation requires 2GB of space, and it will involve a long download time.]"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-Animated-optical-illusions",
            "id": "animated-optical-illusion",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Animated-optical-illusions",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Animated-optical-illusions"
            ],
            "install_type": "git-clone",
            "description": "Animated optical illusions in ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "Qwen-2.5 in ComfyUI",
            "id": "qwen",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Qwen",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-Qwen"
            ],
            "install_type": "git-clone",
            "description": "Using Qwen-2.5 in ComfyUI"
        },
        {
            "author": "ZHO-ZHO-ZHO",
            "title": "ComfyUI-UltraEdit-ZHO",
            "reference": "https://github.com/ZHO-ZHO-ZHO/ComfyUI-UltraEdit-ZHO",
            "files": [
                "https://github.com/ZHO-ZHO-ZHO/ComfyUI-UltraEdit-ZHO"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/UltraEdit](https://mirror.ghproxy.com/https://github.com/HaozheZhao/UltraEdit) (Diffusers) for ComfyUI"
        },
        {
            "author": "kenjiqq",
            "title": "qq-nodes-comfyui",
            "reference": "https://github.com/kenjiqq/qq-nodes-comfyui",
            "files": [
                "https://github.com/kenjiqq/qq-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Any List, Image Accumulator Start, Image Accumulator End, Load Lines From Text File, XY Grid Helper, Slice List, Axis To String/Int/Float/Model, ..."
        },
        {
            "author": "80sVectorz",
            "title": "ComfyUI-Static-Primitives",
            "reference": "https://github.com/80sVectorz/ComfyUI-Static-Primitives",
            "files": [
                "https://github.com/80sVectorz/ComfyUI-Static-Primitives"
            ],
            "install_type": "git-clone",
            "description": "Adds Static Primitives to ComfyUI. Mostly to work with reroute nodes"
        },
        {
            "author": "AbdullahAlfaraj",
            "title": "Comfy-Photoshop-SD",
            "reference": "https://github.com/AbdullahAlfaraj/Comfy-Photoshop-SD",
            "files": [
                "https://github.com/AbdullahAlfaraj/Comfy-Photoshop-SD"
            ],
            "install_type": "git-clone",
            "description": "Nodes: load Image with metadata, get config data, load image from base64 string, Load Loras From Prompt, Generate Latent Noise, Combine Two Latents Into Batch, General Purpose Controlnet Unit, ControlNet Script, Content Mask Latent, Auto-Photoshop-SD Seed, Expand and Blur the Mask"
        },
        {
            "author": "zhuanqianfish",
            "title": "EasyCaptureNode for ComfyUI",
            "reference": "https://github.com/zhuanqianfish/ComfyUI-EasyNode",
            "files": [
                "https://github.com/zhuanqianfish/ComfyUI-EasyNode"
            ],
            "install_type": "git-clone",
            "description": "Capture window content from other programs, easyway combined with LCM for real-time painting"
        },
        {
            "author": "discopixel-studio",
            "title": "PhotoRoom Nodes by Discopixel",
            "reference": "https://github.com/discopixel-studio/comfyui-discopixel",
            "files": [
                "https://github.com/discopixel-studio/comfyui-discopixel"
            ],
            "install_type": "git-clone",
            "description": "A small collection of custom nodes for use with ComfyUI, by [a/Discopixel](https://discopixel.studio)"
        },
        {
            "author": "zcfrank1st",
            "title": "ComfyUI Yolov8",
            "reference": "https://github.com/zcfrank1st/Comfyui-Yolov8",
            "files": [
                "https://github.com/zcfrank1st/Comfyui-Yolov8"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Yolov8Detection, Yolov8Segmentation. Deadly simple yolov8 comfyui plugin"
        },
        {
            "author": "SoftMeng",
            "title": "ComfyUI_Mexx_Styler",
            "reference": "https://github.com/SoftMeng/ComfyUI_Mexx_Styler",
            "files": [
                "https://github.com/SoftMeng/ComfyUI_Mexx_Styler"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ComfyUI Mexx Styler, ComfyUI Mexx Styler Advanced"
        },
        {
            "author": "SoftMeng",
            "title": "ComfyUI_Mexx_Poster",
            "reference": "https://github.com/SoftMeng/ComfyUI_Mexx_Poster",
            "files": [
                "https://github.com/SoftMeng/ComfyUI_Mexx_Poster"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ComfyUI_Mexx_Poster"
        },
        {
            "author": "SoftMeng",
            "title": "ComfyUI_ImageToText",
            "reference": "https://github.com/SoftMeng/ComfyUI_ImageToText",
            "files": [
                "https://github.com/SoftMeng/ComfyUI_ImageToText"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ComfyUI_ImageToText"
        },
        {
            "author": "SoftMeng",
            "title": "ComfyUI-DeepCache-Fix",
            "reference": "https://github.com/SoftMeng/ComfyUI-DeepCache-Fix",
            "files": [
                "https://github.com/SoftMeng/ComfyUI-DeepCache-Fix"
            ],
            "install_type": "git-clone",
            "description": "Accelerate ComfyUI Nodes for Faster Image Generation, Ensuring Consistency Pre and Post-Acceleration, Ideal for Bulk Image Production."
        },
        {
            "author": "SoftMeng",
            "title": "ComfyUI-PIL",
            "reference": "https://github.com/SoftMeng/ComfyUI-PIL",
            "files": [
                "https://github.com/SoftMeng/ComfyUI-PIL"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI is proud to present a new plugin designed to enhance user experience through seamless integration with Pillow, the powerful fork of Python Imaging Library (PIL). This plugin offers a suite of basic image manipulation tools that are easy to use and integrate directly into the ComfyUI framework."
        },
        {
            "author": "wmatson",
            "title": "easy-comfy-nodes",
            "reference": "https://github.com/wmatson/easy-comfy-nodes",
            "files": [
                "https://github.com/wmatson/easy-comfy-nodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of utility nodes primarily for interacting with comfy via automated systems"
        },
        {
            "author": "DrJKL",
            "title": "ComfyUI-Anchors",
            "reference": "https://github.com/DrJKL/ComfyUI-Anchors",
            "files": [
                "https://github.com/DrJKL/ComfyUI-Anchors"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI extension to add spatial anchors/waypoints to better navigate large workflows."
        },
        {
            "author": "vanillacode314",
            "title": "Simple Wildcard",
            "reference": "https://github.com/vanillacode314/SimpleWildcardsComfyUI",
            "files": ["https://github.com/vanillacode314/SimpleWildcardsComfyUI"],
            "install_type": "git-clone",
            "pip": ["pipe"],
            "description": "A simple wildcard node for ComfyUI. Can also be used a style prompt node."
        },
        {
            "author": "WebDev9000",
            "title": "WebDev9000-Nodes",
            "reference": "https://github.com/WebDev9000/WebDev9000-Nodes",
            "files": [
                "https://github.com/WebDev9000/WebDev9000-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Ignore Braces, Settings Switch."
        },
        {
            "author": "Scholar01",
            "title": "SComfyUI-Keyframe",
            "reference": "https://github.com/Scholar01/ComfyUI-Keyframe",
            "files": [
                "https://github.com/Scholar01/ComfyUI-Keyframe"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Keyframe Part, Keyframe Interpolation Part, Keyframe Apply."
        },
        {
            "author": "Haoming02",
            "title": "ComfyUI Old Photo Restoration",
            "reference": "https://github.com/Haoming02/comfyui-old-photo-restoration",
            "files": [
                "https://github.com/Haoming02/comfyui-old-photo-restoration"
            ],
            "install_type": "git-clone",
            "description": "Perform Bringing-Old-Photos-Back-to-Life"
        },
        {
            "author": "Haoming02",
            "title": "Diffusion CG",
            "reference": "https://github.com/Haoming02/comfyui-diffusion-cg",
            "files": [
                "https://github.com/Haoming02/comfyui-diffusion-cg"
            ],
            "install_type": "git-clone",
            "description": "Color Grading for Stable Diffusion"
        },
        {
            "author": "Haoming02",
            "title": "Prompt Format",
            "reference": "https://github.com/Haoming02/comfyui-prompt-format",
            "files": [
                "https://github.com/Haoming02/comfyui-prompt-format"
            ],
            "install_type": "git-clone",
            "description": "Add a button that formats the prompts in textfields"
        },
        {
            "author": "Haoming02",
            "title": "Clear Screen",
            "reference": "https://github.com/Haoming02/comfyui-clear-screen",
            "files": [
                "https://github.com/Haoming02/comfyui-clear-screen"
            ],
            "install_type": "git-clone",
            "description": "Add a button that clears the console"
        },
        {
            "author": "Haoming02",
            "title": "Menu Anchor",
            "reference": "https://github.com/Haoming02/comfyui-menu-anchor",
            "files": [
                "https://github.com/Haoming02/comfyui-menu-anchor"
            ],
            "install_type": "git-clone",
            "description": "Snaps the menu to the corner automatically"
        },
        {
            "author": "Haoming02",
            "title": "Tab Handler",
            "reference": "https://github.com/Haoming02/comfyui-tab-handler",
            "files": [
                "https://github.com/Haoming02/comfyui-tab-handler"
            ],
            "install_type": "git-clone",
            "description": "Use the Tab key to switch between textfields"
        },
        {
            "author": "Haoming02",
            "title": "Floodgate",
            "reference": "https://github.com/Haoming02/comfyui-floodgate",
            "files": [
                "https://github.com/Haoming02/comfyui-floodgate"
            ],
            "install_type": "git-clone",
            "description": "A node that allows you to switch between execution flow"
        },
        {
            "author": "Haoming02",
            "title": "Node Beautify",
            "reference": "https://github.com/Haoming02/comfyui-node-beautify",
            "files": [
                "https://github.com/Haoming02/comfyui-node-beautify"
            ],
            "install_type": "git-clone",
            "description": "Add a button that formats the workflow graph"
        },
        {
            "author": "Haoming02",
            "title": "ComfyUI ReSharpen",
            "reference": "https://github.com/Haoming02/comfyui-resharpen",
            "files": [
                "https://github.com/Haoming02/comfyui-resharpen"
            ],
            "install_type": "git-clone",
            "description": "Manipulate the details of generations."
        },
        {
            "author": "bedovyy",
            "title": "ComfyUI_NAIDGenerator",
            "reference": "https://github.com/bedovyy/ComfyUI_NAIDGenerator",
            "files": [
                "https://github.com/bedovyy/ComfyUI_NAIDGenerator"
            ],
            "install_type": "git-clone",
            "description": "This extension helps generate images through NAI."
        },
        {
            "author": "Off-Live",
            "title": "ComfyUI-off-suite",
            "reference": "https://github.com/Off-Live/ComfyUI-off-suite",
            "files": [
                "https://github.com/Off-Live/ComfyUI-off-suite"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Crop Fit, OFF SEGS to Image, Crop Center wigh SEGS, Watermarking, GW Number Formatting Node."
        },
        {
            "author": "ningxiaoxiao",
            "title": "comfyui-NDI",
            "reference": "https://github.com/ningxiaoxiao/comfyui-NDI",
            "files": [
                "https://github.com/ningxiaoxiao/comfyui-NDI"
            ],
            "pip": ["ndi-python"],
            "install_type": "git-clone",
            "description": "Real-time input output node for ComfyUI by NDI. Leveraging the powerful linking capabilities of NDI, you can access NDI video stream frames and send images generated by the model to NDI video streams."
        },
        {
            "author": "subtleGradient",
            "title": "Touchpad two-finger gesture support for macOS",
            "reference": "https://github.com/subtleGradient/TinkerBot-tech-for-ComfyUI-Touchpad",
            "files": [
                "https://github.com/subtleGradient/TinkerBot-tech-for-ComfyUI-Touchpad"
            ],
            "install_type": "git-clone",
            "description": "Two-finger scrolling (vertical and horizontal) to pan the canvas. Two-finger pinch to zoom in and out. Command-scroll up and down to zoom in and out. Fixes [a/comfyanonymous/ComfyUI#2059](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI/issues/2059)."
        },
        {
            "author": "zcfrank1st",
            "title": "comfyui_visual_anagram",
            "reference": "https://github.com/zcfrank1st/comfyui_visual_anagrams",
            "files": [
                "https://github.com/zcfrank1st/comfyui_visual_anagrams"
            ],
            "install_type": "git-clone",
            "description": "Nodes:visual_anagrams_sample, visual_anagrams_animate"
        },
        {
            "author": "Electrofried",
            "title": "OpenAINode",
            "reference": "https://github.com/Electrofried/ComfyUI-OpenAINode",
            "files": [
                "https://github.com/Electrofried/ComfyUI-OpenAINode"
            ],
            "install_type": "git-clone",
            "description": "A simply node for hooking in to openAI API based servers via comfyUI"
        },
        {
            "author": "AustinMroz",
            "title": "SpliceTools",
            "id": "splicetools",
            "reference": "https://github.com/AustinMroz/ComfyUI-SpliceTools",
            "files": [
                "https://github.com/AustinMroz/ComfyUI-SpliceTools"
            ],
            "install_type": "git-clone",
            "description": "Experimental utility nodes with a focus on manipulation of noised latents"
        },
        {
            "author": "AustinMroz",
            "title": "DynamicOversampling",
            "id": "dynamic-oversampling",
            "reference": "https://github.com/AustinMroz/ComfyUI-DynamicOversampling",
            "files": [
                "https://github.com/AustinMroz/ComfyUI-DynamicOversampling"
            ],
            "install_type": "git-clone",
            "description": "Nodes:DynamicSampler, MeasuredSampler, ResolveMaskPromise"
        },
        {
            "author": "AustinMroz",
            "title": "ComfyUI-WorkflowCheckpointing",
            "id": "workflowcheckpointing",
            "reference": "https://github.com/AustinMroz/ComfyUI-WorkflowCheckpointing",
            "files": [
                "https://github.com/AustinMroz/ComfyUI-WorkflowCheckpointing"
            ],
            "install_type": "git-clone",
            "description": "Automatically creates checkpoints during workflow execution. If If an workflow is canceled or ComfyUI crashes mid-execution, then these checkpoints are used when the workflow is re-queued to resume execution with minimal progress loss."
        },
        {
            "author": "11cafe",
            "title": "ComfyUI Workspace Manager - Comfyspace",
            "reference": "https://github.com/11cafe/comfyui-workspace-manager",
            "files": [
                "https://github.com/11cafe/comfyui-workspace-manager"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node for project management to centralize the management of all your workflows in one place. Seamlessly switch between workflows, create and update them within a single workspace, like Google Docs."
        },
        {
            "author": "knuknX",
            "title": "ComfyUI-Image-Tools",
            "reference": "https://github.com/knuknX/ComfyUI-Image-Tools",
            "files": [
                "https://github.com/knuknX/ComfyUI-Image-Tools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:BatchImageResizeProcessor, SingleImagePathLoader, SingleImageUrlLoader"
        },
        {
            "author": "jtrue",
            "title": "ComfyUI-JaRue",
            "reference": "https://github.com/jtrue/ComfyUI-JaRue",
            "files": [
                "https://github.com/jtrue/ComfyUI-JaRue"
            ],
            "nodename_pattern": "_jru$",
            "install_type": "git-clone",
            "description": "A collection of nodes powering a tensor oracle on a home network with automation"
        },
        {
            "author": "filliptm",
            "title": "ComfyUI_Fill-Nodes",
            "reference": "https://github.com/filliptm/ComfyUI_Fill-Nodes",
            "files": [
                "https://github.com/filliptm/ComfyUI_Fill-Nodes"
            ],
            "install_type": "git-clone",
            "description": "VFX nodes, Shaders, Utilities, Mask tools, Prompting Tools, and much more!"
        },
        {
            "author": "zfkun",
            "title": "ComfyUI_zfkun",
            "reference": "https://github.com/zfkun/ComfyUI_zfkun",
            "files": [
                "https://github.com/zfkun/ComfyUI_zfkun"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes for common tools, including text preview, text translation (multi-platform, multi-language), image loader, webcamera capture."
        },
        {
            "author": "zcfrank1st",
            "title": "Comfyui-Toolbox",
            "reference": "https://github.com/zcfrank1st/Comfyui-Toolbox",
            "files": [
                "https://github.com/zcfrank1st/Comfyui-Toolbox"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Preview Json, Save Json, Test Json Preview, ... preview and save nodes"
        },
        {
            "author": "talesofai",
            "title": "ComfyUI Browser",
            "reference": "https://github.com/talesofai/comfyui-browser",
            "files": [
                "https://github.com/talesofai/comfyui-browser"
            ],
            "install_type": "git-clone",
            "description": "This is an image/video/workflow browser and manager for ComfyUI. You could add image/video/workflow to collections and load it to ComfyUI. You will be able to use your collections everywhere."
        },
        {
            "author": "yolain",
            "title": "ComfyUI Easy Use",
            "reference": "https://github.com/yolain/ComfyUI-Easy-Use",
            "files": [
                "https://github.com/yolain/ComfyUI-Easy-Use"
            ],
            "install_type": "git-clone",
            "description": "To enhance the usability of ComfyUI, optimizations and integrations have been implemented for several commonly used nodes."
        },
        {
            "author": "bruefire",
            "title": "ComfyUI Sequential Image Loader",
            "reference": "https://github.com/bruefire/ComfyUI-SeqImageLoader",
            "files": [
                "https://github.com/bruefire/ComfyUI-SeqImageLoader"
            ],
            "install_type": "git-clone",
            "description": "This is an extension node for ComfyUI that allows you to load frames from a video in bulk and perform masking and sketching on each frame through a GUI."
        },
        {
            "author": "mmaker",
            "title": "mmaker/Color Enhance",
            "reference": "https://git.mmaker.moe/mmaker/sd-webui-color-enhance",
            "files": [
                "https://git.mmaker.moe/mmaker/sd-webui-color-enhance"
            ],
            "install_type": "git-clone",
            "description": "Node: Color Enhance, Color Blend. This is the same algorithm GIMP/GEGL uses for color enhancement. The gist of this implementation is that it converts the color space to CIELCh(ab) and normalizes the chroma (or [colorfulness](https://en.wikipedia.org/wiki/Colorfulness)] component. Original source can be found in the link below."
        },
        {
            "author": "modusCell",
            "title": "Preset Dimensions",
            "reference": "https://github.com/modusCell/ComfyUI-dimension-node-modusCell",
            "files": [
                "https://github.com/modusCell/ComfyUI-dimension-node-modusCell"
            ],
            "install_type": "git-clone",
            "description": "Simple node for sharing latent image size between nodes. Preset dimensions for SD and XL."
        },
        {
            "author": "aria1th",
            "title": "ComfyUI-LogicUtils",
            "reference": "https://github.com/aria1th/ComfyUI-LogicUtils",
            "files": [
                "https://github.com/aria1th/ComfyUI-LogicUtils"
            ],
            "install_type": "git-clone",
            "description": "Logical Utils (compare, string, boolean operations) for ComfyUI"
        },
        {
            "author": "MitoshiroPJ",
            "title": "ComfyUI Slothful Attention",
            "reference": "https://github.com/MitoshiroPJ/comfyui_slothful_attention",
            "files": [
                "https://github.com/MitoshiroPJ/comfyui_slothful_attention"
            ],
            "install_type": "git-clone",
            "description": "This custom node allow controlling output without training. The reducing method is similar to [a/Spatial-Reduction Attention](https://paperswithcode.com/method/spatial-reduction-attention), but generating speed may not be increased on typical image sizes due to overheads. (In some cases, slightly slower)"
        },
        {
            "author": "brianfitzgerald",
            "title": "StyleAligned for ComfyUI",
            "reference": "https://github.com/brianfitzgerald/style_aligned_comfy",
            "files": [
                "https://github.com/brianfitzgerald/style_aligned_comfy"
            ],
            "install_type": "git-clone",
            "description": "Implementation of the [a/StyleAligned](https://style-aligned-gen.github.io/) paper for ComfyUI. This node allows you to apply a consistent style to all images in a batch; by default it will use the first image in the batch as the style reference, forcing all other images to be consistent with it."
        },
        {
            "author": "deroberon",
            "title": "demofusion-comfyui",
            "id": "demofusion",
            "reference": "https://github.com/deroberon/demofusion-comfyui",
            "files": [
                "https://github.com/deroberon/demofusion-comfyui"
            ],
            "install_type": "git-clone",
            "description": "The Demofusion Custom Node is a wrapper that adapts the work and implementation of the [a/DemoFusion](https://ruoyidu.github.io/demofusion/demofusion.html) technique created and implemented by Ruoyi Du to the Comfyui environment."
        },
        {
            "author": "deroberon",
            "title": "StableZero123-comfyui",
            "reference": "https://github.com/deroberon/StableZero123-comfyui",
            "files": [
                "https://github.com/deroberon/StableZero123-comfyui"
            ],
            "install_type": "git-clone",
            "description": "StableZero123 is a node wrapper that uses the model and technique provided [here](https://mirror.ghproxy.com/https://github.com/SUDO-AI-3D/zero123plus/). It uses the Zero123plus model to generate 3D views using just one image."
        },
        {
            "author": "glifxyz",
            "title": "ComfyUI-GlifNodes",
            "id": "glif",
            "reference": "https://github.com/glifxyz/ComfyUI-GlifNodes",
            "files": [
                "https://github.com/glifxyz/ComfyUI-GlifNodes"
            ],
            "install_type": "git-clone",
            "description": "Custom set of nodes used by glif.app. With glif you can build mini apps that are powered by custom comfy workflows."
        },
        {
            "author": "concarne000",
            "title": "ConCarneNode",
            "reference": "https://github.com/concarne000/ConCarneNode",
            "files": [
                "https://github.com/concarne000/ConCarneNode"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Bing Image Grabber, Zephyr chat, Hermes Chat"
        },
        {
            "author": "Aegis72",
            "title": "AegisFlow Utility Nodes",
            "id": "aegis",
            "reference": "https://github.com/aegis72/aegisflow_utility_nodes",
            "files": [
                "https://github.com/aegis72/aegisflow_utility_nodes"
            ],
            "install_type": "git-clone",
            "description": "These nodes will be placed in comfyui/custom_nodes/aegisflow and contains the image passer (accepts an image as either wired or wirelessly, input and passes it through. Latent passer does the same for latents, and the Preprocessor chooser allows a passthrough image and 10 controlnets to be passed in AegisFlow Shima. The inputs on the Preprocessor  chooser should not be renamed if you intend to accept image inputs wirelessly through UE nodes. It can be done, but the send node input regex for each controlnet preprocessor column must also be changed."
        },
        {
            "author": "Aegis72",
            "title": "ComfyUI-styles-all",
            "id": "styles-all",
            "reference": "https://github.com/aegis72/comfyui-styles-all",
            "files": [
                "https://github.com/aegis72/comfyui-styles-all"
            ],
            "install_type": "git-clone",
            "description": "This is a straight clone of Azazeal04's all-in-one styler menu, which was removed from gh on Jan 21, 2024. I have made no changes to the files at all."
        },
        {
            "author": "glibsonoran",
            "title": "Plush-for-ComfyUI",
            "id": "plush",
            "reference": "https://github.com/glibsonoran/Plush-for-ComfyUI",
            "files": [
                "https://github.com/glibsonoran/Plush-for-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Style Prompt, OAI Dall_e Image. Plush contains two OpenAI enabled nodes: Style Prompt: Takes your prompt and the art style you specify and generates a prompt from ChatGPT3 or 4 that Stable Diffusion can use to generate an image in that style. OAI Dall_e 3: Takes your prompt and parameters and produces a Dall_e3 image in ComfyUI."
        },
        {
            "author": "vienteck",
            "title": "ComfyUI-Chat-GPT-Integration",
            "reference": "https://github.com/vienteck/ComfyUI-Chat-GPT-Integration",
            "files": [
                "https://github.com/vienteck/ComfyUI-Chat-GPT-Integration"
            ],
            "install_type": "git-clone",
            "description": "This extension is a reimagined version based on the [a/ComfyUI-QualityOfLifeSuit_Omar92](https://mirror.ghproxy.com/https://github.com/omar92/ComfyUI-QualityOfLifeSuit_Omar92) extension, and it supports integration with ChatGPT through the new OpenAI API.\nNOTE: See detailed installation instructions on the [a/repository](https://mirror.ghproxy.com/https://github.com/vienteck/ComfyUI-Chat-GPT-Integration)."
        },
        {
            "author": "MNeMoNiCuZ",
            "title": "ComfyUI-mnemic-nodes",
            "id": "comfyui-mnemic-nodes",
            "reference": "https://github.com/MNeMoNiCuZ/ComfyUI-mnemic-nodes",
            "files": [
                "https://github.com/MNeMoNiCuZ/ComfyUI-mnemic-nodes"
            ],
            "install_type": "git-clone",
            "description": "Added new models to Groq LLM. Added a new node: Tiktoken Tokenizer Info."
        },
        {
            "author": "AI2lab",
            "title": "comfyUI-tool-2lab",
            "id": "tool-2lab",
            "reference": "https://github.com/AI2lab/comfyUI-tool-2lab",
            "files": [
                "https://github.com/AI2lab/comfyUI-tool-2lab"
            ],
            "install_type": "git-clone",
            "description": "tool set for developing workflow and publish to web api server"
        },
        {
            "author": "AI2lab",
            "title": "comfyUI-DeepSeek-2lab",
            "id": "deepseek",
            "reference": "https://github.com/AI2lab/comfyUI-DeepSeek-2lab",
            "files": [
                "https://github.com/AI2lab/comfyUI-DeepSeek-2lab"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of DeepSeek for ComfyUI"
        },
        {
            "author": "AI2lab",
            "title": "comfyUI-siliconflow-api-2lab",
            "id": "siliconflow",
            "reference": "https://github.com/AI2lab/comfyUI-siliconflow-api-2lab",
            "files": [
                "https://github.com/AI2lab/comfyUI-siliconflow-api-2lab"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of siliconflow API for ComfyUI\nHow to use:apply api key in ：https://cloud.siliconflow.cn/\nadd api key in config.json"
        },
        {
            "author": "AI2lab",
            "title": "comfyUI-kling-api-2lab",
            "reference": "https://github.com/AI2lab/comfyUI-kling-api-2lab",
            "files": [
                "https://github.com/AI2lab/comfyUI-kling-api-2lab"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of KLing for ComfyUI"
        },
        {
            "author": "SpaceKendo",
            "title": "Text to video for Stable Video Diffusion in ComfyUI",
            "id": "svd-txt2vid",
            "reference": "https://github.com/SpaceKendo/ComfyUI-svd_txt2vid",
            "files": [
                "https://github.com/SpaceKendo/ComfyUI-svd_txt2vid"
            ],
            "install_type": "git-clone",
            "description": "This is node replaces the init_image conditioning for the [a/Stable Video Diffusion](https://mirror.ghproxy.com/https://github.com/Stability-AI/generative-models) image to video model with text embeds, together with a conditioning frame. The conditioning frame is a set of latents."
        },
        {
            "author": "NimaNzrii",
            "title": "comfyui-popup_preview",
            "id": "popup-preview",
            "reference": "https://github.com/NimaNzrii/comfyui-popup_preview",
            "files": [
                "https://github.com/NimaNzrii/comfyui-popup_preview"
            ],
            "install_type": "git-clone",
            "description": "popup preview for comfyui"
        },
        {
            "author": "NimaNzrii",
            "title": "comfyui-photoshop",
            "id": "comfy-photoshop",
            "reference": "https://github.com/NimaNzrii/comfyui-photoshop",
            "files": [
                "https://github.com/NimaNzrii/comfyui-photoshop"
            ],
            "install_type": "git-clone",
            "description": "Powerfull bridge to Photoshop by NimaNzrii"
        },
        {
            "author": "Rui",
            "title": "RUI-Nodes",
            "id": "rui-nodes",
            "reference": "https://github.com/rui40000/RUI-Nodes",
            "files": [
                "https://github.com/rui40000/RUI-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Rui's workflow-specific custom node, written using GPT."
        },
        {
            "author": "dmarx",
            "title": "ComfyUI-Keyframed",
            "id": "keyframed",
            "reference": "https://github.com/dmarx/ComfyUI-Keyframed",
            "files": [
                "https://github.com/dmarx/ComfyUI-Keyframed"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to facilitate parameter/prompt keyframing using comfyui nodes for defining and manipulating parameter curves. Essentially provides a ComfyUI interface to the [a/keyframed](https://mirror.ghproxy.com/https://github.com/dmarx/keyframed) library."
        },
        {
            "author": "dmarx",
            "title": "ComfyUI-AudioReactive",
            "id": "audioreactive",
            "reference": "https://github.com/dmarx/ComfyUI-AudioReactive",
            "files": [
                "https://github.com/dmarx/ComfyUI-AudioReactive"
            ],
            "install_type": "git-clone",
            "description": "porting audioreactivity pipeline from vktrs to comfyui."
        },
        {
            "author": "TripleHeadedMonkey",
            "title": "ComfyUI_MileHighStyler",
            "id": "milehighstyler",
            "reference": "https://github.com/TripleHeadedMonkey/ComfyUI_MileHighStyler",
            "files": [
                "https://github.com/TripleHeadedMonkey/ComfyUI_MileHighStyler"
            ],
            "install_type": "git-clone",
            "description": "This extension provides various SDXL Prompt Stylers. See: [a/youtube](https://youtu.be/WBHI-2uww7o?si=dijvDaUI4nmx4VkF)"
        },
        {
            "author": "BennyKok",
            "title": "ComfyUI Deploy",
            "id": "comfy-deploy",
            "reference": "https://github.com/BennyKok/comfyui-deploy",
            "files": [
                "https://github.com/BennyKok/comfyui-deploy"
            ],
            "install_type": "git-clone",
            "description": "Open source comfyui deployment platform, a vercel for generative workflow infra."
        },
        {
            "author": "florestefano1975",
            "title": "comfyui-portrait-master",
            "id": "portrait-master",
            "reference": "https://github.com/florestefano1975/comfyui-portrait-master",
            "files": [
                "https://github.com/florestefano1975/comfyui-portrait-master"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Portrait Master. A node designed to help AI image creators to generate prompts for human portraits."
        },
        {
            "author": "florestefano1975",
            "title": "comfyui-prompt-composer",
            "id": "prompt-composer",
            "reference": "https://github.com/florestefano1975/comfyui-prompt-composer",
            "files": [
                "https://github.com/florestefano1975/comfyui-prompt-composer"
            ],
            "install_type": "git-clone",
            "description": "A suite of tools for prompt management. Combining nodes helps the user sequence strings for prompts, also creating logical groupings if necessary. Individual nodes can be chained together in any order."
        },
        {
            "author": "florestefano1975",
            "title": "ComfyUI StabilityAI Suite",
            "id": "sai-suite",
            "reference": "https://github.com/florestefano1975/ComfyUI-StabilityAI-Suite",
            "files": [
                "https://github.com/florestefano1975/ComfyUI-StabilityAI-Suite"
            ],
            "install_type": "git-clone",
            "description": "This fork of the official StabilityAI repository contains a number of enhancements and implementations."
        },
        {
            "author": "florestefano1975",
            "title": "ComfyUI HiDiffusion",
            "id": "hidiffusion",
            "reference": "https://github.com/florestefano1975/ComfyUI-HiDiffusion",
            "files": [
                "https://github.com/florestefano1975/ComfyUI-HiDiffusion"
            ],
            "install_type": "git-clone",
            "description": "Simple custom nodes for testing and use HiDiffusion technology: https://mirror.ghproxy.com/https://github.com/megvii-research/HiDiffusion/"
        },
        {
            "author": "florestefano1975",
            "title": "Advanced Sequence Seed Generator",
            "id": "adv-seq-seed-gen",
            "reference": "https://github.com/florestefano1975/ComfyUI-Advanced-Sequence-Seed",
            "files": [
                "https://github.com/florestefano1975/ComfyUI-Advanced-Sequence-Seed"
            ],
            "install_type": "git-clone",
            "description": "A simple seed generator based on special number sequences: Fibonacci, Prime, Padovan, Triangular, Catalan, Pell, Lucas"
        },
        {
            "author": "florestefano1975",
            "title": "ComfyUI-CogVideoX",
            "id": "sf-cog-video-x",
            "reference": "https://github.com/florestefano1975/ComfyUI-CogVideoX",
            "files": [
                "https://github.com/florestefano1975/ComfyUI-CogVideoX"
            ],
            "install_type": "git-clone",
            "description": "Experience the CogVideoX model on ComfyUI"
        },
        {
            "author": "mozman",
            "title": "ComfyUI_mozman_nodes",
            "id": "mozman-nodes",
            "reference": "https://github.com/mozman/ComfyUI_mozman_nodes",
            "files": [
                "https://github.com/mozman/ComfyUI_mozman_nodes"
            ],
            "install_type": "git-clone",
            "description": "This extension provides styler nodes for SDXL.\n\nNOTE: Due to the dynamic nature of node name definitions, ComfyUI-Manager cannot recognize the node list from this extension. The Missing nodes and Badge features are not available for this extension."
        },
        {
            "author": "rcsaquino",
            "title": "rcsaquino/comfyui-custom-nodes",
            "id": "rcsaquino-nodes",
            "reference": "https://github.com/rcsaquino/comfyui-custom-nodes",
            "files": [
                "https://github.com/rcsaquino/comfyui-custom-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: VAE Processor, VAE Loader, Background Remover"
        },
        {
            "author": "rcfcu2000",
            "title": "zhihuige-nodes-comfyui",
            "id": "zhihuige-nodes",
            "reference": "https://github.com/rcfcu2000/zhihuige-nodes-comfyui",
            "files": [
                "https://github.com/rcfcu2000/zhihuige-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Combine ZHGMasks, Cover ZHGMasks, ZHG FaceIndex, ZHG SaveImage, ZHG SmoothEdge, ZHG GetMaskArea, ..."
        },
        {
            "author": "IDGallagher",
            "title": "IG Interpolation Nodes",
            "id": "ig-nodes",
            "reference": "https://github.com/IDGallagher/ComfyUI-IG-Nodes",
            "files": [
                "https://github.com/IDGallagher/ComfyUI-IG-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes to aid in the exploration of Latent Space"
        },
        {
            "author": "IDGallagher",
            "title": "ComfyUI-IG-Motion-I2V",
            "id": "comfyui-ig-motion-i2v",
            "reference": "https://github.com/IDGallagher/ComfyUI-IG-Motion-I2V",
            "files": [
                "https://github.com/IDGallagher/ComfyUI-IG-Motion-I2V"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI adaptation of https://mirror.ghproxy.com/https://github.com/G-U-N/Motion-I2V"
        },
        {
            "author": "violet-chen",
            "title": "comfyui-psd2png",
            "id": "psd2png",
            "reference": "https://github.com/violet-chen/comfyui-psd2png",
            "files": [
                "https://github.com/violet-chen/comfyui-psd2png"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Psd2Png."
        },
        {
            "author": "lldacing",
            "title": "comfyui-easyapi-nodes",
            "id": "easyapi",
            "reference": "https://github.com/lldacing/comfyui-easyapi-nodes",
            "files": [
                "https://github.com/lldacing/comfyui-easyapi-nodes"
            ],
            "install_type": "git-clone",
            "description": "Provides some features and nodes related to API calls."
        },
        {
            "author": "CosmicLaca",
            "title": "Primere nodes for ComfyUI",
            "id": "primere",
            "reference": "https://github.com/CosmicLaca/ComfyUI_Primere_Nodes",
            "files": [
                "https://github.com/CosmicLaca/ComfyUI_Primere_Nodes"
            ],
            "install_type": "git-clone",
            "description": "This extension provides various utility nodes. Inputs(prompt, styles, dynamic, merger, ...), Outputs(style pile), Dashboard(selectors, loader, switch, ...), Networks(LORA, Embedding, Hypernetwork), Visuals(visual selectors, )"
        },
        {
            "author": "RenderRift",
            "title": "ComfyUI-RenderRiftNodes",
            "id": "renderrift",
            "reference": "https://github.com/RenderRift/ComfyUI-RenderRiftNodes",
            "files": [
                "https://github.com/RenderRift/ComfyUI-RenderRiftNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:RR_Date_Folder_Format, RR_Image_Metadata_Overlay, RR_VideoPathMetaExtraction, RR_DisplayMetaOptions. This extension provides nodes designed to enhance the Animatediff workflow."
        },
        {
            "author": "OpenArt-AI",
            "title": "ComfyUI Assistant",
            "id": "openart",
            "reference": "https://github.com/OpenArt-AI/ComfyUI-Assistant",
            "files": [
                "https://github.com/OpenArt-AI/ComfyUI-Assistant"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Assistant is your one stop plugin for everything you need to get started with comfy-ui. Now it provides useful courses, tutorials, and basic templates."
        },
        {
            "author": "ttulttul",
            "title": "ComfyUI Iterative Mixing Nodes",
            "id": "itermix",
            "reference": "https://github.com/ttulttul/ComfyUI-Iterative-Mixer",
            "files": [
                "https://github.com/ttulttul/ComfyUI-Iterative-Mixer"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use Florence2 VLM for image vision tasks: object detection, captioning, segmentation and ocr"
        },
        {
            "author": "ttulttul",
            "title": "ComfyUI-Tensor-Operations",
            "id": "tensorop",
            "reference": "https://github.com/ttulttul/ComfyUI-Tensor-Operations",
            "files": [
                "https://github.com/ttulttul/ComfyUI-Tensor-Operations"
            ],
            "install_type": "git-clone",
            "description": "This repo contains nodes for ComfyUI that implement some helpful operations on tensors, such as normalization."
        },
        {
            "author": "jitcoder",
            "title": "LoraInfo",
            "id": "lorainfo",
            "reference": "https://github.com/jitcoder/lora-info",
            "files": [
                "https://github.com/jitcoder/lora-info"
            ],
            "install_type": "git-clone",
            "description": "Shows Lora information from CivitAI and outputs trigger words and example prompt"
        },
        {
            "author": "ceruleandeep",
            "title": "ComfyUI LLaVA Captioner",
            "id": "llava-captioner",
            "reference": "https://github.com/ceruleandeep/ComfyUI-LLaVA-Captioner",
            "files": [
                "https://github.com/ceruleandeep/ComfyUI-LLaVA-Captioner"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI extension for chatting with your images. Runs on your own system, no external services used, no filter. Uses the [a/LLaVA multimodal LLM](https://llava-vl.github.io/) so you can give instructions or ask questions in natural language. It's maybe as smart as GPT3.5, and it can see."
        },
        {
            "author": "styler00dollar",
            "title": "ComfyUI-sudo-latent-upscale",
            "id": "sudo-latent-upscale",
            "reference": "https://github.com/styler00dollar/ComfyUI-sudo-latent-upscale",
            "files": [
                "https://github.com/styler00dollar/ComfyUI-sudo-latent-upscale"
            ],
            "install_type": "git-clone",
            "description": "Directly upscaling inside the latent space. Model was trained for SD1.5 and drawn content. Might add new architectures or update models at some point. This took heavy inspriration from [city96/SD-Latent-Upscaler](https://mirror.ghproxy.com/https://github.com/city96/SD-Latent-Upscaler) and [Ttl/ComfyUi_NNLatentUpscale](https://mirror.ghproxy.com/https://github.com/Ttl/ComfyUi_NNLatentUpscale). "
        },
        {
            "author": "styler00dollar",
            "title": "ComfyUI-deepcache",
            "id": "deepcache",
            "reference": "https://github.com/styler00dollar/ComfyUI-deepcache",
            "files": [
                "https://github.com/styler00dollar/ComfyUI-deepcache"
            ],
            "install_type": "git-clone",
            "description": "This extension provides nodes for [a/DeepCache: Accelerating Diffusion Models for Free](https://arxiv.org/abs/2312.00858)\nNOTE:Original code can be found [a/here](https://gist.github.com/laksjdjf/435c512bc19636e9c9af4ee7bea9eb86). Full credit to laksjdjf for sharing the code. "
        },
        {
            "author": "HarroweD and quadmoon",
            "title": "Harrlogos Prompt Builder Node",
            "id": "harrlogos-prompt-builder",
            "reference": "https://github.com/NotHarroweD/Harronode",
            "nodename_pattern": "Harronode",
            "files": [
                "https://github.com/NotHarroweD/Harronode"
            ],
            "install_type": "git-clone",
            "description": "Harronode is a custom node designed to build prompts easily for use with the Harrlogos SDXL LoRA. This Node simplifies the process of crafting prompts and makes all built in activation terms available at your fingertips."
        },
        {
            "author": "Limitex",
            "title": "ComfyUI-Calculation",
            "id": "calc",
            "reference": "https://github.com/Limitex/ComfyUI-Calculation",
            "files": [
                "https://github.com/Limitex/ComfyUI-Calculation"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Center Calculation. Improved Numerical Calculation for ComfyUI"
        },
        {
            "author": "Limitex",
            "title": "ComfyUI-Diffusers",
            "id": "diffusers",
            "reference": "https://github.com/Limitex/ComfyUI-Diffusers",
            "files": [
                "https://github.com/Limitex/ComfyUI-Diffusers"
            ],
            "install_type": "git-clone",
            "description": "This extension enables the use of the diffuser pipeline in ComfyUI."
        },
        {
            "author": "aiXander",
            "title": "Eden.art nodesuite",
            "id": "eden",
            "reference": "https://github.com/edenartlab/eden_comfy_pipelines",
            "files": [
                "https://github.com/edenartlab/eden_comfy_pipelines"
            ],
            "install_type": "git-clone",
            "description": "Maintained by Eden.art, this is a growing suite of custom nodes for building advanced pipelines."
        },
        {
            "author": "aiXander",
            "title": "Eden.art LoRa Trainer",
            "id": "eden-lora-trainer",
            "reference": "https://github.com/edenartlab/sd-lora-trainer",
            "files": [
                "https://github.com/edenartlab/sd-lora-trainer"
            ],
            "install_type": "git-clone",
            "description": "Maintained by Eden.art, this is a very fast, well tuned trainer for SDXL and SD15"
        },
        {
            "author": "pkpk",
            "title": "ComfyUI-SaveAVIF",
            "id": "saveavif",
            "reference": "https://github.com/pkpkTech/ComfyUI-SaveAVIF",
            "files": [
                "https://github.com/pkpkTech/ComfyUI-SaveAVIF"
            ],
            "install_type": "git-clone",
            "description": "A custom node on ComfyUI that saves images in AVIF format. Workflow can be loaded from images saved at this node."
        },
        {
            "author": "pkpkTech",
            "title": "ComfyUI-ngrok",
            "id": "ngrok",
            "reference": "https://github.com/pkpkTech/ComfyUI-ngrok",
            "files": [
                "https://github.com/pkpkTech/ComfyUI-ngrok"
            ],
            "install_type": "git-clone",
            "description": "Use ngrok to allow external access to ComfyUI.\nNOTE: Need to manually modify a token inside the __init__.py file."
        },
        {
            "author": "pkpk",
            "title": "ComfyUI-TemporaryLoader",
            "id": "temploader",
            "reference": "https://github.com/pkpkTech/ComfyUI-TemporaryLoader",
            "files": [
                "https://github.com/pkpkTech/ComfyUI-TemporaryLoader"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node of ComfyUI that downloads and loads models from the input URL. The model is temporarily downloaded into memory and not saved to storage.\nThis could be useful when trying out models or when using various models on machines with limited storage. Since the model is downloaded into memory, expect higher memory usage than usual."
        },
        {
            "author": "pkpkTech",
            "title": "ComfyUI-SaveQueues",
            "id": "savequeues",
            "reference": "https://github.com/pkpkTech/ComfyUI-SaveQueues",
            "files": [
                "https://github.com/pkpkTech/ComfyUI-SaveQueues"
            ],
            "install_type": "git-clone",
            "description": "Add a button to the menu to save and load the running queue and the pending queues.\nThis is intended to be used when you want to exit ComfyUI with queues still remaining."
        },
        {
            "author": "Crystian",
            "title": "Crystools",
            "id": "crytools",
            "reference": "https://github.com/crystian/ComfyUI-Crystools",
            "files": [
                "https://github.com/crystian/ComfyUI-Crystools"
            ],
            "nodename_pattern": " \\[Crystools\\]$",
            "install_type": "git-clone",
            "description": "With this suit, you can see the resources monitor, progress bar & time elapsed, metadata and compare between two images, compare between two JSONs, show any value to console/display, pipes, and more!\nThis provides better nodes to load/save images, previews, etc, and see \"hidden\" data without loading a new workflow."
        },
        {
            "author": "Crystian",
            "title": "Crystools-save",
            "id": "crytools-save",
            "reference": "https://github.com/crystian/ComfyUI-Crystools-save",
            "files": [
                "https://github.com/crystian/ComfyUI-Crystools-save"
            ],
            "install_type": "git-clone",
            "description": "With this quality of life extension, you can save your workflow with a specific name and include additional details such as the author, a description, and the version (in metadata/json). Important: When you share your workflow (via png/json), others will be able to see your information!"
        },
        {
            "author": "Kangkang625",
            "title": "ComfyUI-Paint-by-Example",
            "id": "paint-by-example",
            "reference": "https://github.com/Kangkang625/ComfyUI-paint-by-example",
            "pip": ["diffusers"],
            "files": [
                "https://github.com/Kangkang625/ComfyUI-paint-by-example"
            ],
            "install_type": "git-clone",
            "description": "This repo is a simple implementation of [a/Paint-by-Example](https://mirror.ghproxy.com/https://github.com/Fantasy-Studio/Paint-by-Example) based on its [a/huggingface pipeline](https://huggingface.co/Fantasy-Studio/Paint-by-Example)."
        },
        {
            "author": "54rt1n",
            "title": "ComfyUI-DareMerge",
            "id": "daremerge",
            "reference": "https://github.com/54rt1n/ComfyUI-DareMerge",
            "files": [
                "https://github.com/54rt1n/ComfyUI-DareMerge"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI powertools for SD1.5 and SDXL model merging."
        },
        {
            "author": "an90ray",
            "title": "ComfyUI_RErouter_CustomNodes",
            "id": "rerouter",
            "reference": "https://github.com/an90ray/ComfyUI_RErouter_CustomNodes",
            "files": [
                "https://github.com/an90ray/ComfyUI_RErouter_CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: RErouter, String (RE), Int (RE)"
        },
        {
            "author": "jesenzhang",
            "title": "ComfyUI_StreamDiffusion",
            "id": "streamdiffusion",
            "reference": "https://github.com/jesenzhang/ComfyUI_StreamDiffusion",
            "files": [
                "https://github.com/jesenzhang/ComfyUI_StreamDiffusion"
            ],
            "install_type": "git-clone",
            "description": "This is a simple implementation StreamDiffusion(A Pipeline-Level Solution for Real-Time Interactive Generation) for ComfyUI"
        },
        {
            "author": "ai-liam",
            "title": "LiamUtil (single node)",
            "id": "liam-util-single",
            "reference": "https://github.com/ai-liam/comfyui_liam_util",
            "files": [
                "https://github.com/ai-liam/comfyui_liam_util"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LiamLoadImage. This node provides the capability to load images from a URL."
        },
        {
            "author": "ai-liam",
            "title": "LiamUtil",
            "id": "liam-util",
            "reference": "https://github.com/ai-liam/comfyui-liam",
            "files": [
                "https://github.com/ai-liam/comfyui-liam"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LiamLibLoadImage, LiamLibImageToGray, LiamLibSaveImg, LiamLibFillImage, PreviewReliefImage, GetBetterDepthImage, LiamLibSaveText"
        },
        {
            "author": "Ryuukeisyou",
            "title": "comfyui_face_parsing",
            "id": "face-parsing",
            "reference": "https://github.com/Ryuukeisyou/comfyui_face_parsing",
            "files": [
                "https://github.com/Ryuukeisyou/comfyui_face_parsing"
            ],
            "install_type": "git-clone",
            "description": "This is a set of custom nodes for ComfyUI. The nodes utilize the [a/face parsing model](https://huggingface.co/jonathandinu/face-parsing) to provide detailed segmantation of face. To improve face segmantation accuracy, [a/yolov8 face model](https://huggingface.co/Bingsu/adetailer/) is used to first extract face from an image. There are also auxiliary nodes for image and mask processing. A guided filter is also provided for skin smoothing."
        },
        {
            "author": "Ryuukeisyou",
            "title": "ComfyUI-SyncTalk",
            "id": "synctalk",
            "reference": "https://github.com/Ryuukeisyou/ComfyUI-SyncTalk",
            "files": [
                "https://github.com/Ryuukeisyou/ComfyUI-SyncTalk"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implemntation for [a/SyncTalk](https://mirror.ghproxy.com/https://github.com/ZiqiaoPeng/SyncTalk)"
        },
        {
            "author": "tocubed",
            "title": "ComfyUI-AudioReactor",
            "id": "audioreactor",
            "reference": "https://github.com/tocubed/ComfyUI-AudioReactor",
            "files": [
                "https://github.com/tocubed/ComfyUI-AudioReactor"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Shadertoy, Load Audio (from Path), Audio Frame Transform (Shadertoy), Audio Frame Transform (Beats)"
        },
        {
            "author": "ntc-ai",
            "title": "ComfyUI - Apply LoRA Stacker with DARE",
            "reference": "https://github.com/ntc-ai/ComfyUI-DARE-LoRA-Merge",
            "files": [
                "https://github.com/ntc-ai/ComfyUI-DARE-LoRA-Merge"
            ],
            "install_type": "git-clone",
            "description": "An experiment about combining multiple LoRAs with [a/DARE](https://arxiv.org/pdf/2311.03099.pdf)"
        },
        {
            "author": "wwwins",
            "title": "ComfyUI-Simple-Aspect-Ratio",
            "reference": "https://github.com/wwwins/ComfyUI-Simple-Aspect-Ratio",
            "files": [
                "https://github.com/wwwins/ComfyUI-Simple-Aspect-Ratio"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SimpleAspectRatio"
        },
        {
            "author": "ownimage",
            "title": "ComfyUI-ownimage",
            "reference": "https://github.com/ownimage/ComfyUI-ownimage",
            "files": [
                "https://github.com/ownimage/ComfyUI-ownimage"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Caching Image Loader."
        },
        {
            "author": "Ryuukeisyou",
            "title": "comfyui_io_helpers",
            "reference": "https://github.com/Ryuukeisyou/comfyui_io_helpers",
            "files": [
                "https://github.com/Ryuukeisyou/comfyui_io_helpers"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageLoadFromBase64, ImageLoadByPath, ImageLoadAsMaskByPath, ImageSaveToPath, ImageSaveAsBase64, VHSFileNamesToStrings(IOHelpers)."
        },
        {
            "author": "flowtyone",
            "title": "ComfyUI-Flowty-LDSR",
            "reference": "https://github.com/flowtyone/ComfyUI-Flowty-LDSR",
            "files": [
                "https://github.com/flowtyone/ComfyUI-Flowty-LDSR"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node that lets you take advantage of Latent Diffusion Super Resolution (LDSR) models inside ComfyUI."
        },
        {
            "author": "flowtyone",
            "title": "ComfyUI-Flowty-TripoSR",
            "reference": "https://github.com/flowtyone/ComfyUI-Flowty-TripoSR",
            "files": [
                "https://github.com/flowtyone/ComfyUI-Flowty-TripoSR"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node that lets you use TripoSR right from ComfyUI.\n[a/TripoSR](https://mirror.ghproxy.com/https://github.com/VAST-AI-Research/TripoSR) is a state-of-the-art open-source model for fast feedforward 3D reconstruction from a single image, collaboratively developed by Tripo AI and Stability AI. (TL;DR it creates a 3d model from an image.)"
        },
        {
            "author": "flowtyone",
            "title": "ComfyUI-Flowty-CRM",
            "reference": "https://github.com/flowtyone/ComfyUI-Flowty-CRM",
            "files": [
                "https://github.com/flowtyone/ComfyUI-Flowty-CRM"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node that lets you use Convolutional Reconstruction Models right from ComfyUI.\n[a/CRM](https://ml.cs.tsinghua.edu.cn/~zhengyi/CRM/) is a high-fidelity feed-forward single image-to-3D generative model."
        },
        {
            "author": "massao000",
            "title": "ComfyUI_aspect_ratios",
            "reference": "https://github.com/massao000/ComfyUI_aspect_ratios",
            "files": [
                "https://github.com/massao000/ComfyUI_aspect_ratios"
            ],
            "install_type": "git-clone",
            "description": "Aspect ratio selector for ComfyUI based on [a/sd-webui-ar](https://mirror.ghproxy.com/https://github.com/alemelis/sd-webui-ar?tab=readme-ov-file)."
        },
        {
            "author": "SiliconFlow",
            "title": "OneDiff Nodes",
            "id": "onddiff",
            "reference": "https://github.com/siliconflow/onediff_comfy_nodes",
            "files": [
                "https://github.com/siliconflow/onediff_comfy_nodes"
            ],
            "install_type": "git-clone",
            "description": "[a/Onediff](https://mirror.ghproxy.com/https://github.com/siliconflow/onediff) ComfyUI Nodes."
        },
        {
            "author": "hinablue",
            "title": "ComfyUI 3D Pose Editor",
            "id": "3d-pose-editor",
            "reference": "https://github.com/hinablue/ComfyUI_3dPoseEditor",
            "files": [
                "https://github.com/hinablue/ComfyUI_3dPoseEditor"
            ],
            "install_type": "git-clone",
            "description": "Nodes:3D Pose Editor"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-CameraCtrl-Wrapper",
            "id": "cameractrl-wrapper",
            "reference": "https://github.com/chaojie/ComfyUI-CameraCtrl-Wrapper",
            "files": [
                "https://github.com/chaojie/ComfyUI-CameraCtrl-Wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-CameraCtrl-Wrapper"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-EasyAnimate",
            "id": "easyanimate",
            "reference": "https://github.com/chaojie/ComfyUI-EasyAnimate",
            "files": [
                "https://github.com/chaojie/ComfyUI-EasyAnimate"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-EasyAnimate"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI_StreamingT2V",
            "id": "streamingt2v",
            "reference": "https://github.com/chaojie/ComfyUI_StreamingT2V",
            "files": [
                "https://github.com/chaojie/ComfyUI_StreamingT2V"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI_StreamingT2V"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Open-Sora-Plan",
            "id": "opensora-plan",
            "reference": "https://github.com/chaojie/ComfyUI-Open-Sora-Plan",
            "files": [
                "https://github.com/chaojie/ComfyUI-Open-Sora-Plan"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node for [a/Open-Sora-Plan](https://mirror.ghproxy.com/https://github.com/PKU-YuanGroup/Open-Sora-Plan)"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-MuseTalk",
            "id": "musetalk-chaojie",
            "reference": "https://github.com/chaojie/ComfyUI-MuseTalk",
            "files": [
                "https://github.com/chaojie/ComfyUI-MuseTalk"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI MuseTalk"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-MuseV",
            "id": "musev",
            "reference": "https://github.com/chaojie/ComfyUI-MuseV",
            "files": [
                "https://github.com/chaojie/ComfyUI-MuseV"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI MuseV"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-AniPortrait",
            "id": "aniportrait",
            "reference": "https://github.com/chaojie/ComfyUI-AniPortrait",
            "files": [
                "https://github.com/chaojie/ComfyUI-AniPortrait"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI [a/AniPortrait](https://mirror.ghproxy.com/https://github.com/Zejun-Yang/AniPortrait)"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Img2Img-Turbo",
            "id": "img2img-turbo",
            "reference": "https://github.com/chaojie/ComfyUI-Img2Img-Turbo",
            "files": [
                "https://github.com/chaojie/ComfyUI-Img2Img-Turbo"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Img2Img-Turbo"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Champ",
            "id": "champ",
            "reference": "https://github.com/chaojie/ComfyUI-Champ",
            "files": [
                "https://github.com/chaojie/ComfyUI-Champ"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Champ"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Open-Sora",
            "id": "opensora",
            "reference": "https://github.com/chaojie/ComfyUI-Open-Sora",
            "files": [
                "https://github.com/chaojie/ComfyUI-Open-Sora"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Open Sora\nNOTE:only supports Linux now"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Trajectory",
            "id": "trajectory",
            "reference": "https://github.com/chaojie/ComfyUI-Trajectory",
            "files": [
                "https://github.com/chaojie/ComfyUI-Trajectory"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Trajectory"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-dust3r",
            "id": "dust3r",
            "reference": "https://github.com/chaojie/ComfyUI-dust3r",
            "files": [
                "https://github.com/chaojie/ComfyUI-dust3r"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI dust3r"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Gemma",
            "id": "gamma",
            "reference": "https://github.com/chaojie/ComfyUI-Gemma",
            "files": [
                "https://github.com/chaojie/ComfyUI-Gemma"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Gemma"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-DynamiCrafter",
            "id": "dynamicrafter-chaojie",
            "reference": "https://github.com/chaojie/ComfyUI-DynamiCrafter",
            "files": [
                "https://github.com/chaojie/ComfyUI-DynamiCrafter"
            ],
            "install_type": "git-clone",
            "description": "Better Dynamic, Higher Resolution, and Stronger Coherence!"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Panda3d",
            "id": "panda3d",
            "reference": "https://github.com/chaojie/ComfyUI-Panda3d",
            "files": [
                "https://github.com/chaojie/ComfyUI-Panda3d"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI 3d engine"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Pymunk",
            "id": "pymunk",
            "reference": "https://github.com/chaojie/ComfyUI-Pymunk",
            "files": [
                "https://github.com/chaojie/ComfyUI-Pymunk"
            ],
            "install_type": "git-clone",
            "description": "Pymunk is a easy-to-use pythonic 2d physics library that can be used whenever you need 2d rigid body physics from Python"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-MotionCtrl",
            "id": "motionctrl",
            "reference": "https://github.com/chaojie/ComfyUI-MotionCtrl",
            "files": [
                "https://github.com/chaojie/ComfyUI-MotionCtrl"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Download the weights of MotionCtrl [a/motionctrl.pth](https://huggingface.co/TencentARC/MotionCtrl/blob/main/motionctrl.pth) and put it to ComfyUI/models/checkpoints"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Motion-Vector-Extractor",
            "id": "motion-vector-extractor",
            "reference": "https://github.com/chaojie/ComfyUI-Motion-Vector-Extractor",
            "files": [
                "https://github.com/chaojie/ComfyUI-Motion-Vector-Extractor"
            ],
            "install_type": "git-clone",
            "description": "Nodes: that we currently provide the package only for x86-64 linux, such as Ubuntu or Debian, and Python 3.8, 3.9, and 3.10."
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-MotionCtrl-SVD",
            "id": "motionctrl-svd",
            "reference": "https://github.com/chaojie/ComfyUI-MotionCtrl-SVD",
            "files": [
                "https://github.com/chaojie/ComfyUI-MotionCtrl-SVD"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Download the weights of MotionCtrl-SVD [a/motionctrl_svd.ckpt](https://huggingface.co/TencentARC/MotionCtrl/blob/main/motionctrl_svd.ckpt) and put it to ComfyUI/models/checkpoints"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-DragAnything",
            "id": "draganything",
            "reference": "https://github.com/chaojie/ComfyUI-DragAnything",
            "files": [
                "https://github.com/chaojie/ComfyUI-DragAnything"
            ],
            "install_type": "git-clone",
            "description": "DragAnything"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-DragNUWA",
            "id": "dragnuwa",
            "reference": "https://github.com/chaojie/ComfyUI-DragNUWA",
            "files": [
                "https://github.com/chaojie/ComfyUI-DragNUWA"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Download the weights of DragNUWA [a/drag_nuwa_svd.pth](https://drive.google.com/file/d/1Z4JOley0SJCb35kFF4PCc6N6P1ftfX4i/view) and put it to ComfyUI/models/checkpoints/drag_nuwa_svd.pth\n[w/Due to changes in the torch package and versions of many other packages, it may disrupt your installation environment.]"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Moore-AnimateAnyone",
            "id": "moore-animateanyone",
            "reference": "https://github.com/chaojie/ComfyUI-Moore-AnimateAnyone",
            "files": [
                "https://github.com/chaojie/ComfyUI-Moore-AnimateAnyone"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Run python tools/download_weights.py first to download weights automatically"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-I2VGEN-XL",
            "id": "i2vgen-xl",
            "reference": "https://github.com/chaojie/ComfyUI-I2VGEN-XL",
            "files": [
                "https://github.com/chaojie/ComfyUI-I2VGEN-XL"
            ],
            "install_type": "git-clone",
            "description": "This is an implementation of [a/i2vgen-xl](https://mirror.ghproxy.com/https://github.com/ali-vilab/i2vgen-xl)"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-LightGlue",
            "id": "lightglue",
            "reference": "https://github.com/chaojie/ComfyUI-LightGlue",
            "files": [
                "https://github.com/chaojie/ComfyUI-LightGlue"
            ],
            "install_type": "git-clone",
            "description": "This is an ComfyUI implementation of LightGlue to generate motion brush"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-RAFT",
            "id": "raft",
            "reference": "https://github.com/chaojie/ComfyUI-RAFT",
            "files": [
                "https://github.com/chaojie/ComfyUI-RAFT"
            ],
            "install_type": "git-clone",
            "description": "This is an ComfyUI implementation of RAFT to generate motion brush"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-LaVIT",
            "id": "lavit",
            "reference": "https://github.com/chaojie/ComfyUI-LaVIT",
            "files": [
                "https://github.com/chaojie/ComfyUI-LaVIT"
            ],
            "install_type": "git-clone",
            "description": "Nodes:VideoLaVITLoader, VideoLaVITT2V, VideoLaVITI2V, VideoLaVITI2VLong, VideoLaVITT2VLong, VideoLaVITI2I"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-SimDA",
            "id": "simda",
            "reference": "https://github.com/chaojie/ComfyUI-SimDA",
            "files": [
                "https://github.com/chaojie/ComfyUI-SimDA"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SimDATrain, SimDALoader, SimDARun, VHS_FILENAMES_STRING_SimDA"
        },
        {
            "author": "chaojie",
            "title": "ComfyUI-Video-Editing-X-Attention",
            "id": "video-editing-x-attention",
            "reference": "https://github.com/chaojie/ComfyUI-Video-Editing-X-Attention",
            "files": [
                "https://github.com/chaojie/ComfyUI-Video-Editing-X-Attention"
            ],
            "install_type": "git-clone",
            "description": "Investigating the Effectiveness of Cross Attention to Unlock Zero-Shot Editing of Text-to-Video Diffusion Models"
        },
        {
            "author": "alexopus",
            "title": "ComfyUI Image Saver",
            "id": "comfyui-image-saver",
            "reference": "https://github.com/alexopus/ComfyUI-Image-Saver",
            "files": [
                "https://github.com/alexopus/ComfyUI-Image-Saver"
            ],
            "install_type": "git-clone",
            "description": "Allows you to save images with their generation metadata compatible with Civitai. Works with png, jpeg and webp. Stores LoRAs, models and embeddings hashes for resource recognition."
        },
        {
            "author": "kft334",
            "title": "Knodes",
            "id": "knodes",
            "reference": "https://github.com/kft334/Knodes",
            "files": [
                "https://github.com/kft334/Knodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Image(s) To Websocket (Base64), Load Image (Base64),Load Images (Base64)"
        },
        {
            "author": "MrForExample",
            "title": "ComfyUI-3D-Pack",
            "id": "3dpack",
            "reference": "https://github.com/MrForExample/ComfyUI-3D-Pack",
            "files": [
                "https://github.com/MrForExample/ComfyUI-3D-Pack"
            ],
            "nodename_pattern": "^\\[Comfy3D\\]",
            "install_type": "git-clone",
            "description": "Make 3D assets generation in ComfyUI good and convenient as it generates image/video!\nThis is an extensive node suite that enables ComfyUI to process 3D inputs (Mesh & UV Texture, etc.) using cutting edge algorithms (3DGS, NeRF, etc.) and models (InstantMesh, CRM, TripoSR, etc.)\nNOTE: Pre-built python wheels can manually download from [a/https://mirror.ghproxy.com/https://github.com/MrForExample/Comfy3D_Pre_Builds](https://mirror.ghproxy.com/https://github.com/MrForExample/Comfy3D_Pre_Builds) if automatic install failed"
        },
        {
            "author": "Mr.ForExample",
            "title": "ComfyUI-AnimateAnyone-Evolved",
            "id": "animateanyone-evolved",
            "reference": "https://github.com/MrForExample/ComfyUI-AnimateAnyone-Evolved",
            "files": [
                "https://github.com/MrForExample/ComfyUI-AnimateAnyone-Evolved"
            ],
            "nodename_pattern": "^\\[AnimateAnyone\\]",
            "install_type": "git-clone",
            "description": "Improved AnimateAnyone implementation that allows you to use the opse image sequence and reference image to generate stylized video.\nThe current goal of this project is to achieve desired pose2video result with 1+FPS on GPUs that are equal to or better than RTX 3080!🚀\n[w/The torch environment may be compromised due to version issues as some torch-related packages are being reinstalled.]"
        },
        {
            "author": "Hangover3832",
            "title": "ComfyUI-Hangover-Nodes",
            "reference": "https://github.com/Hangover3832/ComfyUI-Hangover-Nodes",
            "files": [
                "https://github.com/Hangover3832/ComfyUI-Hangover-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: MS kosmos-2 Interrogator, Save Image w/o Metadata, Image Scale Bounding Box. An implementation of Microsoft [a/kosmos-2](https://huggingface.co/microsoft/kosmos-2-patch14-224) image to text transformer."
        },
        {
            "author": "Hangover3832",
            "title": "ComfyUI-Hangover-Moondream",
            "reference": "https://github.com/Hangover3832/ComfyUI-Hangover-Moondream",
            "files": [
                "https://github.com/Hangover3832/ComfyUI-Hangover-Moondream"
            ],
            "install_type": "git-clone",
            "description": "Moondream is a lightweight multimodal large language model.\n[w/WARN:Additional python code will be downloaded from huggingface and executed. You have to trust this creator if you want to use this node!]"
        },
        {
            "author": "Hangover3832",
            "title": "Recognize Anything Model (RAM) for ComfyUI",
            "reference": "https://github.com/Hangover3832/ComfyUI-Hangover-Recognize_Anything",
            "files": [
                "https://github.com/Hangover3832/ComfyUI-Hangover-Recognize_Anything"
            ],
            "install_type": "git-clone",
            "description": "This is an image recognition node for ComfyUI based on the RAM++ model from [a/xinyu1205](https://huggingface.co/xinyu1205).\nThis node outputs a string of tags with all the recognized objects and elements in the image in English or Chinese language.\nFor image tagging and captioning."
        },
        {
            "author": "tzwm",
            "title": "ComfyUI Profiler",
            "reference": "https://github.com/tzwm/comfyui-profiler",
            "files": [
                "https://github.com/tzwm/comfyui-profiler"
            ],
            "install_type": "git-clone",
            "description": "Calculate the execution time of all nodes."
        },
        {
            "author": "Daniel Lewis",
            "title": "ComfyUI-Llama",
            "reference": "https://github.com/daniel-lewis-ab/ComfyUI-Llama",
            "files": [
                "https://github.com/daniel-lewis-ab/ComfyUI-Llama"
            ],
            "install_type": "git-clone",
            "description": "This is a set of nodes to interact with llama-cpp-python"
        },
        {
            "author": "Daniel Lewis",
            "title": "ComfyUI-TTS",
            "reference": "https://github.com/daniel-lewis-ab/ComfyUI-TTS",
            "files": [
                "https://github.com/daniel-lewis-ab/ComfyUI-TTS"
            ],
            "install_type": "git-clone",
            "description": "Text To Speech (TTS) for ComfyUI"
        },
        {
            "author": "djbielejeski",
            "title": "a-person-mask-generator",
            "reference": "https://github.com/djbielejeski/a-person-mask-generator",
            "files": [
                "https://github.com/djbielejeski/a-person-mask-generator"
            ],
            "install_type": "git-clone",
            "description": "Extension for Automatic1111 and ComfyUI to automatically create masks for Background/Hair/Body/Face/Clothes in Img2Img"
        },
        {
            "author": "smagnetize",
            "title": "kb-comfyui-nodes",
            "reference": "https://github.com/smagnetize/kb-comfyui-nodes",
            "files": [
                "https://github.com/smagnetize/kb-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SingleImageDataUrlLoader"
        },
        {
            "author": "ginlov",
            "title": "segment_to_mask_comfyui",
            "reference": "https://github.com/ginlov/segment_to_mask_comfyui",
            "files": [
                "https://github.com/ginlov/segment_to_mask_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SegToMask"
        },
        {
            "author": "glowcone",
            "title": "Load Image From Base64 URI",
            "reference": "https://github.com/glowcone/comfyui-base64-to-image",
            "files": [
                "https://github.com/glowcone/comfyui-base64-to-image"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LoadImageFromBase64. Loads an image and its transparency mask from a base64-encoded data URI for easy API connection."
        },
        {
            "author": "glowcone",
            "title": "String Converter",
            "reference": "https://github.com/glowcone/comfyui-string-converter",
            "files": [
                "https://github.com/glowcone/comfyui-string-converter"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Convert String To Int, Convert String To Float"
        },
        {
            "author": "AInseven",
            "title": "ComfyUI-fastblend",
            "reference": "https://github.com/AInseven/ComfyUI-fastblend",
            "files": [
                "https://github.com/AInseven/ComfyUI-fastblend"
            ],
            "install_type": "git-clone",
            "description": "fastblend for comfyui, and other nodes that I write for video2video. rebatch image, my openpose"
        },
        {
            "author": "HebelHuber",
            "title": "comfyui-enhanced-save-node",
            "reference": "https://github.com/HebelHuber/comfyui-enhanced-save-node",
            "files": [
                "https://github.com/HebelHuber/comfyui-enhanced-save-node"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Enhanced Save Node"
        },
        {
            "author": "LarryJane491",
            "title": "Lora-Training-in-Comfy",
            "reference": "https://github.com/LarryJane491/Lora-Training-in-Comfy",
            "files": [
                "https://github.com/LarryJane491/Lora-Training-in-Comfy"
            ],
            "install_type": "git-clone",
            "description": "If you see this message, your ComfyUI-Manager is outdated.\nRecent channel provides only the list of the latest nodes. If you want to find the complete node list, please go to the Default channel.\nMaking LoRA has never been easier!"
        },
        {
            "author": "LarryJane491",
            "title": "Image-Captioning-in-ComfyUI",
            "reference": "https://github.com/LarryJane491/Image-Captioning-in-ComfyUI",
            "files": [
                "https://github.com/LarryJane491/Image-Captioning-in-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "The LoRA Caption custom nodes, just like their name suggests, allow you to caption images so they are ready for LoRA training."
        },
        {
            "author": "Layer-norm",
            "title": "Comfyui lama remover",
            "reference": "https://github.com/Layer-norm/comfyui-lama-remover",
            "files": [
                "https://github.com/Layer-norm/comfyui-lama-remover"
            ],
            "install_type": "git-clone",
            "description": "A very simple ComfyUI node to remove item with mask."
        },
        {
            "author": "Taremin",
            "title": "ComfyUI Prompt ExtraNetworks",
            "reference": "https://github.com/Taremin/comfyui-prompt-extranetworks",
            "files": [
                "https://github.com/Taremin/comfyui-prompt-extranetworks"
            ],
            "install_type": "git-clone",
            "description": "Instead of LoraLoader or HypernetworkLoader, it receives a prompt and loads and applies LoRA or HN based on the specifications within the prompt. The main purpose of this custom node is to allow changes without reconnecting the LoraLoader node when the prompt is randomly altered, etc."
        },
        {
            "author": "Taremin",
            "title": "ComfyUI String Tools",
            "reference": "https://github.com/Taremin/comfyui-string-tools",
            "files": [
                "https://github.com/Taremin/comfyui-string-tools"
            ],
            "install_type": "git-clone",
            "description": " This extension provides the StringToolsConcat node, which concatenates multiple texts, and the StringToolsRandomChoice node, which selects one randomly from multiple texts."
        },
        {
            "author": "Taremin",
            "title": "WebUI Monaco Prompt",
            "reference": "https://github.com/Taremin/webui-monaco-prompt",
            "files": [
                "https://github.com/Taremin/webui-monaco-prompt"
            ],
            "install_type": "git-clone",
            "description": "Make it possible to edit the prompt using the Monaco Editor, an editor implementation used in VSCode.\nNOTE: This extension supports both ComfyUI and A1111 simultaneously."
        },
        {
            "author": "foxtrot-roger",
            "title": "RF Nodes",
            "reference": "https://github.com/foxtrot-roger/comfyui-rf-nodes",
            "files": [
                "https://github.com/foxtrot-roger/comfyui-rf-nodes"
            ],
            "install_type": "git-clone",
            "description": "A bunch of nodes that can be useful to manipulate primitive types (numbers, text, ...) Also some helpers to generate text and timestamps."
        },
        {
            "author": "abyz22",
            "title": "image_control",
            "reference": "https://github.com/abyz22/image_control",
            "files": [
                "https://github.com/abyz22/image_control"
            ],
            "install_type": "git-clone",
            "description": "Nodes:abyz22_Padding Image, abyz22_ImpactWildcardEncode, abyz22_setimageinfo, abyz22_SaveImage, abyz22_ImpactWildcardEncode_GetPrompt, abyz22_SetQueue, abyz22_drawmask, abyz22_FirstNonNull, abyz22_blendimages, abyz22_blend_onecolor. Please check workflow in [a/https://mirror.ghproxy.com/https://github.com/abyz22/image_control](https://mirror.ghproxy.com/https://github.com/abyz22/image_control)"
        },
        {
            "author": "HAL41",
            "title": "ComfyUI aichemy nodes",
            "reference": "https://github.com/HAL41/ComfyUI-aichemy-nodes",
            "files": [
                "https://github.com/HAL41/ComfyUI-aichemy-nodes"
            ],
            "install_type": "git-clone",
            "description": "Simple node to handle scaling of YOLOv8 segmentation masks"
        },
        {
            "author": "nkchocoai",
            "title": "ComfyUI-SizeFromPresets",
            "reference": "https://github.com/nkchocoai/ComfyUI-SizeFromPresets",
            "files": [
                "https://github.com/nkchocoai/ComfyUI-SizeFromPresets"
            ],
            "install_type": "git-clone",
            "description": "Add a node that outputs width and height of the size selected from the preset (.csv)."
        },
        {
            "author": "nkchocoai",
            "title": "ComfyUI-PromptUtilities",
            "reference": "https://github.com/nkchocoai/ComfyUI-PromptUtilities",
            "files": [
                "https://github.com/nkchocoai/ComfyUI-PromptUtilities"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Format String, Join String List, Load Preset, Load Preset (Advanced), Const String, Const String (multi line). Add useful nodes related to prompt."
        },
        {
            "author": "nkchocoai",
            "title": "ComfyUI-TextOnSegs",
            "reference": "https://github.com/nkchocoai/ComfyUI-TextOnSegs",
            "files": [
                "https://github.com/nkchocoai/ComfyUI-TextOnSegs"
            ],
            "install_type": "git-clone",
            "description": "Add a node for drawing text with CR Draw Text of ComfyUI_Comfyroll_CustomNodes to the area of SEGS detected by Ultralytics Detector of ComfyUI-Impact-Pack."
        },
        {
            "author": "nkchocoai",
            "title": "ComfyUI-SaveImageWithMetaData",
            "reference": "https://github.com/nkchocoai/ComfyUI-SaveImageWithMetaData",
            "files": [
                "https://github.com/nkchocoai/ComfyUI-SaveImageWithMetaData"
            ],
            "install_type": "git-clone",
            "description": "Add a node to save images with metadata (PNGInfo) extracted from the input values of each node.\nSince the values are extracted dynamically, values output by various extension nodes can be added to metadata."
        },
        {
            "author": "nkchocoai",
            "title": "ComfyUI-Dart",
            "id": "dart",
            "reference": "https://github.com/nkchocoai/ComfyUI-Dart",
            "files": [
                "https://github.com/nkchocoai/ComfyUI-Dart"
            ],
            "install_type": "git-clone",
            "description": "Add nodes that generates danbooru tags by [a/Dart(Danbooru Tags Transformer)](https://huggingface.co/p1atdev/dart-v1-sft)."
        },
        {
            "author": "JaredTherriault",
            "title": "ComfyUI-JNodes",
            "id": "jnodes",
            "reference": "https://github.com/JaredTherriault/ComfyUI-JNodes",
            "files": [
                "https://github.com/JaredTherriault/ComfyUI-JNodes"
            ],
            "install_type": "git-clone",
            "description": "python and web UX improvements for ComfyUI: Lora/Embedding picker, web extension manager (enable/disable any extension without disabling python nodes), control any parameter with text prompts, image and video viewer, metadata viewer, token counter, comments in prompts, font control, and more! \n[w/'ImageFeed.js' from the custom scripts of pythongosssss is not compatible with this suite's ImageDrawer feature. Additionally, 'DynamicPrompts.js' and 'EditAttention.js' from the core, along with 'favicon.js' from the custom scripts of pythongosssss, are incompatible with advanced features of the suite. Please use the JNodes Extension Management setting in Settings > JNodes > Extension Management to disable these extensions by unchecking them to use the full functionality of the suite.]"
        },
        {
            "author": "prozacgod",
            "title": "ComfyUI Multi-Workspace",
            "id": "multi-workspace",
            "reference": "https://github.com/prozacgod/comfyui-pzc-multiworkspace",
            "files": [
                "https://github.com/prozacgod/comfyui-pzc-multiworkspace"
            ],
            "install_type": "git-clone",
            "description": "A simple, quick, and dirty implementation of multiple workspaces within ComfyUI."
        },
        {
            "author": "Siberpone",
            "title": "Lazy Pony Prompter",
            "id": "lazy-pony-prompter",
            "reference": "https://github.com/Siberpone/lazy-pony-prompter",
            "files": [
                "https://github.com/Siberpone/lazy-pony-prompter"
            ],
            "install_type": "git-clone",
            "description": "A booru API powered prompt generator for A1111 and ComfyUI with flexible tag filtering system and customizable prompt templates."
        },
        {
            "author": "dave-palt",
            "title": "comfyui_DSP_imagehelpers",
            "id": "dsp-imagehelpers",
            "reference": "https://github.com/dave-palt/comfyui_DSP_imagehelpers",
            "files": [
                "https://github.com/dave-palt/comfyui_DSP_imagehelpers"
            ],
            "install_type": "git-clone",
            "description": "Nodes: DSP Image Concat"
        },
        {
            "author": "Inzaniak",
            "title": "Ranbooru for ComfyUI",
            "id": "ranbooru",
            "reference": "https://github.com/Inzaniak/comfyui-ranbooru",
            "files": [
                "https://github.com/Inzaniak/comfyui-ranbooru"
            ],
            "install_type": "git-clone",
            "description": "Ranbooru is an extension for the comfyUI. The purpose of this extension is to add a node that gets a random set of tags from boorus pictures. This is mostly being used to help me test my checkpoints on a large variety of"
        },
        {
            "author": "miosp",
            "title": "ComfyUI-FBCNN",
            "id": "fbcnn",
            "reference": "https://github.com/Miosp/ComfyUI-FBCNN",
            "files": [
                "https://github.com/Miosp/ComfyUI-FBCNN"
            ],
            "install_type": "git-clone",
            "description": "A node for JPEG de-artifacting using [a/FBCNN](https://mirror.ghproxy.com/https://github.com/jiaxi-jiang/FBCNN)."
        },
        {
            "author": "JcandZero",
            "title": "ComfyUI_GLM4Node",
            "id": "glm4node",
            "reference": "https://github.com/JcandZero/ComfyUI_GLM4Node",
            "files": [
                "https://github.com/JcandZero/ComfyUI_GLM4Node"
            ],
            "install_type": "git-clone",
            "description": "GLM4 Vision Integration"
        },
        {
            "author": "darkpixel",
            "title": "DarkPrompts",
            "id": "darkprompts",
            "reference": "https://github.com/darkpixel/darkprompts",
            "files": [
                "https://github.com/darkpixel/darkprompts"
            ],
            "install_type": "git-clone",
            "description": "Slightly better random prompt generation tools that allow combining and picking prompts from both file and text input sources."
        },
        {
            "author": "yytdfc",
            "title": "Amazon Bedrock nodes for ComfyUI",
            "id": "bedrock",
            "reference": "https://github.com/aws-samples/comfyui-llm-node-for-amazon-bedrock",
            "files": [
                "https://github.com/aws-samples/comfyui-llm-node-for-amazon-bedrock"
            ],
            "pip": ["boto3"],
            "install_type": "git-clone",
            "description": "Amazon Bedrock is a fully managed service that offers a choice of high-performing foundation models (FMs) from leading AI companies. This repo is the ComfyUI nodes for Bedrock service. You could invoke the foundation model in your ComfyUI pipeline."
        },
        {
            "author": "Qais Malkawi",
            "title": "ComfyUI-Qais-Helper",
            "id": "qais-helper",
            "reference": "https://github.com/QaisMalkawi/ComfyUI-QaisHelper",
            "files": [
                "https://github.com/QaisMalkawi/ComfyUI-QaisHelper"
            ],
            "install_type": "git-clone",
            "description": "This Extension adds a few custom QOL nodes that ComfyUI lacks by default."
        },
        {
            "author": "longgui0318",
            "title": "comfyui-mask-util",
            "id": "mask-util",
            "reference": "https://github.com/longgui0318/comfyui-mask-util",
            "files": [
                "https://github.com/longgui0318/comfyui-mask-util"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Split Masks, Mask Selection Of Masks, Mask Region Info"
        },
        {
            "author": "longgui0318",
            "title": "comfyui-llm-assistant",
            "id": "llm-assistant",
            "reference": "https://github.com/longgui0318/comfyui-llm-assistant",
            "files": [
                "https://github.com/longgui0318/comfyui-llm-assistant"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Generate Stable Diffsution Prompt With LLM, Translate Text With LLM, Chat With LLM"
        },
        {
            "author": "longgui0318",
            "title": "comfyui-magic-clothing",
            "id": "magic-clothing",
            "reference": "https://github.com/longgui0318/comfyui-magic-clothing",
            "files": [
                "https://github.com/longgui0318/comfyui-magic-clothing"
            ],
            "install_type": "git-clone",
            "description": "The comfyui supported version of the [a/Magic Clothing](https://mirror.ghproxy.com/https://github.com/ShineChen1024/MagicClothing) project, not the diffusers version, allows direct integration with modules such as ipadapter.[w/comfyui-oms-diffusion is renamed to comfyui-magic-clothing. You may need to reinstall this.]"
        },
        {
            "author": "longgui0318",
            "title": "comfyui-common-util",
            "id": "common-util",
            "reference": "https://github.com/longgui0318/comfyui-common-util",
            "files": [
                "https://github.com/longgui0318/comfyui-common-util"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Init Layer Info Array, Added Layer Info To Array, Layer Info Array Fuse, Layer Image Seleted, Layer Images IPAdapter Advanced, Enhanced Random Light Source"
        },
        {
            "author": "DimaChaichan",
            "title": "LAizypainter-Exporter-ComfyUI",
            "reference": "https://github.com/DimaChaichan/LAizypainter-Exporter-ComfyUI",
            "files": [
                "https://github.com/DimaChaichan/LAizypainter-Exporter-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This exporter is a plugin for ComfyUI, which can export tasks for [a/LAizypainter](https://mirror.ghproxy.com/https://github.com/DimaChaichan/LAizypainter).\nLAizypainter is a Photoshop plugin with which you can send tasks directly to a Stable Diffusion server. More information about a [a/Task](https://mirror.ghproxy.com/https://github.com/DimaChaichan/LAizypainter?tab=readme-ov-file#task)"
        },
        {
            "author": "adriflex",
            "title": "ComfyUI_Blender_Texdiff",
            "id": "blender-texdiff",
            "reference": "https://github.com/adriflex/ComfyUI_Blender_Texdiff",
            "files": [
                "https://github.com/adriflex/ComfyUI_Blender_Texdiff"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Blender viewport color, Blender Viewport depth"
        },
        {
            "author": "Shraknard",
            "title": "ComfyUI-Remover",
            "id": "remover",
            "reference": "https://github.com/Shraknard/ComfyUI-Remover",
            "files": [
                "https://github.com/Shraknard/ComfyUI-Remover"
            ],
            "install_type": "git-clone",
            "description": "Custom node for ComfyUI that makes parts of the image transparent (face, background...)"
        },
        {
            "author": "FlyingFireCo",
            "title": "tiled_ksampler",
            "reference": "https://github.com/FlyingFireCo/tiled_ksampler",
            "files": [
                "https://github.com/FlyingFireCo/tiled_ksampler"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Tiled KSampler, Asymmetric Tiled KSampler, Circular VAEDecode."
        },
        {
            "author": "Nlar",
            "title": "ComfyUI_CartoonSegmentation",
            "id": "cartoon-seg",
            "reference": "https://github.com/Nlar/ComfyUI_CartoonSegmentation",
            "files": [
                "https://github.com/Nlar/ComfyUI_CartoonSegmentation"
            ],
            "install_type": "git-clone",
            "description": "Front end ComfyUI nodes for CartoonSegmentation Based upon the work of the CartoonSegmentation repository this project will provide a front end to some of the features."
        },
        {
            "author": "godspede",
            "title": "ComfyUI Substring",
            "id": "substring",
            "reference": "https://github.com/godspede/ComfyUI_Substring",
            "files": [
                "https://github.com/godspede/ComfyUI_Substring"
            ],
            "install_type": "git-clone",
            "description": "Just a simple substring node that takes text and length as input, and outputs the first length characters."
        },
        {
            "author": "gokayfem",
            "title": "VLM_nodes",
            "id": "vlm",
            "reference": "https://github.com/gokayfem/ComfyUI_VLM_nodes",
            "files": [
                "https://github.com/gokayfem/ComfyUI_VLM_nodes"
            ],
            "install_type": "git-clone",
            "description": "Custom Nodes for Vision Language Models (VLM) , Large Language Models (LLM), Image Captioning, Automatic Prompt Generation, Creative and Consistent Prompt Suggestion, Keyword Extraction"
        },
        {
            "author": "gokayfem",
            "title": "ComfyUI-Dream-Interpreter",
            "id": "dream-interpreter",
            "reference": "https://github.com/gokayfem/ComfyUI-Dream-Interpreter",
            "files": [
                "https://github.com/gokayfem/ComfyUI-Dream-Interpreter"
            ],
            "install_type": "git-clone",
            "description": "Tell your dream and it interprets it and puts you inside your dream"
        },
        {
            "author": "gokayfem",
            "title": "ComfyUI-Depth-Visualization",
            "id": "delpth-visualization",
            "reference": "https://github.com/gokayfem/ComfyUI-Depth-Visualization",
            "files": [
                "https://github.com/gokayfem/ComfyUI-Depth-Visualization"
            ],
            "install_type": "git-clone",
            "description": "Works with any Depth Map and visualizes the applied version it inside ComfyUI"
        },
        {
            "author": "gokayfem",
            "title": "ComfyUI-Texture-Simple",
            "id": "texture-simple",
            "reference": "https://github.com/gokayfem/ComfyUI-Texture-Simple",
            "files": [
                "https://github.com/gokayfem/ComfyUI-Texture-Simple"
            ],
            "install_type": "git-clone",
            "description": "Visualize your textures inside ComfyUI"
        },
        {
            "author": "Hiero207",
            "title": "Hiero-Nodes",
            "id": "hiero",
            "reference": "https://github.com/Hiero207/ComfyUI-Hiero-Nodes",
            "files": [
                "https://github.com/Hiero207/ComfyUI-Hiero-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Post to Discord w/ Webhook"
        },
        {
            "author": "azure-dragon-ai",
            "title": "ComfyUI-ClipScore-Nodes",
            "id": "clipscore",
            "reference": "https://github.com/azure-dragon-ai/ComfyUI-ClipScore-Nodes",
            "files": [
                "https://github.com/azure-dragon-ai/ComfyUI-ClipScore-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageScore, Loader, Image Processor, Real Image Processor, Fake Image Processor, Text Processor. ComfyUI Nodes for ClipScore"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Whisper",
            "id": "whisper",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Whisper",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Whisper"
            ],
            "install_type": "git-clone",
            "description": "Transcribe audio and add subtitles to videos using Whisper in ComfyUI"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI-Pronodes",
            "id": "pronodes",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Pronodes",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Pronodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of nice utility nodes for ComfyUI"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI-Vsgan",
            "id": "vsgan",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Vsgan",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Vsgan"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Upscale Video Tensorrt"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Depth Anything TensorRT",
            "id": "depth-anything-tensorrt",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Depth-Anything-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Depth-Anything-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This extension provides a ComfyUI Custom Node implementation of the [a/Depth-Anything-Tensorrt](https://mirror.ghproxy.com/https://github.com/spacewalk01/depth-anything-tensorrt) in Python for ultra fast depth map generation"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI PiperTTS",
            "id": "pipertts",
            "reference": "https://github.com/yuvraj108c/ComfyUI-PiperTTS",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-PiperTTS"
            ],
            "install_type": "git-clone",
            "description": "Convert Text-to-Speech inside ComfyUI using [a/Piper](https://mirror.ghproxy.com/https://github.com/rhasspy/piper)"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Upscaler TensorRT",
            "id": "upscaler-tensorrt",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Upscaler-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Upscaler-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This project provides a Tensorrt implementation for fast image upscaling inside ComfyUI (3-4x faster)"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI YoloNasPose Tensorrt",
            "id": "yolonaspose-tensorrt",
            "reference": "https://github.com/yuvraj108c/ComfyUI-YoloNasPose-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-YoloNasPose-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This repo provides a ComfyUI Custom Node implementation of [a/YOLO-NAS-POSE](https://mirror.ghproxy.com/https://github.com/Deci-AI/super-gradients), powered by TensorRT for ultra fast pose estimation. It has been adapted to work with openpose controlnet (experimental)"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Dwpose TensorRT",
            "id": "dwpose-tensorrt",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Dwpose-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Dwpose-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This project provides a Tensorrt implementation of Dwpose for ultra fast pose estimation inside ComfyUI"
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Facerestore TensorRT",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Facerestore-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Facerestore-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This project provides an experimental Tensorrt implementation for ultra fast face restoration inside ComfyUI.\nNote: This project doesn't do pre/post processing. It only works on cropped faces for now."
        },
        {
            "author": "yuvraj108c",
            "title": "ComfyUI Rife TensorRT",
            "reference": "https://github.com/yuvraj108c/ComfyUI-Rife-Tensorrt",
            "files": [
                "https://github.com/yuvraj108c/ComfyUI-Rife-Tensorrt"
            ],
            "install_type": "git-clone",
            "description": "This project provides a TensorRT implementation of [a/RIFE](https://mirror.ghproxy.com/https://github.com/hzwer/ECCV2022-RIFE) for ultra fast frame interpolation inside ComfyUI"
        },
        {
            "author": "blepping",
            "title": "ComfyUI-bleh",
            "id": "bleh",
            "reference": "https://github.com/blepping/ComfyUI-bleh",
            "files": [
                "https://github.com/blepping/ComfyUI-bleh"
            ],
            "install_type": "git-clone",
            "description": "Better TAESD previews, BlehHyperTile."
        },
        {
            "author": "blepping",
            "title": "ComfyUI-sonar",
            "id": "sonar",
            "reference": "https://github.com/blepping/ComfyUI-sonar",
            "files": [
                "https://github.com/blepping/ComfyUI-sonar"
            ],
            "install_type": "git-clone",
            "description": "A janky implementation of Sonar sampling (momentum-based sampling) for ComfyUI."
        },
        {
            "author": "blepping",
            "title": "comfyui_jankhidiffusion",
            "id": "jank-hidiffusion",
            "reference": "https://github.com/blepping/comfyui_jankhidiffusion",
            "files": [
                "https://github.com/blepping/comfyui_jankhidiffusion"
            ],
            "install_type": "git-clone",
            "description": "Janky implementation of [a/HiDiffusion](https://mirror.ghproxy.com/https://github.com/megvii-research/HiDiffusion) for ComfyUI. Enables generating at resolutions higher than what the model was trained for. Only supports SD 1.x (maybe 2.x) and SDXL."
        },
        {
            "author": "blepping",
            "title": "comfyui_jankdiffusehigh",
            "id": "jank-diffusehigh",
            "reference": "https://github.com/blepping/comfyui_jankdiffusehigh",
            "files": [
                "https://github.com/blepping/comfyui_jankdiffusehigh"
            ],
            "install_type": "git-clone",
            "description": "Janky implementation of [a/DiffuseHigh](https://mirror.ghproxy.com/https://github.com/yhyun225/DiffuseHigh/) for ComfyUI. Enables generating at resolutions higher than what the model was trained for without requiring model patches."
        },
        {
            "author": "blepping",
            "title": "comfyui_overly_complicated_sampling",
            "reference": "https://github.com/blepping/comfyui_overly_complicated_sampling",
            "files": [
                  "https://github.com/blepping/comfyui_overly_complicated_sampling"
            ],
            "install_type": "git-clone",
            "description": "Experimental and mathematically unsound (but fun!) sampling for ComfyUI.\nFeel free create a question in Discussions for usage help: OCS Q&A Discussion[w/Status: In flux, may be useful but likely to change/break workflows frequently. Mainly for advanced users.]"
        },
        {
            "author": "JerryOrbachJr",
            "title": "Random Size",
            "reference": "https://github.com/JerryOrbachJr/ComfyUI-RandomSize",
            "files": [
                "https://github.com/JerryOrbachJr/ComfyUI-RandomSize"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node that randomly selects a height and width pair from a list in a config file"
        },
        {
            "author": "jamal-alkharrat",
            "title": "ComfyUI_rotate_image",
            "reference": "https://github.com/jamal-alkharrat/ComfyUI_rotate_image",
            "files": [
                "https://github.com/jamal-alkharrat/ComfyUI_rotate_image"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Custom Node to Rotate Images, Img2Img node."
        },
        {
            "author": "mape",
            "title": "mape's helpers",
            "id": "mape-helpers",
            "reference": "https://github.com/mape/ComfyUI-mape-Helpers",
            "files": [
                "https://github.com/mape/ComfyUI-mape-Helpers"
            ],
            "install_type": "git-clone",
            "description": "Multi-monitor image preview, Variable Assigment/Wireless Nodes, Prompt Tweaking, Command Palette, Pinned favourite nodes, Node navigation, Fuzzy search, Node time tracking, Organizing and Error management. For more info visit: [a/https://comfyui.ma.pe/](https://comfyui.ma.pe/)"
        },
        {
            "author": "zhongpei",
            "title": "ComfyUI for InstructIR",
            "id": "instructir",
            "reference": "https://github.com/zhongpei/ComfyUI-InstructIR",
            "files": [
                "https://github.com/zhongpei/ComfyUI-InstructIR"
            ],
            "install_type": "git-clone",
            "description": "Enhancing Image Restoration. (ref:[a/InstructIR](https://mirror.ghproxy.com/https://github.com/mv-lab/InstructIR))"
        },
        {
            "author": "Loewen-Hob",
            "title": "Rembg Background Removal Node for ComfyUI (Better)",
            "id": "rembg-better",
            "reference": "https://github.com/Loewen-Hob/rembg-comfyui-node-better",
            "files": [
                "https://github.com/Loewen-Hob/rembg-comfyui-node-better"
            ],
            "install_type": "git-clone",
            "description": "This custom node is based on the [a/rembg-comfyui-node](https://mirror.ghproxy.com/https://github.com/Jcd1230/rembg-comfyui-node) but provides additional functionality to select ONNX models."
        },
        {
            "author": "HaydenReeve",
            "title": "ComfyUI Better Strings",
            "id": "better-string",
            "reference": "https://github.com/HaydenReeve/ComfyUI-Better-Strings",
            "files": [
                "https://github.com/HaydenReeve/ComfyUI-Better-Strings"
            ],
            "install_type": "git-clone",
            "description": "Strings should be easy, and simple. This extension aims to provide a set of nodes that make working with strings in ComfyUI a little bit easier."
        },
        {
            "author": "StartHua",
            "title": "ComfyUI_Seg_VITON",
            "id": "seg-viton",
            "reference": "https://github.com/StartHua/ComfyUI_Seg_VITON",
            "files": [
                "https://github.com/StartHua/ComfyUI_Seg_VITON"
            ],
            "install_type": "git-clone",
            "description": "Nodes:segformer_clothes, segformer_agnostic, segformer_remove_bg, stabel_vition. Nodes for model dress up."
        },
        {
            "author": "StartHua",
            "title": "Comfyui_joytag",
            "id": "joytag",
            "reference": "https://github.com/StartHua/Comfyui_joytag",
            "files": [
                "https://github.com/StartHua/Comfyui_joytag"
            ],
            "install_type": "git-clone",
            "description": "JoyTag is a state of the art AI vision model for tagging images, with a focus on sex positivity and inclusivity. It uses the Danbooru tagging schema, but works across a wide range of images, from hand drawn to photographic.\nDownload the weight and put it under checkpoints: [a/https://huggingface.co/fancyfeast/joytag/tree/main](https://huggingface.co/fancyfeast/joytag/tree/main)"
        },
        {
            "author": "StartHua",
            "title": "comfyui_segformer_b2_clothes",
            "id": "segformer-b2-clothes",
            "reference": "https://github.com/StartHua/Comfyui_segformer_b2_clothes",
            "files": [
                "https://github.com/StartHua/Comfyui_segformer_b2_clothes"
            ],
            "install_type": "git-clone",
            "description": "SegFormer model fine-tuned on ATR dataset for clothes segmentation but can also be used for human segmentation!\nDownload the weight and put it under checkpoints: [a/https://huggingface.co/mattmdjaga/segformer_b2_clothes](https://huggingface.co/mattmdjaga/segformer_b2_clothes)"
        },
        {
            "author": "StartHua",
            "title": "ComfyUI_OOTDiffusion_CXH",
            "id": "ootdiffusion-cxh",
            "reference": "https://github.com/StartHua/ComfyUI_OOTDiffusion_CXH",
            "files": [
                "https://github.com/StartHua/ComfyUI_OOTDiffusion_CXH"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Ood_hd_CXH, Ood_hd_CXH. [a/OOTDiffusion](https://mirror.ghproxy.com/https://github.com/levihsu/OOTDiffusion)"
        },
        {
            "author": "StartHua",
            "title": "ComfyUI_PCDMs",
            "id": "pcdms",
            "reference": "https://github.com/StartHua/ComfyUI_PCDMs",
            "files": [
                "https://github.com/StartHua/ComfyUI_PCDMs"
            ],
            "install_type": "git-clone",
            "description": "Original project: [a/link](https://mirror.ghproxy.com/https://github.com/tencent-ailab/PCDMs)\nBased on testing, the author's original images work very well, but using my own images generally requires some luck!"
        },
        {
            "author": "StartHua",
            "title": "Comfyui_CXH_joy_caption",
            "reference": "https://github.com/StartHua/Comfyui_CXH_joy_caption",
            "files": [
                "https://github.com/StartHua/Comfyui_CXH_joy_caption"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Joy_caption_load, Joy_caption"
        },
        {
            "author": "StartHua",
            "title": "Comfyui_CXH_Phi_3.5",
            "reference": "https://github.com/StartHua/Comfyui_CXH_Phi_3.5",
            "files": [
                "https://github.com/StartHua/Comfyui_CXH_Phi_3.5"
            ],
            "install_type": "git-clone",
            "description": "Phi-3.5-vision-instruct fast talk with image !\nFast , Fast ,Fast!\n1.Phi-3.5-vision-instruct"
        },
        {
            "author": "StartHua",
            "title": "Comfyui_CXH_DeepLX",
            "reference": "https://github.com/StartHua/Comfyui_CXH_DeepLX",
            "files": [
                "https://github.com/StartHua/Comfyui_CXH_DeepLX"
            ],
            "install_type": "git-clone",
            "description": "NODES:CXH_DeepLX_Free, CXH_DeepLX_translate"
        },
        {
            "author": "StartHua",
            "title": "Comfyui_CXH_FluxLoraMerge",
            "reference": "https://github.com/StartHua/Comfyui_CXH_FluxLoraMerge",
            "files": [
                "https://github.com/StartHua/Comfyui_CXH_FluxLoraMerge"
            ],
            "install_type": "git-clone",
            "description": "flux lora merge.\nadaptive Merge (uses tensor norms and weight), manual Merge (uses fixed weights you specify), additive Merge (uses 100% of the first and adds a percentage of the second)"
        },
        {
            "author": "ricklove",
            "title": "comfyui-ricklove",
            "id": "ricklove",
            "reference": "https://github.com/ricklove/comfyui-ricklove",
            "files": [
                "https://github.com/ricklove/comfyui-ricklove"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Image Crop and Resize by Mask, Image Uncrop, Image Shadow, Optical Flow (Dip), Warp Image with Flow, Image Threshold (Channels), Finetune Variable, Finetune Analyze, Finetune Analyze Batch, ... Misc ComfyUI nodes by Rick Love"
        },
        {
            "author": "nosiu",
            "title": "ComfyUI InstantID Faceswapper",
            "id": "instantid-faceswapper",
            "reference": "https://github.com/nosiu/comfyui-instantId-faceswap",
            "files": [
                "https://github.com/nosiu/comfyui-instantId-faceswap"
            ],
            "install_type": "git-clone",
            "description": "Implementation of [a/faceswap](https://mirror.ghproxy.com/https://github.com/nosiu/InstantID-faceswap/tree/main) based on [a/InstantID](https://mirror.ghproxy.com/https://github.com/InstantID/InstantID) for ComfyUI. Allows usage of [a/LCM Lora](https://huggingface.co/latent-consistency/lcm-lora-sdxl) which can produce good results in only a few generation steps.\nNOTE:Works ONLY with SDXL checkpoints."
        },
        {
            "author": "LyazS",
            "title": "Anime Character Segmentation node for comfyui",
            "reference": "https://github.com/LyazS/comfyui-anime-seg",
            "files": [
                "https://github.com/LyazS/comfyui-anime-seg"
            ],
            "install_type": "git-clone",
            "description": "A Anime Character Segmentation node for comfyui, based on [this hf space](https://huggingface.co/spaces/skytnt/anime-remove-background)."
        },
        {
            "author": "LyazS",
            "title": "net tool node for comfyui",
            "reference": "https://github.com/LyazS/comfyui-nettools",
            "files": [
                "https://github.com/LyazS/comfyui-nettools"
            ],
            "install_type": "git-clone",
            "description": "A net tool node for comfyui, rewrite from [comfyui-tooling-nodes](https://mirror.ghproxy.com/https://github.com/Acly/comfyui-tooling-nodes) but support more big data sending."
        },
        {
            "author": "Chan-0312",
            "title": "ComfyUI-IPAnimate",
            "reference": "https://github.com/Chan-0312/ComfyUI-IPAnimate",
            "files": [
                "https://github.com/Chan-0312/ComfyUI-IPAnimate"
            ],
            "install_type": "git-clone",
            "description": "This is a project that generates videos frame by frame based on IPAdapter+ControlNet. Unlike [a/Steerable-motion](https://mirror.ghproxy.com/https://github.com/banodoco/Steerable-Motion), we do not rely on AnimateDiff. This decision is primarily due to the fact that the videos generated by AnimateDiff are often blurry. Through frame-by-frame control using IPAdapter+ControlNet, we can produce higher definition and more controllable videos."
        },
        {
            "author": "Chan-0312",
            "title": "ComfyUI-EasyDeforum",
            "reference": "https://github.com/Chan-0312/ComfyUI-EasyDeforum",
            "files": [
                "https://github.com/Chan-0312/ComfyUI-EasyDeforum"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Easy2DDeforum (Chan)"
        },
        {
            "author": "trumanwong",
            "title": "ComfyUI-NSFW-Detection",
            "reference": "https://github.com/trumanwong/ComfyUI-NSFW-Detection",
            "files": [
                "https://github.com/trumanwong/ComfyUI-NSFW-Detection"
            ],
            "install_type": "git-clone",
            "description": "An implementation of NSFW Detection for ComfyUI"
        },
        {
            "author": "TemryL",
            "title": "ComfyS3",
            "reference": "https://github.com/TemryL/ComfyS3",
            "files": [
                "https://github.com/TemryL/ComfyS3"
            ],
            "install_type": "git-clone",
            "description": "ComfyS3 seamlessly integrates with [a/Amazon S3](https://aws.amazon.com/en/s3/) in ComfyUI. This open-source project provides custom nodes for effortless loading and saving of images, videos, and checkpoint models directly from S3 buckets within the ComfyUI graph interface."
        },
        {
            "author": "MaraScott",
            "title": "🐰 MaraScott Nodes",
            "id": "marascott-nodes",
            "reference": "https://github.com/MaraScott/ComfyUI_MaraScott_Nodes",
            "files": [
                "https://github.com/MaraScott/ComfyUI_MaraScott_Nodes"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes including a universal bus, an Inpainting By Mask and a large Upscaler/Refiner\n[AnyBus,McInpainty,McBoaty]"
        },
        {
            "author": "yffyhk",
            "title": "comfyui_auto_danbooru",
            "reference": "https://github.com/yffyhk/comfyui_auto_danbooru",
            "files": [
                "https://github.com/yffyhk/comfyui_auto_danbooru"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Get Danbooru, Tag Encode"
        },
        {
            "author": "dfl",
            "title": "CLIP with BREAK syntax",
            "reference": "https://github.com/dfl/comfyui-clip-with-break",
            "files": [
                "https://github.com/dfl/comfyui-clip-with-break"
            ],
            "install_type": "git-clone",
            "description": "Clip text encoder with BREAK formatting like A1111 (uses conditioning concat)"
        },
        {
            "author": "dfl",
            "title": "ComfyUI-TCD-scheduler",
            "id": "dfl-tcd",
            "reference": "https://github.com/dfl/comfyui-tcd-scheduler",
            "files": [
                "https://github.com/dfl/comfyui-tcd-scheduler"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Custom Sampler nodes that implement Zheng et al.'s Trajectory Consistency Distillation based on [a/https://mhh0318.github.io/tcd](https://mhh0318.github.io/tcd)"
        },
        {
            "author": "antrobot",
            "title": "antrobots ComfyUI Nodepack",
            "reference": "https://github.com/antrobot1234/antrobots-comfyUI-nodepack",
            "files": [
                "https://github.com/antrobot1234/antrobots-comfyUI-nodepack"
            ],
            "install_type": "git-clone",
            "description": "A small node pack containing various things I felt like ought to be in base comfy-UI. Currently includes Some image handling nodes to help with inpainting, a version of KSampler (advanced) that allows for denoise, and a node that can swap it's inputs. Remember to make an issue if you experience any bugs or errors!"
        },
        {
            "author": "bilal-arikan",
            "title": "ComfyUI_TextAssets",
            "reference": "https://github.com/bilal-arikan/ComfyUI_TextAssets",
            "files": [
                "https://github.com/bilal-arikan/ComfyUI_TextAssets"
            ],
            "install_type": "git-clone",
            "description": "With this node you can upload text files to input folder from your local computer."
        },
        {
            "author": "kadirnar",
            "title": "ComfyUI-Transformers",
            "id": "comfy-transformers",
            "reference": "https://github.com/kadirnar/ComfyUI-Transformers",
            "files": [
                "https://github.com/kadirnar/ComfyUI-Transformers"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Transformers is a cutting-edge project combining the power of computer vision and natural language processing to create intuitive and user-friendly interfaces. Our goal is to make technology more accessible and engaging."
        },
        {
            "author": "kadirnar",
            "title": "ComfyUI-YOLO",
            "id": "comfy-yolo",
            "reference": "https://github.com/kadirnar/ComfyUI-YOLO",
            "files": [
                "https://github.com/kadirnar/ComfyUI-YOLO"
            ],
            "install_type": "git-clone",
            "description": "Ultralytics-Powered Object Recognition for ComfyUI"
        },
        {
            "author": "digitaljohn",
            "title": "ComfyUI-ProPost",
            "reference": "https://github.com/digitaljohn/comfyui-propost",
            "files": [
                "https://github.com/digitaljohn/comfyui-propost"
            ],
            "install_type": "git-clone",
            "description": "A set of custom ComfyUI nodes for performing basic post-processing effects including Film Grain and Vignette. These effects can help to take the edge off AI imagery and make them feel more natural."
        },
        {
            "author": "DonBaronFactory",
            "title": "ComfyUI-Cre8it-Nodes",
            "reference": "https://github.com/DonBaronFactory/ComfyUI-Cre8it-Nodes",
            "files": [
                "https://github.com/DonBaronFactory/ComfyUI-Cre8it-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:CRE8IT Serial Prompter, CRE8IT Apply Serial Prompter, CRE8IT Image Sizer. A few simple nodes to facilitate working wiht ComfyUI Workflows"
        },
        {
            "author": "deforum",
            "title": "Deforum Nodes",
            "reference": "https://github.com/XmYx/deforum-comfy-nodes",
            "files": [
                "https://github.com/XmYx/deforum-comfy-nodes"
            ],
            "install_type": "git-clone",
            "description": "Official Deforum animation pipeline tools that provide a unique way to create frame-by-frame generative motion art."
        },
        {
            "author": "adbrasi",
            "title": "ComfyUI-TrashNodes-DownloadHuggingface",
            "reference": "https://github.com/adbrasi/ComfyUI-TrashNodes-DownloadHuggingface",
            "files": [
                "https://github.com/adbrasi/ComfyUI-TrashNodes-DownloadHuggingface"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-TrashNodes-DownloadHuggingface is a ComfyUI node designed to facilitate the download of models you have just trained and uploaded to Hugging Face. This node is particularly useful for users who employ Google Colab for training and need to quickly download their models for deployment."
        },
        {
            "author": "mbrostami",
            "title": "ComfyUI-HF",
            "reference": "https://github.com/mbrostami/ComfyUI-HF",
            "files": [
                "https://github.com/mbrostami/ComfyUI-HF"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Node to work with Hugging Face repositories"
        },
        {
            "author": "Billius-AI",
            "title": "ComfyUI-Path-Helper",
            "reference": "https://github.com/Billius-AI/ComfyUI-Path-Helper",
            "files": [
                "https://github.com/Billius-AI/ComfyUI-Path-Helper"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Create Project Root, Add Folder, Add Folder Advanced, Add File Name Prefix, Add File Name Prefix Advanced, ShowPath"
        },
        {
            "author": "Franck-Demongin",
            "title": "NX_PromptStyler",
            "reference": "https://github.com/Franck-Demongin/NX_PromptStyler",
            "files": [
                "https://github.com/Franck-Demongin/NX_PromptStyler"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI to create a prompt based on a list of keywords saved in CSV files."
        },
        {
            "author": "Franck-Demongin",
            "title": "NX_HuggingFace_Flux",
            "reference": "https://github.com/Franck-Demongin/NX_HuggingFace_Flux",
            "files": [
                "https://github.com/Franck-Demongin/NX_HuggingFace_Flux"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Hugging Face Flux"
        },
        {
            "author": "Franck-Demongin",
            "title": "NX_Translator",
            "reference": "https://github.com/Franck-Demongin/NX_Translator",
            "files": [
                "https://github.com/Franck-Demongin/NX_Translator"
            ],
            "install_type": "git-clone",
            "description": "A custom node for translating prompts with Google Translate or DeeplL directly in ComfyUI."
        },
        {
            "author": "xiaoxiaodesha",
            "title": "hd-nodes-comfyui",
            "reference": "https://github.com/xiaoxiaodesha/hd_node",
            "files": [
                "https://github.com/xiaoxiaodesha/hd_node"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Combine HDMasks, Cover HDMasks, HD FaceIndex, HD SmoothEdge, HD GetMaskArea, HD Image Levels, HD Ultimate SD Upscale"
        },
        {
            "author": "ShmuelRonen",
            "title": "ComfyUI-SVDResizer",
            "id": "svdresizer",
            "reference": "https://github.com/ShmuelRonen/ComfyUI-SVDResizer",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI-SVDResizer"
            ],
            "install_type": "git-clone",
            "description": "SVDResizer is a helper for resizing the source image, according to the sizes enabled in Stable Video Diffusion. The rationale behind the possibility of changing the size of the image in steps between the ranges of 576 and 1024, is the use of the greatest common denominator of these two numbers which is 64. SVD is lenient with resizing that adheres to this rule, so the chance of coherent video that is not the standard size of 576X1024 is greater. It is advisable to keep the value 1024 constant and play with the second size to maintain the stability of the result."
        },
        {
            "author": "ShmuelRonen",
            "title": "Wav2Lip Node for ComfyUI",
            "id": "wav2lip",
            "reference": "https://github.com/ShmuelRonen/ComfyUI_wav2lip",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI_wav2lip"
            ],
            "install_type": "git-clone",
            "description": "The Wav2Lip node is a custom node for ComfyUI that allows you to perform lip-syncing on videos using the Wav2Lip model. It takes an input video and an audio file and generates a lip-synced output video."
        },
        {
            "author": "ShmuelRonen",
            "title": "ComfyUI_Gemini_Flash",
            "id": "gemini-flash",
            "reference": "https://github.com/ShmuelRonen/ComfyUI_Gemini_Flash",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI_Gemini_Flash"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI_Gemini_Flash is a custom node for ComfyUI, integrating the capabilities of the Gemini 1.5 Flash model. This node supports text and vision-based prompts, allowing users to analyze and adapt images to text prompts for text2image tasks."
        },
        {
            "author": "ShmuelRonen",
            "title": "ComfyUI_pixtral_vision",
            "reference": "https://github.com/ShmuelRonen/ComfyUI_pixtral_vision",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI_pixtral_vision"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI_pixtral_vision is a powerful ComfyUI node designed to integrate seamlessly with the Mistral Pixtral API. It facilitates the analysis of images through deep learning models, interpreting and describing the visual content. Users can input an image directly and provide prompts for context, utilizing an API key for authentication."
        },
        {
            "author": "ShmuelRonen",
            "title": "ComfyUI-FreeMemory",
            "reference": "https://github.com/ShmuelRonen/ComfyUI-FreeMemory",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI-FreeMemory"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-FreeMemory is a custom node extension for ComfyUI that provides advanced memory management capabilities within your image generation workflows. It aims to help prevent out-of-memory errors and optimize resource usage during complex operations."
        },
        {
            "author": "redhottensors",
            "title": "ComfyUI-Prediction",
            "id": "prediction",
            "reference": "https://github.com/redhottensors/ComfyUI-Prediction",
            "files": [
                "https://github.com/redhottensors/ComfyUI-Prediction"
            ],
            "install_type": "git-clone",
            "description": "Fully customizable Classifier Free Guidance for ComfyUI."
        },
        {
            "author": "Mamaaaamooooo",
            "title": "Batch Rembg for ComfyUI",
            "id": "batch-rembg",
            "reference": "https://github.com/Mamaaaamooooo/batchImg-rembg-ComfyUI-nodes",
            "files": [
                "https://github.com/Mamaaaamooooo/batchImg-rembg-ComfyUI-nodes"
            ],
            "install_type": "git-clone",
            "description": "Remove background of plural images."
        },
        {
            "author": "jordoh",
            "title": "ComfyUI Deepface",
            "id": "deepface",
            "reference": "https://github.com/jordoh/ComfyUI-Deepface",
            "files": [
                "https://github.com/jordoh/ComfyUI-Deepface"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes wrapping the [a/deepface](https://mirror.ghproxy.com/https://github.com/serengil/deepface) library."
        },
        {
            "author": "al-swaiti",
            "title": "ComfyUI-CascadeResolutions",
            "id": "cascade-resolution",
            "reference": "https://github.com/al-swaiti/ComfyUI-CascadeResolutions",
            "files": [
                "https://github.com/al-swaiti/ComfyUI-CascadeResolutions"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Cascade Resolutions"
        },
        {
            "author": "al-swaiti",
            "title": "All-IN-ONE-style",
            "id": "all-in-one-style",
            "reference": "https://github.com/al-swaiti/All-IN-ONE-style",
            "files": [
                "https://github.com/al-swaiti/All-IN-ONE-style"
            ],
            "install_type": "git-clone",
            "description": "all art styles"
        },
        {
            "author": "al-swaiti",
            "title": "GeminiOllama ComfyUI Extension",
            "reference": "https://github.com/al-swaiti/ComfyUI-OllamaGemini",
            "files": [
                "https://github.com/al-swaiti/ComfyUI-OllamaGemini"
            ],
            "install_type": "git-clone",
            "description": "This extension integrates Google's Gemini API and Ollama into ComfyUI, allowing users to leverage these powerful language models directly within their ComfyUI workflows."
        },
        {
            "author": "mirabarukaso",
            "title": "ComfyUI_Mira",
            "id": "mira",
            "reference": "https://github.com/mirabarukaso/ComfyUI_Mira",
            "files": [
                "https://github.com/mirabarukaso/ComfyUI_Mira"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Integer Multiplier, Float Multiplier, Convert Numeral to String, Create Canvas Advanced, Create Canvas, Create PNG Mask, Color Mask to HEX String, Color Mask to INT RGB, Color Masks to List"
        },
        {
            "author": "1038lab",
            "title": "ComfyUI-GPT2P",
            "id": "gpt2p",
            "reference": "https://github.com/1038lab/ComfyUI-GPT2P",
            "files": [
                "https://github.com/1038lab/ComfyUI-GPT2P"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Node - Hugging Face repositories GTP2 Prompt"
        },
        {
            "author": "Klinter",
            "title": "Klinter_nodes",
            "id": "klinter",
            "reference": "https://github.com/klinter007/klinter_nodes",
            "files": [
                "https://github.com/klinter007/klinter_nodes"
            ],
            "install_type": "git-clone",
            "description": "Concat_strings atm - celebrating first_node"
        },
        {
            "author": "Ludobico",
            "title": "ComfyUI-ScenarioPrompt",
            "id": "scenarioprompt",
            "reference": "https://github.com/Ludobico/ComfyUI-ScenarioPrompt",
            "files": [
                "https://github.com/Ludobico/ComfyUI-ScenarioPrompt"
            ],
            "install_type": "git-clone",
            "description": "ScenarioPrompt is a custom node that helps you understand what you're prompting for each property as you build your prompts"
        },
        {
            "author": "logtd",
            "title": "InstanceDiffusion Nodes",
            "id": "instancediffusion",
            "reference": "https://github.com/logtd/ComfyUI-InstanceDiffusion",
            "files": [
                "https://github.com/logtd/ComfyUI-InstanceDiffusion"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes to perform multi-object prompting with InstanceDiffusion"
        },
        {
            "author": "logtd",
            "title": "Tracking Nodes for Videos",
            "id": "tracking",
            "reference": "https://github.com/logtd/ComfyUI-TrackingNodes",
            "files": [
                "https://github.com/logtd/ComfyUI-TrackingNodes"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes to track objects through videos using YOLO and other processors."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-InversedNoise",
            "id": "inversed-noise",
            "reference": "https://github.com/logtd/ComfyUI-InversedNoise",
            "files": [
                "https://github.com/logtd/ComfyUI-InversedNoise"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Inversed Euler Sampler, Mix Noise with Latent, Combine Latent Noise"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-RefSampling",
            "id": "refsampling",
            "reference": "https://github.com/logtd/ComfyUI-RefSampling",
            "files": [
                "https://github.com/logtd/ComfyUI-RefSampling"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Apply Ref UNet, Ref Sampler, Ref Sampler Custom"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-FLATTEN",
            "id": "flatten",
            "reference": "https://github.com/logtd/ComfyUI-FLATTEN",
            "files": [
                "https://github.com/logtd/ComfyUI-FLATTEN"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/FLATTEN: optical FLow-guided ATTENtion for consistent text-to-video editing](https://mirror.ghproxy.com/https://github.com/yrcong/flatten)."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-RAVE Attention",
            "id": "rave-attn",
            "reference": "https://github.com/logtd/ComfyUI-RAVE_ATTN",
            "files": [
                "https://github.com/logtd/ComfyUI-RAVE_ATTN"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use RAVE attention as a temporal attention mechanism.\nThis differs from other implementations in that it does not concatenate the images together, but within the UNet's Self-Attention mechanism performs the RAVE technique. By not altering the images/latents throughout the UNet, this method does not affect other temporal techniques, style mechanisms, or other UNet modifications.\nFor example, it can be combined with AnimateDiff, ModelScope/ZeroScope, or FLATTEN."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-RefUNet",
            "id": "refunet",
            "reference": "https://github.com/logtd/ComfyUI-RefUNet",
            "files": [
                "https://github.com/logtd/ComfyUI-RefUNet"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes to use Reference UNets"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-SEGAttention",
            "id": "segattention",
            "reference": "https://github.com/logtd/ComfyUI-SEGAttention",
            "files": [
                "https://github.com/logtd/ComfyUI-SEGAttention"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use [a/Smoothed Energy Guidance](https://mirror.ghproxy.com/https://github.com/SusungHong/SEG-SDXL) for ComfyUI."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-SSREncoder",
            "id": "ssrencoder",
            "reference": "https://github.com/logtd/ComfyUI-SSREncoder",
            "files": [
                "https://github.com/logtd/ComfyUI-SSREncoder"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Nodes to use [a/SSR Encoder:Encoding Selective Subject Representation for Subject-Driven Generation](https://mirror.ghproxy.com/https://github.com/Xiaojiu-z/SSR_Encoder)."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-SeeCoder",
            "id": "seecoder-logtd",
            "reference": "https://github.com/logtd/ComfyUI-SeeCoder",
            "files": [
                "https://github.com/logtd/ComfyUI-SeeCoder"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use the SeeCoder from [a/Prompt-Free-Diffusion](https://mirror.ghproxy.com/https://github.com/SHI-Labs/Prompt-Free-Diffusion)"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-4DHumans",
            "id": "comfyui-4dhumans",
            "reference": "https://github.com/logtd/ComfyUI-4DHumans",
            "files": [
                "https://github.com/logtd/ComfyUI-4DHumans"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/4D-Humans](ComfyUI nodes to use 4D-Humans)"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-ReNoise",
            "reference": "https://github.com/logtd/ComfyUI-ReNoise",
            "files": [
                "https://github.com/logtd/ComfyUI-ReNoise"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use ReNoise"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-MotionThiefExperiment",
            "reference": "https://github.com/logtd/ComfyUI-MotionThiefExperiment",
            "files": [
                "https://github.com/logtd/ComfyUI-MotionThiefExperiment"
            ],
            "install_type": "git-clone",
            "description": "experimental node pack to test using reference videos for their motion."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-DiLightNet",
            "reference": "https://github.com/logtd/ComfyUI-DiLightNet",
            "files": [
                "https://github.com/logtd/ComfyUI-DiLightNet"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/DiLightNet](https://mirror.ghproxy.com/https://github.com/iamNCJ/DiLightNet).\nThese nodes can run DiLightNet, but the Dust3r or BlenderPy implementations to create lighting are not included. Expect those to be added to seperate repos when time allows."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-ViewCrafter",
            "reference": "https://github.com/logtd/ComfyUI-ViewCrafter",
            "files": [
                "https://github.com/logtd/ComfyUI-ViewCrafter"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/ViewCrafter](https://mirror.ghproxy.com/https://github.com/Drexubery/ViewCrafter/tree/main) for novel view synthesis."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-APGScaling",
            "reference": "https://github.com/logtd/ComfyUI-APGScaling",
            "files": [
                "https://github.com/logtd/ComfyUI-APGScaling"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use [a/APG scaling](https://huggingface.co/papers/2410.02416) for CFG, allowing for better image quality with higher CFG."
        },
        {
            "author": "logtd",
            "title": "ComfyUI-Fluxtapoz",
            "reference": "https://github.com/logtd/ComfyUI-Fluxtapoz",
            "files": [
                "https://github.com/logtd/ComfyUI-Fluxtapoz"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for image editing with Flux, such as RF-Inversion and more"
        },
        {
            "author": "logtd",
            "title": "ComfyUI-MochiEdit",
            "reference": "https://github.com/logtd/ComfyUI-MochiEdit",
            "files": [
                "https://github.com/logtd/ComfyUI-MochiEdit"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to edit videos using Genmo Mochi"
        },
        {
            "author": "Big-Idea-Technology",
            "title": "ComfyUI-Book-Tools Nodes for ComfyUI",
            "id": "booktool",
            "reference": "https://github.com/Big-Idea-Technology/ComfyUI-Book-Tools",
            "files": [
                "https://github.com/Big-Idea-Technology/ComfyUI-Book-Tools"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Book-Tools is a set o new nodes for ComfyUI that allows users to easily add text overlays to images within their ComfyUI projects. This Node leverages Python Imaging Library (PIL) and PyTorch to dynamically render text on images, supporting a wide range of customization options including font size, alignment, color, and padding. Loop with any parameters (*), prompt batch schedule with prompt selector, end queue for automatic ending current queue."
        },
        {
            "author": "Big Idea Technology",
            "title": "LLM Node for ComfyUI",
            "reference": "https://github.com/Big-Idea-Technology/ComfyUI_LLM_Node",
            "files": [
                "https://github.com/Big-Idea-Technology/ComfyUI_LLM_Node"
            ],
            "install_type": "git-clone",
            "description": "The LLM_Node enhances ComfyUI by integrating advanced language model capabilities, enabling a wide range of NLP tasks such as text generation, content summarization, question answering, and more. This flexibility is powered by various transformer model architectures from the transformers library, allowing for the deployment of models like T5, GPT-2, and others based on your project's needs."
        },
        {
            "author": "Guillaume-Fgt",
            "title": "ComfyUI_StableCascadeLatentRatio",
            "id": "cascade-latent-ratio",
            "reference": "https://github.com/Guillaume-Fgt/ComfyUI_StableCascadeLatentRatio",
            "files": [
                "https://github.com/Guillaume-Fgt/ComfyUI_StableCascadeLatentRatio"
            ],
            "install_type": "git-clone",
            "description": "A custom node to create empty latents for Stable Cascade.\nfeatures: width and height incrementation of 64 by default, possibility to lock the aspect ratio, switch width/height at execution"
        },
        {
            "author": "AuroBit",
            "title": "ComfyUI OOTDiffusion",
            "id": "ootdiffusion",
            "reference": "https://github.com/AuroBit/ComfyUI-OOTDiffusion",
            "files": [
                "https://github.com/AuroBit/ComfyUI-OOTDiffusion"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node that simply integrates the [a/OOTDiffusion](https://mirror.ghproxy.com/https://github.com/levihsu/OOTDiffusion) functionality."
        },
        {
            "author": "AuroBit",
            "title": "ComfyUI-AnimateAnyone-reproduction",
            "id": "animateanyone-reproduction",
            "reference": "https://github.com/AuroBit/ComfyUI-AnimateAnyone-reproduction",
            "files": [
                "https://github.com/AuroBit/ComfyUI-AnimateAnyone-reproduction"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node that simply integrates the [a/animate-anyone-reproduction](https://mirror.ghproxy.com/https://github.com/bendanzzc/AnimateAnyone-reproduction) functionality."
        },
        {
            "author": "czcz1024",
            "title": "Face Compare",
            "id": "facecompare",
            "reference": "https://github.com/czcz1024/Comfyui-FaceCompare",
            "files": [
                "https://github.com/czcz1024/Comfyui-FaceCompare"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FaceCompare"
        },
        {
            "author": "TheBill2001",
            "title": "comfyui-upscale-by-model",
            "reference": "https://github.com/TheBill2001/comfyui-upscale-by-model",
            "files": [
                "https://github.com/TheBill2001/comfyui-upscale-by-model"
            ],
            "install_type": "git-clone",
            "description": "This custom node allow upscaling an image by a factor using a model."
        },
        {
            "author": "TheBill2001",
            "title": "Save Images with Captions",
            "reference": "https://github.com/TheBill2001/ComfyUI-Save-Image-Caption",
            "files": [
                "https://github.com/TheBill2001/ComfyUI-Save-Image-Caption"
            ],
            "install_type": "git-clone",
            "description": "Provide two custom nodes to load and save images with captions as separate files."
        },
        {
            "author": "leoleelxh",
            "title": "ComfyUI-LLMs",
            "reference": "https://github.com/leoleelxh/ComfyUI-LLMs",
            "files": [
                "https://github.com/leoleelxh/ComfyUI-LLMs"
            ],
            "install_type": "git-clone",
            "description": "A minimalist node that calls LLMs, combined with one API, can call all language models, including local models."
        },
        {
            "author": "hughescr",
            "title": "OpenPose Keypoint Extractor",
            "reference": "https://github.com/hughescr/ComfyUI-OpenPose-Keypoint-Extractor",
            "files": [
                "https://github.com/hughescr/ComfyUI-OpenPose-Keypoint-Extractor"
            ],
            "install_type": "git-clone",
            "description": "This is a single node which can take the POSE_KEYPOINT output from the OpenPose extractor node, parse it, and return x,y,width,height bounding boxes around any elements of the OpenPose skeleton"
        },
        {
            "author": "jkrauss82",
            "title": "ULTools for ComfyUI",
            "reference": "https://github.com/jkrauss82/ultools-comfyui",
            "files": [
                "https://github.com/jkrauss82/ultools-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SaveImgAdv, CLIPTextEncodeWithStats. Collection of tools supporting txt2img generation in ComfyUI and other tasks."
        },
        {
            "author": "hiforce",
            "title": "Comfyui HiFORCE Plugin",
            "reference": "https://github.com/hiforce/comfyui-hiforce-plugin",
            "files": [
                "https://github.com/hiforce/comfyui-hiforce-plugin"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes pack provided by [a/HiFORCE](https://www.hiforce.net/) for ComfyUI. This custom node helps to conveniently enhance images through Sampler, Upscaler, Mask, and more.\nNOTE:You should install [a/ComfyUI-Impact-Pack](https://mirror.ghproxy.com/https://github.com/ltdrdata/ComfyUI-Impact-Pack). Many optimizations are built upon the foundation of ComfyUI-Impact-Pack."
        },
        {
            "author": "kuschanow",
            "title": "Advanced Latent Control",
            "reference": "https://github.com/RomanKuschanow/ComfyUI-Advanced-Latent-Control",
            "files": [
                "https://github.com/RomanKuschanow/ComfyUI-Advanced-Latent-Control"
            ],
            "install_type": "git-clone",
            "description": "This custom node helps to transform latent in different ways."
        },
        {
            "author": "guill",
            "title": "abracadabra-comfyui",
            "reference": "https://github.com/guill/abracadabra-comfyui",
            "files": [
                "https://github.com/guill/abracadabra-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Abracadabra Summary, Abracadabra"
        },
        {
            "author": "cerspense",
            "title": "cspnodes",
            "reference": "https://github.com/cerspense/ComfyUI_cspnodes",
            "files": [
                "https://github.com/cerspense/ComfyUI_cspnodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Dir Iterator, Modelscopet2v, Modelscopev2v, Vid Dir Iterator, Image Dir Iterator, Text File Line Iterator, Remap Range, Split Image Channels, Resize By Image, Increment Every N."
        },
        {
            "author": "qwixiwp",
            "title": "queuetools",
            "reference": "https://github.com/qwixiwp/queuetools",
            "files": [
                "https://github.com/qwixiwp/queuetools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:load images (queue tools). tools made for queueing in comfyUI"
        },
        {
            "author": "Chan-0312",
            "title": "ComfyUI-Prompt-Preview",
            "reference": "https://github.com/Chan-0312/ComfyUI-Prompt-Preview",
            "files": [
                "https://github.com/Chan-0312/ComfyUI-Prompt-Preview"
            ],
            "install_type": "git-clone",
            "description": "Welcome to ComfyUI Prompt Preview, where you can visualize the styles from [sdxl_prompt_styler](https://mirror.ghproxy.com/https://github.com/twri/sdxl_prompt_styler)."
        },
        {
            "author": "munkyfoot",
            "title": "ComfyUI-TextOverlay",
            "id": "textoverlay-munkyfoot",
            "reference": "https://github.com/Munkyfoot/ComfyUI-TextOverlay",
            "files": [
                "https://github.com/Munkyfoot/ComfyUI-TextOverlay"
            ],
            "install_type": "git-clone",
            "description": "This extension provides a node that allows you to overlay text on an image or a batch of images with support for custom fonts and styles."
        },
        {
            "author": "holchan",
            "title": "ComfyUI-ModelDownloader",
            "reference": "https://github.com/holchan/ComfyUI-ModelDownloader",
            "files": [
                "https://github.com/holchan/ComfyUI-ModelDownloader"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI node to download models(Checkpoints and LoRA) from external links and act as an output standalone node."
        },
        {
            "author": "CC-BryanOttho",
            "title": "ComfyUI_API_Manager",
            "reference": "https://github.com/CC-BryanOttho/ComfyUI_API_Manager",
            "files": [
                "https://github.com/CC-BryanOttho/ComfyUI_API_Manager"
            ],
            "install_type": "git-clone",
            "description": "This package provides three custom nodes designed to streamline workflows involving API requests, dynamic text manipulation based on API responses, and image posting to APIs. These nodes are particularly useful for automating interactions with APIs, enhancing text-based workflows with dynamic data, and facilitating image uploads."
        },
        {
            "author": "maracman",
            "title": "ComfyUI-SubjectStyle-CSV",
            "reference": "https://github.com/maracman/ComfyUI-SubjectStyle-CSV",
            "files": [
                "https://github.com/maracman/ComfyUI-SubjectStyle-CSV"
            ],
            "install_type": "git-clone",
            "description": "Store a CSV of prompts where the style can change for each subject. The CSV node initialises with the column (style) and row (subject) names for easy interpretability."
        },
        {
            "author": "438443467",
            "title": "ComfyUI-GPT4V-Image-Captioner",
            "reference": "https://github.com/438443467/ComfyUI-GPT4V-Image-Captioner",
            "files": [
                "https://github.com/438443467/ComfyUI-GPT4V-Image-Captioner"
            ],
            "install_type": "git-clone",
            "description": "Nodes:GPT4V-Image-Captioner"
        },
        {
            "author": "uetuluk",
            "title": "comfyui-webcam-node",
            "id": "webcam",
            "reference": "https://github.com/uetuluk/comfyui-webcam-node",
            "files": [
                "https://github.com/uetuluk/comfyui-webcam-node"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Webcam Capture"
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI-layerdiffuse (layerdiffusion)",
            "id": "layerdiffuse",
            "reference": "https://github.com/huchenlei/ComfyUI-layerdiffuse",
            "files": [
                "https://github.com/huchenlei/ComfyUI-layerdiffuse"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation of [a/LayerDiffusion](https://mirror.ghproxy.com/https://github.com/layerdiffusion/LayerDiffusion)."
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI_DanTagGen",
            "id": "dantangen",
            "reference": "https://github.com/huchenlei/ComfyUI_DanTagGen",
            "files": [
                "https://github.com/huchenlei/ComfyUI_DanTagGen"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node of [a/Kohaku's DanTagGen Demo](https://huggingface.co/KBlueLeaf/DanTagGen?not-for-all-audiences=true)."
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI-openpose-editor",
            "reference": "https://github.com/huchenlei/ComfyUI-openpose-editor",
            "files": [
                "https://github.com/huchenlei/ComfyUI-openpose-editor"
            ],
            "install_type": "git-clone",
            "description": "Port of [a/https://mirror.ghproxy.com/https://github.com/huchenlei/sd-webui-openpose-editor](https://mirror.ghproxy.com/https://github.com/huchenlei/sd-webui-openpose-editor) in ComfyUI"
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI-IC-Light-Native",
            "id": "ic-light-native",
            "reference": "https://github.com/huchenlei/ComfyUI-IC-Light-Native",
            "files": [
                "https://github.com/huchenlei/ComfyUI-IC-Light-Native"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI native implementation of [a/IC-Light](https://mirror.ghproxy.com/https://github.com/lllyasviel/IC-Light)."
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI DenseDiffusion",
            "id": "densediffusion",
            "reference": "https://github.com/huchenlei/ComfyUI_densediffusion",
            "files": [
                "https://github.com/huchenlei/ComfyUI_densediffusion"
            ],
            "install_type": "git-clone",
            "description": "[a/DenseDiffusion](https://mirror.ghproxy.com/https://github.com/naver-ai/DenseDiffusion) custom node for ComfyUI."
        },
        {
            "author": "huchenlei",
            "title": "ComfyUI_omost",
            "id": "omost",
            "reference": "https://github.com/huchenlei/ComfyUI_omost",
            "files": [
                "https://github.com/huchenlei/ComfyUI_omost"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation of [a/Omost](https://mirror.ghproxy.com/https://github.com/lllyasviel/Omost), and everything about regional prompt.\nNOTE: You need to install ComfyUI_densediffusion to use this node."
        },
        {
            "author": "nathannlu",
            "title": "ComfyUI Pets",
            "id": "pets",
            "reference": "https://github.com/nathannlu/ComfyUI-Pets",
            "files": [
                "https://github.com/nathannlu/ComfyUI-Pets"
            ],
            "install_type": "git-clone",
            "description": "Play with your pet while your workflow generates!"
        },
        {
            "author": "nathannlu",
            "title": "Comfy Cloud",
            "id": "cloud",
            "reference": "https://github.com/nathannlu/ComfyUI-Cloud",
            "files": [
                "https://github.com/nathannlu/ComfyUI-Cloud"
            ],
            "install_type": "git-clone",
            "description": "Run your workflow using cloud GPU resources, from your local ComfyUI.\nNOTE:After you first install the plugin...\nThe first time you click 'generate', you will be prompted to log into your account.Subsequent generations after the first is faster (the first run it takes a while to process your workflow). Once those two steps have been completed, you will be able to seamlessly generate your workflow on the cloud!"
        },
        {
            "author": "11dogzi",
            "title": "Comfyui-ergouzi-Nodes",
            "id": "ergouzi-nodes",
            "reference": "https://github.com/11dogzi/Comfyui-ergouzi-Nodes",
            "files": [
                "https://github.com/11dogzi/Comfyui-ergouzi-Nodes"
            ],
            "install_type": "git-clone",
            "description": "This is a node group kit that covers multiple nodes such as local refinement, tag management, random prompt words, text processing, image processing, mask processing, etc"
        },
        {
            "author": "11dogzi",
            "title": "Comfyui-ergouzi-samplers",
            "id": "ergouzi-samplers",
            "reference": "https://github.com/11dogzi/Comfyui-ergouzi-samplers",
            "files": [
                "https://github.com/11dogzi/Comfyui-ergouzi-samplers"
            ],
            "install_type": "git-clone",
            "description": "Partial redraw sampler and variant seed sampler"
        },
        {
            "author": "11dogzi",
            "title": "Comfyui-ergouzi-kaiguan",
            "id": "ergouzi-kaiguan",
            "reference": "https://github.com/11dogzi/Comfyui-ergouzi-kaiguan",
            "files": [
                "https://github.com/11dogzi/Comfyui-ergouzi-kaiguan"
            ],
            "install_type": "git-clone",
            "description": "Group switching control, one click control to ignore and disable multiple groups, as well as wired switch combination nodes, allowing for arbitrary switching of annotation names"
        },
        {
            "author": "11dogzi",
            "title": "ComfUI-EGAdapterMadAssistant",
            "id": "madassistant",
            "reference": "https://github.com/11dogzi/ComfUI-EGAdapterMadAssistant",
            "files": [
                "https://github.com/11dogzi/ComfUI-EGAdapterMadAssistant"
            ],
            "install_type": "git-clone",
            "description": "This is a hierarchical auxiliary project of the IPAdapter project, which uses a slider to quickly control the hierarchical weights and add fully random and semi random modes"
        },
        {
            "author": "BXYMartin",
            "title": "ComfyUI-InstantIDUtils",
            "id": "instantid-utils",
            "reference": "https://github.com/BXYMartin/ComfyUI-InstantIDUtils",
            "files": [
                "https://github.com/BXYMartin/ComfyUI-InstantIDUtils"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Multi-ControlNet Converter, List of Images, Convert PIL to Tensor (NHWC), Convert Tensor (NHWC) to (NCHW), Convert Tensor (NHWC) to PIL"
        },
        {
            "author": "cdb-boop",
            "title": "comfyui-image-round",
            "id": "image-round",
            "reference": "https://github.com/cdb-boop/comfyui-image-round",
            "files": [
                "https://github.com/cdb-boop/comfyui-image-round"
            ],
            "install_type": "git-clone",
            "description": "A simple node to round an input image up (pad) or down (crop) to the nearest integer multiple. Padding offset from left/bottom and the padding value are adjustable."
        },
        {
            "author": "cdb-boop",
            "title": "ComfyUI Bringing Old Photos Back to Life",
            "reference": "https://github.com/cdb-boop/ComfyUI-Bringing-Old-Photos-Back-to-Life",
            "files": [
                "https://github.com/cdb-boop/ComfyUI-Bringing-Old-Photos-Back-to-Life"
            ],
            "install_type": "git-clone",
            "description": "Enhance old or low-quality images in ComfyUI. Optional features include automatic scratch removal and face enhancement. Based on Microsoft's Bringing-Old-Photos-Back-to-Life. Requires installing models, so see instructions here: https://mirror.ghproxy.com/https://github.com/cdb-boop/ComfyUI-Bringing-Old-Photos-Back-to-Life."
        },
        {
            "author": "atmaranto",
            "title": "SaveAsScript",
            "id": "saveasscript",
            "reference": "https://github.com/atmaranto/ComfyUI-SaveAsScript",
            "files": [
                "https://github.com/atmaranto/ComfyUI-SaveAsScript"
            ],
            "install_type": "git-clone",
            "description": "A version of ComfyUI-to-Python-Extension that works as a custom node. Adds a button in the UI that saves the current workflow as a Python file, a CLI for converting workflows, and slightly better custom node support."
        },
        {
            "author": "meshmesh-io",
            "title": "mm-comfyui-megamask",
            "id": "megamask",
            "reference": "https://github.com/meshmesh-io/mm-comfyui-megamask",
            "files": [
                "https://github.com/meshmesh-io/mm-comfyui-megamask"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ColorListMaskToImage, FlattenAndCombineMaskImages"
        },
        {
            "author": "meshmesh-io",
            "title": "mm-comfyui-loopback",
            "id": "mm-loopback",
            "reference": "https://github.com/meshmesh-io/mm-comfyui-loopback",
            "files": [
                "https://github.com/meshmesh-io/mm-comfyui-loopback"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Loop, LoopStart, LoopEnd, LoopStart_SEGIMAGE, LoopEnd_SEGIMAGE"
        },
        {
            "author": "meshmesh-io",
            "title": "ComfyUI-MeshMesh",
            "id": "meshmesh",
            "reference": "https://github.com/meshmesh-io/ComfyUI-MeshMesh",
            "files": [
                "https://github.com/meshmesh-io/ComfyUI-MeshMesh"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Masks to Colored Masks, Color Picker"
        },
        {
            "author": "CozyMantis",
            "title": "Cozy Human Parser",
            "id": "humanparser",
            "reference": "https://github.com/cozymantis/human-parser-comfyui-node",
            "files": [
                "https://github.com/cozymantis/human-parser-comfyui-node"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI node to automatically extract masks for body regions and clothing/fashion items. Made with 💚 by the CozyMantis squad."
        },
        {
            "author": "CozyMantis",
            "title": "Cozy Reference Pose Generator",
            "id": "posegen",
            "reference": "https://github.com/cozymantis/pose-generator-comfyui-node",
            "files": [
                "https://github.com/cozymantis/pose-generator-comfyui-node"
            ],
            "install_type": "git-clone",
            "description": "Generate OpenPose face/body reference poses in ComfyUI with ease. Made with 💚 by the CozyMantis squad."
        },
        {
            "author": "CozyMantis",
            "title": "Cozy Utils",
            "id": "cozy-utils",
            "reference": "https://github.com/cozymantis/cozy-utils-comfyui-nodes",
            "files": [
                "https://github.com/cozymantis/cozy-utils-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Various cozy nodes, made with 💚 by the CozyMantis squad."
        },
        {
            "author": "vivax3794",
            "title": "ComfyUI-Vivax-Nodes",
            "reference": "https://github.com/vivax3794/ComfyUI-Vivax-Nodes",
            "files": [
                "https://github.com/vivax3794/ComfyUI-Vivax-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Inspect, Any String, Model From URL"
        },
        {
            "author": "vivax3794",
            "title": "ComfyUI-Sub-Nodes",
            "reference": "https://github.com/vivax3794/ComfyUI-Sub-Nodes",
            "files": [
                "https://github.com/vivax3794/ComfyUI-Sub-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Creating subgraph and Calling subgraphs."
        },
        {
            "author": "victorchall",
            "title": "Comfyui Webcam capture node",
            "reference": "https://github.com/victorchall/comfyui_webcamcapture",
            "files": [
                "https://github.com/victorchall/comfyui_webcamcapture"
            ],
            "install_type": "git-clone",
            "description": "This node captures images one at a time from your webcam when you click generate.\nThis is particular useful for img2img or controlnet workflows.\nNOTE:This node will take over your webcam, so if you have another program using it, you may need to close that program first. Likewise, you may need to close Comfyui or close the workflow to release the webcam."
        },
        {
            "author": "ljleb",
            "title": "Mecha Merge Node Pack",
            "id": "mecha",
            "reference": "https://github.com/ljleb/comfy-mecha",
            "files": [
                "https://github.com/ljleb/comfy-mecha"
            ],
            "install_type": "git-clone",
            "description": "model merging nodes powered by sd-mecha, a memory efficient state dict recipe merger."
        },
        {
            "author": "diSty",
            "title": "ComfyUI Frame Maker",
            "id": "frame-maker",
            "reference": "https://github.com/diStyApps/ComfyUI_FrameMaker",
            "files": [
                "https://github.com/diStyApps/ComfyUI_FrameMaker"
            ],
            "install_type": "git-clone",
            "description": "This node creates a sequence of frames by moving and scaling a subject image over a background image."
        },
        {
            "author": "diSty",
            "title": "Flow - Streamlined Way to ComfyUI",
            "reference": "https://github.com/diStyApps/ComfyUI-disty-Flow",
            "files": [
                "https://github.com/diStyApps/ComfyUI-disty-Flow"
            ],
            "install_type": "git-clone",
            "description": "Flow is a custom node designed to provide a more user-friendly interface for ComfyUI by acting as an alternative user interface for running workflows. It is not a replacement for workflow creation.\nFlow is currently in the early stages of development, so expect bugs and ongoing feature enhancements. With your support and feedback, Flow will settle into a steady stream."
        },
        {
            "author": "hackkhai",
            "title": "ComfyUI-Image-Matting",
            "id": "image-matting",
            "reference": "https://github.com/hackkhai/ComfyUI-Image-Matting",
            "files": [
                "https://github.com/hackkhai/ComfyUI-Image-Matting"
            ],
            "install_type": "git-clone",
            "description": "This node improves the quality of the image mask. more suitable for image composite matting"
        },
        {
            "author": "Pos13",
            "title": "Cyclist",
            "id": "cyclist",
            "reference": "https://github.com/Pos13/comfyui-cyclist",
            "files": [
                "https://github.com/Pos13/comfyui-cyclist"
            ],
            "install_type": "git-clone",
            "description": "This extension provides tools to iterate generation results between runs. In general, it's for cycles."
        },
        {
            "author": "ExponentialML",
            "title": "ComfyUI_ModelScopeT2V",
            "id": "modelscopet2v",
            "reference": "https://github.com/ExponentialML/ComfyUI_ModelScopeT2V",
            "files": [
                "https://github.com/ExponentialML/ComfyUI_ModelScopeT2V"
            ],
            "install_type": "git-clone",
            "description": "Allows native usage of ModelScope based Text To Video Models in ComfyUI"
        },
        {
            "author": "ExponentialML",
            "title": "ComfyUI - Native DynamiCrafter",
            "id": "dynamicrafter",
            "reference": "https://github.com/ExponentialML/ComfyUI_Native_DynamiCrafter",
            "files": [
                "https://github.com/ExponentialML/ComfyUI_Native_DynamiCrafter"
            ],
            "install_type": "git-clone",
            "description": "DynamiCrafter that works natively with ComfyUI's nodes, optimizations, ControlNet, and more."
        },
        {
            "author": "ExponentialML",
            "title": "ComfyUI_VisualStylePrompting",
            "id": "visual-style-prompting",
            "reference": "https://github.com/ExponentialML/ComfyUI_VisualStylePrompting",
            "files": [
                "https://github.com/ExponentialML/ComfyUI_VisualStylePrompting"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Version of '[a/Visual Style Prompting with Swapping Self-Attention](https://mirror.ghproxy.com/https://github.com/naver-ai/Visual-Style-Prompting)'"
        },
        {
            "author": "angeloshredder",
            "title": "StableCascadeResizer",
            "reference": "https://github.com/angeloshredder/StableCascadeResizer",
            "files": [
                "https://github.com/angeloshredder/StableCascadeResizer"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Cascade_Resizer"
        },
        {
            "author": "stavsap",
            "title": "ComfyUI Ollama",
            "id": "ollama",
            "reference": "https://github.com/stavsap/comfyui-ollama",
            "files": [
                "https://github.com/stavsap/comfyui-ollama"
            ],
            "install_type": "git-clone",
            "description": "Custom ComfyUI Nodes for interacting with [a/Ollama](https://ollama.com/) using the [a/ollama python client](https://mirror.ghproxy.com/https://github.com/ollama/ollama-python).\nIntegrate the power of LLMs into CompfyUI workflows easily."
        },
        {
            "author": "dchatel",
            "title": "comfyui_davcha",
            "reference": "https://github.com/dchatel/comfyui_davcha",
            "files": [
                "https://github.com/dchatel/comfyui_davcha"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SmartMask, ResizeCropFit, Percent Padding, SoftErosion, StringScheduleHelper, DStack, DavchaConditioningConcat, DavchaModelMergeSimple, DavchaCLIPMergeSimple, DavchaModelMergeSD1, DavchaModelMergeSDXL, ConditioningCompress... Some personal QoL and experimental nodes"
        },
        {
            "author": "dchatel",
            "title": "comfyui_facetools",
            "id": "facetools",
            "reference": "https://github.com/dchatel/comfyui_facetools",
            "files": [
                "https://github.com/dchatel/comfyui_facetools"
            ],
            "install_type": "git-clone",
            "description": "These custom nodes provide a rotation aware face extraction, paste back, and various face related masking options."
        },
        {
            "author": "prodogape",
            "title": "Comfyui-Minio",
            "id": "minio",
            "reference": "https://github.com/prodogape/ComfyUI-Minio",
            "files": [
                "https://github.com/prodogape/ComfyUI-Minio"
            ],
            "install_type": "git-clone",
            "description": "This plugin is mainly based on Minio, implementing the ability to read images from Minio, save images, facilitating expansion and connection across multiple machines."
        },
        {
            "author": "prodogape",
            "title": "ComfyUI-EasyOCR",
            "id": "easyocr",
            "reference": "https://github.com/prodogape/ComfyUI-EasyOCR",
            "files": [
                "https://github.com/prodogape/ComfyUI-EasyOCR"
            ],
            "install_type": "git-clone",
            "description": "This node is primarily based on Easy-OCR to implement OCR text recognition functionality."
        },
        {
            "author": "prodogape",
            "title": "ComfyUI-OmDet",
            "id": "omdet",
            "reference": "https://github.com/prodogape/ComfyUI-OmDet",
            "files": [
                "https://github.com/prodogape/ComfyUI-OmDet"
            ],
            "install_type": "git-clone",
            "description": "This node is mainly based on [a/OmDet](https://mirror.ghproxy.com/https://github.com/om-ai-lab/OmDet) for object detection, and it outputs related images, masks, and Labelme JSON information."
        },
        {
            "author": "prodogape",
            "title": "Comfyui-Yolov8-JSON",
            "reference": "https://github.com/prodogape/Comfyui-Yolov8-JSON",
            "files": [
                "https://github.com/prodogape/Comfyui-Yolov8-JSON"
            ],
            "install_type": "git-clone",
            "description": "This node is mainly based on the Yolov8 model for object detection, and it outputs related images, masks, and JSON information.[w/Repository url is changed. Please remove previous one and reinstall.]"
        },
        {
            "author": "kingzcheung",
            "title": "ComfyUI_kkTranslator_nodes",
            "id": "kktranslator",
            "reference": "https://github.com/AIGCTeam/ComfyUI_kkTranslator_nodes",
            "files": [
                "https://github.com/AIGCTeam/ComfyUI_kkTranslator_nodes"
            ],
            "install_type": "git-clone",
            "description": "These nodes are mainly used to translate prompt words from other languages into English. PromptTranslateToText implements prompt word translation based on Helsinki NLP translation model.It doesn't require internet connection。"
        },
        {
            "author": "vsevolod-oparin",
            "title": "Kandinsky 2.2 ComfyUI Plugin",
            "id": "kandinsky",
            "reference": "https://github.com/vsevolod-oparin/comfyui-kandinsky22",
            "files": [
                "https://github.com/vsevolod-oparin/comfyui-kandinsky22"
            ],
            "install_type": "git-clone",
            "description": "Nodes provide an options to combine prior and decoder models of Kandinsky 2.2."
        },
        {
            "author": "Xyem",
            "title": "Xycuno Oobabooga",
            "id": "xycuno-oobabooga",
            "reference": "https://github.com/Xyem/Xycuno-Oobabooga",
            "files": [
                "https://github.com/Xyem/Xycuno-Oobabooga"
            ],
            "install_type": "git-clone",
            "description": "Xycuno Oobabooga provides custom nodes for ComfyUI, for sending requests to an [a/Oobabooga](https://mirror.ghproxy.com/https://github.com/oobabooga/text-generation-webui) instance to assist in creating prompt texts."
        },
        {
            "author": "shi3z",
            "title": "ComfyUI_Memeplex_DALLE",
            "id": "memeplex-dalle",
            "reference": "https://github.com/shi3z/ComfyUI_Memeplex_DALLE",
            "files": [
                "https://github.com/shi3z/ComfyUI_Memeplex_DALLE"
            ],
            "install_type": "git-clone",
            "description": "You can use memeplex and DALL-E thru ComfyUI. You need API keys."
        },
        {
            "author": "if-ai",
            "title": "ComfyUI-IF_AI_tools",
            "id": "if-ai-tools",
            "reference": "https://github.com/if-ai/ComfyUI-IF_AI_tools",
            "files": [
                "https://github.com/if-ai/ComfyUI-IF_AI_tools"
            ],
            "install_type": "git-clone",
            "description": "Various AI tools to use in Comfy UI. Starting with VL and prompt making tools using Ollma as backend will evolve as I find time."
        },
        {
            "author": "if-ai",
            "title": "ComfyUI-IF_AI_WishperSpeechNode",
            "id": "if-ai-whisper-speech",
            "reference": "https://github.com/if-ai/ComfyUI-IF_AI_WishperSpeechNode",
            "files": [
                "https://github.com/if-ai/ComfyUI-IF_AI_WishperSpeechNode"
            ],
            "install_type": "git-clone",
            "description": "This repository hosts a Text-to-Speech (TTS) application that leverages Whisper Speech for voice synthesis, allowing users to train a voice model on-the-fly. It is built on ComfyUI and supports rapid training and inference processes."
        },
        {
            "author": "if-ai",
            "title": "ComfyUI-IF_AI_HFDownloaderNode",
            "id": "if-ai-hfdownloader",
            "reference": "https://github.com/if-ai/ComfyUI-IF_AI_HFDownloaderNode",
            "files": [
                "https://github.com/if-ai/ComfyUI-IF_AI_HFDownloaderNode"
            ],
            "install_type": "git-clone",
            "description": "Talking avatars Heads for the IF_AI tools integrates dreamtalk in ComfyUI"
        },
        {
            "author": "dmMaze",
            "title": "Sketch2Manga",
            "id": "sketch2manga",
            "reference": "https://github.com/dmMaze/sketch2manga",
            "files": [
                "https://github.com/dmMaze/sketch2manga"
            ],
            "install_type": "git-clone",
            "description": "Apply screentone to line drawings or colored illustrations with diffusion models."
        },
        {
            "author": "olduvai-jp",
            "title": "ComfyUI-HfLoader",
            "id": "hfloader",
            "reference": "https://github.com/olduvai-jp/ComfyUI-HfLoader",
            "files": [
                "https://github.com/olduvai-jp/ComfyUI-HfLoader"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Lora Loader From HF"
        },
        {
            "author": "AiMiDi",
            "title": "ComfyUI-Aimidi-nodes",
            "id": "aimidi-nodes",
            "reference": "https://github.com/AiMiDi/ComfyUI-Aimidi-nodes",
            "files": [
                "https://github.com/AiMiDi/ComfyUI-Aimidi-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Merge Tag, Clear Tag, Add Tag, Load Images Pair Batch, Save Images Pair"
        },
        {
            "author": "ForeignGods",
            "title": "ComfyUI-Mana-Nodes",
            "id": "mana-nodes",
            "reference": "https://github.com/ForeignGods/ComfyUI-Mana-Nodes",
            "files": [
                "https://github.com/ForeignGods/ComfyUI-Mana-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Font Animation, Speech Recognition, Caption Generator, TTS"
        },
        {
            "author": "Cornea Valentin",
            "title": "ControlNet Auxiliar",
            "id": "controlnet-aux-valentin",
            "reference": "https://github.com/madtunebk/ComfyUI-ControlnetAux",
            "files": [
                "https://github.com/madtunebk/ComfyUI-ControlnetAux"
            ],
            "install_type": "git-clone",
            "description": "This ComfyUI custom node, named ControlNet Auxiliar, is designed to provide auxiliary functionalities for image processing tasks. It is particularly useful for various image manipulation and enhancement operations. The node is integrated with functionalities for converting images between different formats and applying various image processing techniques."
        },
        {
            "author": "MarkoCa1",
            "title": "ComfyUI-Text",
            "reference": "https://github.com/MarkoCa1/ComfyUI-Text",
            "files": [
                "https://github.com/MarkoCa1/ComfyUI-Text"
            ],
            "install_type": "git-clone",
            "description": "Text."
        },
        {
            "author": "MarkoCa1",
            "title": "ComfyUI_Segment_Mask",
            "id": "seg-mask",
            "reference": "https://github.com/MarkoCa1/ComfyUI_Segment_Mask",
            "files": [
                "https://github.com/MarkoCa1/ComfyUI_Segment_Mask"
            ],
            "install_type": "git-clone",
            "description": "Mask cutout based on Segment Anything."
        },
        {
            "author": "Shadetail",
            "title": "Eagleshadow Custom Nodes",
            "id": "eagleshadow",
            "reference": "https://github.com/Shadetail/ComfyUI_Eagleshadow",
            "files": [
                "https://github.com/Shadetail/ComfyUI_Eagleshadow"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for ComfyUI by Eagleshadow."
        },
        {
            "author": "Jannchie",
            "title": "ComfyUI-J",
            "reference": "https://github.com/Jannchie/ComfyUI-J",
            "files": [
                "https://github.com/Jannchie/ComfyUI-J"
            ],
            "install_type": "git-clone",
            "description": "This is a completely different set of nodes than Comfy's own KSampler series. This set of nodes is based on Diffusers, which makes it easier to import models, apply prompts with weights, inpaint, reference only, controlnet, etc."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-JDCN",
            "id": "jdcn",
            "reference": "https://github.com/daxcay/ComfyUI-JDCN",
            "files": [
                "https://github.com/daxcay/ComfyUI-JDCN"
            ],
            "install_type": "git-clone",
            "description": "Jerry Davos Custom Nodes for Saving Latents in Directory (BatchLatentSave) , Importing Latent from directory (BatchLatentLoadFromDir) , List to string, string to list, get any file list from directory which give filepath, filename, move any files from any directory to any other directory, VHS Video combine file mover, rebatch list of strings, batch image load from any dir, load image batch from any directory and other custom nodes."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-DataSet",
            "reference": "https://github.com/daxcay/ComfyUI-DataSet",
            "files": [
                "https://github.com/daxcay/ComfyUI-DataSet"
            ],
            "install_type": "git-clone",
            "description": "Data research, preparation, and manipulation nodes for model trainers and artists."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-Nexus",
            "reference": "https://github.com/daxcay/ComfyUI-Nexus",
            "files": [
                "https://github.com/daxcay/ComfyUI-Nexus"
            ],
            "install_type": "git-clone",
            "description": "Node to enable seamless multiuser workflow collaboration, run on local and remote comfy servers."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-WA",
            "reference": "https://github.com/daxcay/ComfyUI-WA",
            "files": [
                "https://github.com/daxcay/ComfyUI-WA"
            ],
            "install_type": "git-clone",
            "description": "Node to enable WhatsApp in ComfyUI."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-TG",
            "reference": "https://github.com/daxcay/ComfyUI-TG",
            "files": [
                "https://github.com/daxcay/ComfyUI-TG"
            ],
            "install_type": "git-clone",
            "description": "Node to enable Telegram in ComfyUI."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-NODEJS",
            "reference": "https://github.com/daxcay/ComfyUI-NODEJS",
            "files": [
                "https://github.com/daxcay/ComfyUI-NODEJS"
            ],
            "install_type": "git-clone",
            "description": "This node allows the execution of Node.js application within ComfyUI by leveraging the ComfyUI-NODEJS, which starts alongside ComfyUI and facilitates the installation of Node.js. The integration enables Python subprocesses to execute Node.js scripts."
        },
        {
            "author": "daxcay",
            "title": "ComfyUI-YouTubeVideoPlayer",
            "reference": "https://github.com/daxcay/ComfyUI-YouTubeVideoPlayer",
            "files": [
                "https://github.com/daxcay/ComfyUI-YouTubeVideoPlayer"
            ],
            "install_type": "git-clone",
            "description": "Plays youtube videos in comfy. Use this node to share tutorials or renders. Youtube Playlists mode is also in Future Development in which you can add multiple youtube links and form a playlist which would be ideal for chained tutorials or lisitening and sharing songs playlists with others."
        },
        {
            "author": "Seedsa",
            "title": "ComfyUI Fooocus Nodes",
            "id": "fooocus-nodes",
            "reference": "https://github.com/Seedsa/Fooocus_Nodes",
            "files": [
                "https://github.com/Seedsa/Fooocus_Nodes"
            ],
            "install_type": "git-clone",
            "description": "This extension provides image generation features based on Fooocus."
        },
        {
            "author": "zhangp365",
            "title": "zhangp365/ComfyUI-utils-nodes",
            "reference": "https://github.com/zhangp365/ComfyUI-utils-nodes",
            "files": [
                "https://github.com/zhangp365/ComfyUI-utils-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LoadImageWithSwitch, ImageBatchOneOrMore, ModifyTextGender, ImageCompositeMaskedWithSwitch, ColorCorrectOfUtils, SplitMask, MaskFastGrow, CheckpointLoaderSimpleWithSwitch, ImageResizeTo8x, MatchImageRatioToPreset etc."
        },
        {
            "author": "zhangp365",
            "title": "ComfyUI_photomakerV2_native",
            "id": "comfyui_photomakerV2_native",
            "reference": "https://github.com/zhangp365/ComfyUI_photomakerV2_native",
            "files": [
                "https://github.com/zhangp365/ComfyUI_photomakerV2_native"
            ],
            "install_type": "git-clone",
            "description": "Nodes: PhotoMakerLoaderV2,PhotoMakerEncodeV2"
        },
        {
            "author": "ratulrafsan",
            "title": "Comfyui-SAL-VTON",
            "id": "sal-vton",
            "reference": "https://github.com/ratulrafsan/Comfyui-SAL-VTON",
            "files": [
                "https://github.com/ratulrafsan/Comfyui-SAL-VTON"
            ],
            "install_type": "git-clone",
            "description": "Dressup your models!\nThis is my quick implementation of the SAL-VTON node for ComfyUI.\nBased on the paper [a/Keyu Y. Tingwei G. et al. (2023). Linking Garment with Person via Semantically Associated Landmakrs for Virtual Try-On](https://openaccess.thecvf.com/content/CVPR2023/papers/Yan_Linking_Garment_With_Person_via_Semantically_Associated_Landmarks_for_Virtual_CVPR_2023_paper.pdf)"
        },
        {
            "author": "Nevysha",
            "title": "ComfyUI-nevysha-top-menu",
            "id": "nevysha-top-menu",
            "reference": "https://github.com/Nevysha/ComfyUI-nevysha-top-menu",
            "files": [
                "https://github.com/Nevysha/ComfyUI-nevysha-top-menu"
            ],
            "install_type": "git-clone",
            "description": "A simple sidebar tweak to force fixe the ComfyUI menu to the top of the screen. Reaaally quick and dirty. May break with some ComfyUI setup."
        },
        {
            "author": "alisson-anjos",
            "title": "ComfyUI-Ollama-Describer",
            "id": "ollama-describer",
            "reference": "https://github.com/alisson-anjos/ComfyUI-Ollama-Describer",
            "files": [
                "https://github.com/alisson-anjos/ComfyUI-Ollama-Describer"
            ],
            "install_type": "git-clone",
            "description": "This is an extension for ComfyUI that makes it possible to use some LLM models provided by Ollama, such as Gemma, Llava (multimodal), Llama2, Llama3 or Mistral. Speaking specifically of the LLaVa - Large Language and Vision Assistant model, although trained on a relatively small dataset, it demonstrates exceptional capabilities in understanding images and answering questions about them. This model presents similar behaviors to multimodal models such as GPT-4, even when presented with invisible images and instructions."
        },
        {
            "author": "chaosaiart",
            "title": "Chaosaiart-Nodes",
            "id": "chaosaiart",
            "reference": "https://github.com/chaosaiart/Chaosaiart-Nodes",
            "files": [
                "https://github.com/chaosaiart/Chaosaiart-Nodes"
            ],
            "install_type": "git-clone",
            "description": "LowVRAM Animation : txt2video - img2video - video2video , Frame by Frame, compatible with LowVRAM GPUs\nIncluded : Prompt Switch, Checkpoint Switch, Cache, Number Count by Frame, Ksampler txt2img & img2img ..."
        },
        {
            "author": "viperyl",
            "title": "ComfyUI-BiRefNet",
            "id": "comfyui-birefnet",
            "reference": "https://github.com/viperyl/ComfyUI-BiRefNet",
            "files": [
                "https://github.com/viperyl/ComfyUI-BiRefNet"
            ],
            "install_type": "git-clone",
            "description": "Bilateral Reference Network achieves SOTA result in multi Salient Object Segmentation dataset, this repo pack BiRefNet as ComfyUI nodes, and make this SOTA model easier use for everyone."
        },
        {
            "author": "viperyl",
            "title": "ComfyUI-RGT",
            "id": "rgt",
            "reference": "https://github.com/viperyl/ComfyUI-RGT",
            "pip": ["loguru"],
            "files": [
                "https://github.com/viperyl/ComfyUI-RGT"
            ],
            "install_type": "git-clone",
            "description": "This repo cast Recursive Generalization Transformer for Image Super-Resolution to ComfyUI, the original [a/paper link](https://arxiv.org/abs/2303.06373) and [a/github link](https://mirror.ghproxy.com/https://github.com/zhengchen1999/RGT)"
        },
        {
            "author": "SuperBeastsAI",
            "title": "ComfyUI-SuperBeasts",
            "id": "superbeasts",
            "reference": "https://github.com/SuperBeastsAI/ComfyUI-SuperBeasts",
            "files": [
                "https://github.com/SuperBeastsAI/ComfyUI-SuperBeasts"
            ],
            "install_type": "git-clone",
            "description": "Nodes:HDR Effects (SuperBeasts.AI). This repository contains custom nodes for ComfyUI created and used by SuperBeasts.AI (@SuperBeasts.AI on Instagram)"
        },
        {
            "author": "hay86",
            "title": "ComfyUI Dreamtalk",
            "id": "dreamtalk",
            "reference": "https://github.com/hay86/ComfyUI_Dreamtalk",
            "files": [
                "https://github.com/hay86/ComfyUI_Dreamtalk"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/dreamtalk](https://mirror.ghproxy.com/https://github.com/ali-vilab/dreamtalk) for ComfyUI"
        },
        {
            "author": "hay86",
            "title": "ComfyUI Hallo",
            "id": "hallo-hay86",
            "reference": "https://github.com/hay86/ComfyUI_Hallo",
            "files": [
                "https://github.com/hay86/ComfyUI_Hallo"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/hallo](https://mirror.ghproxy.com/https://github.com/fudan-generative-vision/hallo) for ComfyUI"
        },
        {
            "author": "hay86",
            "title": "ComfyUI OpenVoice",
            "id": "openvoice-hay86",
            "reference": "https://github.com/hay86/ComfyUI_OpenVoice",
            "files": [
                "https://github.com/hay86/ComfyUI_OpenVoice"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/OpenVoice](https://mirror.ghproxy.com/https://github.com/myshell-ai/OpenVoice) for ComfyUI"
        },
        {
            "author": "hay86",
            "title": "ComfyUI DDColor",
            "id": "ddcolor-hay86",
            "reference": "https://github.com/hay86/ComfyUI_DDColor",
            "files": [
                "https://github.com/hay86/ComfyUI_DDColor"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/DDColor](https://mirror.ghproxy.com/https://github.com/piddnad/DDColor) for ComfyUI"
        },
        {
            "author": "hay86",
            "title": "ComfyUI MiniCPM-V",
            "id": "minicpm-v",
            "reference": "https://github.com/hay86/ComfyUI_MiniCPM-V",
            "files": [
                "https://github.com/hay86/ComfyUI_MiniCPM-V"
            ],
            "install_type": "git-clone",
            "description": "Unofficial implementation of [a/MiniCPM-V](https://mirror.ghproxy.com/https://github.com/OpenBMB/MiniCPM-V) for ComfyUI"
        },
        {
            "author": "hay86",
            "title": "ComfyUI AceNodes",
            "reference": "https://github.com/hay86/ComfyUI_AceNodes",
            "files": [
                "https://github.com/hay86/ComfyUI_AceNodes"
            ],
            "install_type": "git-clone",
            "description": "Some useful custom nodes that are not included in ComfyUI core yet."
        },
        {
            "author": "shinich39",
            "title": "comfyui-load-image-in-seq",
            "reference": "https://github.com/shinich39/comfyui-load-image-in-seq",
            "files": [
                "https://github.com/shinich39/comfyui-load-image-in-seq"
            ],
            "install_type": "git-clone",
            "description": "This node is load png image sequentially with metadata. Only supported for PNG format that has been created by ComfyUI.[w/renamed from comfyui-load-image-39. You need to remove previous one and reinstall to this.]"
        },
        {
            "author": "shinich39",
            "title": "comfyui-random-node",
            "id": "random-node",
            "reference": "https://github.com/shinich39/comfyui-ramdom-node",
            "files": [
                "https://github.com/shinich39/comfyui-ramdom-node"
            ],
            "install_type": "git-clone",
            "description": "Shuffle nodes after queue added. [w/Repository name has been changed from comfyui-random-node-39 to comfyui-random-node. Please remove and reinstall it.]"
        },
        {
            "author": "shinich39",
            "title": "comfyui-local-db",
            "reference": "https://github.com/shinich39/comfyui-local-db",
            "files": [
                "https://github.com/shinich39/comfyui-local-db"
            ],
            "install_type": "git-clone",
            "description": "Store text to Key-Values pair json."
        },
        {
            "author": "shinich39",
            "title": "comfyui-model-db",
            "reference": "https://github.com/shinich39/comfyui-model-db",
            "files": [
                "https://github.com/shinich39/comfyui-model-db"
            ],
            "install_type": "git-clone",
            "description": "Store settings by model."
        },
        {
            "author": "shinich39",
            "title": "comfyui-load-image-with-cmd",
            "reference": "https://github.com/shinich39/comfyui-load-image-with-cmd",
            "files": [
                "https://github.com/shinich39/comfyui-load-image-with-cmd"
            ],
            "install_type": "git-clone",
            "description": "Load image and partially workflow with javascript."
        },
        {
            "author": "shinich39",
            "title": "connect-from-afar",
            "reference": "https://github.com/shinich39/comfyui-connect-from-afar",
            "files": [
                "https://github.com/shinich39/comfyui-connect-from-afar"
            ],
            "install_type": "git-clone",
            "description": "Connect a new link from out of screen."
        },
        {
            "author": "shinich39",
            "title": "comfyui-target-search",
            "reference": "https://github.com/shinich39/comfyui-target-search",
            "files": [
                "https://github.com/shinich39/comfyui-target-search"
            ],
            "install_type": "git-clone",
            "description": "Move canvas to target on dragging connection."
        },
        {
            "author": "shinich39",
            "title": "comfyui-group-selection",
            "reference": "https://github.com/shinich39/comfyui-group-selection",
            "files": [
                "https://github.com/shinich39/comfyui-group-selection"
            ],
            "install_type": "git-clone",
            "description": "Create a new group of nodes."
        },
        {
            "author": "shinich39",
            "title": "comfyui-textarea-keybindings",
            "reference": "https://github.com/shinich39/comfyui-textarea-keybindings",
            "files": [
                "https://github.com/shinich39/comfyui-textarea-keybindings"
            ],
            "install_type": "git-clone",
            "description": "Add keybindings to textarea."
        },
        {
            "author": "shinich39",
            "title": "comfyui-put-image",
            "reference": "https://github.com/shinich39/comfyui-put-image",
            "files": [
                "https://github.com/shinich39/comfyui-put-image"
            ],
            "install_type": "git-clone",
            "description": "Load image from directory."
        },
        {
            "author": "shinich39",
            "title": "comfyui-parse-image",
            "reference": "https://github.com/shinich39/comfyui-parse-image",
            "files": [
                "https://github.com/shinich39/comfyui-parse-image"
            ],
            "install_type": "git-clone",
            "description": "Extract metadata from image."
        },
        {
            "author": "wei30172",
            "title": "comfygen",
            "reference": "https://github.com/wei30172/comfygen",
            "files": [
                "https://github.com/wei30172/comfygen"
            ],
            "install_type": "git-clone",
            "description": "Setting Up a Web Interface Using ComfyUI.\nNOTE:When installed, you can access it via http://127.0.0.1:8188/comfygen."
        },
        {
            "author": "zombieyang",
            "title": "SD-PPP",
            "reference": "https://github.com/zombieyang/sd-ppp",
            "files": [
                "https://github.com/zombieyang/sd-ppp"
            ],
            "install_type": "git-clone",
            "description": "getting/sending picture from/to Photoshop with a simple connection. Make Photoshop become the workspace of your ComfyUI"
        },
        {
            "author": "KytraScript",
            "title": "ComfyUI_KytraWebhookHTTP",
            "reference": "https://github.com/KytraScript/ComfyUI_KytraWebhookHTTP",
            "files": [
                "https://github.com/KytraScript/ComfyUI_KytraWebhookHTTP"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI node that utilizes Moviepy to convert and send your images or videos to a webhook endpoint directly from ComfyUI."
        },
        {
            "author": "1mckw",
            "title": "Comfyui-Gelbooru",
            "reference": "https://github.com/1mckw/Comfyui-Gelbooru",
            "files": [
                "https://github.com/1mckw/Comfyui-Gelbooru"
            ],
            "install_type": "git-clone",
            "description": "Get random images from gelbooru, support multiple tag searches, exclude tags, etc. user and api key are optional."
        },
        {
            "author": "NeuralSamurAI",
            "title": "SuperPrompter Node for ComfyUI",
            "reference": "https://github.com/NeuralSamurAI/Comfyui-Superprompt-Unofficial",
            "files": [
                "https://github.com/NeuralSamurAI/Comfyui-Superprompt-Unofficial"
            ],
            "install_type": "git-clone",
            "description": "The SuperPrompter node is a ComfyUI node that uses the SuperPrompt-v1 model from Hugging Face to generate text based on a given prompt. It provides various parameters to control the text generation process."
        },
        {
            "author": "NeuralSamurAI",
            "title": "Dimensional Latent Perlin for ComfyUI",
            "reference": "https://github.com/NeuralSamurAI/ComfyUI-Dimensional-Latent-Perlin",
            "files": [
                "https://github.com/NeuralSamurAI/ComfyUI-Dimensional-Latent-Perlin"
            ],
            "install_type": "git-clone",
            "description": "Dimensional Latent Perlin is a custom node for ComfyUI that generates Perlin noise in the latent space. This node is designed to work seamlessly with various diffusion models and can be used as an alternative or complement to standard random noise generators in image generation pipelines."
        },
        {
            "author": "NeuralSamurAI",
            "title": "PromptJSON Node for ComfyUI",
            "reference": "https://github.com/NeuralSamurAI/ComfyUI-PromptJSON",
            "files": [
                "https://github.com/NeuralSamurAI/ComfyUI-PromptJSON"
            ],
            "install_type": "git-clone",
            "description": "PromptJSON is a custom node for ComfyUI that structures natural language prompts and generates prompts for external LLM nodes in image generation workflows. It aids in creating consistent, schema-based image descriptions."
        },
        {
            "author": "NeuralSamurAI",
            "title": "FluxPseudoNegative",
            "reference": "https://github.com/NeuralSamurAI/ComfyUI-FluxPseudoNegativePrompt",
            "files": [
                "https://github.com/NeuralSamurAI/ComfyUI-FluxPseudoNegativePrompt"
            ],
            "install_type": "git-clone",
            "description": "FluxPseudoNegative is an advanced custom node for ComfyUI that converts negative prompts into positive ones. It's designed to enhance prompt engineering for image generation models that don't natively support negative prompts or where using negative prompts significantly increases generation time. So instead of hacking CFG we simply invert your negative words and find their antonyms!"
        },
        {
            "author": "MokkaBoss1",
            "title": "Node Pack mostly for manipulating strings and integers",
            "reference": "https://github.com/MokkaBoss1/ComfyUI_Mokkaboss1/wiki/Documentation-for-the-ComfyUI-Nodes-in-this-Node-Pack",
            "files": [
                "https://github.com/MokkaBoss1/ComfyUI_Mokkaboss1"
            ],
            "install_type": "git-clone",
            "description": "Node Pack mostly for manipulating strings and integers"
        },
        {
            "author": "jiaxiangc",
            "title": "ResAdapter for ComfyUI",
            "reference": "https://github.com/jiaxiangc/ComfyUI-ResAdapter",
            "files": [
                "https://github.com/jiaxiangc/ComfyUI-ResAdapter"
            ],
            "install_type": "git-clone",
            "description": "We provide ComfyUI-ResAdapter node to help users to use [a/ResAdapter](https://mirror.ghproxy.com/https://github.com/bytedance/res-adapter) in ComfyUI."
        },
        {
            "author": "ParisNeo",
            "title": "lollms_nodes_suite",
            "reference": "https://github.com/ParisNeo/lollms_nodes_suite",
            "files": [
                "https://github.com/ParisNeo/lollms_nodes_suite"
            ],
            "install_type": "git-clone",
            "description": "lollms_nodes_suite is a set of nodes for comfyui that harnesses the power of lollms, a state-of-the-art AI text generation tool, to improve the quality of image generation."
        },
        {
            "author": "IsItDanOrAi",
            "title": "ComfyUI-Stereopsis",
            "reference": "https://github.com/IsItDanOrAi/ComfyUI-Stereopsis",
            "files": [
                "https://github.com/IsItDanOrAi/ComfyUI-Stereopsis"
            ],
            "install_type": "git-clone",
            "description": "This initiative represents a solo venture dedicated to integrating a stereopsis effect within ComfyUI (Stable Diffusion). Presently, the project is focused on the refinement of node categorization within a unified framework, as it is in the early stages of development. However, it has achieved functionality in a fundamental capacity. By processing a video through the Side-by-Side (SBS) node and applying Frame Delay to one of the inputs, it facilitates the creation of a stereopsis effect. This effect is compatible with any Virtual Reality headset that supports SBS video playback, offering a practical application in immersive media experiences."
        },
        {
            "author": "nickve28",
            "title": "ComfyUI Nich Utils",
            "reference": "https://github.com/nickve28/ComfyUI-Nich-Utils",
            "files": [
                "https://github.com/nickve28/ComfyUI-Nich-Utils"
            ],
            "install_type": "git-clone",
            "description": "Several utility nodes for use with ComfyUI."
        },
        {
            "author": "FrankChieng",
            "title": "ComfyUI_Aniportrait",
            "reference": "https://github.com/frankchieng/ComfyUI_Aniportrait",
            "files": [
                "https://github.com/frankchieng/ComfyUI_Aniportrait"
            ],
            "install_type": "git-clone",
            "description": "implementation of [a/AniPortrait](https://mirror.ghproxy.com/https://github.com/Zejun-Yang/AniPortrait) generating of videos, includes self driven, face reenacment and audio driven with a reference image"
        },
        {
            "author": "FrankChieng",
            "title": "ComfyUI_MagicClothing",
            "reference": "https://github.com/frankchieng/ComfyUI_MagicClothing",
            "files": [
                "https://github.com/frankchieng/ComfyUI_MagicClothing"
            ],
            "install_type": "git-clone",
            "description": "implementation of MagicClothing with garment and prompt in ComfyUI"
        },
        {
            "author": "BlakeOne",
            "title": "ComfyUI SchedulerMixer",
            "reference": "https://github.com/BlakeOne/ComfyUI-SchedulerMixer",
            "files": [
                "https://github.com/BlakeOne/ComfyUI-SchedulerMixer"
            ],
            "install_type": "git-clone",
            "description": "Create a custom scheduler from a weighted average of the built-in schedulers"
        },
        {
            "author": "BlakeOne",
            "title": "ComfyUI CustomScheduler",
            "reference": "https://github.com/BlakeOne/ComfyUI-CustomScheduler",
            "files": [
                "https://github.com/BlakeOne/ComfyUI-CustomScheduler"
            ],
            "install_type": "git-clone",
            "description": "Simple node for setting the sigma values directly. Note, for a full denoise the last sigma should be zero."
        },
        {
            "author": "BlakeOne",
            "title": "ComfyUI NodePresets",
            "id": "nodepresets",
            "reference": "https://github.com/BlakeOne/ComfyUI-NodePresets",
            "files": [
                "https://github.com/BlakeOne/ComfyUI-NodePresets"
            ],
            "install_type": "git-clone",
            "description": "An extension for ComyUI that enables saving and loading node presets using the node's context menu.\nRight click a node and choose 'Presets' from its context menu to access the node's presets."
        },
        {
            "author": "BlakeOne",
            "title": "ComfyUI NodeReset",
            "id": "nodereset",
            "reference": "https://github.com/BlakeOne/ComfyUI-NodeReset",
            "files": [
                "https://github.com/BlakeOne/ComfyUI-NodeReset"
            ],
            "install_type": "git-clone",
            "description": "An extension for ComyUI to allow resetting a node's inputs to their default values.\nNOTE:Right click any node and choose 'Reset' from the context menu."
        },
        {
            "author": "kale4eat",
            "title": "ComfyUI_demucus",
            "id": "demucus",
            "reference": "https://github.com/kale4eat/ComfyUI-path-util",
            "files": [
                "https://github.com/kale4eat/ComfyUI-path-util"
            ],
            "install_type": "git-clone",
            "description": "Path utility for ComfyUI"
        },
        {
            "author": "kale4eat",
            "title": "ComfyUI-string-util",
            "reference": "https://github.com/kale4eat/ComfyUI-string-util",
            "files": [
                "https://github.com/kale4eat/ComfyUI-string-util"
            ],
            "install_type": "git-clone",
            "description": "String utility for ComfyUI"
        },
        {
            "author": "kale4eat",
            "title": "ComfyUI-text-file-util",
            "reference": "https://github.com/kale4eat/ComfyUI-text-file-util",
            "files": [
                "https://github.com/kale4eat/ComfyUI-text-file-util"
            ],
            "install_type": "git-clone",
            "description": "Text file utility for ComfyUI"
        },
        {
            "author": "kale4eat",
            "title": "ComfyUI-speech-dataset-toolkit",
            "reference": "https://github.com/kale4eat/ComfyUI-speech-dataset-toolkit",
            "files": [
                "https://github.com/kale4eat/ComfyUI-speech-dataset-toolkit"
            ],
            "install_type": "git-clone",
            "description": "Basic audio tools using torchaudio for ComfyUI. It is assumed to assist in the speech dataset creation for ASR, TTS, etc."
        },
        {
            "author": "DrMWeigand",
            "title": "ComfyUI Color Detection Nodes",
            "reference": "https://github.com/DrMWeigand/ComfyUI_ColorImageDetection",
            "files": [
                "https://github.com/DrMWeigand/ComfyUI_ColorImageDetection"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes for detecting color in images, leveraging RGB and LAB color spaces. These nodes aim to distinguish colored images from black and white, including those with color tints."
        },
        {
            "author": "DrMWeigand",
            "title": "StereoVision Plugin for ComfyUI",
            "reference": "https://github.com/DrMWeigand/ComfyUI-StereoVision",
            "files": [
                "https://github.com/DrMWeigand/ComfyUI-StereoVision"
            ],
            "install_type": "git-clone",
            "description": "The StereoVision plugin for ComfyUI enables the creation of stereoscopic and autostereoscopic images and videos using depth maps. It supports both traditional stereoscopic image generation and autostereogram (Magic Eye) creation."
        },
        {
            "author": "bobmagicii",
            "title": "ComfyKit Custom Nodes",
            "reference": "https://github.com/bobmagicii/comfykit-custom-nodes",
            "files": [
                "https://github.com/bobmagicii/comfykit-custom-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:LoraWithMetadata, TypecasterImage."
        },
        {
            "author": "TJ16th",
            "title": "comfyUI_TJ_NormalLighting",
            "reference": "https://github.com/TJ16th/comfyUI_TJ_NormalLighting",
            "files": [
                "https://github.com/TJ16th/comfyUI_TJ_NormalLighting"
            ],
            "install_type": "git-clone",
            "description": "Custom Node for comfyUI for virtual lighting based on normal map.\nYou can use normal maps to add virtual lighting effects to your images."
        },
        {
            "author": "A4P7J1N7M05OT",
            "title": "ComfyUI-PixelOE-Wrapper",
            "reference": "https://github.com/A4P7J1N7M05OT/ComfyUI-PixelOE-Wrapper",
            "files": [
                "https://github.com/A4P7J1N7M05OT/ComfyUI-PixelOE-Wrapper"
            ],
            "install_type": "git-clone",
            "description": "A barebones ComfyUI wrapper for [a/PixelOE](https://mirror.ghproxy.com/https://github.com/KohakuBlueleaf/PixelOE).\nI cannot promise any support, if there is someone who wants to make a proper node, please do."
        },
        {
            "author": "A4P7J1N7M05OT",
            "title": "ComfyUI-AutoColorGimp",
            "reference": "https://github.com/A4P7J1N7M05OT/ComfyUI-AutoColorGimp",
            "files": [
                "https://github.com/A4P7J1N7M05OT/ComfyUI-AutoColorGimp"
            ],
            "install_type": "git-clone",
            "description": "Shamelessly copied the code to auto color correct the image like in gimp from this answer: [a/https://stackoverflow.com/a/56365560/4561887](https://stackoverflow.com/a/56365560/4561887)"
        },
        {
            "author": "ronniebasak",
            "title": "ComfyUI-Tara-LLM-Integration",
            "id": "tarallm",
            "reference": "https://github.com/ronniebasak/ComfyUI-Tara-LLM-Integration",
            "files": [
                "https://github.com/ronniebasak/ComfyUI-Tara-LLM-Integration"
            ],
            "install_type": "git-clone",
            "description": "Tara is a powerful node for ComfyUI that integrates Large Language Models (LLMs) to enhance and automate workflow processes. With Tara, you can create complex, intelligent workflows that refine and generate content, manage API keys, and seamlessly integrate various LLMs into your projects."
        },
        {
            "author": "Sida Liu",
            "title": "ComfyUI-Debug",
            "id": "debug",
            "reference": "https://github.com/liusida/ComfyUI-Debug",
            "files": [
                "https://github.com/liusida/ComfyUI-Debug"
            ],
            "install_type": "git-clone",
            "description": "Attach a debug node to an output to obtain more detailed information. Uncover the details of your models in ComfyUI with ease."
        },
        {
            "author": "Sida Liu",
            "title": "ComfyUI-Login",
            "id": "login",
            "reference": "https://github.com/liusida/ComfyUI-Login",
            "files": [
                "https://github.com/liusida/ComfyUI-Login"
            ],
            "install_type": "git-clone",
            "description": "A simple password to protect ComfyUI."
        },
        {
            "author": "Sida Liu",
            "title": "ComfyUI-AutoCropFaces",
            "id": "autocropfaces",
            "reference": "https://github.com/liusida/ComfyUI-AutoCropFaces",
            "files": [
                "https://github.com/liusida/ComfyUI-AutoCropFaces"
            ],
            "install_type": "git-clone",
            "description": "Use RetinaFace to detect and automatically crop faces."
        },
        {
            "author": "Sida Liu",
            "title": "ComfyUI-SD3-nodes",
            "id": "sd3-nodes",
            "reference": "https://github.com/liusida/ComfyUI-SD3-nodes",
            "files": [
                "https://github.com/liusida/ComfyUI-SD3-nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes that support Stable Diffusion 3 Medium better."
        },
        {
            "author": "Sida Liu",
            "title": "ComfyUI-B-LoRA",
            "id": "b-lora",
            "reference": "https://github.com/liusida/ComfyUI-B-LoRA",
            "files": [
                "https://github.com/liusida/ComfyUI-B-LoRA"
            ],
            "install_type": "git-clone",
            "description": "Load and apply B-LoRA models, currently B-LoRA models only works with SDXL (sdxl_base_1.0)."
        },
        {
            "author": "jtydhr88",
            "title": "ComfyUI-Workflow-Encrypt",
            "id": "workflow-encrypt",
            "reference": "https://github.com/jtydhr88/ComfyUI-Workflow-Encrypt",
            "files": [
                "https://github.com/jtydhr88/ComfyUI-Workflow-Encrypt"
            ],
            "install_type": "git-clone",
            "description": "Encrypt your comfyui workflow, and share it with key"
        },
        {
            "author": "jtydhr88",
            "title": "ComfyUI LayerDivider",
            "id": "layer-divider",
            "reference": "https://github.com/jtydhr88/ComfyUI-LayerDivider",
            "files": [
                "https://github.com/jtydhr88/ComfyUI-LayerDivider"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI LayerDivider is custom nodes that generating layered psd files inside ComfyUI[w/Please follow readme and run install_windows_portable_win_py311_cu121 for ComfyUI embedded python.]"
        },
        {
            "author": "SeaArtLab",
            "title": "ComfyUI-Long-CLIP",
            "reference": "https://github.com/SeaArtLab/ComfyUI-Long-CLIP",
            "files": [
                "https://github.com/SeaArtLab/ComfyUI-Long-CLIP"
            ],
            "install_type": "git-clone",
            "description": "This project implements the comfyui for long-clip, currently supporting the replacement of clip-l. For SD1.5, the SeaArtLongClip module can be used to replace the original clip in the model, expanding the token length from 77 to 248."
        },
        {
            "author": "tsogzark",
            "title": "ComfyUI-load-image-from-url",
            "reference": "https://github.com/tsogzark/ComfyUI-load-image-from-url",
            "files": [
                "https://github.com/tsogzark/ComfyUI-load-image-from-url"
            ],
            "install_type": "git-clone",
            "description": "A simple node to load image from local path or http url.\nYou can find this node from 'image' category."
        },
        {
            "author": "discus0434",
            "title": "ComfyUI Caching Embeddings",
            "id": "caching-embeddings",
            "reference": "https://github.com/discus0434/comfyui-caching-embeddings",
            "files": [
                "https://github.com/discus0434/comfyui-caching-embeddings"
            ],
            "install_type": "git-clone",
            "description": "This repository simply caches the CLIP embeddings and subtly accelerates the inference process by bypassing unnecessary computations."
        },
        {
            "author": "discus0434",
            "title": "ComfyUI Aesthetic Predictor V2.5",
            "id": "aesthetic-predictor",
            "reference": "https://github.com/discus0434/comfyui-aesthetic-predictor-v2-5",
            "files": [
                "https://github.com/discus0434/comfyui-aesthetic-predictor-v2-5"
            ],
            "install_type": "git-clone",
            "description": "Simple ComfyUI node that predicts the score of an aesthetic image with SigLIP-based predictor."
        },
        {
            "author": "discus0434",
            "title": "ComfyUI Flux Accelerator",
            "reference": "https://github.com/discus0434/comfyui-flux-accelerator",
            "files": [
                "https://github.com/discus0434/comfyui-flux-accelerator"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Flux Accelerator is a custom node for ComfyUI that accelerates Flux.1 image generation, just by using this node."
        },
        {
            "author": "AIFSH",
            "title": "StyleShot-ComfyUI",
            "id": "styleshot",
            "reference": "https://github.com/AIFSH/StyleShot-ComfyUI",
            "files": [
                "https://github.com/AIFSH/StyleShot-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/StyleShot](https://mirror.ghproxy.com/https://github.com/open-mmlab/StyleShot.git)"
        },
        {
            "author": "AIFSH",
            "title": "VocalSeparation-ComfyUI",
            "id": "vocalseparation",
            "reference": "https://github.com/AIFSH/VocalSeparation-ComfyUI",
            "files": [
                "https://github.com/AIFSH/VocalSeparation-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for separation vocals from music based on [a/ZFTurbo/Music-Source-Separation-Training](https://mirror.ghproxy.com/https://github.com/ZFTurbo/Music-Source-Separation-Training)"
        },
        {
            "author": "AIFSH",
            "title": "DiffMorpher-ComfyUI",
            "id": "diffmorpher",
            "reference": "https://github.com/AIFSH/DiffMorpher-ComfyUI",
            "files": [
                "https://github.com/AIFSH/DiffMorpher-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/DiffMorpher](https://mirror.ghproxy.com/https://github.com/Kevin-thu/DiffMorpher),you can find base workflow in [a/doc](https://mirror.ghproxy.com/https://github.com/AIFSH/DiffMorpher-ComfyUI/blob/main/doc)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-UVR5",
            "id": "uvr5",
            "reference": "https://github.com/AIFSH/ComfyUI-UVR5",
            "files": [
                "https://github.com/AIFSH/ComfyUI-UVR5"
            ],
            "install_type": "git-clone",
            "description": "the custom code for [a/UVR5](https://mirror.ghproxy.com/https://github.com/Anjok07/ultimatevocalremovergui) to separate vocals and background music"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-IP_LAP",
            "id": "iplap",
            "reference": "https://github.com/AIFSH/ComfyUI-IP_LAP",
            "files": [
                "https://github.com/AIFSH/ComfyUI-IP_LAP"
            ],
            "install_type": "git-clone",
            "description": "Nodes:IP_LAP Node, Video Loader, PreView Video, Combine Audio Video. the comfyui custom node of [a/IP_LAP](https://mirror.ghproxy.com/https://github.com/Weizhi-Zhong/IP_LAP) to make audio driven videos!"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-GPT_SoVITS",
            "id": "sovits",
            "reference": "https://github.com/AIFSH/ComfyUI-GPT_SoVITS",
            "files": [
                "https://github.com/AIFSH/ComfyUI-GPT_SoVITS"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/GPT-SoVITS](https://mirror.ghproxy.com/https://github.com/RVC-Boss/GPT-SoVITS)! you can voice cloning and tts in comfyui now\n[w/NOTE:make sure ffmpeg is worked in your commandline]"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-MuseTalk_FSH",
            "id": "musetalk-fsh",
            "reference": "https://github.com/AIFSH/ComfyUI-MuseTalk_FSH",
            "files": [
                "https://github.com/AIFSH/ComfyUI-MuseTalk_FSH"
            ],
            "install_type": "git-clone",
            "description": "the comfyui custom node of [a/MuseTalk](https://mirror.ghproxy.com/https://github.com/TMElyralab/MuseTalk) to make audio driven videos!"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-WhisperX",
            "id": "whisperx",
            "reference": "https://github.com/AIFSH/ComfyUI-WhisperX",
            "files": [
                "https://github.com/AIFSH/ComfyUI-WhisperX"
            ],
            "install_type": "git-clone",
            "description": "a comfyui cuatom node for audio subtitling based on [a/whisperX](https://mirror.ghproxy.com/https://github.com/m-bain/whisperX.git) and [a/translators](https://mirror.ghproxy.com/https://github.com/UlionTse/translators)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-RVC",
            "id": "aifsh-rvc",
            "reference": "https://github.com/AIFSH/ComfyUI-RVC",
            "files": [
                "https://github.com/AIFSH/ComfyUI-RVC"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/Retrieval-based-Voice-Conversion-WebUI](https://mirror.ghproxy.com/https://github.com/RVC-Project/Retrieval-based-Voice-Conversion-WebUI.git), you can Voice-Conversion in comfyui now!\nNOTE: make sure ffmpeg is worked in your commandline for Linux"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-XTTS",
            "id": "xtts",
            "reference": "https://github.com/AIFSH/ComfyUI-XTTS",
            "files": [
                "https://github.com/AIFSH/ComfyUI-XTTS"
            ],
            "install_type": "git-clone",
            "description": "a custom comfyui node for [a/coqui-ai/TTS](https://mirror.ghproxy.com/https://github.com/coqui-ai/TTS.git)'s xtts module! support 17 languages voice cloning and tts"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-Live2DViewer",
            "id": "live2dviewer",
            "reference": "https://github.com/AIFSH/ComfyUI-Live2DViewer",
            "files": [
                "https://github.com/AIFSH/ComfyUI-Live2DViewer"
            ],
            "install_type": "git-clone",
            "description": "a comfyui node for viewing Live2D model"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-FishSpeech",
            "id": "fishspeech",
            "reference": "https://github.com/AIFSH/ComfyUI-FishSpeech",
            "files": [
                "https://github.com/AIFSH/ComfyUI-FishSpeech"
            ],
            "install_type": "git-clone",
            "description": "a custom comfyui node for [a/fish-speech](https://mirror.ghproxy.com/https://github.com/fishaudio/fish-speech.git)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI_V-Express",
            "id": "v-express-aifsh",
            "reference": "https://github.com/AIFSH/ComfyUI_V-Express",
            "files": [
                "https://github.com/AIFSH/ComfyUI_V-Express"
            ],
            "install_type": "git-clone",
            "description": "the comfyui custom node of [a/V-Express](https://mirror.ghproxy.com/https://github.com/tencent-ailab/V-Express) to make audio driven videos!"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-MimicBrush",
            "id": "mimicbrush",
            "reference": "https://github.com/AIFSH/ComfyUI-MimicBrush",
            "files": [
                "https://github.com/AIFSH/ComfyUI-MimicBrush"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/MimicBrush](https://mirror.ghproxy.com/https://github.com/ali-vilab/MimicBrush),then inpainting with reference image."
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-Hallo",
            "id": "hallo",
            "reference": "https://github.com/AIFSH/ComfyUI-Hallo",
            "files": [
                "https://github.com/AIFSH/ComfyUI-Hallo"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/hallo](https://mirror.ghproxy.com/https://github.com/fudan-generative-vision/hallo)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-UniAnimate",
            "id": "unianimate",
            "reference": "https://github.com/AIFSH/ComfyUI-UniAnimate",
            "files": [
                "https://github.com/AIFSH/ComfyUI-UniAnimate"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/UniAnimate](https://mirror.ghproxy.com/https://github.com/ali-vilab/UniAnimate)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-3d-photo-inpainting",
            "id": "3d-photo-inpainting",
            "reference": "https://github.com/AIFSH/ComfyUI-3d-photo-inpainting",
            "files": [
                "https://github.com/AIFSH/ComfyUI-3d-photo-inpainting"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/3d-photo-inpainting](https://mirror.ghproxy.com/https://github.com/vt-vl-lab/3d-photo-inpainting),then you can render one image to zoom-in/dolly zoom/swing motion/circle motion video"
        },
        {
            "author": "AIFSH",
            "title": "AIFSH/ComfyUI-AuraSR",
            "id": "aurasr-aifsh",
            "reference": "https://github.com/AIFSH/ComfyUI-AuraSR",
            "files": [
                "https://github.com/AIFSH/ComfyUI-AuraSR"
            ],
            "install_type": "git-clone",
            "description": "a node for [a/AuraSR](https://mirror.ghproxy.com/https://github.com/fal-ai/aura-sr)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-MARS5-TTS",
            "id": "mars5-tts",
            "reference": "https://github.com/AIFSH/ComfyUI-MARS5-TTS",
            "files": [
                "https://github.com/AIFSH/ComfyUI-MARS5-TTS"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/MARS5-TTS](https://mirror.ghproxy.com/https://github.com/Camb-ai/MARS5-TTS)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-I2V-Adapter",
            "id": "i2v-adapter",
            "reference": "https://github.com/AIFSH/ComfyUI-I2V-Adapter",
            "files": [
                "https://github.com/AIFSH/ComfyUI-I2V-Adapter"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/I2V-Adapter](https://mirror.ghproxy.com/https://github.com/KwaiVGI/I2V-Adapter)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-MimicMotion",
            "id": "mimicmotion-aifsh",
            "reference": "https://github.com/AIFSH/ComfyUI-MimicMotion",
            "files": [
                "https://github.com/AIFSH/ComfyUI-MimicMotion"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/MimicMotion](https://mirror.ghproxy.com/https://github.com/Tencent/MimicMotion)"
        },
        {
            "author": "AIFSH",
            "title": "ComfyUI-DiffSynth-Studio",
            "id": "diffsynth-studio",
            "reference": "https://github.com/AIFSH/ComfyUI-DiffSynth-Studio",
            "files": [
                "https://github.com/AIFSH/ComfyUI-DiffSynth-Studio"
            ],
            "install_type": "git-clone",
            "description": "make [a/DiffSynth-Studio](https://mirror.ghproxy.com/https://github.com/modelscope/DiffSynth-Studio) available in ComfyUI"
        },
        {
            "author": "AIFSH",
            "title": "CosyVoice-ComfyUI",
            "id": "cosyvoice",
            "reference": "https://github.com/AIFSH/CosyVoice-ComfyUI",
            "files": [
                "https://github.com/AIFSH/CosyVoice-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/CosyVoice](https://mirror.ghproxy.com/https://github.com/FunAudioLLM/CosyVoice)"
        },
        {
            "author": "AIFSH",
            "title": "AniTalker-ComfyUI",
            "id": "anitalker",
            "reference": "https://github.com/AIFSH/AniTalker-ComfyUI",
            "files": [
                "https://github.com/AIFSH/AniTalker-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/AniTalker](https://mirror.ghproxy.com/https://github.com/X-LANCE/AniTalker)"
        },
        {
            "author": "AIFSH",
            "title": "DHLive-ComfyUI",
            "id": "dhlive",
            "reference": "https://github.com/AIFSH/DHLive-ComfyUI",
            "files": [
                "https://github.com/AIFSH/DHLive-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/DH_live](https://mirror.ghproxy.com/https://github.com/kleinlee/DH_live)"
        },
        {
            "author": "AIFSH",
            "title": "GSTTS-ComfyUI",
            "id": "gstts",
            "reference": "https://github.com/AIFSH/GSTTS-ComfyUI",
            "files": [
                "https://github.com/AIFSH/GSTTS-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node for [a/GPT-SoVITS](https://mirror.ghproxy.com/https://github.com/RVC-Boss/GPT-SoVITS)"
        },
        {
            "author": "AIFSH",
            "title": "FancyVideo-ComfyUI",
            "reference": "https://github.com/AIFSH/FancyVideo-ComfyUI",
            "files": [
                "https://github.com/AIFSH/FancyVideo-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/FancyVideo](https://mirror.ghproxy.com/https://github.com/360CVGroup/FancyVideo)"
        },
        {
            "author": "AIFSH",
            "title": "VideoSys-ComfyUI",
            "reference": "https://github.com/AIFSH/VideoSys-ComfyUI",
            "files": [
                "https://github.com/AIFSH/VideoSys-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "NODES:TextNode, PreViewVideo, VideoSysNode."
        },
        {
            "author": "AIFSH",
            "title": "HivisionIDPhotos-ComfyUI",
            "reference": "https://github.com/AIFSH/HivisionIDPhotos-ComfyUI",
            "files": [
                "https://github.com/AIFSH/HivisionIDPhotos-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/HivisionIDPhotos](https://mirror.ghproxy.com/https://github.com/Zeyi-Lin/HivisionIDPhotos)"
        },
        {
            "author": "AIFSH",
            "title": "DiffSynth-ComfyUI",
            "reference": "https://github.com/AIFSH/DiffSynth-ComfyUI",
            "files": [
                "https://github.com/AIFSH/DiffSynth-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/DiffSynth-Studio](https://mirror.ghproxy.com/https://github.com/modelscope/DiffSynth-Studio)"
        },
        {
            "author": "AIFSH",
            "title": "RealisDance-ComfyUI",
            "reference": "https://github.com/AIFSH/RealisDance-ComfyUI",
            "files": [
                "https://github.com/AIFSH/RealisDance-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/RealisDance](https://mirror.ghproxy.com/https://github.com/damo-cv/RealisDance)"
        },
        {
            "author": "AIFSH",
            "title": "ViewCrafter-ComfyUI",
            "reference": "https://github.com/AIFSH/ViewCrafter-ComfyUI",
            "files": [
                "https://github.com/AIFSH/ViewCrafter-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/ViewCrafter](https://mirror.ghproxy.com/https://github.com/Drexubery/ViewCrafter)"
        },
        {
            "author": "AIFSH",
            "title": "SenseVoice-ComfyUI",
            "reference": "https://github.com/AIFSH/SenseVoice-ComfyUI",
            "files": [
                "https://github.com/AIFSH/SenseVoice-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for SenseVoice"
        },
        {
            "author": "AIFSH",
            "title": "EzAudio-ComfyUI",
            "reference": "https://github.com/AIFSH/EzAudio-ComfyUI",
            "files": [
                "https://github.com/AIFSH/EzAudio-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/EzAudio](https://mirror.ghproxy.com/https://github.com/haidog-yaqub/EzAudio)"
        },
        {
            "author": "AIFSH",
            "title": "PyramidFlow-ComfyUI",
            "reference": "https://github.com/AIFSH/PyramidFlow-ComfyUI",
            "files": [
                "https://github.com/AIFSH/PyramidFlow-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/Pyramid-Flow](https://mirror.ghproxy.com/https://github.com/jy0205/Pyramid-Flow)"
        },
        {
            "author": "AIFSH",
            "title": "JoyHallo-ComfyUI",
            "reference": "https://github.com/AIFSH/JoyHallo-ComfyUI",
            "files": [
                "https://github.com/AIFSH/JoyHallo-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/JoyHallo](https://mirror.ghproxy.com/https://github.com/jdh-algo/JoyHallo)"
        },
        {
            "author": "AIFSH",
            "title": "F5-TTS-ComfyUI",
            "reference": "https://github.com/AIFSH/F5-TTS-ComfyUI",
            "files": [
                "https://github.com/AIFSH/F5-TTS-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/F5-TTS](https://mirror.ghproxy.com/https://github.com/SWivid/F5-TTS)"
        },
        {
            "author": "AIFSH",
            "title": "FireRedTTS-ComfyUI",
            "reference": "https://github.com/AIFSH/FireRedTTS-ComfyUI",
            "files": [
                "https://github.com/AIFSH/FireRedTTS-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/FireRedTTS](https://mirror.ghproxy.com/https://github.com/FireRedTeam/FireRedTTS)"
        },
        {
            "author": "AIFSH",
            "title": "IMAGDressing-ComfyUI",
            "reference": "https://github.com/AIFSH/IMAGDressing-ComfyUI",
            "files": [
                "https://github.com/AIFSH/IMAGDressing-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom nodde for [a/IMAGDressing](https://mirror.ghproxy.com/https://github.com/muzishen/IMAGDressing)"
        },
        {
            "author": "AIFSH",
            "title": "OmniGen-ComfyUI",
            "reference": "https://github.com/AIFSH/OmniGen-ComfyUI",
            "files": [
                "https://github.com/AIFSH/OmniGen-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a custom node for [a/OmniGen](https://mirror.ghproxy.com/https://github.com/VectorSpaceLab/OmniGen)"
        },
        {
            "author": "Koishi-Star",
            "title": "Euler-Smea-Dyn-Sampler",
            "id": "smea",
            "reference": "https://github.com/Koishi-Star/Euler-Smea-Dyn-Sampler",
            "files": [
                "https://github.com/Koishi-Star/Euler-Smea-Dyn-Sampler"
            ],
            "install_type": "git-clone",
            "description": "СomfyUI version of [a/Euler Smea Dyn Sampler](https://mirror.ghproxy.com/https://github.com/Koishi-Star/Euler-Smea-Dyn-Sampler). It adds samplers directly to KSampler nodes."
        },
        {
            "author": "sdfxai",
            "title": "SDFXBridgeForComfyUI - ComfyUI Custom Node for SDFX Integration",
            "id": "sdfx",
            "reference": "https://github.com/sdfxai/SDFXBridgeForComfyUI",
            "files": [
                "https://github.com/sdfxai/SDFXBridgeForComfyUI"
            ],
            "install_type": "git-clone",
            "description": "SDFXBridgeForComfyUI is a custom node designed for seamless integration between ComfyUI and SDFX. This custom node allows users to make ComfyUI compatible with SDFX when running the ComfyUI instance on their local machines."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_FoleyCrafter",
            "id": "comfyui_foleycrafter",
            "reference": "https://github.com/smthemex/ComfyUI_FoleyCrafter",
            "files": [
                "https://github.com/smthemex/ComfyUI_FoleyCrafter"
            ],
            "install_type": "git-clone",
            "description": "FoleyCrafter is a video-to-audio generation framework which can produce realistic sound effects semantically relevant and synchronized with videos."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Diffree",
            "id": "comfyui_diffree",
            "reference": "https://github.com/smthemex/ComfyUI_Diffree",
            "files": [
                "https://github.com/smthemex/ComfyUI_Diffree"
            ],
            "install_type": "git-clone",
            "description": "using diffree: Text-Guided Shape Free Object Inpainting with Diffusion Model"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Stable_Makeup",
            "id": "Stable_Makeup",
            "reference": "https://github.com/smthemex/ComfyUI_Stable_Makeup",
            "files": [
                "https://github.com/smthemex/ComfyUI_Stable_Makeup"
            ],
            "install_type": "git-clone",
            "description": "You can apply makeup to the characters in comfyui\nStable_Makeup From: [a/Stable_Makeup](https://mirror.ghproxy.com/https://github.com/Xiaojiu-z/Stable-Makeup)"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_FollowYourEmoji",
            "id": "FollowYourEmoji",
            "reference": "https://github.com/smthemex/ComfyUI_FollowYourEmoji",
            "files": [
                "https://github.com/smthemex/ComfyUI_FollowYourEmoji"
            ],
            "install_type": "git-clone",
            "description": "You can make emoji from a video and a image in comfyui"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_EchoMimic",
            "id": "EchoMimic",
            "reference": "https://github.com/smthemex/ComfyUI_EchoMimic",
            "files": [
                "https://github.com/smthemex/ComfyUI_EchoMimic"
            ],
            "install_type": "git-clone",
            "description": "You can using [a/EchoMimic](https://mirror.ghproxy.com/https://github.com/BadToBest/EchoMimic/tree/main) in comfyui,whitch Lifelike Audio-Driven Portrait Animations through Editable Landmark Conditioning "
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_ChatGLM_API",
            "id": "chatglm-api",
            "reference": "https://github.com/smthemex/ComfyUI_ChatGLM_API",
            "files": [
                "https://github.com/smthemex/ComfyUI_ChatGLM_API"
            ],
            "install_type": "git-clone",
            "description": "You can call Chatglm's API in comfyUI to translate and describe pictures, and the API similar to OpenAI."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_StableAudio_Open",
            "id": "stable-audio-open-1.0",
            "reference": "https://github.com/smthemex/ComfyUI_StableAudio_Open",
            "files": [
                "https://github.com/smthemex/ComfyUI_StableAudio_Open"
            ],
            "install_type": "git-clone",
            "description": "You can use stable-audio-open-1.0 in comfyUI"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_AnyDoor",
            "id": "ComfyUI_AnyDoor",
            "reference": "https://github.com/smthemex/ComfyUI_AnyDoor",
            "files": [
                "https://github.com/smthemex/ComfyUI_AnyDoor"
            ],
            "install_type": "git-clone",
            "description": "you can using anydoor ,change clothes,object"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_HiDiffusion_Pro",
            "id": "hidiffusion-pro",
            "reference": "https://github.com/smthemex/ComfyUI_HiDiffusion_Pro",
            "files": [
                "https://github.com/smthemex/ComfyUI_HiDiffusion_Pro"
            ],
            "install_type": "git-clone",
            "description": "A HiDiffusion node for ComfyUI."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_StoryDiffusion",
            "id": "StoryDiffusion",
            "reference": "https://github.com/smthemex/ComfyUI_StoryDiffusion",
            "files": [
                "https://github.com/smthemex/ComfyUI_StoryDiffusion"
            ],
            "install_type": "git-clone",
            "description": "A StoryDiffusion node for ComfyUI."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_MS_Diffusion",
            "id": "MS_Diffusion",
            "reference": "https://github.com/smthemex/ComfyUI_MS_Diffusion",
            "files": [
                "https://github.com/smthemex/ComfyUI_MS_Diffusion"
            ],
            "install_type": "git-clone",
            "description": "you can make story in comfyUI using MS-diffusion"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Pic2Story",
            "id": "pic2story",
            "reference": "https://github.com/smthemex/ComfyUI_Pic2Story",
            "files": [
                "https://github.com/smthemex/ComfyUI_Pic2Story"
            ],
            "install_type": "git-clone",
            "description": "you can using pic2story in comfyUI"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Llama3_8B",
            "id": "llama3-8b",
            "reference": "https://github.com/smthemex/ComfyUI_Llama3_8B",
            "files": [
                "https://github.com/smthemex/ComfyUI_Llama3_8B"
            ],
            "install_type": "git-clone",
            "description": "Llama3_8B for comfyUI， using pipeline workflow."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_ParlerTTS",
            "id": "parlertts",
            "reference": "https://github.com/smthemex/ComfyUI_ParlerTTS",
            "files": [
                "https://github.com/smthemex/ComfyUI_ParlerTTS"
            ],
            "install_type": "git-clone",
            "description": "Parler-TTS is a lightweight text-to-speech (TTS) model that can generate high-quality, natural sounding speech in the style of a given speaker (gender, pitch, speaking style, etc)"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Pipeline_Tool",
            "id": "pipeline-tool",
            "reference": "https://github.com/smthemex/ComfyUI_Pipeline_Tool",
            "files": [
                "https://github.com/smthemex/ComfyUI_Pipeline_Tool"
            ],
            "install_type": "git-clone",
            "description": "A tool for novice users in Chinese Mainland to call the huggingface hub and download the huggingface models."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_ID_Animator",
            "id": "id-animator",
            "reference": "https://github.com/smthemex/ComfyUI_ID_Animator",
            "files": [
                "https://github.com/smthemex/ComfyUI_ID_Animator"
            ],
            "install_type": "git-clone",
            "description": "This node allows you to use ID_Animator, the zero shot video generation model"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_CustomNet",
            "id": "customnet",
            "reference": "https://github.com/smthemex/ComfyUI_CustomNet",
            "files": [
                "https://github.com/smthemex/ComfyUI_CustomNet"
            ],
            "install_type": "git-clone",
            "description": "you can using customnet in comfyUI"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Pops",
            "id": "pops",
            "reference": "https://github.com/smthemex/ComfyUI_Pops",
            "files": [
                "https://github.com/smthemex/ComfyUI_Pops"
            ],
            "install_type": "git-clone",
            "description": "You can use [a/Popspaper](https://popspaper.github.io/pOps/) method in comfyUI"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Streamv2v_Plus",
            "id": "streamv2v",
            "reference": "https://github.com/smthemex/ComfyUI_Streamv2v_Plus",
            "files": [
                "https://github.com/smthemex/ComfyUI_Streamv2v_Plus"
            ],
            "install_type": "git-clone",
            "description": "using [a/StreamV2V](https://mirror.ghproxy.com/https://github.com/Jeff-LiangF/streamv2v) in ComfyUI"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_MooER",
            "id": "comfyui_mooer",
            "reference": "https://github.com/smthemex/ComfyUI_MooER",
            "files": [
                "https://github.com/smthemex/ComfyUI_MooER"
            ],
            "install_type": "git-clone",
            "description": "MooER is an LLM-based Speech Recognition and Translation Model from Moore Threads.You can use MooER when install ComfyUI_MooER node"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_CSGO_Wrapper",
            "id": "comfyui_csgo_wrapper",
            "reference": "https://github.com/smthemex/ComfyUI_CSGO_Wrapper",
            "files": [
                "https://github.com/smthemex/ComfyUI_CSGO_Wrapper"
            ],
            "install_type": "git-clone",
            "description": "using InstantX's CSGO in comfyUI for style"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_DeepFakeDefenders",
            "id": "comfyui_deepfakedefenders",
            "reference": "https://github.com/smthemex/ComfyUI_DeepFakeDefenders",
            "files": [
                "https://github.com/smthemex/ComfyUI_DeepFakeDefenders"
            ],
            "install_type": "git-clone",
            "description": "ou can using DeepFakeDefenders in comfyUI to Prediction image is a DeepFake img or not."
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Sapiens",
            "reference": "https://github.com/smthemex/ComfyUI_Sapiens",
            "files": [
                "https://github.com/smthemex/ComfyUI_Sapiens"
            ],
            "install_type": "git-clone",
            "description": "You can call Using Sapiens to get seg,normal,pose,depth,mask maps. Sapiens From: [a/facebookresearch/sapiens](https://mirror.ghproxy.com/https://github.com/facebookresearch/sapiens)"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_Hallo2",
            "reference": "https://github.com/smthemex/ComfyUI_Hallo2",
            "files": [
                "https://github.com/smthemex/ComfyUI_Hallo2"
            ],
            "install_type": "git-clone",
            "description": "Long-Duration and High-Resolution Audio-driven Portrait Image Animation,"
        },
        {
            "author": "smthemex",
            "title": "ComfyUI_OmniParser",
            "reference": "https://github.com/smthemex/ComfyUI_OmniParser",
            "files": [
                "https://github.com/smthemex/ComfyUI_OmniParser"
            ],
            "install_type": "git-clone",
            "description": "Try [a/OmniParser](https://mirror.ghproxy.com/https://github.com/microsoft/OmniParser) in ComfyUI which a simple screen parsing tool towards pure vision based GUI agent."
        },
        {
            "author": "choey",
            "title": "Comfy-Topaz",
            "id": "topaz",
            "reference": "https://github.com/choey/Comfy-Topaz",
            "files": [
                "https://github.com/choey/Comfy-Topaz"
            ],
            "install_type": "git-clone",
            "description": "Comfy-Topaz is a custom node for ComfyUI, which integrates with Topaz Photo AI to enhance (upscale, sharpen, denoise, etc.) images, allowing this traditionally asynchronous step to become a part of ComfyUI workflows.\nNOTE:Licensed installation of Topaz Photo AI"
        },
        {
            "author": "ALatentPlace",
            "title": "ComfyUI_yanc",
            "id": "yanc-alatentplace",
            "reference": "https://github.com/ALatentPlace/ComfyUI_yanc",
            "files": [
                "https://github.com/ALatentPlace/ComfyUI_yanc"
            ],
            "install_type": "git-clone",
            "description": "Yet Another Node Collection. Adds some useful nodes, check out the GitHub page for more details."
        },
        {
            "author": "Wicloz",
            "title": "ComfyUI Simply Nodes",
            "reference": "https://github.com/Wicloz/ComfyUI-Simply-Nodes",
            "files": [
                "https://github.com/Wicloz/ComfyUI-Simply-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Simple nodes to help clean up your workflow, mostly focussed on text operations."
        },
        {
            "author": "wandbrandon",
            "title": "comfyui-pixel",
            "id": "pixel",
            "reference": "https://github.com/wandbrandon/comfyui-pixel",
            "files": [
                "https://github.com/wandbrandon/comfyui-pixel"
            ],
            "install_type": "git-clone",
            "description": "pixel art workshop nodes for comfyui."
        },
        {
            "author": "nullquant",
            "title": "BrushNet",
            "id": "brushnet",
            "reference": "https://github.com/nullquant/ComfyUI-BrushNet",
            "files": [
                "https://github.com/nullquant/ComfyUI-BrushNet"
            ],
            "install_type": "git-clone",
            "description": "These are custom nodes for ComfyUI native implementation of [a/BrushNet](https://arxiv.org/abs/2403.06976) (inpaint), PowerPaint (inpaint, object removal) and HiDiffusion (higher resolution for SD15 and SDXL)"
        },
        {
            "author": "pamparamm",
            "title": "Perturbed-Attention Guidance",
            "id": "pag",
            "reference": "https://github.com/pamparamm/sd-perturbed-attention",
            "files": [
                "https://github.com/pamparamm/sd-perturbed-attention"
            ],
            "install_type": "git-clone",
            "description": "Perturbed-Attention Guidance with advanced parameters for ComfyUI. (PAG)"
        },
        {
            "author": "pamparamm",
            "title": "ComfyUI Vectorscope CC",
            "id": "vectorscope",
            "reference": "https://github.com/pamparamm/ComfyUI-vectorscope-cc",
            "files": [
                "https://github.com/pamparamm/ComfyUI-vectorscope-cc"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI port of Vectorscope CC and Diffusion Color Grading by Haoming02. Makes it possible to adjust Brightness/Contrast/Saturation/Hue during image generation."
        },
        {
            "author": "pamparamm",
            "title": "ComfyUI-ppm",
            "id": "comfyui-ppm",
            "reference": "https://github.com/pamparamm/ComfyUI-ppm",
            "files": [
                "https://github.com/pamparamm/ComfyUI-ppm"
            ],
            "install_type": "git-clone",
            "description": "Fixed AttentionCouple, NegPip(negative weights in prompts) for SDXL and FLUX, more CFG++ and SMEA DY samplers, etc."
        },
        {
            "author": "unwdef",
            "title": "unwdef-nodes",
            "reference": "https://github.com/unwdef/unwdef-nodes-comfyui",
            "files": [
                "https://github.com/unwdef/unwdef-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for ComfyUI by unwdef."
        },
        {
            "author": "fevre27",
            "title": "Self-Guidance nodes",
            "id": "self-guidance",
            "reference": "https://github.com/forever22777/comfyui-self-guidance",
            "files": [
                "https://github.com/forever22777/comfyui-self-guidance"
            ],
            "install_type": "git-clone",
            "description": "Unofficial ComfyUI implementation of Self-Guidance."
        },
        {
            "author": "aburahamu",
            "title": "ComfyUI-RequestPoster",
            "id": "request-poster",
            "reference": "https://github.com/aburahamu/ComfyUI-RequestsPoster",
            "files": [
                "https://github.com/aburahamu/ComfyUI-RequestsPoster"
            ],
            "install_type": "git-clone",
            "description": "This extension can send HTTP Requests. You can request image generation to StableDiffusion3 and post images to X (Twitter) and Discord."
        },
        {
            "author": "aburahamu",
            "title": "ComfyUI-IsNiceParts",
            "id": "isniceparts",
            "reference": "https://github.com/aburahamu/ComfyUI-IsNiceParts",
            "files": [
                "https://github.com/aburahamu/ComfyUI-IsNiceParts"
            ],
            "install_type": "git-clone",
            "description": "This custom node detects body parts (currently only hands) from the received image and outputs the image if the skeleton can be estimated."
        },
        {
            "author": "Sorcerio",
            "title": "MBM's Music Visualizer",
            "reference": "https://github.com/Sorcerio/MBM-Music-Visualizer",
            "files": [
                "https://github.com/Sorcerio/MBM-Music-Visualizer"
            ],
            "install_type": "git-clone",
            "description": "An image generation based music visualizer integrated into comfyanonymous/ComfyUI as custom nodes."
        },
        {
            "author": "quadmoon",
            "title": "quadmoon's ComfyUI nodes",
            "reference": "https://github.com/traugdor/ComfyUI-quadMoons-nodes",
            "files": [
                "https://github.com/traugdor/ComfyUI-quadMoons-nodes"
            ],
            "install_type": "git-clone",
            "description": "These are just some nodes I wanted and couldn't find where anyone else had made them yet."
        },
        {
            "author": "quadme7macoon",
            "title": "ComfyUI-ShadertoyGL",
            "reference": "https://github.com/e7mac/ComfyUI-ShadertoyGL",
            "files": [
                "https://github.com/e7mac/ComfyUI-ShadertoyGL"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Shadertoy, Shader, ColorChannelOffset."
        },
        {
            "author": "royceschultz",
            "title": "ComfyUI-TranscriptionTools",
            "id": "transcription-tools",
            "reference": "https://github.com/royceschultz/ComfyUI-TranscriptionTools",
            "files": [
                "https://github.com/royceschultz/ComfyUI-TranscriptionTools"
            ],
            "install_type": "git-clone",
            "description": "Transcribe audio and video files in ComfyUI."
        },
        {
            "author": "kunieone",
            "title": "ComfyUI_alkaid",
            "id": "alkadi",
            "reference": "https://github.com/kunieone/ComfyUI_alkaid",
            "files": [
                "https://github.com/kunieone/ComfyUI_alkaid"
            ],
            "install_type": "git-clone",
            "description": "Nodes:A_Face3DSwapper, A_FaceCrop, A_FacePaste, A_OpenPosePreprocessor, A_EmptyLatentImageLongside, A_GetImageSize, AlkaidLoader, AdapterFaceLoader, AdapterStyleLoader, ..."
        },
        {
            "author": "txt2any",
            "title": "ComfyUI-PromptOrganizer",
            "id": "prompt-organizer",
            "reference": "https://github.com/txt2any/ComfyUI-PromptOrganizer",
            "files": [
                "https://github.com/txt2any/ComfyUI-PromptOrganizer"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node for ComfyUI that automatically saves your AI-generated images specifically to [a/www.txt2any.com](http://www.txt2any.com/)."
        },
        {
            "author": "kealiu",
            "title": "ComfyUI Load and Save file to S3",
            "id": "savefile-to-s3",
            "reference": "https://github.com/kealiu/ComfyUI-S3-Tools",
            "files": [
                "https://github.com/kealiu/ComfyUI-S3-Tools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Load From S3, Save To S3."
        },
        {
            "author": "kealiu",
            "title": "ComfyUI-ZeroShot-MTrans",
            "id": "zeroshot-mtrans",
            "reference": "https://github.com/kealiu/ComfyUI-ZeroShot-MTrans",
            "files": [
                "https://github.com/kealiu/ComfyUI-ZeroShot-MTrans"
            ],
            "install_type": "git-clone",
            "description": "An unofficial ComfyUI custom node for [a/Zero-Shot Material Transfer from a Single Image](https://ttchengab.github.io/zest), Given an input image (e.g., a photo of an apple) and a single material exemplar image (e.g., a golden bowl), ZeST can transfer the gold material from the exemplar onto the apple with accurate lighting cues while making everything else consistent."
        },
        {
            "author": "kealiu",
            "title": "ComfyUI-Zero123-Porting",
            "id": "zero123-porting",
            "reference": "https://github.com/kealiu/ComfyUI-Zero123-Porting",
            "files": [
                "https://github.com/kealiu/ComfyUI-Zero123-Porting"
            ],
            "install_type": "git-clone",
            "description": "Zero-1-to-3: Zero-shot One Image to 3D Object, unofficial porting of original [Zero123](https://mirror.ghproxy.com/https://github.com/cvlab-columbia/zero123)"
        },
        {
            "author": "TashaSkyUp",
            "title": "ComfyUI_LiteLLM",
            "id": "litellm",
            "reference": "https://github.com/Hopping-Mad-Games/ComfyUI_LiteLLM",
            "files": [
                "https://github.com/Hopping-Mad-Games/ComfyUI_LiteLLM"
            ],
            "install_type": "git-clone",
            "description": "Nodes for calling LLMs, enabled by LiteLLM"
        },
        {
            "author": "AonekoSS",
            "title": "ComfyUI-SimpleCounter",
            "id": "simplecounter",
            "reference": "https://github.com/AonekoSS/ComfyUI-SimpleCounter",
            "files": [
                "https://github.com/AonekoSS/ComfyUI-SimpleCounter"
            ],
            "install_type": "git-clone",
            "description": "Node: utils/Simple Counter\nThis node is a simple counter, when pressing 'Queue Prompt' resets the count."
        },
        {
            "author": "AonekoSS",
            "title": "ComfyUI-LoRA-Tuner",
            "id": "lora-tuner",
            "reference": "https://github.com/AonekoSS/ComfyUI-LoRA-Tuner",
            "files": [
                "https://github.com/AonekoSS/ComfyUI-LoRA-Tuner"
            ],
            "install_type": "git-clone",
            "description": "Nodes: LoRA-Tuner. For using multiple LoRA easily."
        },
        {
            "author": "heshengtao",
            "title": "comfyui_LLM_party",
            "id": "llm-party",
            "reference": "https://github.com/heshengtao/comfyui_LLM_party",
            "files": [
                "https://github.com/heshengtao/comfyui_LLM_party"
            ],
            "install_type": "git-clone",
            "description": "A set of block-based LLM agent node libraries designed for ComfyUI.This project aims to develop a complete set of nodes for LLM workflow construction based on comfyui. It allows users to quickly and conveniently build their own LLM workflows and easily integrate them into their existing SD workflows."
        },
        {
            "author": "heshengtao",
            "title": "comfyui_LLM_schools",
            "reference": "https://github.com/heshengtao/comfyui_LLM_schools",
            "files": [
                "https://github.com/heshengtao/comfyui_LLM_schools"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node library for fine-tuning LLMs"
        },
        {
            "author": "VAST-AI-Research",
            "title": "Tripo for ComfyUI",
            "id": "tripo",
            "reference": "https://github.com/VAST-AI-Research/ComfyUI-Tripo",
            "files": [
                "https://github.com/VAST-AI-Research/ComfyUI-Tripo"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for using [a/Tripo](https://www.tripo3d.ai/) in ComfyUI to create 3D from text and image prompts."
        },
        {
            "author": "JettHu",
            "title": "ComfyUI_TGate",
            "id": "tgate",
            "reference": "https://github.com/JettHu/ComfyUI_TGate",
            "files": [
                "https://github.com/JettHu/ComfyUI_TGate"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI reference implementation for [a/T-GATE](https://mirror.ghproxy.com/https://github.com/HaozheLiu-ST/T-GATE)."
        },
        {
            "author": "JettHu",
            "title": "ComfyUI-TCD",
            "id": "jetthu-tcd",
            "reference": "https://github.com/JettHu/ComfyUI-TCD",
            "files": [
                "https://github.com/JettHu/ComfyUI-TCD"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation for [a/TCD](https://mirror.ghproxy.com/https://github.com/jabir-zheng/TCD)."
        },
        {
            "author": "sugarkwork",
            "title": "comfyui_tag_filter",
            "id": "tag-filter",
            "reference": "https://github.com/sugarkwork/comfyui_tag_fillter",
            "files": [
                "https://github.com/sugarkwork/comfyui_tag_fillter"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node of ComfyUI that categorizes tags outputted by tools like WD14Tagger, filters them by each category, and returns the filtered results."
        },
        {
            "author": "Intersection98",
            "title": "ComfyUI-MX-post-processing-nodes",
            "reference": "https://github.com/Intersection98/ComfyUI_MX_post_processing-nodes",
            "files": [
                "https://github.com/Intersection98/ComfyUI_MX_post_processing-nodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of post processing nodes for ComfyUI, dds image post-processing adjustment capabilities to the ComfyUI."
        },
        {
            "author": "TencentQQGYLab",
            "title": "ComfyUI-ELLA",
            "id": "ella",
            "reference": "https://github.com/TencentQQGYLab/ComfyUI-ELLA",
            "files": [
                "https://github.com/TencentQQGYLab/ComfyUI-ELLA"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation for [a/ELLA](https://mirror.ghproxy.com/https://github.com/TencentQQGYLab/ELLA)."
        },
        {
            "author": "DarKDinDoN",
            "title": "ComfyUI Checkpoint Automatic Config",
            "id": "checkpoint-autoconfig",
            "reference": "https://github.com/mech-tools/comfyui-checkpoint-automatic-config",
            "files": [
                "https://github.com/mech-tools/comfyui-checkpoint-automatic-config"
            ],
            "install_type": "git-clone",
            "description": "This node was designed to help with checkpoint configuration. Fee free to add new checkpoint configurations!"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-Prompt-MZ",
            "id": "prompt-mz",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-Prompt-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-Prompt-MZ"
            ],
            "install_type": "git-clone",
            "description": "Use llama.cpp to help generate some nodes for prompt word related work"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-StylizePhoto-MZ",
            "id": "stylizephoto",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-StylizePhoto-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-StylizePhoto-MZ"
            ],
            "install_type": "git-clone",
            "description": "A stylized node with simple operation. The effect is achieved by I2I and lora. The clay style is currently implemented.Comes with watermark function."
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-TrainTools-MZ",
            "id": "traintools",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-TrainTools-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-TrainTools-MZ"
            ],
            "install_type": "git-clone",
            "description": "Nodes for fine-tuning lora in ComfyUI, dependent on training tools such as kohya-ss/sd-scripts"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-Kolors-MZ",
            "id": "kolors-mz",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-Kolors-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-Kolors-MZ"
            ],
            "install_type": "git-clone",
            "description": "Implementation of Kolors on ComfyUI\nReference from [a/https://mirror.ghproxy.com/https://github.com/kijai/ComfyUI-KwaiKolorsWrapper](https://mirror.ghproxy.com/https://github.com/kijai/ComfyUI-KwaiKolorsWrapper)\nUsing ComfyUI Native Sampling"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-Flux1Quantize-MZ",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-Flux1Quantize-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-Flux1Quantize-MZ"
            ],
            "pip": ["git+https://mirror.ghproxy.com/https://github.com/IST-DASLab/marlin"],
            "install_type": "git-clone",
            "description": "Quantization tools are from [a/https://mirror.ghproxy.com/https://github.com/casper-hansen/AutoAWQ](https://mirror.ghproxy.com/https://github.com/casper-hansen/AutoAWQ) and [a/https://mirror.ghproxy.com/https://github.com/IST-DASLab/marlin](https://mirror.ghproxy.com/https://github.com/IST-DASLab/marlin)\nOnly applicable to graphics cards with sm_80 and above (30 series and above)\nNeed to install marlin dependencies first"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-FluxExt-MZ",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-FluxExt-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-FluxExt-MZ"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MZ_Flux1PartialLoad_Patch. Tool nodes related to flux1"
        },
        {
            "author": "MinusZoneAI",
            "title": "ComfyUI-CogVideoX-MZ",
            "reference": "https://github.com/MinusZoneAI/ComfyUI-CogVideoX-MZ",
            "files": [
                "https://github.com/MinusZoneAI/ComfyUI-CogVideoX-MZ"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MZ_CogVideoXLoader"
        },
        {
            "author": "blueraincoatli",
            "title": "comfyUI_SillyNodes",
            "id": "silly",
            "reference": "https://github.com/blueraincoatli/comfyUI_SillyNodes",
            "files": [
                "https://github.com/blueraincoatli/comfyUI_SillyNodes"
            ],
            "install_type": "git-clone",
            "description": "Using rgthree's fast_group_muter and bookmark nodes, introduce the pyautogui library to simulate clicks and hotkeys, and run groups in sequence. screen manipulation is involved"
        },
        {
            "author": "ty0x2333",
            "title": "ComfyUI-Dev-Utils",
            "id": "dev-utils",
            "reference": "https://github.com/ty0x2333/ComfyUI-Dev-Utils",
            "files": [
                "https://github.com/ty0x2333/ComfyUI-Dev-Utils"
            ],
            "install_type": "git-clone",
            "description": "Execution Time Analysis, Reroute Enhancement, Node collection for developers."
        },
        {
            "author": "lquesada",
            "title": "ComfyUI-Prompt-Combinator",
            "id": "prompt-combinator",
            "reference": "https://github.com/lquesada/ComfyUI-Prompt-Combinator",
            "files": [
                "https://github.com/lquesada/ComfyUI-Prompt-Combinator"
            ],
            "install_type": "git-clone",
            "description": "'🔢 Prompt Combinator' is a node that generates all possible combinations of prompts from several lists of strings.\n'🔢 Prompt Combinator Merger' is a node that enables merging the output of two different '🔢 Prompt Combinator' nodes."
        },
        {
            "author": "lquesada",
            "title": "ComfyUI-Inpaint-CropAndStitch",
            "id": "crop-and-stitch",
            "reference": "https://github.com/lquesada/ComfyUI-Inpaint-CropAndStitch",
            "files": [
                "https://github.com/lquesada/ComfyUI-Inpaint-CropAndStitch"
            ],
            "install_type": "git-clone",
            "description": "'✂️ Inpaint Crop' is a node that crops an image before sampling. The context area can be specified via the mask, expand pixels and expand factor or via a separate (optional) mask.\n'✂️ Inpaint Stitch' is a node that stitches the inpainted image back into the original image without altering unmasked areas."
        },
        {
            "author": "randjtw",
            "title": "advance-aesthetic-score",
            "reference": "https://github.com/randjtw/advance-aesthetic-score",
            "files": [
                "https://github.com/randjtw/advance-aesthetic-score"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Advance Aesthetic Score"
        },
        {
            "author": "FredBill1",
            "title": "comfyui-fb-utils",
            "id": "fb-utils",
            "reference": "https://github.com/FredBill1/comfyui-fb-utils",
            "files": [
                "https://github.com/FredBill1/comfyui-fb-utils"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FBStringJoin, FBStringSplit, FBMultilineStringList, FBMultilineString"
        },
        {
            "author": "jeffy5",
            "title": "Faceless Node for ComfyUI",
            "id": "faceless",
            "reference": "https://github.com/jeffy5/comfyui-faceless-node",
            "files": [
                "https://github.com/jeffy5/comfyui-faceless-node"
            ],
            "install_type": "git-clone",
            "description": "A facefusion custom node for ComfyUI. Swap or restore faces for image or video"
        },
        {
            "author": "TaiTair",
            "title": "Simswap Node for ComfyUI",
            "id": "simswap",
            "reference": "https://github.com/TaiTair/comfyui-simswap",
            "files": [
                "https://github.com/TaiTair/comfyui-simswap"
            ],
            "install_type": "git-clone",
            "description": "A hacky implementation of Simswap based on [a/Comfyui ReActor Node 0.5.1](https://mirror.ghproxy.com/https://github.com/Gourieff/comfyui-reactor-node) and [a/Simswap](https://mirror.ghproxy.com/https://github.com/neuralchen/SimSwap)."
        },
        {
            "author": "fofr",
            "title": "ComfyUI-HyperSDXL1StepUnetScheduler (ByteDance)",
            "id": "hypersdxl",
            "reference": "https://github.com/fofr/ComfyUI-HyperSDXL1StepUnetScheduler",
            "files": [
                "https://github.com/fofr/ComfyUI-HyperSDXL1StepUnetScheduler"
            ],
            "install_type": "git-clone",
            "description": "Original author is ByteDance.\nComfyUI sampler for HyperSDXL UNet\nPorted from: [a/https://huggingface.co/ByteDance/Hyper-SD](https://huggingface.co/ByteDance/Hyper-SD)"
        },
        {
            "author": "fofr",
            "title": "ComfyUI-Prompter-fofrAI",
            "id": "prompter-fofr",
            "reference": "https://github.com/fofr/ComfyUI-Prompter-fofrAI",
            "files": [
                "https://github.com/fofr/ComfyUI-Prompter-fofrAI"
            ],
            "install_type": "git-clone",
            "description": "A prompt helper for ComfyUI, based on [a/prompter.fofr.ai](https://prompter.fofr.ai)"
        },
        {
            "author": "fofr",
            "title": "comfyui-fofr-toolkit",
            "id": "fofr-toolkit",
            "reference": "https://github.com/fofr/comfyui-fofr-toolkit",
            "files": [
                "https://github.com/fofr/comfyui-fofr-toolkit"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Incrementer, Width and height from aspect ratio, Width and height for scaling image to ideal resolutio. A simple set of tooling nodes."
        },
        {
            "author": "fofr",
            "title": "ComfyUI-Replicate",
            "reference": "https://github.com/replicate/comfyui-replicate",
            "files": [
                "https://github.com/replicate/comfyui-replicate"
            ],
            "install_type": "git-clone",
            "description": "Run [a/Replicate models](https://replicate.com/explore) in ComfyUI."
        },
        {
            "author": "cfreilich",
            "title": "Virtuoso Nodes for ComfyUI",
            "id": "virtuoso",
            "reference": "https://github.com/chrisfreilich/virtuoso-nodes",
            "files": [
                "https://github.com/chrisfreilich/virtuoso-nodes"
            ],
            "install_type": "git-clone",
            "description": "Photoshop type functions and adjustment layers: 30 blend modes, Selective Color, Blend If, Color Balance, Solid Color Images, Black and White, Hue/Saturation, Levels, and RGB Splitting and Merging."
        },
        {
            "author": "da2el-ai",
            "title": "D2 Nodes ComfyUI",
            "id": "d2-nodes-comfyui",
            "reference": "https://github.com/da2el-ai/D2-nodes-ComfyUI",
            "files": [
                "https://github.com/da2el-ai/D2-nodes-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This is a collection of custom nodes that make ComfyUI slightly more convenient.[w/This node pack integrates ComfyUI-d2-size-selector, ComfyUI-d2-steps, and ComfyUI-d2-xyplot-utils. To avoid conflicts, please delete the existing node packs if you install this one.]"
        },
        {
            "author": "da2el-ai",
            "title": "D2 Steps",
            "id": "d2steps",
            "reference": "https://github.com/da2el-ai/ComfyUI-d2-steps",
            "files": [
                "https://github.com/da2el-ai/ComfyUI-d2-steps"
            ],
            "install_type": "git-clone",
            "description": "A handy custom node for using Refiner (switching to a different checkpoint midway) When you specify the end of the base checkpoint, you can extract refiner_start which is end + 1. The output is fixed as an INT, so it can be passed to the handy custom node, Anything Everywhere? Since it only outputs a numerical value, it can also be used for other purposes."
        },
        {
            "author": "da2el-ai",
            "title": "D2 Size Selector",
            "id": "size-selector",
            "reference": "https://github.com/da2el-ai/ComfyUI-d2-size-selector",
            "files": [
                "https://github.com/da2el-ai/ComfyUI-d2-size-selector"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node that allows you to easily call up and set image size presets. Settings can be made by editing the included config.yaml. It is almost identical to Comfyroll Studio's CR AspectRatio. I created it because I wanted to easily edit the presets."
        },
        {
            "author": "da2el-ai",
            "title": "D2 Send Eagle",
            "id": "d2-send-eagle",
            "reference": "https://github.com/da2el-ai/ComfyUI-d2-send-eagle",
            "files": [
                "https://github.com/da2el-ai/ComfyUI-d2-send-eagle"
            ],
            "install_type": "git-clone",
            "description": "Send images generated by ComfyUI to Eagle image management software"
        },
        {
            "author": "da2el-ai",
            "title": "D2 XYPlot Utils",
            "reference": "https://github.com/da2el-ai/ComfyUI-d2-xyplot-utils",
            "files": [
                "https://github.com/da2el-ai/ComfyUI-d2-xyplot-utils"
            ],
            "install_type": "git-clone",
            "description": "Custom node for using Prompt S/R in XY Plot\nAlso includes nodes for listing generic parameters like seed and cfg\nEasy to manipulate as elements are separated by line breaks\nDesigned for use with the XY Plot custom node qq-nodes-comfyui, but may work with other custom nodes as well"
        },
        {
            "author": "nat-chan",
            "title": "ComfyUI-Transceiver📡",
            "id": "transceiver",
            "reference": "https://github.com/nat-chan/comfyui-transceiver",
            "files": [
                "https://github.com/nat-chan/comfyui-transceiver"
            ],
            "install_type": "git-clone",
            "description": "Transceiver is a python library that swiftly exchanges fundamental data structures, specifically numpy arrays, between processes, optimizing AI inference tasks that utilize ComfyUI."
        },
        {
            "author": "nat-chan",
            "title": "ComfyUI-graphToPrompt",
            "id": "graph2prompt",
            "reference": "https://github.com/nat-chan/ComfyUI-graphToPrompt",
            "files": [
                "https://github.com/nat-chan/ComfyUI-graphToPrompt"
            ],
            "install_type": "git-clone",
            "description": "workflow.json -> workflow_api.json"
        },
        {
            "author": "web3nomad",
            "title": "ComfyUI Invisible Watermark",
            "id": "invisible-watermark",
            "reference": "https://github.com/web3nomad/ComfyUI_Invisible_Watermark",
            "files": [
                "https://github.com/web3nomad/ComfyUI_Invisible_Watermark"
            ],
            "install_type": "git-clone",
            "description": "Nodes: InvisibleWatermarkEncode"
        },
        {
            "author": "GentlemanHu",
            "title": "ComfyUI Suno API",
            "id": "suno-api",
            "reference": "https://github.com/GentlemanHu/ComfyUI-SunoAI",
            "files": [
                "https://github.com/GentlemanHu/ComfyUI-SunoAI"
            ],
            "install_type": "git-clone",
            "description": "An unofficial Python library for [a/Suno AI](https://www.suno.ai/) API"
        },
        {
            "author": "TemryL",
            "title": "ComfyUI-IDM-VTON [WIP]",
            "id": "idm-vton",
            "reference": "https://github.com/TemryL/ComfyUI-IDM-VTON",
            "files": [
                "https://github.com/TemryL/ComfyUI-IDM-VTON"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI adaptation of [a/IDM-VTON](https://mirror.ghproxy.com/https://github.com/yisol/IDM-VTON) for virtual try-on."
        },
        {
            "author": "NStor",
            "title": "ComfyUI-RUS localization",
            "reference": "https://github.com/Nestorchik/NStor-ComfyUI-Translation",
            "files": [
                "https://github.com/Nestorchik/NStor-ComfyUI-Translation"
            ],
            "install_type": "git-clone",
            "description": "Russian localization of ComfyUI, ComafyUI-Manager & more..."
        },
        {
            "author": "jax-explorer",
            "title": "fast_video_comfyui",
            "reference": "https://github.com/jax-explorer/fast_video_comfyui",
            "files": [
                "https://github.com/jax-explorer/fast_video_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FastImageListToImageBatch"
        },
        {
            "author": "sugarkwork",
            "title": "comfyui_cohere",
            "id": "cohere",
            "reference": "https://github.com/sugarkwork/comfyui_cohere",
            "files": [
                "https://github.com/sugarkwork/comfyui_cohere"
            ],
            "install_type": "git-clone",
            "description": "This is a node for using cohere (Command R+) from ComfyUI. You need to edit the startup .bat file of ComfyUI and describe the API key obtained from Cohere as follows."
        },
        {
            "author": "alessandrozonta",
            "title": "Bounding Box Crop Node for ComfyUI",
            "id": "bbox-crop",
            "reference": "https://github.com/alessandrozonta/ComfyUI-CenterNode",
            "files": [
                "https://github.com/alessandrozonta/ComfyUI-CenterNode"
            ],
            "install_type": "git-clone",
            "description": "This extension contains a custom node for ComfyUI. The node, called 'Bounding Box Crop', is designed to compute the top-left coordinates of a cropped bounding box based on input coordinates and dimensions of the final cropped image. It does so computing the center of the cropping area and then computing where the top-left coordinates would be."
        },
        {
            "author": "alessandrozonta",
            "title": "Save Layers Node for ComfyUI",
            "id": "layers",
            "reference": "https://github.com/alessandrozonta/ComfyUI-Layers",
            "files": [
                "https://github.com/alessandrozonta/ComfyUI-Layers"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI allows you to create layers of an image based on input masks and save them into a PSD file."
        },
        {
            "author": "alessandrozonta",
            "title": "OpenPose Node",
            "id": "openpose-alessandrozonta",
            "reference": "https://github.com/alessandrozonta/ComfyUI-OpenPose",
            "files": [
                "https://github.com/alessandrozonta/ComfyUI-OpenPose"
            ],
            "install_type": "git-clone",
            "description": "This extension contains a custom node for ComfyUI. The node, called 'Bounding Box Crop', is designed to compute the top-left coordinates of a cropped bounding box based on input coordinates and dimensions of the final cropped image. It does so computing the center of the cropping area and then computing where the top-left coordinates would be."
        },
        {
            "author": "curiousjp",
            "title": "ComfyUI-MaskBatchPermutations",
            "id": "maskbatch-permutations",
            "reference": "https://github.com/curiousjp/ComfyUI-MaskBatchPermutations",
            "files": [
                "https://github.com/curiousjp/ComfyUI-MaskBatchPermutations"
            ],
            "install_type": "git-clone",
            "description": "Permutes a mask batch to present possible additive combinations. Passing a mask batch (e.g. out of [a/SEGS to Mask Batch](https://mirror.ghproxy.com/https://github.com/ltdrdata/ComfyUI-Impact-Pack)) will return a new mask batch representing all the possible combinations of the included masks. So, a mask batch with two mask sections, 'A' and 'B', will return a batch containing an empty mask, an empty mask & A, an empty mask & B, and an empty mask & A & B."
        },
        {
            "author": "BAIS1C",
            "title": "ComfyUI_RSS_Feed_Reader",
            "id": "rssfeed",
            "reference": "https://github.com/BAIS1C/ComfyUI_RSS_Feed_Reader",
            "files": [
                "https://github.com/BAIS1C/ComfyUI_RSS_Feed_Reader"
            ],
            "install_type": "git-clone",
            "description": "A Simple Python RSS Feed Reader to create Prompts in Comfy UI"
        },
        {
            "author": "runtime44",
            "title": "Runtime44 ComfyUI Nodes",
            "reference": "https://github.com/runtime44/comfyui_r44_nodes",
            "files": [
                "https://github.com/runtime44/comfyui_r44_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Runtime44Upscaler, Runtime44ColorMatch, Runtime44DynamicKSampler, Runtime44ImageOverlay, Runtime44ImageResizer, Runtime44ImageToNoise, Runtime44MaskSampler, Runtime44TiledMaskSampler, Runtime44IterativeUpscaleFactor, Runtime44ImageEnhance, Runtime44FilmGrain"
        },
        {
            "author": "osiworx",
            "title": "ComfyUI_Prompt-Quill",
            "reference": "https://github.com/osi1880vr/prompt_quill_comfyui",
            "files": [
                "https://github.com/osi1880vr/prompt_quill_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Use Prompt Quill in Comfyui"
        },
        {
            "author": "philz1337x",
            "title": "✨ Clarity AI - Creative Image Upscaler and Enhancer for ComfyUI",
            "reference": "https://github.com/philz1337x/ComfyUI-ClarityAI",
            "files": [
                "https://github.com/philz1337x/ComfyUI-ClarityAI"
            ],
            "install_type": "git-clone",
            "description": "[a/Clarity AI](https://clarityai.cc) is a creative image enhancer and is able to upscale to high resolution. [w/NOTE: This is a Magnific AI alternative for ComfyUI.] \nCreate an API key on [a/ClarityAI.cc/api](https://clarityai.cc/api) and add to environment variable 'CAI_API_KEY'\nAlternatively you can write your API key to file 'cai_platform_key.txt'\nYou can also use and/or override the above by entering your API key in the 'api_key_override' field of the node."
        },
        {
            "author": "KoreTeknology",
            "title": "ComfyUI Universal Styler",
            "id": "universal-styler",
            "reference": "https://github.com/KoreTeknology/ComfyUI-Universal-Styler",
            "files": [
                "https://github.com/KoreTeknology/ComfyUI-Universal-Styler"
            ],
            "install_type": "git-clone",
            "description": "A research Node based project on Artificial Intelligence using ComfyUI visual editor with Stable diffusion Local processing focus in mind. This custom node is intended to serve the purpose to offer a large palette of prompting scenrarios, based on Public Checkpoint Models OR/AND Private custom Models and LoRas. It includes an integrated learning machine process as well as a set of workflows."
        },
        {
            "author": "KoreTeknology",
            "title": "ComfyUI Compositing Nodes Pack",
            "reference": "https://github.com/KoreTeknology/ComfyUI-Compositing-Nodes-Pack",
            "files": [
                "https://github.com/KoreTeknology/ComfyUI-Compositing-Nodes-Pack"
            ],
            "install_type": "git-clone",
            "description": "This is set of custom nodes for your ComfyUI1 local installation. It offers the very basic nodes that are missing in the official 'Vanilla' package. It is a research Node based project on Artificial Intelligence using ComfyUI visual editor. This repository also includes a set of workflows to test the nodes."
        },
        {
            "author": "ZeDarkAdam",
            "title": "ComfyUI-Embeddings-Tools",
            "id": "embeddings-tools",
            "reference": "https://github.com/ZeDarkAdam/ComfyUI-Embeddings-Tools",
            "files": [
                "https://github.com/ZeDarkAdam/ComfyUI-Embeddings-Tools"
            ],
            "install_type": "git-clone",
            "description": "EmbeddingsNameLoader, EmbendingList"
        },
        {
            "author": "chenpx976",
            "title": "ComfyUI-RunRunRun",
            "id": "runrunrun",
            "reference": "https://github.com/chenpx976/ComfyUI-RunRunRun",
            "files": [
                "https://github.com/chenpx976/ComfyUI-RunRunRun"
            ],
            "install_type": "git-clone",
            "description": "add http api http://127.0.0.1:8188/comfyui-run/run use in other llm project."
        },
        {
            "author": "githubYiheng",
            "title": "ComfyUI_GetFileNameFromURL",
            "id": "getfilename-from-url",
            "reference": "https://github.com/githubYiheng/ComfyUI_GetFileNameFromURL",
            "files": [
                "https://github.com/githubYiheng/ComfyUI_GetFileNameFromURL"
            ],
            "install_type": "git-clone",
            "description": "GetFileNameFromURL is a ComfyUI custom node that extracts the filename from a URL. It can handle various URLs and is capable of handling redirects."
        },
        {
            "author": "githubYiheng",
            "title": "comfyui_kmeans_filter",
            "id": "kmeans-filter",
            "reference": "https://github.com/githubYiheng/comfyui_kmeans_filter",
            "files": [
                "https://github.com/githubYiheng/comfyui_kmeans_filter"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Apply Kmeans Filter"
        },
        {
            "author": "githubYiheng",
            "title": "ComfyUI_Change_IMAGE_BOREDER",
            "id": "change-image-border",
            "reference": "https://github.com/githubYiheng/ComfyUI_Change_IMAGE_BOREDER",
            "files": [
                "https://github.com/githubYiheng/ComfyUI_Change_IMAGE_BOREDER"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Change Image Border"
        },
        {
            "author": "githubYiheng",
            "title": "comfyui_meanshift_filter",
            "id": "meanshift-filter",
            "reference": "https://github.com/githubYiheng/comfyui_meanshift_filter",
            "files": [
                "https://github.com/githubYiheng/comfyui_meanshift_filter"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Apply Meanshift Filter"
        },
        {
            "author": "githubYiheng",
            "title": "comfyui_private_postprocessor",
            "id": "githubyiheng-private-postprocessor",
            "reference": "https://github.com/githubYiheng/comfyui_private_postprocessor",
            "files": [
                "https://github.com/githubYiheng/comfyui_private_postprocessor"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Private ImageCPostprocessor"
        },
        {
            "author": "Fihade",
            "title": "IC-Light-ComfyUI-Node",
            "reference": "https://github.com/Fihade/IC-Light-ComfyUI-Node",
            "files": [
                "https://github.com/Fihade/IC-Light-ComfyUI-Node"
            ],
            "install_type": "git-clone",
            "description": "Original repo: [a/https://mirror.ghproxy.com/https://github.com/lllyasviel/IC-Light](https://mirror.ghproxy.com/https://github.com/lllyasviel/IC-Light)\nModels: [a/https://huggingface.co/lllyasviel/ic-light/tree/main](https://huggingface.co/lllyasviel/ic-light/tree/main), [a/https://huggingface.co/digiplay/Photon_v1/tree/main](https://huggingface.co/digiplay/Photon_v1/tree/main)\nmodels go into ComfyUI/models/unet"
        },
        {
            "author": "KewkLW",
            "title": "ComfyUI-kewky_tools",
            "id": "kewky-tools",
            "reference": "https://github.com/KewkLW/ComfyUI-kewky_tools",
            "files": [
                "https://github.com/KewkLW/ComfyUI-kewky_tools"
            ],
            "install_type": "git-clone",
            "description": "text_append_node, vramdebugplus, tensordebugplus, animation_schedule_output"
        },
        {
            "author": "ITurchenko",
            "title": "ComfyUI-SizeFromArray",
            "id": "sizefromarray",
            "reference": "https://github.com/ITurchenko/ComfyUI-SizeFromArray",
            "files": [
                "https://github.com/ITurchenko/ComfyUI-SizeFromArray"
            ],
            "install_type": "git-clone",
            "description": "Nodes:SizeFromArray"
        },
        {
            "author": "Suplex",
            "title": "Suplex Misc ComfyUI Nodes",
            "id": "suplex",
            "reference": "https://github.com/saftle/uber_comfy_nodes",
            "files": [
                "https://github.com/saftle/uber_comfy_nodes"
            ],
            "install_type": "git-clone",
            "description": "Misc Nodes: ControlNet Selector Node, Load Optional ControlNet Model, Diffusers Selector, Save Image JPG No Meta, Multi Input Variable Rewrite"
        },
        {
            "author": "mephisto83",
            "title": "petty-paint-comfyui-node",
            "id": "petty-paint",
            "reference": "https://github.com/mephisto83/petty-paint-comfyui-node",
            "files": [
                "https://github.com/mephisto83/petty-paint-comfyui-node"
            ],
            "install_type": "git-clone",
            "description": "An integration between comfy ui and petty paint"
        },
        {
            "author": "fsdymy1024",
            "title": "ComfyUI_fsdymy",
            "id": "fsdymy",
            "reference": "https://github.com/fsdymy1024/ComfyUI_fsdymy",
            "files": [
                "https://github.com/fsdymy1024/ComfyUI_fsdymy"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Save Image Without Metadata"
        },
        {
            "author": "ray",
            "title": "Light Gradient for ComfyUI",
            "id": "light-gradient",
            "reference": "https://github.com/huagetai/ComfyUI_LightGradient",
            "files": [
                "https://github.com/huagetai/ComfyUI_LightGradient"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Gradient,Mask Gradient"
        },
        {
            "author": "ray",
            "title": "comfyui's gaffer(ComfyUI native implementation of IC-Light. )",
            "id": "gaffer",
            "reference": "https://github.com/huagetai/ComfyUI-Gaffer",
            "files": [
                "https://github.com/huagetai/ComfyUI-Gaffer"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Load ICLight Model,Apply ICLight,Simple Light Source,Calculate Normal Map"
        },
        {
            "author": "YFG",
            "title": "😸 YFG Comical Nodes",
            "id": "comical",
            "reference": "https://github.com/gonzalu/ComfyUI_YFG_Comical",
            "files": [
                "https://github.com/gonzalu/ComfyUI_YFG_Comical"
            ],
            "install_type": "git-clone",
            "description": "Image Historgram Generator - Outputs a set of images displaying the Histogram of the input image. Nodes: img2histograms, img2histogramsSelf"
        },
        {
            "author": "ruiqutech",
            "title": "RuiquNodes for ComfyUI",
            "id": "RuiquNodes",
            "reference": "https://github.com/ruiqutech/ComfyUI-RuiquNodes",
            "files": [
                "https://github.com/ruiqutech/ComfyUI-RuiquNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes of EvaluateMultiple1, EvaluateMultiple3...\nSupport the execution of any fragment of Python code, generating multiple outputs from multiple inputs."
        },
        {
            "author": "teward",
            "title": "ComfyUI-Helper-Nodes",
            "id": "helper-nodes",
            "reference": "https://github.com/teward/ComfyUI-Helper-Nodes",
            "files": [
                "https://github.com/teward/ComfyUI-Helper-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: HelperNodes_MultilineStringLiteral, HelperNodes_StringLiteral, HelperNodes_Steps, HelperNodes_CfgScale, HelperNodes_WidthHeight, HelperNodes_SchedulerSelector, HelperNodes_SamplerSelector, ..."
        },
        {
            "author": "fmatray",
            "title": "ComfyUI_BattlemapGrid",
            "id": "battlemap-grid",
            "reference": "https://github.com/fmatray/ComfyUI_BattlemapGrid",
            "files": [
                "https://github.com/fmatray/ComfyUI_BattlemapGrid"
            ],
            "install_type": "git-clone",
            "description": "Nodes for ComfyUI in order to generate battelmaps"
        },
        {
            "author": "christian-byrne",
            "title": "img2txt-comfyui-nodes",
            "id": "img2txt-nodes",
            "reference": "https://github.com/christian-byrne/img2txt-comfyui-nodes",
            "files": [
                "https://github.com/christian-byrne/img2txt-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Get general description or specify questions to ask about images (medium, art style, background, etc.). Supports Chinese 🇨🇳 questions via MiniCPM model."
        },
        {
            "author": "christian-byrne",
            "title": "Img2color - Extract Colors from Image",
            "id": "img2colors-comfyui-node",
            "reference": "https://github.com/christian-byrne/img2colors-comfyui-node",
            "files": [
                "https://github.com/christian-byrne/img2colors-comfyui-node"
            ],
            "install_type": "git-clone",
            "description": "Extract the most common colors from an image, up to any number. Convert colors to plain English names using various color naming systems."
        },
        {
            "author": "christian-byrne",
            "title": "Node - Size Matcher",
            "id": "sizematcher",
            "reference": "https://github.com/christian-byrne/size-match-compositing-nodes",
            "files": [
                "https://github.com/christian-byrne/size-match-compositing-nodes"
            ],
            "install_type": "git-clone",
            "description": "Match image/mask sizes"
        },
        {
            "author": "christian-byrne",
            "title": "comfyui-search-navigation",
            "reference": "https://github.com/christian-byrne/comfyui-search-navigation",
            "files": [
                "https://github.com/christian-byrne/comfyui-search-navigation"
            ],
            "install_type": "git-clone",
            "description": "Search navigation extension."
        },
        {
            "author": "christian-byrne",
            "title": "audio-separation-nodes-comfyui",
            "reference": "https://github.com/christian-byrne/audio-separation-nodes-comfyui",
            "files": [
                "https://github.com/christian-byrne/audio-separation-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Separate audio track into stems (vocals, bass, drums, other). Along with tools to recombine, tempo match, slice/crop audio."
        },
        {
            "author": "christian-byrne",
            "title": "comfyui-default-values-manager",
            "reference": "https://github.com/christian-byrne/comfyui-default-values-manager",
            "files": [
                "https://github.com/christian-byrne/comfyui-default-values-manager"
            ],
            "install_type": "git-clone",
            "description": "comfyui-default-values-manager"
        },
        {
            "author": "christian-byrne",
            "title": "youtube-dl-comfyui",
            "reference": "https://github.com/christian-byrne/youtube-dl-comfyui",
            "files": [
                "https://github.com/christian-byrne/youtube-dl-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Download youtube videos/playlists"
        },
        {
            "author": "oztrkoguz",
            "title": "ComfyUI StoryCreater",
            "id": "storycreater",
            "reference": "https://github.com/oztrkoguz/ComfyUI_StoryCreator",
            "files": [
                "https://github.com/oztrkoguz/ComfyUI_StoryCreator"
            ],
            "install_type": "git-clone",
            "description": "Nodes:story_sampler_simple, text2, kosmos2_sampler.\nI created a dataset for generating short stories [a/Short-Story](https://huggingface.co/datasets/oztrkoguz/Short-Story) and used it to fine-tune my own model using Phi-3."
        },
        {
            "author": "GraftingRayman",
            "title": "GraftingRayman",
            "id": "graftingrayman",
            "reference": "https://github.com/GraftingRayman/ComfyUI_GraftingRayman",
            "files": [
                "https://github.com/GraftingRayman/ComfyUI_GraftingRayman"
            ],
            "install_type": "git-clone",
            "description": "Image Manipulation Nodes"
        },
        {
            "author": "royceschultz",
            "title": "ComfyUI-Notifications",
            "reference": "https://github.com/royceschultz/ComfyUI-Notifications",
            "files": [
                "https://github.com/royceschultz/ComfyUI-Notifications"
            ],
            "install_type": "git-clone",
            "description": "Send notifications when a workflow completes."
        },
        {
            "author": "katalist-ai",
            "title": "comfyUI-nsfw-detection",
            "id": "nsfw-detection",
            "reference": "https://github.com/katalist-ai/comfyUI-nsfw-detection",
            "files": [
                "https://github.com/katalist-ai/comfyUI-nsfw-detection"
            ],
            "install_type": "git-clone",
            "description": "Nodes: NudenetDetector"
        },
        {
            "author": "kaanyalova",
            "title": "Extended Image Formats for ComfyUI",
            "id": "extended-image-format",
            "reference": "https://github.com/kaanyalova/ComfyUI_ExtendedImageFormats",
            "files": [
                "https://github.com/kaanyalova/ComfyUI_ExtendedImageFormats"
            ],
            "install_type": "git-clone",
            "description": "Adds a custom node for saving images in webp, jpeg, avif, jxl (no metadata) and supports loading workflows from saved images"
        },
        {
            "author": "badayvedat",
            "title": "ComfyUI-fal-Connector",
            "id": "fal",
            "reference": "https://github.com/badayvedat/ComfyUI-fal-Connector",
            "files": [
                "https://github.com/badayvedat/ComfyUI-fal-Connector"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI-fal-Connector is a tool designed to provide an integration between ComfyUI and fal. This extension allows users to execute their ComfyUI workflows directly on [a/fal.ai](https://fal.ai/). This enables users to leverage the computational power and resources provided by fal.ai for running their ComfyUI workflows."
        },
        {
            "author": "TheMistoAI",
            "title": "Anyline",
            "id": "anyline",
            "reference": "https://github.com/TheMistoAI/ComfyUI-Anyline",
            "files": [
                "https://github.com/TheMistoAI/ComfyUI-Anyline"
            ],
            "install_type": "git-clone",
            "description": "A Fast, Accurate, and Detailed Line Detection Preprocessor.\nAnyline is a ControlNet line preprocessor that accurately extracts object edges, image details, and textual content from most images. Users can input any type of image to quickly obtain line drawings with clear edges, sufficient detail preservation, and high fidelity text, which are then used as input for conditional generation in Stable Diffusion."
        },
        {
            "author": "mbrostami",
            "title": "ComfyUI-TITrain",
            "id": "titrain",
            "reference": "https://github.com/mbrostami/ComfyUI-TITrain",
            "files": [
                "https://github.com/mbrostami/ComfyUI-TITrain"
            ],
            "install_type": "git-clone",
            "description": "Nodes:TextualInversionTrainingSDXL, TextualInversionTraining"
        },
        {
            "author": "ArcherFMY",
            "title": "Diffusion360_ComfyUI",
            "id": "diffusion360",
            "reference": "https://github.com/ArcherFMY/Diffusion360_ComfyUI",
            "files": [
                "https://github.com/ArcherFMY/Diffusion360_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI plugin of [a/SD-T2I-360PanoImage](https://mirror.ghproxy.com/https://github.com/ArcherFMY/SD-T2I-360PanoImage).\nbase t2i-pipeline for generating 512*1024 panorama image from text input"
        },
        {
            "author": "Makeezi",
            "title": "ComfyUI-promptLAB",
            "id": "promptlab",
            "reference": "https://github.com/Makeezi/ComfyUI-promptLAB",
            "files": [
                "https://github.com/Makeezi/ComfyUI-promptLAB"
            ],
            "install_type": "git-clone",
            "description": "connection nodes for api requests, fully supports promptLAB"
        },
        {
            "author": "portu-sim",
            "title": "comfyui_bmab",
            "id": "bmab",
            "reference": "https://github.com/portu-sim/comfyui_bmab",
            "files": [
                "https://github.com/portu-sim/comfyui_bmab"
            ],
            "install_type": "git-clone",
            "description": "BMAB for ComfyUI. BMAB is an custom nodes of ComfyUI and has the function of post-processing the generated image according to settings. If necessary, you can find and redraw people, faces, and hands, or perform functions such as resize, resample, and add noise. You can composite two images or perform the Upscale function."
        },
        {
            "author": "griptape-ai",
            "title": "ComfyUI Griptape Nodes",
            "id": "griptape",
            "reference": "https://github.com/griptape-ai/ComfyUI-Griptape",
            "files": [
                "https://github.com/griptape-ai/ComfyUI-Griptape"
            ],
            "install_type": "git-clone",
            "description": "This repo creates a series of nodes that enable you to utilize the [a/Griptape Python Framework](https://mirror.ghproxy.com/https://github.com/griptape-ai/griptape/) with ComfyUI, integrating AI into your workflow. This repo creates a series of nodes that enable you to utilize the Griptape Python Framework with ComfyUI, integrating AI into your workflow."
        },
        {
            "author": "cavinHuang",
            "title": "comfyui-nodes-docs",
            "id": "nodedocs",
            "reference": "https://github.com/CavinHuang/comfyui-nodes-docs",
            "files": [
                "https://github.com/CavinHuang/comfyui-nodes-docs"
            ],
            "install_type": "git-clone",
            "description": "This is a plugin for displaying documentation for each comfyui node. "
        },
        {
            "author": "icesun963",
            "title": "HFDownLoad Node for ComfyUI",
            "id": "HFDownLoad-ic",
            "reference": "https://github.com/icesun963/ComfyUI_HFDownLoad",
            "files": [
                "https://github.com/icesun963/ComfyUI_HFDownLoad"
            ],
            "install_type": "git-clone",
            "description": "Download the model from huggingface and put it in any directory."
        },
        {
            "author": "conquestace",
            "title": "Image Uploader",
            "id": "image-uploader",
            "reference": "https://github.com/conquestace/ComfyUI-ImageUploader",
            "files": [
                "https://github.com/conquestace/ComfyUI-ImageUploader"
            ],
            "install_type": "git-clone",
            "description": "Upload images automatically to image hosting sites."
        },
        {
            "author": "chandlergis",
            "title": "ComfyUI-IMG_Query",
            "id": "img-query",
            "reference": "https://github.com/chandlergis/ComfyUI-IMG_Query",
            "files": [
                "https://github.com/chandlergis/ComfyUI-IMG_Query"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageRequestNode"
        },
        {
            "author": "Isaac Emesowum",
            "title": "Isaac's Nodes",
            "id": "isaac",
            "reference": "https://github.com/iemesowum/ComfyUI_IsaacNodes",
            "files": [
                "https://github.com/iemesowum/ComfyUI_IsaacNodes"
            ],
            "install_type": "git-clone",
            "description": "This extension offers automatic drums extraction from audio files, as well as a few helper nodes to support my audio synchronization AnimateDiff workflows."
        },
        {
            "author": "fexploit",
            "title": "ComfyUI-AutoTrimBG",
            "id": "autotrimbg",
            "reference": "https://github.com/fexploit/ComfyUI-AutoTrimBG",
            "files": [
                "https://github.com/fexploit/ComfyUI-AutoTrimBG"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-AutoCropBgTrim is a powerful tool designed to automatically clean up the background of your images. This tool trims unnecessary spaces and pixels, leaving only the main subject of the image. It generates both a mask and an image output, making it easy to focus on the essential elements. Perfect for enhancing your photos and preparing them for professional use."
        },
        {
            "author": "fexploit",
            "title": "ComfyUI-AutoLabel",
            "id": "autolabel",
            "reference": "https://github.com/fexploit/ComfyUI-AutoLabel",
            "files": [
                "https://github.com/fexploit/ComfyUI-AutoLabel"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-AutoLabel is a custom node for ComfyUI that uses BLIP (Bootstrapping Language-Image Pre-training) to generate detailed descriptions of the main object in an image. This node leverages the power of BLIP to provide accurate and context-aware captions for images."
        },
        {
            "author": "fexploit",
            "title": "ComfyUI-Classifier",
            "id": "classifier",
            "reference": "https://github.com/fexploit/ComfyUI-Classifier",
            "files": [
                "https://github.com/fexploit/ComfyUI-Classifier"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Classifier is a custom node for ComfyUI that uses a zero-shot classification model to classify text inputs based on a set of candidate labels. This node leverages the power of Hugging Face Transformers to provide accurate and flexible text classification."
        },
        {
            "author": "linshier",
            "title": "comfyui-remote-tools",
            "id": "remote-tools",
            "reference": "https://github.com/linshier/comfyui-remote-tools",
            "files": [
                "https://github.com/linshier/comfyui-remote-tools"
            ],
            "install_type": "git-clone",
            "description": "Node:SendBase64ToRemote. To connect to another ComfyUI server."
        },
        {
            "author": "Fantaxico",
            "title": "ComfyUI-GCP-Storage",
            "id": "gcp-storage",
            "reference": "https://github.com/Fantaxico/ComfyUI-GCP-Storage",
            "files": [
                "https://github.com/Fantaxico/ComfyUI-GCP-Storage"
            ],
            "install_type": "git-clone",
            "description": "Node:GCP Storage Node. Support google-cloud-storage."
        },
        {
            "author": "daniabib",
            "title": "ComfyUI ProPainter Nodes",
            "id": "propainter",
            "reference": "https://github.com/daniabib/ComfyUI_ProPainter_Nodes",
            "files": [
                "https://github.com/daniabib/ComfyUI_ProPainter_Nodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom node implementation of [a/ProPainter](https://mirror.ghproxy.com/https://github.com/sczhou/ProPainter) framework for video inpainting."
        },
        {
            "author": "iFREEGROUP",
            "title": "comfyui-undistort",
            "id": "undistort",
            "reference": "https://github.com/iFREEGROUP/comfyui-undistort",
            "files": [
                "https://github.com/iFREEGROUP/comfyui-undistort"
            ],
            "install_type": "git-clone",
            "description": "Node:Load Checkerboard Images for Calibrate Camera, Matrix and distortion coefficient to text, Undistort"
        },
        {
            "author": "Auttasak-L",
            "title": "ComfyUI-ImageCropper",
            "id": "imagecropper",
            "reference": "https://github.com/Auttasak-L/ComfyUI-ImageCropper",
            "files": [
                "https://github.com/Auttasak-L/ComfyUI-ImageCropper"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image cropping tool"
        },
        {
            "author": "muzi12888",
            "title": "PoseKeypoint Mask",
            "id": "posekeypoint-mask",
            "reference": "https://github.com/muzi12888/ComfyUI-PoseKeypoint-Mask",
            "files": [
                "https://github.com/muzi12888/ComfyUI-PoseKeypoint-Mask"
            ],
            "install_type": "git-clone",
            "description": "Convert PoseKeypoint to mask, please refer to the example workflow for usage instructions."
        },
        {
            "author": "muzi12888",
            "title": "m9-prompts-comfyui",
            "id": "m9-prompts-comfyui",
            "reference": "https://github.com/MarcusNyne/m9-prompts-comfyui",
            "files": [
                "https://github.com/MarcusNyne/m9-prompts-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes for modifying a prompt to create prompt variations.\nScramblePrompts [m9]: Reorder prompts, remove prompts, modify weights\nTweakWeights [m9]: Modify the weights of prompts matching keywords"
        },
        {
            "author": "xuhongming251",
            "title": "ComfyUI-GPEN",
            "id": "gpen",
            "reference": "https://github.com/xuhongming251/ComfyUI-GPEN",
            "files": [
                "https://github.com/xuhongming251/ComfyUI-GPEN"
            ],
            "install_type": "git-clone",
            "description": "Nodes:FaceEnhancement. Based on modelscope pipeline."
        },
        {
            "author": "xuhongming251",
            "title": "ComfyUI-MuseTalkUtils",
            "id": "musetalk-utils",
            "reference": "https://github.com/xuhongming251/ComfyUI-MuseTalkUtils",
            "files": [
                "https://github.com/xuhongming251/ComfyUI-MuseTalkUtils"
            ],
            "install_type": "git-clone",
            "description": "MuseTalk ComfyUI Preprocess and Postprocess Nodes"
        },
        {
            "author": "Thomas Ward",
            "title": "TW-CUI-Util",
            "id": "tw-cui-util",
            "reference": "https://github.com/TW-CUI/TW-CUI-Util",
            "files": [
                "https://github.com/TW-CUI/TW-CUI-Util"
            ],
            "install_type": "git-clone",
            "description": "A collection of custom nodes to help with saving images, providing generation parameters, static literal nodes, and other useful nodes."
        },
        {
            "author": "lks-ai",
            "title": "ComfyUI AnyNode: Any Node you ask for",
            "id": "anynode",
            "reference": "https://github.com/lks-ai/anynode",
            "files": [
                "https://github.com/lks-ai/anynode"
            ],
            "install_type": "git-clone",
            "description": "Nodes: AnyNode. Nodes that can be anything you ask. Auto-Generate functional nodes using LLMs. Create impossible workflows. API Compatibility: (OpenAI, LocalLLMs, Gemini)."
        },
        {
            "author": "lks-ai",
            "title": "ComfyUI Stable Audio Open 1.0 Sampler",
            "id": "stableaudiosampler",
            "reference": "https://github.com/lks-ai/ComfyUI-StableAudioSampler",
            "files": [
                "https://github.com/lks-ai/ComfyUI-StableAudioSampler"
            ],
            "install_type": "git-clone",
            "description": "Nodes: StableAudioSampler. Wraps the new Stable Audio Open Model in the sampler that dropped Jun 5th. See Github for Features"
        },
        {
            "author": "SayanoAI",
            "title": "Comfy-RVC",
            "id": "sayano-rvc",
            "reference": "https://github.com/SayanoAI/Comfy-RVC",
            "files": [
                "https://github.com/SayanoAI/Comfy-RVC"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom nodes for RVC related inference and image generation"
        },
        {
            "author": "nirbhay-faaya",
            "title": "ImgProcessing_ComfyUI",
            "id": "imgprocessing",
            "reference": "https://github.com/nirbhay-faaya/ImgProcessing_ComfyUI",
            "files": [
                "https://github.com/nirbhay-faaya/ImgProcessing_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Custom Image processing ComfyUI Nodes"
        },
        {
            "author": "larsupb",
            "title": "LoRA Power-Merger ComfyUI",
            "id": "lora-powermerger",
            "reference": "https://github.com/larsupb/LoRA-Merger-ComfyUI",
            "files": [
                "https://github.com/larsupb/LoRA-Merger-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "An extension for merging LoRAs. Offers a wide range of LoRA merge techniques (including dare) and XY plots. XY plots require efficiency nodes."
        },
        {
            "author": "Mason-McGough",
            "title": "Mosaica",
            "id": "mosaica",
            "reference": "https://github.com/Mason-McGough/ComfyUI-Mosaica",
            "files": [
                "https://github.com/Mason-McGough/ComfyUI-Mosaica"
            ],
            "install_type": "git-clone",
            "description": "Create colorful mosaic images in ComfyUI by computing label images and applying lookup tables."
        },
        {
            "author": "cuongloveit",
            "title": "comfy_http_request",
            "reference": "https://github.com/cuongloveit/comfy_http_request",
            "files": [
                "https://github.com/cuongloveit/comfy_http_request"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Send Http Request. You can use this node to save full size images through the websocket."
        },
        {
            "author": "Ron-Digital",
            "title": "ComfyUI-SceneGenerator",
            "id": "scenegenerator",
            "reference": "https://github.com/Ron-Digital/ComfyUI-SceneGenerator",
            "files": [
                "https://github.com/Ron-Digital/ComfyUI-SceneGenerator"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-SceneGenerator is a ComfyUI plugin used to generate scene preview photos from JSON files. This plugin creates scenes based on the provided JSON configuration and produces two different image outputs: one containing only the products and the other containing both the products and the props."
        },
        {
            "author": "xliry",
            "title": "ComfyUI_SendDiscord",
            "id": "senddiscord",
            "reference": "https://github.com/xliry/ComfyUI_SendDiscord",
            "files": [
                "https://github.com/xliry/ComfyUI_SendDiscord"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Send Video to Discord"
        },
        {
            "author": "xliry",
            "title": "color2rgb",
            "reference": "https://github.com/vxinhao/color2rgb",
            "files": [
                "https://github.com/vxinhao/color2rgb/raw/main/color2rgb.py"
            ],
            "install_type": "copy",
            "description": "Nodes:color2RGB"
        },
        {
            "author": "moyi7712",
            "title": "ComfyUI_Seamless_Patten",
            "id": "seamless-pattern",
            "reference": "https://github.com/moyi7712/ComfyUI_Seamless_Patten",
            "files": [
                "https://github.com/moyi7712/ComfyUI_Seamless_Patten"
            ],
            "install_type": "git-clone",
            "description": "It make any text2image create seamless patten"
        },
        {
            "author": "nirex0",
            "title": "ComfyUI_pytorch_openpose",
            "id": "pytorch-openpose",
            "reference": "https://github.com/nirex0/ComfyUI_pytorch_openpose",
            "files": [
                "https://github.com/nirex0/ComfyUI_pytorch_openpose"
            ],
            "install_type": "git-clone",
            "description": "All Credits go to the original Repo: [a/Hzzone/pytorch-openpose](https://mirror.ghproxy.com/https://github.com/Hzzone/pytorch-openpose)."
        },
        {
            "author": "AshMartian",
            "title": "Dir Gir",
            "id": "dir-gir",
            "reference": "https://github.com/AshMartian/ComfyUI-DirGir",
            "files": [
                "https://github.com/AshMartian/ComfyUI-DirGir"
            ],
            "install_type": "git-clone",
            "description": "A collection of ComfyUI directory automation utility nodes. Directory Get It Right adds a GUI directory browser, and smart directory loop/iteration node that supports regex and file extension filtering."
        },
        {
            "author": "SozeInc",
            "title": "ComfyUI-Mobile",
            "id": "comfyui-mobile",
            "reference": "https://github.com/SozeInc/ComfyUI-Mobile",
            "files": [
                "https://github.com/SozeInc/ComfyUI-Mobile"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Ultimate Concat (Mobile), Send Notification (Mobile), Settings Launcher (Mobile), Settings Launcher Outputs (Mobile)"
        },
        {
            "author": "audioscavenger",
            "title": "ComfyUI Thumbnails",
            "id": "thumbnails",
            "reference": "https://github.com/audioscavenger/ComfyUI-Thumbnails",
            "files": [
                "https://github.com/audioscavenger/ComfyUI-Thumbnails"
            ],
            "install_type": "git-clone",
            "description": "Load Image thumbnails for ComfyUI"
        },
        {
            "author": "goktug",
            "title": "Save Image Plus for ComfyUI",
            "id": "saveimage-plus",
            "reference": "https://github.com/Goktug/comfyui-saveimage-plus",
            "files": [
                "https://github.com/Goktug/comfyui-saveimage-plus"
            ],
            "install_type": "git-clone",
            "description": "Save Image Plus is a custom node for ComfyUI that allows you to save images in JPEG and WEBP formats with optional metadata embedding."
        },
        {
            "author": "wujm424606",
            "title": "ComfyUi-Ollama-YN",
            "id": "ollama-YN",
            "reference": "https://github.com/wujm424606/ComfyUi-Ollama-YN",
            "files": [
                "https://github.com/wujm424606/ComfyUi-Ollama-YN"
            ],
            "install_type": "git-clone",
            "description": "Custom ComfyUI Nodes for interacting with [a/Ollama](https://ollama.com/) using the [a/ollama python client](https://mirror.ghproxy.com/https://github.com/ollama/ollama-python).\n Meanwhile it will provide better prompt descriptor for stable diffusion."
        },
        {
            "author": "tmagara",
            "title": "ComfyUI-Prediction-Boost",
            "id": "prediction-boost",
            "reference": "https://github.com/tmagara/ComfyUI-Prediction-Boost",
            "files": [
                "https://github.com/tmagara/ComfyUI-Prediction-Boost"
            ],
            "install_type": "git-clone",
            "description": "prediction boost custom node for ComfyUI"
        },
        {
            "author": "chesnokovivan",
            "title": "ComfyUI-Novakid",
            "id": "novakid",
            "reference": "https://github.com/chesnokovivan/ComfyUI-Novakid",
            "files": [
                "https://github.com/chesnokovivan/ComfyUI-Novakid"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI: Novakid. A node."
        },
        {
            "author": "Jin Liu",
            "title": "ComfyUI-Photopea",
            "id": "photopea",
            "reference": "https://github.com/coolzilj/ComfyUI-Photopea",
            "files": [
                "https://github.com/coolzilj/ComfyUI-Photopea"
            ],
            "install_type": "git-clone",
            "description": "Edit images in the Photopea editor directly within ComfyUI."
        },
        {
            "author": "bitaffinity",
            "title": "ComfyUI_HF_Inference",
            "id": "hf-inference",
            "reference": "https://github.com/bitaffinity/ComfyUI_HF_Inference",
            "files": [
                "https://github.com/bitaffinity/ComfyUI_HF_Inference"
            ],
            "install_type": "git-clone",
            "description": "Unofficial support for Hugging Face's hosted inference."
        },
        {
            "author": "claussteinmassl",
            "title": "CS Transform Node for ComfyUI",
            "id": "cs-transform",
            "reference": "https://github.com/claussteinmassl/ComfyUI-CS-CustomNodes",
            "files": [
                "https://github.com/claussteinmassl/ComfyUI-CS-CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "The CS Transform node is a custom node for ComfyUI that applies a series of transformations to an input image and mask. The transformations include scaling, rotation, and translation, all centered around a specified pivot point. The node ensures that the transformed image is properly accommodated within a canvas, which can be expanded if needed."
        },
        {
            "author": "MariusKM",
            "title": "ComfyUI-BadmanNodes",
            "id": "badman",
            "reference": "https://github.com/MariusKM/ComfyUI-BadmanNodes",
            "files": [
                "https://github.com/MariusKM/ComfyUI-BadmanNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Badman_Blend, Badman_HexGenerator, Badman_String, Badman_Concat_String, Badman_Print, BadmanIO, BadmanIntUtil"
        },
        {
            "author": "TMElyralab",
            "title": "Comfyui-MusePose",
            "id": "musepose",
            "reference": "https://github.com/TMElyralab/Comfyui-MusePose",
            "files": [
                "https://github.com/TMElyralab/Comfyui-MusePose"
            ],
            "install_type": "git-clone",
            "description": "[a/MusePose](https://mirror.ghproxy.com/https://github.com/TMElyralab/MusePose) is an image-to-video generation framework for virtual human under control signal such as pose.\nNOTE: You need to download weigths manually from: [a/https://huggingface.co/TMElyralab/MusePose](https://huggingface.co/TMElyralab/MusePose).[w/The repository name has changed. If you are not receiving updates, please delete the existing node and reinstall it.]"
        },
        {
            "author": "PnthrLeo",
            "title": "comfyUI-image-search",
            "id": "image-search",
            "reference": "https://github.com/PnthrLeo/comfyUI-image-search",
            "files": [
                "https://github.com/PnthrLeo/comfyUI-image-search"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Close Images Searcher"
        },
        {
            "author": "l20richo",
            "title": "ComfyUI-Azure-Blob-Storage",
            "id": "azure-blob-storage",
            "reference": "https://github.com/l20richo/ComfyUI-Azure-Blob-Storage",
            "files": [
                "https://github.com/l20richo/ComfyUI-Azure-Blob-Storage"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Azure-Blob-Storage seamlessly integrates with [a/Azure Blob Storage](https://azure.microsoft.com/en-us/products/storage/blobs/) in ComfyUI. This open-source project provides custom nodes for effortless loading and saving of images, videos, and checkpoint models directly from Azure blob containers within the ComfyUI graph interface."
        },
        {
            "author": "AARG-FAN",
            "title": "Image-vector-for-ComfyUI",
            "id": "image-vector",
            "reference": "https://github.com/AARG-FAN/Image-Vector-for-ComfyUI",
            "files": [
                "https://github.com/AARG-FAN/Image-Vector-for-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a wrap-up of ComfyUI nodes for converting pixels to raster, sent out to [a/Vtracer](https://mirror.ghproxy.com/https://github.com/visioncortex/vtracer)!"
        },
        {
            "author": "Smirnov75",
            "title": "ComfyUI-mxToolkit",
            "id": "mxtoolkit",
            "reference": "https://github.com/Smirnov75/ComfyUI-mxToolkit",
            "files": [
                "https://github.com/Smirnov75/ComfyUI-mxToolkit"
            ],
            "install_type": "git-clone",
            "description": "A set of useful nodes for convenient use of ComfyUI, including: Seed randomization before the generation process starts, with saving of the last used values and the ability to automatically interrupt the current generation; A function to pause the generation process; Slider nodes for convenient control of input parameters; An alternative version of the standard Reroute node."
        },
        {
            "author": "humgate",
            "title": "simplecomfy",
            "reference": "https://github.com/humgate/simplecomfy",
            "files": [
                "https://github.com/humgate/simplecomfy"
            ],
            "install_type": "git-clone",
            "description": "Simple JS application based on ComfyUI which takes prompt and style picture from user and runs hardcoded workflow inference returning generated image to user."
        },
        {
            "author": "vanche1212",
            "title": "ZMG PLUGIN",
            "id": "zmg",
            "reference": "https://github.com/vanche1212/ComfyUI-ZMG-Nodes",
            "files": [
                "https://github.com/vanche1212/ComfyUI-ZMG-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ApiRequestNode, LoadVideoNode, JsonParserNode, OllamaRequestNode, OldPhotoColorizationNode."
        },
        {
            "author": "hben35096",
            "title": "ComfyUI-ToolBox",
            "id": "hben-toolbox",
            "reference": "https://github.com/hben35096/ComfyUI-ToolBox",
            "files": [
                "https://github.com/hben35096/ComfyUI-ToolBox"
            ],
            "install_type": "git-clone",
            "description": "Nodes:commonly_node."
        },
        {
            "author": "hben35096",
            "title": "ComfyUI-ReplenishNodes",
            "reference": "https://github.com/hben35096/ComfyUI-ReplenishNodes",
            "files": [
                "https://github.com/hben35096/ComfyUI-ReplenishNodes"
            ],
            "install_type": "git-clone",
            "description": "NODES:Batch Image Blend, Mask Levels Adjust, Get Batch Count, Load Lora Name, Load Sampler Name, Load Scheduler Name, Load Ckpt Name....\nThe nodes in this repository are only used as secondary nodes."
        },
        {
            "author": "tiankuan93",
            "title": "V-Express: Conditional Dropout for Progressive Training of Portrait Video Generation",
            "id": "v-express",
            "reference": "https://github.com/tiankuan93/ComfyUI-V-Express",
            "files": [
                "https://github.com/tiankuan93/ComfyUI-V-Express"
            ],
            "install_type": "git-clone",
            "description": "[Original] In the field of portrait video generation, the use of single images to generate portrait videos has become increasingly prevalent. A common approach involves leveraging generative models to enhance adapters for controlled generation. However, control signals can vary in strength, including text, audio, image reference, pose, depth map, etc. Among these, weaker conditions often struggle to be effective due to interference from stronger conditions, posing a challenge in balancing these conditions. In our work on portrait video generation, we identified audio signals as particularly weak, often overshadowed by stronger signals such as pose and original image. However, direct training with weak signals often leads to difficulties in convergence. To address this, we propose V-Express, a simple method that balances different control signals through a series of progressive drop operations. Our method gradually enables effective control by weak conditions, thereby achieving generation capabilities that simultaneously take into account pose, input image, and audio.\nNOTE: You need to downdload [a/model_ckpts](https://huggingface.co/tk93/V-Express/tree/main) manually."
        },
        {
            "author": "CMonk",
            "title": "Stable Projectorz Bridge",
            "id": "projectorz",
            "reference": "https://github.com/tianlang0704/ComfyUI-StableProjectorzBridge",
            "files": [
                "https://github.com/tianlang0704/ComfyUI-StableProjectorzBridge"
            ],
            "install_type": "git-clone",
            "description": "This custom nodes enables Stable Projectorz to work with ComfyUI Directly."
        },
        {
            "author": "Scorpinaus",
            "title": "ComfyUI-DiffusersLoader",
            "id": "comfyui-diffusersloader",
            "reference": "https://github.com/Scorpinaus/ComfyUI-DiffusersLoader",
            "files": [
                "https://github.com/Scorpinaus/ComfyUI-DiffusersLoader"
            ],
            "install_type": "git-clone",
            "description": "This node pack allows loading of SD checkpoints that uses diffusers format in comfyUI."
        },
        {
            "author": "chakib-belgaid",
            "title": "ComfyUI Style Plugin",
            "id": "style-plugin",
            "reference": "https://github.com/chakib-belgaid/Comfyui_Prompt_styler",
            "files": [
                "https://github.com/chakib-belgaid/Comfyui_Prompt_styler"
            ],
            "install_type": "git-clone",
            "description": "This is a simple plugin for ComfyUI that allows you to import A1111 CSV styles into ComfyUI prompts."
        },
        {
            "author": "chakib-belgaid",
            "title": "ComfyUI-autosize",
            "id": "autosize",
            "reference": "https://github.com/chakib-belgaid/ComfyUI-autosize",
            "files": [
                "https://github.com/chakib-belgaid/ComfyUI-autosize"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI utility plugin designed to optimize the latent space for generating high-quality results. It approximates the closest size model for better generation results."
        },
        {
            "author": "ThereforeGames",
            "title": "ComfyUI-Unprompted",
            "id": "unprompted",
            "reference": "https://github.com/ThereforeGames/ComfyUI-Unprompted",
            "files": [
                "https://github.com/ThereforeGames/ComfyUI-Unprompted"
            ],
            "install_type": "git-clone",
            "description": "A node that processes input text with the [a/Unprompted templating language](https://mirror.ghproxy.com/https://github.com/ThereforeGames/unprompted)."
        },
        {
            "author": "Tool Of North america",
            "title": "Easy automatic (square) image cropper using Yolo",
            "id": "tooldigital",
            "reference": "https://github.com/tooldigital/ComfyUI-Yolo-Cropper",
            "files": [
                "https://github.com/tooldigital/ComfyUI-Yolo-Cropper"
            ],
            "install_type": "git-clone",
            "description": "A very simple and easy to use node to automaticaaly create (square) image crops and masks using YoloV8. This can be very useful when using controlnet and ip adapters"
        },
        {
            "author": "luandev",
            "title": "ComfyUI CrewAI",
            "id": "crewai",
            "reference": "https://github.com/luandev/ComfyUI-CrewAI",
            "files": [
                "https://github.com/luandev/ComfyUI-CrewAI"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-CrewAI aims to integrate Crew AI's multi-agent collaboration framework into the ComfyUI environment. By combining the strengths of Crew AI's role-based, collaborative AI agent system with ComfyUI's intuitive interface, we will create a robust platform for managing and executing complex AI tasks seamlessly"
        },
        {
            "author": "chandlergis",
            "title": "ComfyUI_EmojiOverlay",
            "id": "emoji-overlay",
            "reference": "https://github.com/chandlergis/ComfyUI_EmojiOverlay",
            "files": [
                "https://github.com/chandlergis/ComfyUI_EmojiOverlay"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Emoji Overlay"
        },
        {
            "author": "risunobushi",
            "title": "comfyUI_FrequencySeparation_RGB-HSV",
            "id": "freq-sep",
            "reference": "https://github.com/risunobushi/comfyUI_FrequencySeparation_RGB-HSV",
            "files": [
                "https://github.com/risunobushi/comfyUI_FrequencySeparation_RGB-HSV"
            ],
            "install_type": "git-clone",
            "description": "A collection of simple nodes for Frequency Separation / Frequency Recombine with RGB and HSV methods"
        },
        {
            "author": "zohac",
            "title": "ComfyUI_ZC_DrawShape",
            "id": "drawshape",
            "reference": "https://github.com/zohac/ComfyUI_ZC_DrawShape",
            "files": [
                "https://github.com/zohac/ComfyUI_ZC_DrawShape"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ZC DrawShape Node"
        },
        {
            "author": "DataCTE",
            "title": "Prompt Injection Node for ComfyUI",
            "id": "prompt-injection",
            "reference": "https://github.com/DataCTE/prompt_injection",
            "files": [
                "https://github.com/DataCTE/prompt_injection"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI allows you to inject specific prompts at specific blocks of the Stable Diffusion UNet, providing fine-grained control over the generated image. It is based on the concept that the content/subject understanding of the model is primarily contained within the MID0 and MID1 blocks, as demonstrated in the B-Lora (Content Style implicit separation) paper. Features.\nInject different prompts into specific UNet blocks Three different node variations for flexible workflow integration Customize the learning rate of specific blocks to focus on content, lighting, style, or other aspects Potential for developing a 'Mix of Experts' approach by swapping blocks on-the-fly based on prompt content"
        },
        {
            "author": "FrankChieng",
            "title": "ComfyUI_llm_easyanimiate",
            "id": "llm-easyanimate",
            "nodename_pattern": "^FrankChiengEasyAnimate",
            "reference": "https://github.com/frankchieng/ComfyUI_llm_easyanimiate",
            "files": [
                "https://github.com/frankchieng/ComfyUI_llm_easyanimiate"
            ],
            "install_type": "git-clone",
            "description": "implementation easyanimate with llama3-8b-6bit instruction LLM generation prompt help"
        },
        {
            "author": "nuanarchy",
            "title": "ComfyUI-NuA-FlashFace",
            "id": "nua-flashface",
            "reference": "https://github.com/nuanarchy/ComfyUI-NuA-FlashFace",
            "files": [
                "https://github.com/nuanarchy/ComfyUI-NuA-FlashFace"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation of [a/FlashFace: Human Image Personalization with High-fidelity Identity Preservation](https://mirror.ghproxy.com/https://github.com/ali-vilab/FlashFace)\nNOTE: You need to downalod models manually."
        },
        {
            "author": "nuanarchy",
            "title": "ComfyUI-NuA-BIRD",
            "id": "nua-bird",
            "reference": "https://github.com/nuanarchy/ComfyUI-NuA-BIRD",
            "files": [
                "https://github.com/nuanarchy/ComfyUI-NuA-BIRD"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation of '[a/Blind Image Restoration via Fast Diffusion Inversion](https://mirror.ghproxy.com/https://github.com/hamadichihaoui/BIRD)' Original [a/article](https://arxiv.org/abs/2405.19572)"
        },
        {
            "author": "denfrost",
            "title": "Den_ComfyUI_Workflows",
            "id": "den",
            "reference": "https://github.com/denfrost/Den_ComfyUI_Workflow",
            "files": [
                "https://github.com/denfrost/Den_ComfyUI_Workflow"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes make easy Advanced Workflows. Focus on Image/Video and ControlNet efficiency and performances. Manipulation of Latent Space, Automatic pipeline with a bit efforts."
        },
        {
            "author": "marduk191",
            "title": "marduk191 workflow settings",
            "id": "marnodes",
            "reference": "https://github.com/marduk191/comfyui-marnodes",
            "files": [
                "https://github.com/marduk191/comfyui-marnodes"
            ],
            "install_type": "git-clone",
            "description": "A node to set workflow settings."
        },
        {
            "author": "marduk191",
            "title": "Flux Prompt Enhance Node for ComfyUI",
            "id": "fluxpromptenhancer",
            "reference": "https://github.com/marduk191/ComfyUI-Fluxpromptenhancer",
            "files": [
                "https://github.com/marduk191/ComfyUI-Fluxpromptenhancer"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI integrates the Flux-Prompt-Enhance model, allowing you to enhance your prompts directly within your ComfyUI workflows."
        },
        {
            "author": "haohaocreates",
            "title": "ComfyUI-HH-Image-Selector",
            "id": "hh-image-selector",
            "reference": "https://github.com/haohaocreates/ComfyUI-HH-Image-Selector",
            "files": [
                "https://github.com/haohaocreates/ComfyUI-HH-Image-Selector"
            ],
            "install_type": "git-clone",
            "description": "comfy ui custom node that returns an image from a batch based on selected criteria such as RGB value, brightness, etc (credits to chris goringe's custom nodes tutorial )."
        },
        {
            "author": "exdysa",
            "title": "comfyui-selector",
            "reference": "https://github.com/exdysa/comfyui-selector",
            "files": [
                "https://github.com/exdysa/comfyui-selector"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Selector. Quick and dirty parameter generator node for ComfyUI."
        },
        {
            "author": "Jin Liu",
            "title": "ComfyUI-LJNodes",
            "id": "ComfyUI-LJNodes",
            "reference": "https://github.com/coolzilj/ComfyUI-LJNodes",
            "files": [
                "https://github.com/coolzilj/ComfyUI-LJNodes"
            ],
            "install_type": "git-clone",
            "description": "A variety of custom nodes to enhance ComfyUI for a buttery smooth experience."
        },
        {
            "author": "GavChap",
            "title": "ComfyUI-SD3LatentSelectRes",
            "id": "sd3latent-select-res",
            "reference": "https://github.com/GavChap/ComfyUI-SD3LatentSelectRes",
            "files": [
                "https://github.com/GavChap/ComfyUI-SD3LatentSelectRes"
            ],
            "install_type": "git-clone",
            "description": "You'll get a new node called SD3 Latent Select Resolution, you can pick the x and y sizes from a list."
        },
        {
            "author": "BenNarum",
            "title": "SigmaWaveFormNodes",
            "id": "sigmawaveform",
            "reference": "https://github.com/BenNarum/SigmaWaveFormNode",
            "files": [
                "https://github.com/BenNarum/SigmaWaveFormNode"
            ],
            "install_type": "git-clone",
            "description": "A set of tools for generating and altering sigmas in ComfyUI."
        },
        {
            "author": "shobhitic",
            "title": "PlusMinusTextClip - Single node for Positive and Negative Prompts",
            "id": "plusminustextclip",
            "reference": "https://github.com/shobhitic/ComfyUI-PlusMinusTextClip",
            "files": [
                "https://github.com/shobhitic/ComfyUI-PlusMinusTextClip"
            ],
            "install_type": "git-clone",
            "description": "This adds a node that has both the positive and negative prompts as input in one node. You can just add one node and be done with both Positive and Negative prompts, in place of adding two different nodes for them."
        },
        {
            "author": "Late Night Labs",
            "title": "LNL Frame Selector",
            "id": "lnlframeselector",
            "reference": "https://github.com/latenightlabs/ComfyUI-LNL",
            "files": [
                "https://github.com/latenightlabs/ComfyUI-LNL"
            ],
            "install_type": "git-clone",
            "description": "Frame Selector & Sequence Selection Node for ComfyUI."
        },
        {
            "author": "Michael Standen",
            "title": "Ollama Prompt Encode",
            "id": "ollamapromptencode",
            "reference": "https://github.com/ScreamingHawk/comfyui-ollama-prompt-encode",
            "files": [
                "https://github.com/ScreamingHawk/comfyui-ollama-prompt-encode"
            ],
            "install_type": "git-clone",
            "description": "A prompt generator and CLIP encoder using AI provided by Ollama."
        },
        {
            "author": "NvidiaGameWorksAdmin",
            "title": "ComfyUI-RTX-Remix",
            "id": "comfyui-rtx-remix",
            "reference": "https://github.com/NVIDIAGameWorks/ComfyUI-RTX-Remix",
            "files": [
                "https://github.com/NVIDIAGameWorks/ComfyUI-RTX-Remix"
            ],
            "install_type": "git-clone",
            "description": "Use ComfyUI with RTX Remix to remaster classic games [a/https://mirror.ghproxy.com/https://github.com/NVIDIAGameWorks/rtx-remix](https://mirror.ghproxy.com/https://github.com/NVIDIAGameWorks/rtx-remix)"
        },
        {
            "author": "toxicwind",
            "title": "TTools for ComfyUI",
            "id": "ttools",
            "reference": "https://github.com/toxicwind/ComfyUI-TTools",
            "files": [
                "https://github.com/toxicwind/ComfyUI-TTools"
            ],
            "install_type": "git-clone",
            "description": "Text Randomization and Formatting, JSON Extraction and Processing, SD3 Resolution Solver"
        },
        {
            "author": "Yanick112",
            "title": "ComfyUI-ToSVG",
            "id": "tosvg",
            "reference": "https://github.com/Yanick112/ComfyUI-ToSVG",
            "files": [
                "https://github.com/Yanick112/ComfyUI-ToSVG"
            ],
            "install_type": "git-clone",
            "description": "This project converts raster images into SVG format using the [a/VTracer](https://mirror.ghproxy.com/https://github.com/visioncortex/vtracer) library. It's a handy tool for designers and developers who need to work with vector graphics programmatically."
        },
        {
            "author": "dicksondickson",
            "title": "ComfyUI-Dickson-Nodes",
            "id": "dicksonnodes",
            "reference": "https://github.com/dicksondickson/ComfyUI-Dickson-Nodes",
            "files": [
                "https://github.com/dicksondickson/ComfyUI-Dickson-Nodes"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes that I've either written myself or adapted from other authors"
        },
        {
            "author": "juehackr",
            "title": "comfyui_fk_server",
            "id": "fk-server",
            "reference": "https://github.com/juehackr/comfyui_fk_server",
            "files": [
                "https://github.com/juehackr/comfyui_fk_server"
            ],
            "install_type": "git-clone",
            "description": "🤗🤗🤗Comfyui Universal Translation Plugin (no longer requires adding various nodes, directly add translation function on the existing nodes), allowing Comfyui to support Chinese input and automatic translation for any long text input box, while adding error translation function (calling Baidu Translate), achieving translation freedom!"
        },
        {
            "author": "G-370",
            "title": "ComfyUI-SD3-Powerlab",
            "id": "sd3-powerlab",
            "reference": "https://github.com/G-370/ComfyUI-SD3-Powerlab",
            "files": [
                "https://github.com/G-370/ComfyUI-SD3-Powerlab"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Render SD3 Attention, SD3 Attention To Image, SD3 Image Into Attention."
        },
        {
            "author": "TylerZoro",
            "title": "SD3-Scaling",
            "id": "sd3-scaling",
            "reference": "https://github.com/TylerZoro/SD3-Scaling",
            "files": [
                "https://github.com/TylerZoro/SD3-Scaling"
            ],
            "install_type": "git-clone",
            "description": "Tools for scaling images and latents appropriate to SD3 in ComfyUI."
        },
        {
            "author": "baicai99",
            "title": "ComfyUI-FrameSkipping",
            "id": "FrameSkipping",
            "reference": "https://github.com/baicai99/ComfyUI-FrameSkipping",
            "files": [
                "https://github.com/baicai99/ComfyUI-FrameSkipping"
            ],
            "install_type": "git-clone",
            "description": "Used to process video redrawing, frame skipping, frame ending early, etc."
        },
        {
            "author": "SuperMasterBlasterLaser",
            "title": "ComfyUI_YOLO_Classifiers",
            "id": "yolo-classifier",
            "reference": "https://github.com/SuperMasterBlasterLaser/ComfyUI_YOLO_Classifiers",
            "files": [
                "https://github.com/SuperMasterBlasterLaser/ComfyUI_YOLO_Classifiers"
            ],
            "install_type": "git-clone",
            "description": "Nodes:YOLO Classifier Model Loader, YOLO Classify."
        },
        {
            "author": "SamKhoze",
            "title": "DeepFuze",
            "id": "deepfuze",
            "reference": "https://github.com/SamKhoze/ComfyUI-DeepFuze",
            "files": [
                "https://github.com/SamKhoze/ComfyUI-DeepFuze"
            ],
            "install_type": "git-clone",
            "description": "DeepFuze is a state-of-the-art deep learning tool that seamlessly integrates with ComfyUI to revolutionize facial transformations, lipsyncing, video generation, voice cloning, face swapping, and lipsync translation. Leveraging advanced algorithms, DeepFuze enables users to combine audio and video with unparalleled realism, ensuring perfectly synchronized facial movements. This innovative solution is ideal for content creators, animators, developers, and anyone seeking to elevate their video editing projects with sophisticated AI-driven features."
        },
        {
            "author": "superyoman",
            "title": "comfyui_lumaAPI",
            "id": "luma",
            "reference": "https://github.com/superyoman/comfyui_lumaAPI",
            "files": [
                "https://github.com/superyoman/comfyui_lumaAPI"
            ],
            "install_type": "git-clone",
            "description": "Unofficial Luma API-ComfyUI version.[w/WARN: This project is for learning purpose only!]"
        },
        {
            "author": "chris-the-wiz",
            "title": "EmbeddingsCurveEditor_ComfyUI",
            "id": "embeddings-curve-editor",
            "reference": "https://github.com/chris-the-wiz/EmbeddingsCurveEditor_ComfyUI",
            "files": [
                "https://github.com/chris-the-wiz/EmbeddingsCurveEditor_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Edit embeddings with a curve. Actually should work on any 1D input tensor. Tested with IPAdapter-Plus."
        },
        {
            "author": "zhulu111",
            "title": "ComfyUI_Bxb",
            "id": "ComfyUI_Bxb",
            "reference": "https://github.com/zhulu111/ComfyUI_Bxb",
            "files": [
                "https://github.com/zhulu111/ComfyUI_Bxb"
            ],
            "install_type": "git-clone",
            "description": "sdBxb, a tool that converts ComfyUI workflows into WeChat Mini Program, Douyin Mini Program, and H5 with one click, and supports payments."
        },
        {
            "author": "lordgasmic",
            "title": "comfyui_wildcards",
            "reference": "https://github.com/lordgasmic/comfyui_wildcards",
            "files": [
                "https://github.com/lordgasmic/comfyui_wildcards"
            ],
            "install_type": "git-clone",
            "description": "This is an attempt to recreate the wildcards plugin for Automatic1111 but for ComfyUI."
        },
        {
            "author": "lordgasmic",
            "title": "comfyui_save_image_with_options",
            "reference": "https://github.com/lordgasmic/comfyui_save_image_with_options",
            "files": [
                "https://github.com/lordgasmic/comfyui_save_image_with_options"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Save Image with Options"
        },
        {
            "author": "opvelll",
            "title": "Comfy UI Text List Product",
            "id": "listproduct",
            "reference": "https://github.com/opvelll/ComfyUI_TextListProduct",
            "files": [
                "https://github.com/opvelll/ComfyUI_TextListProduct"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node for Comfy UI. It mainly wraps itertools.product and can be used to create patterns by combining prompts. It is recommended to install this custom node in combination with the nodes from the WAS Node Suite."
        },
        {
            "author": "jakechai",
            "title": "ComfyUI-JakeUpgrade",
            "id": "jkupgrade",
            "reference": "https://github.com/jakechai/ComfyUI-JakeUpgrade",
            "files": [
                "https://github.com/jakechai/ComfyUI-JakeUpgrade"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI workflow customization by Jake."
        },
        {
            "author": "celsojr2013",
            "title": "ComfyUI SimpleTools Suit",
            "reference": "https://github.com/celsojr2013/comfyui_simpletools",
            "files": [
                "https://github.com/celsojr2013/comfyui_simpletools"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Simple Google Translator, Simple Resolution Solver.\nThis is a small set of simple nodes that help your workflow on ComfyUI."
        },
        {
            "author": "celsojr2013",
            "title": "comfyui_jamworks_client",
            "reference": "https://github.com/celsojr2013/comfyui_jamworks_client",
            "files": [
                "https://github.com/celsojr2013/comfyui_jamworks_client"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Jamworks_Login, Jamworks_Download, Shell_Command.\nA Simple Client for Jamworks Platform DAM Integration"
        },
        {
            "author": "fablestudio",
            "title": "ComfyUI-Showrunner-Utils",
            "reference": "https://github.com/fablestudio/ComfyUI-Showrunner-Utils",
            "files": [
                "https://github.com/fablestudio/ComfyUI-Showrunner-Utils"
            ],
            "install_type": "git-clone",
            "description": "Comfyui Custom Nodes for Showrunner"
        },
        {
            "author": "MilitantHitchhiker",
            "title": "MilitantHitchhiker-SwitchbladePack",
            "id": "hitchhiker",
            "reference": "https://github.com/MilitantHitchhiker/MilitantHitchhiker-SwitchbladePack",
            "files": [
                "https://github.com/MilitantHitchhiker/MilitantHitchhiker-SwitchbladePack"
            ],
            "install_type": "git-clone",
            "description": "Militant Hitchhiker's Switchblade Pack is a collection of custom nodes for ComfyUI that provide various multi-function capabilities."
        },
        {
            "author": "slyt",
            "title": "comfyui-ollama-nodes",
            "reference": "https://github.com/slyt/comfyui-ollama-nodes",
            "files": [
                "https://github.com/slyt/comfyui-ollama-nodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom nodes for working with [a/Ollama](https://mirror.ghproxy.com/https://github.com/ollama/ollama).\nNOTE:Assumes that an Ollama server is running at http://127.0.0.1:11434 and accessible by the ComfyUI backend."
        },
        {
            "author": "zwng",
            "title": "ComfyUI_Zwng_Nodes",
            "id": "zwng",
            "reference": "https://github.com/za-wa-n-go/ComfyUI_Zwng_Nodes",
            "files": [
                "https://github.com/za-wa-n-go/ComfyUI_Zwng_Nodes"
            ],
            "install_type": "git-clone",
            "description": "Simple nodes for loading image files.Nodes that include a simple remote connection to Photoshop, a node that can overlay and preview an image with a mask, and a node that can load images directly from a file path."
        },
        {
            "author": "RedRayz",
            "title": "ComfyUI-Danbooru-To-WD",
            "id": "danbooru2wd",
            "reference": "https://github.com/RedRayz/ComfyUI-Danbooru-To-WD",
            "files": [
                "https://github.com/RedRayz/ComfyUI-Danbooru-To-WD"
            ],
            "install_type": "git-clone",
            "description": "Converts booru tags to a format suitable for Waifu Diffusion(or Danbooru based models)."
        },
        {
            "author": "Shibiko-AI",
            "title": "Shibiko AI ComfyUI Tools",
            "id": "shibiko-ai-tools",
            "reference": "https://github.com/Shibiko-AI/ShibikoAI-ComfyUI-Tools",
            "files": [
                "https://github.com/Shibiko-AI/ShibikoAI-ComfyUI-Tools"
            ],
            "install_type": "git-clone",
            "description": "This is a collection of tools that I use to make my life easier when developing ComfyUI applications. It is a collection of tools that I have created to help me with my development process. I have decided to share these tools with the community in the hopes that they will be useful to others as well. I use this tools to further develop features for [a/https://shibiko.ai](https://shibiko.ai)"
        },
        {
            "author": "SaltAI",
            "title": "SaltAI_AudioViz",
            "id": "saltai-audioviz",
            "reference": "https://github.com/get-salt-AI/SaltAI_AudioViz",
            "files": [
                "https://github.com/get-salt-AI/SaltAI_AudioViz"
            ],
            "install_type": "git-clone",
            "description": "SaltAI AudioViz contains ComfyUI nodes for generating complex audio reactive visualizations"
        },
        {
            "author": "SherryXieYuchen",
            "title": "ComfyUI-Image-Inpainting",
            "id": "image-inpainting",
            "reference": "https://github.com/SherryXieYuchen/ComfyUI-Image-Inpainting",
            "files": [
                "https://github.com/SherryXieYuchen/ComfyUI-Image-Inpainting"
            ],
            "install_type": "git-clone",
            "description": "Nodes:VAE Encode Inpaint, VAE Decode Inpaint, ColorCorrection Inpaint, ImagePreprocess Inpaint, ImagePostprocess Inpaint, Load Model Inpaint, Inpainting (using Model)"
        },
        {
            "author": "zeroxoxo",
            "title": "ComfyUI-Fast-Style-Transfer",
            "id": "fast-style-transfer",
            "reference": "https://github.com/zeroxoxo/ComfyUI-Fast-Style-Transfer",
            "files": [
                "https://github.com/zeroxoxo/ComfyUI-Fast-Style-Transfer"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node for fast neural style transfer. This is a simple conversion based on this: [a/https://mirror.ghproxy.com/https://github.com/rrmina/fast-neural-style-pytorch](https://mirror.ghproxy.com/https://github.com/rrmina/fast-neural-style-pytorch) Only basic inference functionality is ported for now."
        },
        {
            "author": "iwanders",
            "title": "iwanders/ComfyUI_nodes",
            "id": "iwanders-nodes",
            "reference": "https://github.com/iwanders/ComfyUI_nodes",
            "files": [
                "https://github.com/iwanders/ComfyUI_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:IW SaveString, IW PrintString, IW ReplaceString, IW StringNode, IW StringConcat, IW TokenizerVocab, IW JsonPickItem."
        },
        {
            "author": "rhdunn",
            "title": "comfyui-bus-plugin",
            "id": "bus",
            "reference": "https://github.com/rhdunn/comfyui-bus-plugin",
            "files": [
                "https://github.com/rhdunn/comfyui-bus-plugin"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes for rerouting multiple I/O lines together in a bus."
        },
        {
            "author": "rhdunn",
            "title": "comfyui-audio-processing",
            "reference": "https://github.com/rhdunn/comfyui-audio-processing",
            "files": [
                "https://github.com/rhdunn/comfyui-audio-processing"
            ],
            "install_type": "git-clone",
            "description": "This plugin is compatible with the ComfyUI audio nodes."
        },
        {
            "author": "hyunamy",
            "title": "Comfy-UI on-complete-email-me",
            "id": "hyunamy",
            "reference": "https://github.com/hyunamy/comfy-ui-on-complete-email-me",
            "files": [
                "https://github.com/hyunamy/comfy-ui-on-complete-email-me"
            ],
            "install_type": "git-clone",
            "description": "A feature that sends an email via Gmail once image generation is completed in Comfy-ui."
        },
        {
            "author": "veighnsche",
            "title": "comfyui_gr85",
            "id": "gr85",
            "reference": "https://github.com/veighnsche/comfyui_gr85",
            "files": [
                "https://github.com/veighnsche/comfyui_gr85"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Dimension Resizer, Image Sizer, Random Ratio, Show Text, Random Title Character, Random Wildcard Tag Picker, Random Show Atm Loc Outfit, Contains Word, Elements Concatenator, ..."
        },
        {
            "author": "leiweiqiang",
            "title": "ComfyUI-TRA",
            "id": "tra",
            "reference": "https://github.com/leiweiqiang/ComfyUI-TRA",
            "files": [
                "https://github.com/leiweiqiang/ComfyUI-TRA"
            ],
            "install_type": "git-clone",
            "description": "Nodes:TCL EbSynth, TCL Extract Frames (From File), TCL Extract Frames (From Video), TCL Combine Frames, TCL Save Video (From Frames)"
        },
        {
            "author": "hwhaocool",
            "title": "ComfyUI-Select-Any",
            "id": "select-any",
            "reference": "https://github.com/hwhaocool/ComfyUI-Select-Any",
            "files": [
                "https://github.com/hwhaocool/ComfyUI-Select-Any"
            ],
            "install_type": "git-clone",
            "description": "a comfyui custom node, which can select value from inputs"
        },
        {
            "author": "GreenLandisaLie",
            "title": "AuraSR-ComfyUI",
            "id": "aurasr-greenlandisalie",
            "reference": "https://github.com/GreenLandisaLie/AuraSR-ComfyUI",
            "files": [
                "https://github.com/GreenLandisaLie/AuraSR-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI implementation of [a/Aura-SR](https://mirror.ghproxy.com/https://github.com/fal-ai/aura-sr)"
        },
        {
            "author": "licyk",
            "title": "ComfyUI-Restart-Sampler",
            "id": "restart-sampler-licyk",
            "reference": "https://github.com/licyk/ComfyUI-Restart-Sampler",
            "files": [
                "https://github.com/licyk/ComfyUI-Restart-Sampler"
            ],
            "install_type": "git-clone",
            "description": "This extension is a node that directly expands the functionality of KSampler, rather than being in the form of a custom node. [w/Workflows created using this feature are not compatible with other users.]"
        },
        {
            "author": "my-opencode",
            "title": "ComfyUI_IndustrialMagick",
            "id": "industrialmagick",
            "reference": "https://github.com/my-opencode/ComfyUI_IndustrialMagick",
            "files": [
                "https://github.com/my-opencode/ComfyUI_IndustrialMagick"
            ],
            "install_type": "git-clone",
            "description": "[a/ImageMagick](https://imagemagick.org/index.php) nodes for ComfyUI. Adds nodes to call ImageMagick subprocesses from ComfyUI.\nRequirements: [a/ImagMagick7](https://imagemagick.org/script/download.php), 'magick' command in your CLI environment."
        },
        {
            "author": "my-opencode",
            "title": "ComfyUI_KSamplerTimer",
            "id": "ksamplertimer",
            "reference": "https://github.com/my-opencode/ComfyUI_KSamplerTimer",
            "files": [
                "https://github.com/my-opencode/ComfyUI_KSamplerTimer"
            ],
            "install_type": "git-clone",
            "description": "A custom node that returns the generation time of the KSampler. Intended for benchmarking or debugging."
        },
        {
            "author": "SEkINVR",
            "title": "ComfyUI SaveAS",
            "id": "saveas",
            "reference": "https://github.com/SEkINVR/ComfyUI-SaveAs",
            "files": [
                "https://github.com/SEkINVR/ComfyUI-SaveAs"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI allows you to save images in multiple formats, including PNG, JPG, WebP, and ICO.\n[w/ComfyUI-Save-Multi-Format is renamed to SaveAs. Remove previous one and reinstall to this.]"
        },
        {
            "author": "MrSamSeen",
            "title": "ComfyUI_SSStereoscope",
            "id": "ssstereoscope",
            "reference": "https://github.com/MrSamSeen/ComfyUI_SSStereoscope",
            "files": [
                "https://github.com/MrSamSeen/ComfyUI_SSStereoscope"
            ],
            "install_type": "git-clone",
            "description": "Side by Side 3D Stereoscope generation node for ComfyUI by SamSeen."
        },
        {
            "author": "jroc22",
            "title": "ComfyUI-CSV-prompt-builder",
            "id": "csv-prompt-builder",
            "reference": "https://github.com/jroc22/ComfyUI-CSV-prompt-builder",
            "files": [
                "https://github.com/jroc22/ComfyUI-CSV-prompt-builder"
            ],
            "install_type": "git-clone",
            "description": "This is a simple node for creating prompts using a .csv file. I created this node as an easy way to output different prompts each time a workflow is run."
        },
        {
            "author": "DeJoker",
            "title": "Pipeline Parallel ComfyUI",
            "reference": "https://github.com/DeJoker/pipeline-parallel-comfy",
            "files": [
                "https://github.com/DeJoker/pipeline-parallel-comfy"
            ],
            "install_type": "git-clone",
            "description": "provide extra api to run prompt request with parallel execution of independent node"
        },
        {
            "author": "yiwangsimple",
            "title": "ComfyUI_DW_Chat",
            "reference": "https://github.com/yiwangsimple/ComfyUI_DW_Chat",
            "files": [
                "https://github.com/yiwangsimple/ComfyUI_DW_Chat"
            ],
            "install_type": "git-clone",
            "description": "Content generation with open source models in comfyui via graq api implementation.\n[w/This repo is renamed from ComfyUI_GroqChat to ComfyUI_DW_CHAT. Please remove previous one and reinstall to this.]"
        },
        {
            "author": "yiwangsimple",
            "title": "florence_dw",
            "reference": "https://github.com/yiwangsimple/florence_dw",
            "files": [
                "https://github.com/yiwangsimple/florence_dw"
            ],
            "install_type": "git-clone",
            "description": "Based on the original repository [a/https://mirror.ghproxy.com/https://github.com/spacepxl/ComfyUI-Florence-2](https://mirror.ghproxy.com/https://github.com/spacepxl/ComfyUI-Florence-2), the model loading and storage methods have been improved, and sd3 has been newly added with enhanced speed and accuracy."
        },
        {
            "author": "Tritant",
            "title": "ComfyUI-CreaPrompt",
            "id": "creaprompt",
            "reference": "https://github.com/tritant/ComfyUI_CreaPrompt",
            "files": [
                "https://github.com/tritant/ComfyUI_CreaPrompt"
            ],
            "install_type": "git-clone",
            "description": "Generate random prompts easily."
        },
        {
            "author": "metncelik",
            "title": "comfyui_met_suite",
            "reference": "https://github.com/metncelik/comfyui_met_suite",
            "files": [
                "https://github.com/metncelik/comfyui_met_suite"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Primitive BBOX, BBOX Padding, BBOX Resize, ImageResize KeepRatio."
        },
        {
            "author": "Smuzzies",
            "title": "comfyui_meme_maker",
            "reference": "https://github.com/Smuzzies/comfyui_meme_maker",
            "files": [
                "https://github.com/Smuzzies/comfyui_meme_maker"
            ],
            "install_type": "git-clone",
            "description": "Meme Maker Node for ComfyUI."
        },
        {
            "author": "bluevisor",
            "title": "ComfyUI_PS_Blend_Node",
            "reference": "https://github.com/bluevisor/ComfyUI_PS_Blend_Node",
            "files": [
                "https://github.com/bluevisor/ComfyUI_PS_Blend_Node"
            ],
            "install_type": "git-clone",
            "description": "This repository contains a simple custom node for ComfyUI that implements familiar PS-style blend modes using PyTorch. The PSBlendNode allows you to blend two images together using a variety of blend modes and an opacity parameter."
        },
        {
            "author": "wTechArtist",
            "title": "ComfyUI-CustomNodes",
            "reference": "https://github.com/wTechArtist/ComfyUI-CustomNodes",
            "files": [
                "https://github.com/wTechArtist/ComfyUI-CustomNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Image Blending Mode Mask, Load Image With Bool, IPAdapter Mad Scientist Weight_Type, IPAdapter FaceID With Bool"
        },
        {
            "author": "mullakhmetov",
            "title": "comfyui_dynamic_util_nodes",
            "reference": "https://github.com/mullakhmetov/comfyui_dynamic_util_nodes",
            "files": [
                "https://github.com/mullakhmetov/comfyui_dynamic_util_nodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyS3 helpful util nodes for dynamic workflows"
        },
        {
            "author": "HECer",
            "title": "ComfyUI-FilePathCreator",
            "reference": "https://github.com/HECer/ComfyUI-FilePathCreator",
            "files": [
                "https://github.com/HECer/ComfyUI-FilePathCreator"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI-FilePathCreator is a custom node extension for ComfyUI designed to generate dynamic filenames based on user-defined parameters. This node helps streamline the process of creating organized and timestamped filenames, ideal for saving output files in a structured manner."
        },
        {
            "author": "adigayung",
            "title": "ComfyUI-Translator",
            "reference": "https://github.com/adigayung/ComfyUI-Translator",
            "files": [
                "https://github.com/adigayung/ComfyUI-Translator"
            ],
            "install_type": "git-clone",
            "description": "Auto translate all languages ​​to english"
        },
        {
            "author": "ZZXYWQ",
            "title": "ZZX Nodes",
            "id": "ZZXYWQ",
            "reference": "https://github.com/ZZXYWQ/ComfyUI-ZZXYWQ",
            "files": [
                "https://github.com/ZZXYWQ/ComfyUI-ZZXYWQ"
            ],
            "install_type": "git-clone",
            "description": "Nodes: StreamRecorder, VideoFormatConverter, ZZX_PaintsUndo"
        },
        {
            "author": "SiliconFlow",
            "title": "☁️BizyAir Nodes",
            "id": "bizyair",
            "reference": "https://github.com/siliconflow/BizyAir",
            "files": [
                "https://github.com/siliconflow/BizyAir"
            ],
            "install_type": "git-clone",
            "description": "[a/BizyAir](https://mirror.ghproxy.com/https://github.com/siliconflow/BizyAir) Comfy Nodes that can run in any environment."
        },
        {
            "author": "BenNarum",
            "title": "ComfyUI_CAS",
            "reference": "https://github.com/BenNarum/ComfyUI_CAS",
            "files": [
                "https://github.com/BenNarum/ComfyUI_CAS"
            ],
            "install_type": "git-clone",
            "description": "This extension provides nodes that allow experimentation with various elements (samplers, latent, activators, attenuator, scheulders, ...) of Stable Diffusion."
        },
        {
            "author": "Indra's Mirror",
            "title": "ComfyUI-Documents",
            "reference": "https://github.com/Excidos/ComfyUI-Documents",
            "files": [
                "https://github.com/Excidos/ComfyUI-Documents"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Documents is a powerful extension for ComfyUI that enhances workflows with advanced document processing capabilities. It includes nodes for loading and parsing various document types (PDF, TXT, DOC, DOCX), converting PDF pages to images, splitting PDFs into individual pages, and selecting specific images from batches. Features include text extraction, image conversion, and seamless integration with existing ComfyUI projects."
        },
        {
            "author": "Indra's Mirror",
            "title": "ComfyUI-Lumina-Next-SFT-DiffusersWrapper",
            "reference": "https://github.com/Excidos/ComfyUI-Lumina-Next-SFT-DiffusersWrapper",
            "files": [
                "https://github.com/Excidos/ComfyUI-Lumina-Next-SFT-DiffusersWrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Lumina-Next-SFT-DiffusersWrapper is a custom node for ComfyUI that integrates the advanced Lumina-Next-SFT model. It offers high-quality image generation with features like time-aware scaling, optional ODE sampling, and support for high-resolution outputs. This node brings the power of the Lumina text-to-image pipeline directly into ComfyUI workflows, allowing for flexible and powerful image generation capabilities."
        },
        {
            "author": "Expo",
            "title": "LM Studio Image to Text Node for ComfyUI",
            "id": "comfyui-lmstudio-image-to-text-node",
            "reference": "https://github.com/mattjohnpowell/comfyui-lmstudio-image-to-text-node",
            "files": [
                "https://github.com/mattjohnpowell/comfyui-lmstudio-image-to-text-node"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI that integrates LM Studio's vision models to generate text descriptions of images. It provides a flexible and customizable way to add image-to-text capabilities to your ComfyUI workflows, working with LM Studio's local API."
        },
        {
            "author": "injet-zhou",
            "title": "comfyui_extra_api",
            "reference": "https://github.com/injet-zhou/comfyui_extra_api",
            "files": [
                "https://github.com/injet-zhou/comfyui_extra_api"
            ],
            "install_type": "git-clone",
            "description": "Add more endpoints to make easy for utilizing ComfyUI API."
        },
        {
            "author": "leestuartx",
            "title": "ComfyUI-GG",
            "reference": "https://github.com/leestuartx/ComfyUI-GG",
            "files": [
                "https://github.com/leestuartx/ComfyUI-GG"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-GG is a collection of ComfyUI nodes designed to enhance productivity in image processing workflows. This plugin provides a set of custom nodes that perform various image manipulations and metadata extractions to streamline your tasks."
        },
        {
            "author": "mgfxer",
            "title": "ComfyUI-FrameFX",
            "reference": "https://github.com/mgfxer/ComfyUI-FrameFX",
            "files": [
                "https://github.com/mgfxer/ComfyUI-FrameFX"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes for frame interpolation and video processing in ComfyUI."
        },
        {
            "author": "Cyberschorsch",
            "title": "ComfyUI Checkpoint Loader Config",
            "reference": "https://github.com/Cyberschorsch/ComfyUI-checkpoint-config-loader",
            "files": [
                "https://github.com/Cyberschorsch/ComfyUI-checkpoint-config-loader"
            ],
            "install_type": "git-clone",
            "description": "Provides a custom node to load config for sampler nodes from a yaml file."
        },
        {
            "author": "fearnworks",
            "title": "Fearnworks Nodes",
            "id": "fearnworks",
            "reference": "https://github.com/fearnworks/ComfyUI_FearnworksNodes",
            "files": [
                "https://github.com/fearnworks/ComfyUI_FearnworksNodes"
            ],
            "install_type": "git-clone",
            "description": "This extension provides various nodes to support multimodal workflows."
        },
        {
            "author": "807502278",
            "title": "ComfyUI-3D-MeshTool",
            "id": "3D-MeshTool",
            "reference": "https://github.com/807502278/ComfyUI-3D-MeshTool",
            "files": [
                "https://github.com/807502278/ComfyUI-3D-MeshTool"
            ],
            "install_type": "git-clone",
            "description": "A simple 3D model processing tool within ComfyUI."
        },
        {
            "author": "807502278",
            "title": "ComfyUI-WJNodes",
            "reference": "https://github.com/807502278/ComfyUI-WJNodes",
            "files": [
                "https://github.com/807502278/ComfyUI-WJNodes"
            ],
            "install_type": "git-clone",
            "description": "Ready to use upon download. No need to install dependencies for the time being.\nIf there are new functions or suggestions, please provide feedback.\nAttention! The delfile node is not recommended for use on servers. I am not responsible for any losses incurred."
        },
        {
            "author": "JackEllie",
            "title": "ComfyUI-AI-Assistant",
            "id": "AI-Assistant",
            "reference": "https://github.com/JackEllie/ComfyUI_AI_Assistant",
            "files": [
                "https://github.com/JackEllie/ComfyUI_AI_Assistant"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI native implementation of [a/AI-Assistant](https://mirror.ghproxy.com/https://github.com/tori29umai0123/AI-Assistant)."
        },
        {
            "author": "APZmedia",
            "title": "APZmedia Clean Name",
            "id": "clean-filename",
            "reference": "https://github.com/APZmedia/ComfyUI-APZmedia-cleanName-from-string",
            "files": [
                "https://github.com/APZmedia/ComfyUI-APZmedia-cleanName-from-string"
            ],
            "install_type": "git-clone",
            "description": "A custom node to sanitize text to make clean file names from strings."
        },
        {
            "author": "APZmedia",
            "title": "APZmedia Fast Image Save Node",
            "reference": "https://github.com/APZmedia/APZmedia-comfyui-fast-image-save",
            "files": [
                "https://github.com/APZmedia/APZmedia-comfyui-fast-image-save"
            ],
            "install_type": "git-clone",
            "description": "This node for ComfyUI allows saving images with an optional alpha channel (transparency). It supports saving images in formats like PNG, JPEG, and WebP."
        },
        {
            "author": "N3rd00d",
            "title": "ComfyUI-Paint3D-Nodes",
            "id": "paint3d",
            "reference": "https://github.com/N3rd00d/ComfyUI-Paint3D-Nodes",
            "files": [
                "https://github.com/N3rd00d/ComfyUI-Paint3D-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Paint3D Nodes is a custom ComfyUI node for 3D model texture inpainting based on [a/Paint3D](https://arxiv.org/pdf/2312.13913)."
        },
        {
            "author": "sn0w12",
            "title": "ComfyUI-Sn0w-Scripts",
            "reference": "https://github.com/sn0w12/ComfyUI-Sn0w-Scripts",
            "files": [
                "https://github.com/sn0w12/ComfyUI-Sn0w-Scripts"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes and improvements created for general ease and lora management. These are just nodes I made and found useful, they should work with most other nodes. Most nodes that take in a prompt are made with booru tags in mind and might not work as expected with other prompts."
        },
        {
            "author": "MiaoshouAI",
            "title": "ComfyUI-Miaoshouai-Tagger",
            "id": "miaoshouai-tagger",
            "reference": "https://github.com/miaoshouai/ComfyUI-Miaoshouai-Tagger",
            "files": [
                "https://github.com/miaoshouai/ComfyUI-Miaoshouai-Tagger"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use Florence2 VLM for image tagging and captioning"
        },
        {
            "author": "Patricio Gonzalez Vivo",
            "title": "GLSL Nodes",
            "reference": "https://github.com/patriciogonzalezvivo/comfyui_glslnodes",
            "files": [
                "https://github.com/patriciogonzalezvivo/comfyui_glslnodes"
            ],
            "install_type": "git-clone",
            "description": "A collections of nodes to support GLSL shaders inside a workflow."
        },
        {
            "author": "2kpr",
            "title": "ComfyUI-UltraPixel",
            "id": "comfyui-ultrapixel",
            "reference": "https://github.com/2kpr/ComfyUI-UltraPixel",
            "files": [
                "https://github.com/2kpr/ComfyUI-UltraPixel"
            ],
            "install_type": "git-clone",
            "description": "Implementation of UltraPixel on ComfyUI"
        },
        {
            "author": "LightSketch-ai",
            "title": "ComfyUI-LivePortraitNode (Replicate API)",
            "id": "lightsketchlp",
            "reference": "https://github.com/LightSketch-ai/ComfyUI-LivePortraitNode",
            "files": [
                "https://github.com/LightSketch-ai/ComfyUI-LivePortraitNode"
            ],
            "install_type": "git-clone",
            "description": "Two simple to install nodes to get Live Portrait working in ComfyUI without the need for a fancy GPU (Replicate account needed)."
        },
        {
            "author": "aaronchm",
            "title": "z-a1111-sd-webui-DanTagGen",
            "id": "z-a1111-sd-webui-DanTagGen",
            "reference": "https://github.com/Aaron-CHM/ComfyUI-z-a1111-sd-webui-DanTagGen",
            "files": [
                "https://github.com/Aaron-CHM/ComfyUI-z-a1111-sd-webui-DanTagGen"
            ],
            "install_type": "git-clone",
            "description": "Improved DanTagGen implementation that posesses all functionality of the A1111 webui extension."
        },
        {
            "author": "mikebilly",
            "title": "Transparent-background-comfyUI",
            "id": "transparent-background-comfyui",
            "reference": "https://github.com/mikebilly/Transparent-background-comfyUI",
            "files": [
                "https://github.com/mikebilly/Transparent-background-comfyUI"
            ],
            "install_type": "git-clone",
            "description": "Removes background using Transparent Background"
        },
        {
            "author": "un-seen",
            "title": "comfyui-tensorop",
            "id": "comfyui-tensorop",
            "reference": "https://github.com/un-seen/comfyui-tensorops",
            "files": [
                "https://github.com/un-seen/comfyui-tensorops"
            ],
            "install_type": "git-clone",
            "description": "Nodes to perform tensor operations in ComfyUI"
        },
        {
            "author": "un-seen",
            "title": "ComfyUI Segment Anything",
            "reference": "https://github.com/un-seen/comfyui_segment_anything_plus",
            "files": [
                "https://github.com/un-seen/comfyui_segment_anything_plus"
            ],
            "install_type": "git-clone",
            "description": "This project is a ComfyUI version of [a/sd-webui-segment-anything](https://mirror.ghproxy.com/https://github.com/continue-revolution/sd-webui-segment-anything). At present, only the most core functionalities have been implemented. I would like to express my gratitude to [a/continue-revolution](https://mirror.ghproxy.com/https://github.com/continue-revolution) for their preceding work on which this is based."
        },
        {
            "author": "john-mnz",
            "title": "ComfyUI-Inspyrenet-Rembg",
            "id": "inspyrenet",
            "reference": "https://github.com/john-mnz/ComfyUI-Inspyrenet-Rembg",
            "files": [
                "https://github.com/john-mnz/ComfyUI-Inspyrenet-Rembg"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node for background removal, implementing [a/InSPyReNet](https://mirror.ghproxy.com/https://github.com/plemeri/InSPyReNet)"
        },
        {
            "author": "Koushakur",
            "title": "ComfyUI-DenoiseChooser",
            "id": "denoise-chooser",
            "reference": "https://github.com/Koushakur/ComfyUI-DenoiseChooser",
            "files": [
                "https://github.com/Koushakur/ComfyUI-DenoiseChooser"
            ],
            "install_type": "git-clone",
            "description": "The latent gets passed straight through unaltered, if it's empty (i.e from a 'Empty Latent Image' node) FLOAT outputs the first value, otherwise it outputs the second value"
        },
        {
            "author": "ycchanau",
            "title": "ComfyUI Preview Magnifier",
            "id": "magnifier",
            "reference": "https://github.com/ycchanau/ComfyUI_Preview_Magnifier",
            "files": [
                "https://github.com/ycchanau/ComfyUI_Preview_Magnifier"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes that preview image with a magnifier."
        },
        {
            "author": "lrzjason",
            "title": "Comfyui Kolors Utils",
            "reference": "https://github.com/lrzjason/Comfyui-Kolors-Utils",
            "files": [
                "https://github.com/lrzjason/Comfyui-Kolors-Utils"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Save Weight As Kolors Unet, Save Kolors"
        },
        {
            "author": "lrzjason",
            "title": "ComfyUIJasonNode",
            "reference": "https://github.com/lrzjason/ComfyUIJasonNode",
            "files": [
                "https://github.com/lrzjason/ComfyUIJasonNode/raw/main/SDXLMixSampler.py",
                "https://github.com/lrzjason/ComfyUIJasonNode/raw/main/LatentByRatio.py",
                ""
            ],
            "install_type": "copy",
            "description": "Nodes:SDXLMixSampler, LatentByRatio"
        },
        {
            "author": "amorano",
            "title": "Cozy Communication",
            "id": "cozy_comm",
            "reference": "https://github.com/cozy-comfyui/cozy_comm",
            "files": [
                "https://github.com/cozy-comfyui/cozy_comm"
            ],
            "nodename_pattern": " \\(cozy\\)",
            "install_type": "git-clone",
            "description": "Post images and video to Discord. Nodes to facilitate communication using REST."
        },
        {
            "author": "ShmuelRonen",
            "title": "ComfyUI-AstralAnimator",
            "id": "astralanimator",
            "reference": "https://github.com/ShmuelRonen/ComfyUI-AstralAnimator",
            "files": [
                "https://github.com/ShmuelRonen/ComfyUI-AstralAnimator"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI that enables smooth, keyframe-based animations for image generation. Create dynamic sequences with control over motion, zoom, rotation, and easing effects. Ideal for AI-assisted animation and video content creation."
        },
        {
            "author": "RhizoNymph",
            "title": "ComfyUI-Latte",
            "id": "latte",
            "reference": "https://github.com/RhizoNymph/ComfyUI-Latte",
            "files": [
                "https://github.com/RhizoNymph/ComfyUI-Latte"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use [a/latte](https://mirror.ghproxy.com/https://github.com/Vchitect/Latte) for text to video generation"
        },
        {
            "author": "RhizoNymph",
            "title": "ComfyUI-CLIPSlider",
            "id": "clipslider",
            "reference": "https://github.com/RhizoNymph/ComfyUI-CLIPSlider",
            "files": [
                "https://github.com/RhizoNymph/ComfyUI-CLIPSlider"
            ],
            "install_type": "git-clone",
            "description": "A node to replicate [a/https://huggingface.co/spaces/latentexplorers/latentnavigation-flux](A node to replicate https://huggingface.co/spaces/latentexplorers/latentnavigation-flux)"
        },
        {
            "author": "RhizoNymph",
            "title": "ComfyUI-ColorWheel",
            "reference": "https://github.com/RhizoNymph/ComfyUI-ColorWheel",
            "files": [
                "https://github.com/RhizoNymph/ComfyUI-ColorWheel"
            ],
            "install_type": "git-clone",
            "description": "NODES:Color Wheel Generator"
        },
        {
            "author": "Marksusu",
            "title": "ComfyUI_MTCLIPEncode",
            "id": "mtclipencode",
            "reference": "https://github.com/Marksusu/ComfyUI_MTCLIPEncode",
            "files": [
                "https://github.com/Marksusu/ComfyUI_MTCLIPEncode"
            ],
            "install_type": "git-clone",
            "description": "MTCLIPEncode: An extension for ComfyUI's CLIPTextEncode node, offering multilingual translation (using MarianMT) and prompt enhancement (using Ollama). Seamlessly translate your native language prompts into English and further optimize them for generating your desired images with Stable Diffusion. Supports Krita AI Diffusion."
        },
        {
            "author": "fssorc",
            "title": "ComfyUI_FaceShaper",
            "id": "facesharper",
            "reference": "https://github.com/fssorc/ComfyUI_FaceShaper",
            "files": [
                "https://github.com/fssorc/ComfyUI_FaceShaper"
            ],
            "install_type": "git-clone",
            "description": "Match two faces' shape before using other face swap nodes\nFace-swapping tools typically only replace facial features during the swap, without altering the facial shape. When there is a significant difference in facial shape between the target person and the person in the original photo, the result of the face swap is less satisfactory.\nThis project is a small script that can first liquefy and stretch the face in the original photo according to the horizontal and vertical proportions of the target person's facial contour. The resulting image can be used as input for other face-swapping nodes."
        },
        {
            "author": "fssorc",
            "title": "ComfyUI_pose_inter",
            "reference": "https://github.com/fssorc/ComfyUI_pose_inter",
            "files": [
                "https://github.com/fssorc/ComfyUI_pose_inter"
            ],
            "install_type": "git-clone",
            "description": "Generate transition frames between two character posture images. The prerequisite for running is to have installed comfyui_controlnet_aux, using its Open Pose or DWPose preprocessor"
        },
        {
            "author": "fssorc",
            "title": "ComfyUI_FFT",
            "reference": "https://github.com/fssorc/ComfyUI_FFT",
            "files": [
                "https://github.com/fssorc/ComfyUI_FFT"
            ],
            "install_type": "git-clone",
            "description": "Perform a Fast Fourier Transform on the image, and then users can freely select the filtering range to filter the image. The main function is to remove the grid patterns on the image, and it can also perform high-pass filtering and low-pass filtering. The detailed workflow is shown in the figure below. The PNG file contains the ComfyUI workflow.The working principle is similar to the FFT filter in Photoshop."
        },
        {
            "author": "BetaDoggo",
            "title": "ComfyUI YetAnotherSafetyChecker",
            "id": "yetanothersafetychecker",
            "reference": "https://github.com/BetaDoggo/ComfyUI-YetAnotherSafetyChecker",
            "files": [
                "https://github.com/BetaDoggo/ComfyUI-YetAnotherSafetyChecker"
            ],
            "install_type": "git-clone",
            "description": "Just a simple node to filter out NSFW outputs. This node utilizes [a/AdamCodd/vit-base-nsfw-detector](https://huggingface.co/AdamCodd/vit-base-nsfw-detector) to score the outputs. I chose this model because it's small, fast, and performed very well in my testing. Nudity tends to be scored in the 0.95+ range, but I've set the default to 0.8 as a safe baseline."
        },
        {
            "author": "BetaDoggo",
            "title": "neggles/ComfyUI-WDV-Nodes [gist-wrapper]",
            "reference": "https://github.com/BetaDoggo/ComfyUI-WDV-Nodes",
            "files": [
                "https://github.com/BetaDoggo/ComfyUI-WDV-Nodes"
            ],
            "install_type": "git-clone",
            "description": "100% of code taken from [a/https://gist.github.com/neggles/ecb6327251a9e274428d07636c727eb9](https://gist.github.com/neggles/ecb6327251a9e274428d07636c727eb9)."
        },
        {
            "author": "BetaDoggo",
            "title": "ComfyUI Video Player",
            "id": "videoplayer",
            "reference": "https://github.com/BetaDoggo/ComfyUI-VideoPlayer",
            "files": [
                "https://github.com/BetaDoggo/ComfyUI-VideoPlayer"
            ],
            "install_type": "git-clone",
            "description": "A silly POC Video Player for ComfyUI"
        },
        {
            "author": "BetaDoggo",
            "title": "Gatcha Embeddings",
            "reference": "https://github.com/BetaDoggo/ComfyUI-Gatcha-Embedding",
            "files": [
                "https://github.com/BetaDoggo/ComfyUI-Gatcha-Embedding"
            ],
            "install_type": "git-clone",
            "description": "A revolutionary technique for increasing output variety."
        },
        {
            "author": "BetaDoggo",
            "title": "ComfyUI-FastSDCPU",
            "reference": "https://github.com/BetaDoggo/ComfyUI-FastSDCPU",
            "files": [
                "https://github.com/BetaDoggo/ComfyUI-FastSDCPU"
            ],
            "install_type": "git-clone",
            "description": "A set of nodes for interfacing with the FastSDCPU webserver."
        },
        {
            "author": "WX-NPS1598",
            "title": "Auto Crop By NPS",
            "id": "autocrop-nps",
            "reference": "https://github.com/WX-NPS1598/ComfyUI-Auto_Crop_By_NPS",
            "files": [
                "https://github.com/WX-NPS1598/ComfyUI-Auto_Crop_By_NPS"
            ],
            "install_type": "git-clone",
            "description": "A very useful automatic cropping tool! It can realize cropping, expansion and rotation functions in the form of a slider. "
        },
        {
            "author": "googincheng",
            "title": "ComfyUX",
            "id": "comfyux",
            "reference": "https://github.com/googincheng/ComfyUX",
            "files": [
                "https://github.com/googincheng/ComfyUX"
            ],
            "install_type": "git-clone",
            "description": "Better user experience plugin for ComfyUI."
        },
        {
            "author": "wootwootwootwoot",
            "title": "ComfyUI-RK-Sampler",
            "id": "rk_sampler",
            "reference": "https://github.com/wootwootwootwoot/ComfyUI-RK-Sampler",
            "files": [
                "https://github.com/wootwootwootwoot/ComfyUI-RK-Sampler"
            ],
            "install_type": "git-clone",
            "description": "Batched Runge-Kutta Samplers for ComfyUI"
        },
        {
            "author": "TechnoByteJS",
            "title": "TechNodes",
            "id": "technodes",
            "reference": "https://github.com/TechnoByteJS/ComfyUI-TechNodes",
            "files": [
                "https://github.com/TechnoByteJS/ComfyUI-TechNodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for merging, testing and more.\nNOTE: SDNext Merge, VAE Merge, MBW Layers, Repeat VAE, Quantization."
        },
        {
            "author": "Fantasy AI Studio",
            "title": "FAI-Node",
            "id": "FAI-Node",
            "reference": "https://github.com/alanhuang67/ComfyUI-FAI-Node",
            "files": [
                "https://github.com/alanhuang67/ComfyUI-FAI-Node"
            ],
            "install_type": "git-clone",
            "description": "Various custom nodes for ComfyUI"
        },
        {
            "author": "MuziekMagie",
            "title": "ComfyUI-Matchering",
            "id": "matchering",
            "reference": "https://github.com/MuziekMagie/ComfyUI-Matchering",
            "files": [
                "https://github.com/MuziekMagie/ComfyUI-Matchering"
            ],
            "install_type": "git-clone",
            "description": "A [a/Matchering](https://mirror.ghproxy.com/https://github.com/sergree/matchering)-node for ComfyUI.\nNOTE: You take TWO audio files and feed them into Matchering"
        },
        {
            "author": "filliptm",
            "title": "ComfyUI_FL-Trainer",
            "reference": "https://github.com/filliptm/ComfyUI_FL-Trainer",
            "files": [
                "https://github.com/filliptm/ComfyUI_FL-Trainer"
            ],
            "install_type": "git-clone",
            "description": "Train Image Loras on both sd1.5 and SDXL. This repo git clones the pieces needed to train. It pops open a second terminal window do do the training. It will also display the inference samples in the node itself so you can track the results."
        },
        {
            "author": "Mintbeer96",
            "title": "ComfyUI-KerasOCR",
            "reference": "https://github.com/Mintbeer96/ComfyUI-KerasOCR",
            "files": [
                "https://github.com/Mintbeer96/ComfyUI-KerasOCR"
            ],
            "install_type": "git-clone",
            "description": "An OCR node for detect text in image and returns covering mask."
        },
        {
            "author": "pikenrover",
            "title": "ComfyUI_PRNodes",
            "reference": "https://github.com/pikenrover/ComfyUI_PRNodes",
            "files": [
                "https://github.com/pikenrover/ComfyUI_PRNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:RandomPrompt, RandomPromptMixed, ImageScaleTo, EmptyLatentImageScaleBy, LoraLoaderExtended, Save Image w/Metadata, CheckpointLoaderSimpleExtended"
        },
        {
            "author": "EnragedAntelope",
            "title": "ComfyUI-Doubutsu-Describer",
            "reference": "https://github.com/EnragedAntelope/ComfyUI-Doubutsu-Describer",
            "files": [
                "https://github.com/EnragedAntelope/ComfyUI-Doubutsu-Describer"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI allows you to use the Doubutsu small VLM model to describe images. Credit and further information on Doubutsu: [a/https://huggingface.co/qresearch/doubutsu-2b-pt-756](https://huggingface.co/qresearch/doubutsu-2b-pt-756)"
        },
        {
            "author": "jn-jairo",
            "title": "JNComfy",
            "reference": "https://github.com/jn-jairo/jn_comfyui",
            "files": [
                "https://github.com/jn-jairo/jn_comfyui"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI extension with patches and nodes.\nPatches:Preview device, Extension device, Temperature, Memory estimation, Optimizations, Easy generic inputs, Easy multiple inputs.\nNODES: Image nodes, Image/Area nodes, Image/Blip nodes, Image/Face nodes, Sampling nodes, Patch nodes, Primitive nodes, Primitive/Conversion nodes, Primitive/Process nodes, Workflow nodes, etc..."
        },
        {
            "author": "akierson",
            "title": "comfyui-colornodes",
            "reference": "https://github.com/akierson/comfyui-colornodes",
            "files": [
                "https://github.com/akierson/comfyui-colornodes"
            ],
            "install_type": "git-clone",
            "description": "Basic Color Nodes for ComfyUI"
        },
        {
            "author": "akierson",
            "title": "ComfyUI-textnodes",
            "reference": "https://github.com/akierson/ComfyUI-textnodes",
            "files": [
                "https://github.com/akierson/ComfyUI-textnodes"
            ],
            "install_type": "git-clone",
            "description": "Misc Text Nodes for Comfy UI"
        },
        {
            "author": "ai-shizuka",
            "title": "ComfyUI-tbox",
            "reference": "https://github.com/ai-shizuka/ComfyUI-tbox",
            "files": [
                "https://github.com/ai-shizuka/ComfyUI-tbox"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageLoader, ImageSaver, ImagesSaver, ImageResize, ImageSize."
        },
        {
            "author": "neverbiasu",
            "title": "ComfyUI-Image-Captioner",
            "id": "image-captioner",
            "reference": "https://github.com/neverbiasu/ComfyUI-Image-Captioner",
            "files": [
                "https://github.com/neverbiasu/ComfyUI-Image-Captioner"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI extension for generating captions for your images. Runs on your own system, no external services used, no filter.\nUses various VLMs with APIs to generate captions for images. You can give instructions or ask questions in natural language."
        },
        {
            "author": "neverbiasu",
            "title": "ComfyUI SAM2(Segment Anything 2)",
            "id": "sam2",
            "reference": "https://github.com/neverbiasu/ComfyUI-SAM2",
            "files": [
                "https://github.com/neverbiasu/ComfyUI-SAM2"
            ],
            "install_type": "git-clone",
            "description": "This project adapts the SAM2 to incorporate functionalities from [a/comfyui_segment_anything](https://mirror.ghproxy.com/https://github.com/storyicon/comfyui_segment_anything). Many thanks to continue-revolution for their foundational work."
        },
        {
            "author": "neverbiasu",
            "title": "ComfyUI-StyleShot",
            "reference": "https://github.com/neverbiasu/ComfyUI-StyleShot",
            "files": [
                "https://github.com/neverbiasu/ComfyUI-StyleShot"
            ],
            "install_type": "git-clone",
            "description": "NODES:StyleShotApply"
        },
        {
            "author": "DriftJohnson",
            "title": "DJZ-Nodes",
            "id": "DJZ-Nodes",
            "reference": "https://github.com/MushroomFleet/DJZ-Nodes",
            "files": [
                "https://github.com/MushroomFleet/DJZ-Nodes"
            ],
            "install_type": "git-clone",
            "description": "AspectSize and other nodes"
        },
        {
            "author": "var1ableX",
            "title": "ComfyUI_Accessories",
            "reference": "https://github.com/var1ableX/ComfyUI_Accessories",
            "files": [
                "https://github.com/var1ableX/ComfyUI_Accessories"
            ],
            "install_type": "git-clone",
            "description": "Get Mask Dimensions"
        },
        {
            "author": "Makki_Shizu",
            "title": "comfyui_reimgsize",
            "id": "reimgsize",
            "reference": "https://github.com/MakkiShizu/comfyui_reimgsize",
            "files": [
                "https://github.com/MakkiShizu/comfyui_reimgsize"
            ],
            "install_type": "git-clone",
            "description": "a simple reimgsize node(s) in comfyui."
        },
        {
            "author": "JosefKuchar",
            "title": "ComfyUI-AdvancedTiling",
            "reference": "https://github.com/JosefKuchar/ComfyUI-AdvancedTiling",
            "files": [
                "https://github.com/JosefKuchar/ComfyUI-AdvancedTiling"
            ],
            "install_type": "git-clone",
            "description": "Advanced tiling of various shapes for ComfyUI"
        },
        {
            "author": "Fuou Marinas",
            "title": "ComfyUI-EbSynth",
            "id": "comfyEbsynth",
            "reference": "https://github.com/FuouM/ComfyUI-EbSynth",
            "files": [
                "https://github.com/FuouM/ComfyUI-EbSynth"
            ],
            "install_type": "git-clone",
            "description": "Run EbSynth, Fast Example-based Image Synthesis and Style Transfer, in ComfyUI."
        },
        {
            "author": "Parameshvadivel",
            "title": "ComfyUI-SVGview",
            "id": "svgview",
            "reference": "https://github.com/Parameshvadivel/ComfyUI-SVGview",
            "files": [
                "https://github.com/Parameshvadivel/ComfyUI-SVGview"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Preview SVG"
        },
        {
            "author": "Visionatrix",
            "title": "ComfyUI-Visionatrix",
            "id": "visionatrix",
            "reference": "https://github.com/Visionatrix/ComfyUI-Visionatrix",
            "files": [
                "https://github.com/Visionatrix/ComfyUI-Visionatrix"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI-Visionatrix nodes are designed for convenient ComfyUI to [a/Visionatrix](https://mirror.ghproxy.com/https://github.com/Visionatrix/Visionatrix) workflow support migration, in particular to extract prompt input params (input, textarea, checkbox, select, range, file) to be used in simplified Visionatrix UI."
        },
        {
            "author": "liangt",
            "title": "comfyui-loadimagewithsubfolder",
            "reference": "https://github.com/liangt/comfyui-loadimagewithsubfolder",
            "files": [
                "https://github.com/liangt/comfyui-loadimagewithsubfolder"
            ],
            "install_type": "git-clone",
            "description": "Extend LoadImage node with subfolder support"
        },
        {
            "author": "vault-developer",
            "title": "ImageBlender",
            "reference": "https://github.com/vault-developer/comfyui-image-blender",
            "files": [
                "https://github.com/vault-developer/comfyui-image-blender"
            ],
            "install_type": "git-clone",
            "description": "ComfyuiImageBlender is a custom node for ComfyUI. It may be used to blend two images together using a specified blending mode."
        },
        {
            "author": "gisu",
            "title": "foxpack",
            "id": "foxp",
            "reference": "https://github.com/gisu/comfyui-foxpack",
            "files": [
                "https://github.com/gisu/comfyui-foxpack"
            ],
            "install_type": "git-clone",
            "description": "Collection of nodes for the automation of workflows"
        },
        {
            "author": "webfiltered",
            "title": "WTF? - a debug node for ComfyUI",
            "id": "debugnode",
            "reference": "https://github.com/webfiltered/DebugNode-ComfyUI",
            "files": [
                "https://github.com/webfiltered/DebugNode-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This node provides a simple way to view the output of many nodes, without leaving ComfyUI."
        },
        {
            "author": "pzc163",
            "title": "Comfyui-CatVTON",
            "id": "comfyui-catvton",
            "reference": "https://github.com/pzc163/Comfyui-CatVTON",
            "files": [
                "https://github.com/pzc163/Comfyui-CatVTON"
            ],
            "install_type": "git-clone",
            "description": "Comfyui-CatVTON This repository is the modified official Comfyui node of CatVTON, which is a simple and efficient virtual try-on diffusion model with 1) Lightweight Network (899.06M parameters totally), 2) Parameter-Efficient Training (49.57M parameters trainable) 3) Simplified Inference (< 8G VRAM for 1024X768 resolution).\nThe original GitHub project is [a/https://mirror.ghproxy.com/https://github.com/Zheng-Chong/CatVTON](https://mirror.ghproxy.com/https://github.com/Zheng-Chong/CatVTON)"
        },
        {
            "author": "pzc163",
            "title": "Comfyui_MiniCPMv2_6-prompt-generator",
            "id": "Comfyui_MiniCPMv2_6-prompt-generator",
            "reference": "https://github.com/pzc163/Comfyui_MiniCPMv2_6-prompt-generator",
            "files": [
                "https://github.com/pzc163/Comfyui_MiniCPMv2_6-prompt-generator"
            ],
            "install_type": "git-clone",
            "description": "This is an implementation of [MiniCPMv2_6-prompt-generator](https://huggingface.co/pzc163/MiniCPMv2_6-prompt-generator) by [ComfyUI](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI), including support for single-image caption, generate prompt by upload image and batch-images Prompt generation."
        },
        {
            "author": "aisabervisionlab",
            "title": "ComfyUI_merge_ASVL",
            "id": "merge-asvl",
            "reference": "https://github.com/aisabervisionlab/ComfyUI_merge_ASVL",
            "files": [
                "https://github.com/aisabervisionlab/ComfyUI_merge_ASVL"
            ],
            "install_type": "git-clone",
            "description": "This is a simple node for connecting images. For pictures of the same size, users can choose to fill in vertical in the parameter to connect the pictures vertically or fill in horizontal to connect the pictures horizontally."
        },
        {
            "author": "akatz-ai",
            "title": "Akatz Custom Nodes",
            "id": "akatz-ai",
            "reference": "https://github.com/akatz-ai/ComfyUI-AKatz-Nodes",
            "files": [
                "https://github.com/akatz-ai/ComfyUI-AKatz-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Simple custom node pack for nodes I use in my workflows. Includes Dilate Mask Linear for animating masks."
        },
        {
            "author": "akatz-ai",
            "title": "🌊 Depthflow Nodes",
            "id": "depthflow-akatz-ai",
            "reference": "https://github.com/akatz-ai/ComfyUI-Depthflow-Nodes",
            "files": [
                "https://github.com/akatz-ai/ComfyUI-Depthflow-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Implementation of DepthFlow nodes for ComfyUI, adds a 2.5D parallax effect to images and videos. Compatible with Ryan's Flex system."
        },
        {
            "author": "akatz-ai",
            "title": "DepthCrafter Nodes",
            "id": "depthcrafter-akatz-ai",
            "reference": "https://github.com/akatz-ai/ComfyUI-DepthCrafter-Nodes",
            "files": [
                "https://github.com/akatz-ai/ComfyUI-DepthCrafter-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Implementation of DepthCrafter nodes for ComfyUI, create consistent depth maps for your videos."
        },
        {
            "author": "teward",
            "title": "Comfy-Sentry",
            "reference": "https://github.com/teward/Comfy-Sentry",
            "files": [
                "https://github.com/teward/Comfy-Sentry"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom node that activates integration with a Sentry instance for loading. Has no actual nodes."
        },
        {
            "author": "Fuou Marinas",
            "title": "FM_nodes",
            "reference": "https://github.com/FuouM/FM_nodes",
            "files": [
                "https://github.com/FuouM/FM_nodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of ComfyUI nodes. Including: WFEN, RealViFormer, ProPIH"
        },
        {
            "author": "Fuou Marinas",
            "title": "ComfyUI-FirstOrderMM",
            "id": "fomm",
            "reference": "https://github.com/FuouM/ComfyUI-FirstOrderMM",
            "files": [
                "https://github.com/FuouM/ComfyUI-FirstOrderMM"
            ],
            "install_type": "git-clone",
            "description": "Run [a/First Order Motion Model](https://mirror.ghproxy.com/https://github.com/AliaksandrSiarohin/first-order-model) for Image Animation in ComfyUI."
        },
        {
            "author": "Fuou Marinas",
            "title": "ComfyUI-StyleTransferPlus",
            "id": "styletransferplus",
            "reference": "https://github.com/FuouM/ComfyUI-StyleTransferPlus",
            "files": [
                "https://github.com/FuouM/ComfyUI-StyleTransferPlus"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Neural Neighbor, CAST, EFDM, MicroAST, Coral Color Transfer."
        },
        {
            "author": "MiddleKD",
            "title": "ComfyUI-mem-safe-wrapper",
            "reference": "https://github.com/MiddleKD/ComfyUI-mem-safe-wrapper",
            "files": [
                "https://github.com/MiddleKD/ComfyUI-mem-safe-wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI's Smart Memory Management efficiently manages RAM, GPU memory, and garbage collection. This feature keeps frequently used models in memory to increase inference speed, and dynamically releases less important models when memory is low to optimize resources. However, not all ComfyUI custom node developers create nodes that are compatible with Smart memory management. This includes several impressive models. Mem-safe-wrapper is a custom node that wraps these model nodes to enable ComfyUI's Smart memory management capabilities."
        },
        {
            "author": "MiddleKD",
            "title": "ComfyUI-productfix",
            "reference": "https://github.com/MiddleKD/ComfyUI-productfix",
            "files": [
                "https://github.com/MiddleKD/ComfyUI-productfix"
            ],
            "install_type": "git-clone",
            "description": "This is a ComfyUI custom node that helps generate images while preserving the text, logos, and details of e-commerce products."
        },
        {
            "author": "MiddleKD",
            "title": "ComfyUI-denoise-mask-scheduler",
            "reference": "https://github.com/MiddleKD/ComfyUI-denoise-mask-scheduler",
            "files": [
                "https://github.com/MiddleKD/ComfyUI-denoise-mask-scheduler"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-denoise-mask-scheduler experimental approach involves selectively applying a denoise mask at each step during the inpainting inference process in diffusion models."
        },
        {
            "author": "PowerHouseMan",
            "title": "ComfyUI-AdvancedLivePortrait",
            "id": "advancedliveportrait",
            "reference": "https://github.com/PowerHouseMan/ComfyUI-AdvancedLivePortrait",
            "files": [
                "https://github.com/PowerHouseMan/ComfyUI-AdvancedLivePortrait"
            ],
            "install_type": "git-clone",
            "description": "AdvancedLivePortrait with Facial expression editor"
        },
        {
            "author": "cdxOo",
            "title": "Text Node With Comments (@cdxoo)",
            "reference": "https://github.com/cdxOo/comfyui-text-node-with-comments",
            "files": [
                "https://github.com/cdxOo/comfyui-text-node-with-comments"
            ],
            "install_type": "git-clone",
            "description": "multiline text node that strips c-style comments (i.e.'//' and '/* ... */') before passing output string downstream"
        },
        {
            "author": "noarche",
            "title": "noarche/Color Enhance",
            "id": "color-enhance",
            "reference": "https://github.com/noarche/sd-webui-color-enhance",
            "files": [
                "https://github.com/noarche/sd-webui-color-enhance"
            ],
            "install_type": "git-clone",
            "description": "Script for [a/AUTOMATIC1111/stable-diffusion-webui](https://mirror.ghproxy.com/https://github.com/AUTOMATIC1111/stable-diffusion-webui) and node for ComfyUI to enhance colors.\nThis is the same algorithm GIMP/GEGL uses for color enhancement. The gist of this implementation is that it converts the color space to [CIELCh(ab) and normalizes the chroma (or '[a/colorfulness](https://en.wikipedia.org/wiki/Colorfulness)') component. Original source can be found in the link below."
        },
        {
            "author": "emojiiii",
            "title": "ComfyUI_Emojiiii_Custom_Nodes",
            "reference": "https://github.com/emojiiii/ComfyUI_Emojiiii_Custom_Nodes",
            "files": [
                "https://github.com/emojiiii/ComfyUI_Emojiiii_Custom_Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:MultiTextEncode, KolorsMultiTextEncode, Caption, BatchImageProcessor"
        },
        {
            "author": "aonekoss",
            "title": "ComfyUI-Counter",
            "reference": "https://github.com/oleksandr612/ComfyUI-Counter",
            "files": [
                "https://github.com/oleksandr612/ComfyUI-Counter"
            ],
            "install_type": "git-clone",
            "description": "A simple counter, when pressing 'Queue Prompt' resets the count."
        },
        {
            "author": "alpertunga-bile",
            "title": "image-caption-comfyui",
            "reference": "https://github.com/alpertunga-bile/image-caption-comfyui",
            "files": [
                "https://github.com/alpertunga-bile/image-caption-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Using image caption models to extract prompts in ComfyUI"
        },
        {
            "author": "Anibaaal",
            "title": "ComfyUI UX Nodes",
            "reference": "https://github.com/Anibaaal/ComfyUI-UX-Nodes",
            "files": [
                "https://github.com/Anibaaal/ComfyUI-UX-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Easy Resolution Picker, Save Diffusion Model, Load Checkpoint BNB On the fly, Load UNET BNB On the fly"
        },
        {
            "author": "NMWave",
            "title": "Image Captioning and Tagging Assistor Nodes",
            "id": "naderimagecaptionandtag",
            "reference": "https://github.com/NMWave/ComfyUI-Nader-Tagging",
            "files": [
                "https://github.com/NMWave/ComfyUI-Nader-Tagging"
            ],
            "install_type": "git-clone",
            "description": "A small set of useful nodes which aid with the tagging process by splitting tags and strings, alternating tags from multiple sources and removing duplicates."
        },
        {
            "author": "caleboleary",
            "title": "Arc2Face ComfyUI Node Library",
            "reference": "https://github.com/caleboleary/ComfyUI-Arc2Face",
            "files": [
                "https://github.com/caleboleary/ComfyUI-Arc2Face"
            ],
            "install_type": "git-clone",
            "description": "This ComfyUI node library builds upon the work done to train the [a/Arc2Face](https://mirror.ghproxy.com/https://github.com/foivospar/Arc2Face) model by foivospar. It provides a set of nodes for ComfyUI that allow users to extract face embeddings, generate images based on these embeddings, and perform image-to-image transformations."
        },
        {
            "author": "GeekyGhost",
            "title": "ComfyUI-GeekyRemB",
            "reference": "https://github.com/GeekyGhost/ComfyUI-GeekyRemB",
            "files": [
                "https://github.com/GeekyGhost/ComfyUI-GeekyRemB"
            ],
            "install_type": "git-clone",
            "description": "GeekyRemB is a powerful and versatile image processing node for ComfyUI, designed to remove backgrounds from images with advanced customization options. This node leverages the rembg library and offers a wide range of features for fine-tuning the background removal process and enhancing the resulting images."
        },
        {
            "author": "Dobidop",
            "title": "Dobidop ComfyStereo",
            "id": "simple-stereoscopic",
            "reference": "https://github.com/Dobidop/ComfyStereo",
            "files": [
                "https://github.com/Dobidop/ComfyStereo"
            ],
            "install_type": "git-clone",
            "description": "Two simple nodes for stereoscopic image generation. Nodes: Stereo Image Node - a basic port from the Automatic1111 stereo script in thygate/stable-diffusion-webui-depthmap-script, LazyStereo - a naïve stereo image generator"
        },
        {
            "author": "SeniorPioner",
            "title": "SP-Nodes",
            "id": "spnodes",
            "reference": "https://github.com/bananasss00/ComfyUI-SP-Nodes",
            "files": [
                "https://github.com/bananasss00/ComfyUI-SP-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Node Pack: PromptChecker for token toggling, KoboldCPP API, ModelMerging, Telegram-Bot-API, and more"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI_M3Net",
            "reference": "https://github.com/leeguandong/ComfyUI_M3Net",
            "files": [
                "https://github.com/leeguandong/ComfyUI_M3Net"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI for [a/M3Net](https://mirror.ghproxy.com/https://github.com/I2-Multimedia-Lab/M3Net)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI_InternVL2",
            "reference": "https://github.com/leeguandong/ComfyUI_InternVL2",
            "files": [
                "https://github.com/leeguandong/ComfyUI_InternVL2"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI for [a/InternVL](https://mirror.ghproxy.com/https://github.com/OpenGVLab/InternVL)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI_LLaSM",
            "reference": "https://github.com/leeguandong/ComfyUI_LLaSM",
            "files": [
                "https://github.com/leeguandong/ComfyUI_LLaSM"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI for [a/LLaSM](https://huggingface.co/spaces/LinkSoul/LLaSM)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI nodes to use VideoEditing",
            "reference": "https://github.com/leeguandong/ComfyUI_VideoEditing",
            "files": [
                "https://github.com/leeguandong/ComfyUI_VideoEditing"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Load Video to Images, Image to Canny, ControlNet Model Loader, VEDit Model Loader, VEdit Sampler. [a/https://mirror.ghproxy.com/https://github.com/SingleZombie/DiffusersExample/tree/main/ReplaceAttn](https://mirror.ghproxy.com/https://github.com/SingleZombie/DiffusersExample/tree/main/ReplaceAttn)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI nodes to use CrossImageAttention",
            "reference": "https://github.com/leeguandong/ComfyUI_CrossImageAttention",
            "files": [
                "https://github.com/leeguandong/ComfyUI_CrossImageAttention"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI for [a/CrossImageAttention](https://mirror.ghproxy.com/https://github.com/garibida/cross-image-attention)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI nodes to use Style-Aligned",
            "reference": "https://github.com/leeguandong/ComfyUI_Style_Aligned",
            "files": [
                "https://github.com/leeguandong/ComfyUI_Style_Aligned"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI for [a/style-aligned](https://mirror.ghproxy.com/https://github.com/google/style-aligned)"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI_VisualAttentionMap",
            "reference": "https://github.com/leeguandong/ComfyUI_VisualAttentionMap",
            "files": [
                "https://github.com/leeguandong/ComfyUI_VisualAttentionMap"
            ],
            "install_type": "git-clone",
            "description": "NODES:HF ModelLoader, Show Images, Text2Image Inference, Decode Latent, Show CrossAttn Map, Show SelfAttn Map"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI nodes to use MasaCtrl",
            "reference": "https://github.com/leeguandong/ComfyUI_MasaCtrl",
            "files": [
                "https://github.com/leeguandong/ComfyUI_MasaCtrl"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use MasaCtrl"
        },
        {
            "author": "leeguandong",
            "title": "ComfyUI_CompareModelWeights",
            "reference": "https://github.com/leeguandong/ComfyUI_CompareModelWeights",
            "files": [
                "https://github.com/leeguandong/ComfyUI_CompareModelWeights"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes to use CompareModelWeights"
        },
        {
            "author": "lenskikh",
            "title": "Propmt Worker",
            "reference": "https://github.com/lenskikh/ComfyUI-Prompt-Worker",
            "files": [
                "https://github.com/lenskikh/ComfyUI-Prompt-Worker"
            ],
            "install_type": "git-clone",
            "description": "Node:Prompt Worker. A text manipulation node for postprocessing of prompt."
        },
        {
            "author": "kappa54",
            "title": "ComfyUI Usability",
            "id": "comfyui_usability",
            "reference": "https://github.com/kappa54m/ComfyUI_Usability",
            "files": [
                "https://github.com/kappa54m/ComfyUI_Usability"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes intended to improve usability."
        },
        {
            "author": "IuvenisSapiens",
            "title": "ComfyUI_MiniCPM-V-2_6-int4",
            "id": "minicpm-v-2_6-int4",
            "reference": "https://github.com/IuvenisSapiens/ComfyUI_MiniCPM-V-2_6-int4",
            "files": [
                "https://github.com/IuvenisSapiens/ComfyUI_MiniCPM-V-2_6-int4"
            ],
            "install_type": "git-clone",
            "description": "This is an implementation of [a/MiniCPM-V-2_6-int4](https://mirror.ghproxy.com/https://github.com/OpenBMB/MiniCPM-V) by [a/ComfyUI](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI), including support for text-based queries, video queries, single-image queries, and multi-image queries to generate captions or responses."
        },
        {
            "author": "IuvenisSapiens",
            "title": "ComfyUI_Qwen2-Audio-7B-Instruct-Int4",
            "id": "qwen2-audio-7b-instruct-int4",
            "reference": "https://github.com/IuvenisSapiens/ComfyUI_Qwen2-Audio-7B-Instruct-Int4",
            "files": [
                "https://github.com/IuvenisSapiens/ComfyUI_Qwen2-Audio-7B-Instruct-Int4"
            ],
            "install_type": "git-clone",
            "description": "This is an implementation of [a/Qwen2-Audio-7B-Instruct-Int4](https://mirror.ghproxy.com/https://github.com/QwenLM/Qwen2-Audio) by [a/ComfyUI](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI), including support for text-based queries and audio queries to generate captions or responses."
        },
        {
            "author": "mltask",
            "title": "MLTask_ComfyUI",
            "id": "mltask_comfyui",
            "reference": "https://github.com/misterjoessef/MLTask_ComfyUI",
            "files": [
                "https://github.com/misterjoessef/MLTask_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "a set of nodes to help u run ai code using MLTask"
        },
        {
            "author": "smlbiobot",
            "title": "ComfyUI-Flux-Replicate-API",
            "id": "replicate-api",
            "reference": "https://github.com/smlbiobot/ComfyUI-Flux-Replicate-API",
            "files": [
                "https://github.com/smlbiobot/ComfyUI-Flux-Replicate-API"
            ],
            "install_type": "git-clone",
            "description": "Flux Pro via Replicate API\nCreate API key at [a/https://replicate.com/account/api-tokens](https://replicate.com/account/api-tokens)\nCopy config.ini.example to config.ini and put the replicate key there."
        },
        {
            "author": "Jjulianadv",
            "title": "Wild Divide",
            "reference": "https://github.com/Julian-adv/WildDivide",
            "files": [
                "https://github.com/Julian-adv/WildDivide"
            ],
            "install_type": "git-clone",
            "description": "This extension provides the ability to build prompts using wildcards for each region of a split image."
        },
        {
            "author": "goburiin",
            "title": "nsfwrecog-comfyui",
            "reference": "https://github.com/goburiin/nsfwrecog-comfyui",
            "files": [
                "https://github.com/goburiin/nsfwrecog-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:NSFW Detector"
        },
        {
            "author": "eastoc",
            "title": "Semantic-SAM",
            "reference": "https://github.com/eastoc/ComfyUI_SemanticSAM",
            "files": [
                "https://github.com/eastoc/ComfyUI_SemanticSAM"
            ],
            "install_type": "git-clone",
            "description": "Segment and Recognize Anything at Any Granularity."
        },
        {
            "author": "LING-APE",
            "title": "ComfyUI-PixelResolutionCalculator",
            "id": "PixelCalulator",
            "reference": "https://github.com/Ling-APE/ComfyUI-PixelResolutionCalculator",
            "files": [
                "https://github.com/Ling-APE/ComfyUI-PixelResolutionCalculator"
            ],
            "install_type": "git-clone",
            "description": "Simple resuluition calculator to convert pixel resolution and aspect ratio to laten friendlt pixel width and height size."
        },
        {
            "author": "Cyber-Blacat",
            "title": "ComfyUI-Yuan",
            "reference": "https://github.com/Cyber-Blacat/ComfyUI-Yuan",
            "files": [
                "https://github.com/Cyber-Blacat/ComfyUI-Yuan"
            ],
            "install_type": "git-clone",
            "description": "Some simple&practical ComfyUI image processing nodes."
        },
        {
            "author": "blackcodetavern",
            "title": "ComfyUI-Benripack",
            "reference": "https://github.com/blackcodetavern/ComfyUI-Benripack",
            "files": [
                "https://github.com/blackcodetavern/ComfyUI-Benripack"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Benripack is an extension for ComfyUI that provides a CharacterPipe node. This node allows for managing various elements such as images, prompts, and models in a single structure, simplifying the workflow for character-based image generation."
        },
        {
            "author": "MohammadAboulEla",
            "title": "ComfyUI-iTools",
            "reference": "https://github.com/MohammadAboulEla/ComfyUI-iTools",
            "files": [
                "https://github.com/MohammadAboulEla/ComfyUI-iTools"
            ],
            "install_type": "git-clone",
            "description": "The iTools are some quality of life nodes, like read a possible prompt used to create an image, save a prompt to file as a new line, read prompts from a multiline file."
        },
        {
            "author": "Hellrunner2k",
            "title": "Hellrunner's Magical Nodes",
            "reference": "https://github.com/Hellrunner2k/ComfyUI-HellrunnersMagicalNodes",
            "files": [
                "https://github.com/Hellrunner2k/ComfyUI-HellrunnersMagicalNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Magical Save Node, Thermal Latenator. This package contains a collection of neat nodes that are supposed to ease your comfy-flow."
        },
        {
            "author": "caleboleary",
            "title": "Comfyui-calbenodes",
            "reference": "https://github.com/caleboleary/Comfyui-calbenodes",
            "files": [
                "https://github.com/caleboleary/Comfyui-calbenodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:CharacterManagerNode, FilmGrain, FlipFlopperSameArch"
        },
        {
            "author": "Raapys",
            "title": "LatentGC Aggressive",
            "id": "latentgcaggressive",
            "reference": "https://github.com/Raapys/ComfyUI-LatentGC_Aggressive",
            "files": [
                "https://github.com/Raapys/ComfyUI-LatentGC_Aggressive"
            ],
            "install_type": "git-clone",
            "description": "Simple latent-passthrough node for running a full VRAM cleanup between workflow stages."
        },
        {
            "author": "Pheat-AI",
            "title": "Remade_nodes",
            "reference": "https://github.com/Pheat-AI/Remade_nodes",
            "files": [
                "https://github.com/Pheat-AI/Remade_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Batch Image Blend by Mask, Batch Enlarged Overlay, Batch Image Overlay, Remove Black Pixels to Transparent, Canny Shrink and Recenter, ..."
        },
        {
            "author": "OgreLemonSoup",
            "title": "Gallery&Tabs",
            "id": "LoadImageGallery",
            "reference": "https://github.com/OgreLemonSoup/ComfyUI-Load-Image-Gallery",
            "files": [
                "https://github.com/OgreLemonSoup/ComfyUI-Load-Image-Gallery"
            ],
            "install_type": "git-clone",
            "description": "Adds a gallery to the Load Image node and tabs for Load Checkpoint/Lora/etc nodes"
        },
        {
            "author": "OuticNZ",
            "title": "ComfyUI-Simple-Of-Complex",
            "reference": "https://github.com/OuticNZ/ComfyUI-Simple-Of-Complex",
            "files": [
                "https://github.com/OuticNZ/ComfyUI-Simple-Of-Complex"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Text Switch 2 Way, Prompt Tidy, Text With Context."
        },
        {
            "author": "justUmen",
            "title": "Bjornulf_custom_nodes",
            "reference": "https://github.com/justUmen/Bjornulf_custom_nodes",
            "files": [
                "https://github.com/justUmen/Bjornulf_custom_nodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Ollama, Green Screen to Transparency, Save image for Bjornulf LobeChat, Text with random Seed, Random line from input, Combine images (Background+Overlay alpha), Image to grayscale (black & white), Remove image Transparency (alpha), Resize Image, ..."
        },
        {
            "author": "jstit",
            "title": "comfyui_custom_node_image",
            "reference": "https://github.com/jstit/comfyui_custom_node_image",
            "files": [
                "https://github.com/jstit/comfyui_custom_node_image"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ImageCropCircle."
        },
        {
            "author": "jstit",
            "title": "ComfyUI-HeadshotPro",
            "reference": "https://github.com/HeadshotPro/ComfyUI-HeadshotPro",
            "files": [
                "https://github.com/HeadshotPro/ComfyUI-HeadshotPro"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Download Dreambooth Checkpoint, Get Random Value From List, Load Canny Pose Face, Transparent to White Background, Download Flux Lora."
        },
        {
            "author": "Isi-dev",
            "title": "UniAnimate Nodes for ComfyUI",
            "id": "comfyuiunianimatenodes",
            "reference": "https://github.com/Isi-dev/ComfyUI-UniAnimate-W",
            "files": [
                "https://github.com/Isi-dev/ComfyUI-UniAnimate-W"
            ],
            "install_type": "git-clone",
            "description": "These are nodes to animate an image with a reference video using UniAnimate. [w/Name conflict with AIFSH/ComfyUI-UniAnimate. Cannot install simulatenously.]"
        },
        {
            "author": "XLabs-AI",
            "title": "x-flux-comfyui",
            "reference": "https://github.com/XLabs-AI/x-flux-comfyui",
            "files": [
                "https://github.com/XLabs-AI/x-flux-comfyui"
            ],
            "install_type": "git-clone",
            "description": "Nodes:Load Flux LoRA, Load Flux ControlNet, Apply Flux ControlNet, Xlabs Sampler"
        },
        {
            "author": "okgo4",
            "title": "ComfyUI-Mosaic-Mask",
            "reference": "https://github.com/okgo4/ComfyUI-Mosaic-Mask",
            "files": [
                "https://github.com/okgo4/ComfyUI-Mosaic-Mask"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Mosaic-Mask is an automatic tool designed to detect and mask mosaic areas in input images."
        },
        {
            "author": "ChrisColeTech",
            "title": "ComfyUI-Line-counter",
            "reference": "https://github.com/ChrisColeTech/ComfyUI-Line-counter",
            "files": [
                "https://github.com/ChrisColeTech/ComfyUI-Line-counter"
            ],
            "install_type": "git-clone",
            "description": "This custom node package for ComfyUI is designed to streamline your workflow with powerful file-counting capabilities."
        },
        {
            "author": "ChrisColeTech",
            "title": "ComfyUI-Elegant-Resource-Monitor",
            "reference": "https://github.com/ChrisColeTech/ComfyUI-Elegant-Resource-Monitor",
            "files": [
                "https://github.com/ChrisColeTech/ComfyUI-Elegant-Resource-Monitor"
            ],
            "install_type": "git-clone",
            "description": "This custom node for ComfyUI will add a simple and elegant resource monitor."
        },
        {
            "author": "dadoirie",
            "title": "ComfyUI_Dados_Nodes",
            "reference": "https://github.com/dadoirie/ComfyUI_Dados_Nodes",
            "files": [
                "https://github.com/dadoirie/ComfyUI_Dados_Nodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI_Dados_Nodes is a collection of custom nodes for ComfyUI, designed to enhance functionality and provide integration with various services, including Pinterest. This privacy policy explains how these nodes handle user data.\nNOTE: [a/privacy_policy](https://mirror.ghproxy.com/https://github.com/dadoirie/ComfyUI_Dados_Nodes/blob/master/privacy_policy.md)"
        },
        {
            "author": "fanfanfan",
            "title": "chinese_clip_encode",
            "id": "chinese_clip_encode",
            "reference": "https://github.com/yuan199696/chinese_clip_encode",
            "files": [
                "https://github.com/yuan199696/chinese_clip_encode"
            ],
            "install_type": "git-clone",
            "description": "Support input of Chinese prompts."
        },
        {
            "author": "fanfanfan",
            "title": "add_text_2_img",
            "id": "add_text_2_img",
            "reference": "https://github.com/yuan199696/add_text_2_img",
            "files": [
                "https://github.com/yuan199696/add_text_2_img"
            ],
            "install_type": "git-clone",
            "description": "Support adding custom text to the generated images."
        },
        {
            "author": "fairy-root",
            "title": "Ollama and Llava Vision integration for ComfyUI",
            "reference": "https://github.com/fairy-root/comfyui-ollama-llms",
            "files": [
                "https://github.com/fairy-root/comfyui-ollama-llms"
            ],
            "install_type": "git-clone",
            "description": "Ollama and Llava vision integration for ComfyUI"
        },
        {
            "author": "fairy-root",
            "title": "Flux Prompt Generator for ComfyUI",
            "reference": "https://github.com/fairy-root/Flux-Prompt-Generator",
            "files": [
                "https://github.com/fairy-root/Flux-Prompt-Generator"
            ],
            "install_type": "git-clone",
            "description": "A flexible and customizable prompt generator for generating detailed and creative prompts for image generation models for ComfyUI"
        },
        {
            "author": "ryanontheinside",
            "title": "RyanOnTheInside",
            "reference": "https://github.com/ryanontheinside/ComfyUI_RyanOnTheInside",
            "files": [
                "https://github.com/ryanontheinside/ComfyUI_RyanOnTheInside"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes introducing particle simulations, optical flow, audio manipulation & reactivity, and temporal masks"
        },
        {
            "author": "ControlAltAI",
            "title": "ControlAltAI Nodes",
            "id": "controlaltai",
            "reference": "https://github.com/gseth/ControlAltAI-Nodes",
            "files":
            [
                "https://github.com/gseth/ControlAltAI-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Quality of Life ComfyUI nodes starting with Flux Resolution Calculator and Flux Sampler."
        },
        {
            "author": "OliverCrosby",
            "title": "ComfyUI Minimap",
            "id": "minimap",
            "reference": "https://github.com/OliverCrosby/Comfyui-Minimap",
            "files": [
                "https://github.com/OliverCrosby/Comfyui-Minimap"
            ],
            "install_type": "git-clone",
            "description": "A simple minimap in the bottom-right of the window showing the full workflow, left click to navigate"
        },
        {
            "author": "doomy23",
            "title": "ComfyUI-D00MYsNodes",
            "reference": "https://github.com/doomy23/ComfyUI-D00MYsNodes",
            "files": [
                "https://github.com/doomy23/ComfyUI-D00MYsNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Images_Converter, Show_Text, Strings_From_List, Save_Text, Random_Images, Load_Images_From_Paths, JSPaint."
        },
        {
            "author": "Sieyalixnet",
            "title": "ComfyUI_Textarea_Loaders",
            "reference": "https://github.com/Sieyalixnet/ComfyUI_Textarea_Loaders",
            "files": [
                "https://github.com/Sieyalixnet/ComfyUI_Textarea_Loaders"
            ],
            "install_type": "git-clone",
            "description": "An easy custom node that makes the some loaders' input as Text instead of file selector.\nFor example, there are many characters in different loras respectively. If you want to generate different characters' pictures, you have to select corresponding lora, and then edit the prompt. It may cost much time.\nTo solve this problem, You can use it with a chrome extension https://mirror.ghproxy.com/https://github.com/Sieyalixnet/ComfyUI-Prompt-Formatter-Extension that makes the queue prompt easier when you dealing with massive loras and prompt."
        },
        {
            "author": "markuryy",
            "title": "ComfyUI Flux Prompt Saver",
            "reference": "https://github.com/markuryy/ComfyUI-Flux-Prompt-Saver",
            "files": [
                "https://github.com/markuryy/ComfyUI-Flux-Prompt-Saver"
            ],
            "install_type": "git-clone",
            "description": "The Flux Prompt Saver is set of simple nodes for saving images generated with Flux with A1111-style metadata."
        },
        {
            "author": "eesahe",
            "title": "ComfyUI-eesahesNodes",
            "reference": "https://github.com/EeroHeikkinen/ComfyUI-eesahesNodes",
            "files": [
                "https://github.com/EeroHeikkinen/ComfyUI-eesahesNodes"
            ],
            "install_type": "git-clone",
            "description": "InstantX's Flux union ControlNet loader and implementation"
        },
        {
            "author": "anhkhoatranle30",
            "title": "Handy Node ComfyUI",
            "id": "handynode",
            "reference": "https://github.com/anhkhoatranle30/Handy-Nodes-ComfyUI",
            "files": [
                "https://github.com/anhkhoatranle30/Handy-Nodes-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "This is a pack with some handy nodes for ComfyUI."
        },
        {
            "author": "Artiprocher",
            "title": "Dashscope FLUX API for ComfyUI",
            "id": "dashscope_api",
            "reference": "https://github.com/modelscope/comfyscope",
            "files": [
                "https://github.com/modelscope/comfyscope"
            ],
            "install_type": "git-clone",
            "description": "The FLUX model API from DashScope, developed by Black Forest Labs, offers superior image generation capabilities with optimized support for Chinese prompts, achieving a commendable tradeoff between performance and the quality of generated images compared to other open-source models."
        },
        {
            "author": "lucafoscili",
            "title": "LF Nodes",
            "reference": "https://github.com/lucafoscili/comfyui-lf",
            "files": [
                "https://github.com/lucafoscili/comfyui-lf"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes with a touch of extra UX, including: history for primitives, JSON manipulation, logic switches with visual feedback, LLM chat... and more!"
        },
        {
            "author": "JPrevots",
            "title": "ComfyUI-PhyCV",
            "reference": "https://github.com/JPrevots/ComfyUI-PhyCV",
            "files": [
                "https://github.com/JPrevots/ComfyUI-PhyCV"
            ],
            "install_type": "git-clone",
            "description": "Nodes:PhyCV - Phase-Stretch Transform (PST), PhyCV - VEViD, PhyCV - Page."
        },
        {
            "author": "rnbwdsh",
            "title": "Latent Walk",
            "reference": "https://github.com/rnbwdsh/ComfyUI-LatentWalk",
            "files": [
                "https://github.com/rnbwdsh/ComfyUI-LatentWalk"
            ],
            "install_type": "git-clone",
            "description": "Latent space walks for latents, conditionals and noise"
        },
        {
            "author": "kudou-reira",
            "title": "ComfyUI_StringToHex",
            "reference": "https://github.com/kasukanra/ComfyUI_StringToHex",
            "files": [
                "https://github.com/kasukanra/ComfyUI_StringToHex"
            ],
            "install_type": "git-clone",
            "description": "This is a simple ComfyUI node that will take in a string of 'color' (i.e. 'blue') and output a hex color format."
        },
        {
            "author": "phyblas",
            "title": "paint-by-example @ ComfyUI",
            "id": "paintbyexample",
            "reference": "https://github.com/phyblas/paint-by-example_comfyui",
            "files": [
                "https://github.com/phyblas/paint-by-example_comfyui"
            ],
            "install_type": "git-clone",
            "description": "Implementation of paint-by-example on ComfyUI"
        },
        {
            "author": "aidenli",
            "title": "ComfyUI_NYJY",
            "id": "NYJY",
            "reference": "https://github.com/aidenli/ComfyUI_NYJY",
            "files": [
                "https://github.com/aidenli/ComfyUI_NYJY"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Translate, JoyTag, JoyCaption."
        },
        {
            "author": "Pseudotools",
            "title": "Pseudocomfy",
            "id": "pseudocomfy",
            "reference": "https://github.com/Pseudotools/Pseudocomfy",
            "files": [
                "https://github.com/Pseudotools/Pseudocomfy"
            ],
            "install_type": "git-clone",
            "description": "A package designed to enable multi-regional prompting for architectural rendering, integrated with the Rhino Pseudorandom plugin."
        },
        {
            "author": "TTPlanetPig",
            "title": "Comfyui_TTP_Toolset",
            "reference": "https://github.com/TTPlanetPig/Comfyui_TTP_Toolset",
            "files": [
                "https://github.com/TTPlanetPig/Comfyui_TTP_Toolset"
            ],
            "install_type": "git-clone",
            "description": "This is a workflow for my simple logic amazing upscale node for DIT model. it can be common use for Flux,Hunyuan,SD3 It can simple tile the initial image into pieces and then use image-interrogator to get each tile prompts for more accurate upscale process. The condition will be properly handled and the hallucination will be significantly eliminated."
        },
        {
            "author": "TTPlanetPig",
            "title": "for comfyui image proprocessor",
            "reference": "https://github.com/TTPlanetPig/Comfyui_TTP_CN_Preprocessor",
            "files": [
                "https://github.com/TTPlanetPig/Comfyui_TTP_CN_Preprocessor"
            ],
            "install_type": "git-clone",
            "description": "Adapt for Hunyuan now\nNOTE: The files in the repo are not organized, which may lead to update issues."
        },
        {
            "author": "TTPlanetPig",
            "title": "Comfyui_JC2",
            "reference": "https://github.com/TTPlanetPig/Comfyui_JC2",
            "files": [
                "https://github.com/TTPlanetPig/Comfyui_JC2"
            ],
            "install_type": "git-clone",
            "description": "Wrapped Joy Caption alpha 2 node for comfyui from [a/https://huggingface.co/spaces/fancyfeast/joy-caption-alpha-two](https://huggingface.co/spaces/fancyfeast/joy-caption-alpha-two) Easy use, for GPU with less 19G, please use nf4 for better balanced speed and result. This Node also took a reference from /chflame163/ComfyUI_LayerStyle and [a/https://huggingface.co/John6666/joy-caption-alpha-two-cli-mod](https://huggingface.co/John6666/joy-caption-alpha-two-cli-mod)"
        },
        {
            "author": "camenduru",
            "title": "ComfyUI-TostAI",
            "reference": "https://github.com/camenduru/ComfyUI-TostAI",
            "files": [
                "https://github.com/camenduru/ComfyUI-TostAI"
            ],
            "install_type": "git-clone",
            "description": "NODES: SendToTostAI"
        },
        {
            "author": "xlinx",
            "title": "ComfyUI-decadetw-auto-prompt-llm",
            "reference": "https://github.com/xlinx/ComfyUI-decadetw-auto-prompt-llm",
            "files": [
                "https://github.com/xlinx/ComfyUI-decadetw-auto-prompt-llm"
            ],
            "install_type": "git-clone",
            "description": "Auto prompt by LLM and LLM-Vision. (Trigger more details hiding in model)"
        },
        {
            "author": "xlinx",
            "title": "ComfyUI-decadetw-auto-messaging-realtime",
            "reference": "https://github.com/xlinx/ComfyUI-decadetw-auto-messaging-realtime",
            "files": [
                "https://github.com/xlinx/ComfyUI-decadetw-auto-messaging-realtime"
            ],
            "install_type": "git-clone",
            "description": "Auto messging sd-image and sd-info to mobile phone IM realtime. (LINE | Telegram | Discord)"
        },
        {
            "author": "xlinx",
            "title": "ComfyUI-decadetw-spout-syphon-im-vj",
            "reference": "https://github.com/xlinx/ComfyUI-decadetw-spout-syphon-im-vj",
            "files": [
                "https://github.com/xlinx/ComfyUI-decadetw-spout-syphon-im-vj"
            ],
            "install_type": "git-clone",
            "description": "I'm SD-VJ. (share SD-generating-process in realtime by gpu)"
        },
        {
            "author": "wmpmiles",
            "title": "ComfyUI-GTF-Utilities",
            "reference": "https://github.com/wmpmiles/ComfyUI-GTF-Utilities",
            "files": [
                "https://github.com/wmpmiles/ComfyUI-GTF-Utilities"
            ],
            "install_type": "git-clone",
            "description": "Generalised 'image' processing nodes for images, masks, latents, and combinations thereof."
        },
        {
            "author": "nonnonstop",
            "title": "comfyui-faster-loading",
            "reference": "https://github.com/nonnonstop/comfyui-faster-loading",
            "files": [
                "https://github.com/nonnonstop/comfyui-faster-loading"
            ],
            "install_type": "git-clone",
            "description": "This extension applies a patch that limits the model loading speed when using an HDD in a Windows environment. See [a/comfyanonymous/ComfyUI#1992](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI/issues/1992). [w/As this patch is only useful in very limited environments, its installation is not recommended under normal circumstances. Memory usage may increase.]"
        },
        {
            "author": "Dr.Jusseaux",
            "title": "Diffusers-in-ComfyUI",
            "reference": "https://github.com/maepopi/Diffusers-in-ComfyUI",
            "files": [
                "https://github.com/maepopi/Diffusers-in-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "A collection of ComfyUI custom nodes that allow to use most Diffusers pipelines and components in Comfy(Txt2Img, Img2Img, Inpainting, LoRAS, B-LoRAS, ControlNet...)"
        },
        {
            "author": "niknah",
            "title": "Quick Connections",
            "id": "quick-connections",
            "reference": "https://github.com/niknah/quick-connections",
            "files": [
                "https://github.com/niknah/quick-connections"
            ],
            "install_type": "git-clone",
            "description": "Quick connections, Circuit board connections"
        },
        {
            "author": "niknah",
            "title": "ComfyUI-F5-TTS",
            "reference": "https://github.com/niknah/ComfyUI-F5-TTS",
            "files": [
                "https://github.com/niknah/ComfyUI-F5-TTS"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI node for to make text to speech audio using F5-TTS [a/https://mirror.ghproxy.com/https://github.com/SWivid/F5-TTS](https://mirror.ghproxy.com/https://github.com/SWivid/F5-TTS)"
        },
        {
            "author": "daryltucker",
            "title": "ComfyUI-LoadFiles",
            "id": "LoadFiles",
            "reference": "https://github.com/daryltucker/ComfyUI-LoadFiles",
            "files": [
                "https://github.com/daryltucker/ComfyUI-LoadFiles"
            ],
            "install_type": "git-clone",
            "description": "The primary goal of these nodes is to provide a way to access files generated by ComfyUI workflows, local to the machine running ComfyUI. These nodes should always return an updated list of files when triggered."
        },
        {
            "author": "X-T-E-R",
            "title": "ComfyUI Easy Civitai (XTNodes)",
            "reference": "https://github.com/X-T-E-R/ComfyUI-EasyCivitai-XTNodes",
            "files": [
                "https://github.com/X-T-E-R/ComfyUI-EasyCivitai-XTNodes"
            ],
            "install_type": "git-clone",
            "description": "Load your model with image previews, or directly download and import Civitai models via URL. This custom ComfyUI node supports Checkpoint, LoRA, and LoRA Stack models, offering features like bypass options."
        },
        {
            "author": "hyejinlee12",
            "title": "ComfyUI-Fill-Image-for-Outpainting",
            "id": "fill-image-for-outpainting",
            "reference": "https://github.com/Lhyejin/ComfyUI-Fill-Image-for-Outpainting",
            "files": [
                "https://github.com/Lhyejin/ComfyUI-Fill-Image-for-Outpainting"
            ],
            "install_type": "git-clone",
            "description": "This node is to fill image for outpainting(inpainting)\nFill image using cv2 methods(cv2_ns, cv2_telea and edge_pad)"
        },
        {
            "author": "yhayano-ponotech",
            "title": "ComfyUI-Fal-API-Flux",
            "reference": "https://github.com/yhayano-ponotech/ComfyUI-Fal-API-Flux",
            "files": [
                "https://github.com/yhayano-ponotech/ComfyUI-Fal-API-Flux"
            ],
            "install_type": "git-clone",
            "description": "This repository contains custom nodes for ComfyUI that integrate the fal.ai FLUX.1 [dev] with LoRA API, specifically for text-to-image generation. These nodes allow you to use the FLUX.1 model directly within your ComfyUI workflows."
        },
        {
            "author": "Rvage0815",
            "title": "ComfyUI-RvTools",
            "reference": "https://github.com/Rvage0815/ComfyUI-RvTools",
            "files": [
                "https://github.com/Rvage0815/ComfyUI-RvTools"
            ],
            "install_type": "git-clone",
            "description": "With this suit, you have a lot of nodes in one place to build your workflow."
        },
        {
            "author": "erosDiffusion",
            "title": "Compositor Node",
            "reference": "https://github.com/erosDiffusion/ComfyUI-enricos-nodes",
            "files": [
                "https://github.com/erosDiffusion/ComfyUI-enricos-nodes"
            ],
            "install_type": "git-clone",
            "description": "pass up to 8 images and visually place, rotate and scale them to build the perfect composition. group move and group rescale. remember their position and scaling value across generations to easy swap images. use the buffer zone to to park an asset you don't want to use or easily reach transformations controls"
        },
        {
            "author": "Steudio",
            "title": "ComfyUI_Steudio",
            "reference": "https://github.com/Steudio/ComfyUI_Steudio",
            "files": [
                "https://github.com/Steudio/ComfyUI_Steudio"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Make_Tile_Calc, Make_Tiles, Unmake_Tiles, Make_Size, Make_Size_Latent"
        },
        {
            "author": "Assistant",
            "title": "ComfyUI-PromptList",
            "reference": "https://github.com/NakamuraShippo/ComfyUI-PromptList",
            "files": [
                "https://github.com/NakamuraShippo/ComfyUI-PromptList"
            ],
            "install_type": "git-clone",
            "description": "Custom node to manage prompts in YAML format."
        },
        {
            "author": "Assistant",
            "title": "ComfyUI-NS-ManySliders",
            "reference": "https://github.com/NakamuraShippo/ComfyUI-NS-ManySliders",
            "files": [
                "https://github.com/NakamuraShippo/ComfyUI-NS-ManySliders"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-NS-ManySliders is a custom node developed for ComfyUI that allows you to manipulate values using multiple sliders. With this node, you can easily adjust numerous numerical parameters intuitively, making it useful for various purposes."
        },
        {
            "author": "nux1111",
            "title": "ComfyUI_NetDist_Plus",
            "reference": "https://github.com/nux1111/ComfyUI_NetDist_Plus",
            "files": [
                "https://github.com/nux1111/ComfyUI_NetDist_Plus"
            ],
            "install_type": "git-clone",
            "description": "Run ComfyUI workflows on multiple local GPUs/networked machines with options to edit the json values within comfyui.\nOriginal repo: [a/city96/ComfyUI_NetDist](https://mirror.ghproxy.com/https://github.com/city96/ComfyUI_NetDist)"
        },
        {
            "author": "mittimi",
            "title": "ComfyUI_mittimiLoadPreset2",
            "id": "comfyui_mittimi_load_preset2",
            "reference": "https://github.com/mittimi/ComfyUI_mittimiLoadPreset2",
            "files": [
                "https://github.com/mittimi/ComfyUI_mittimiLoadPreset2"
            ],
            "install_type": "git-clone",
            "description": "This node can easily switch between models and prompts by saving presets."
        },
        {
            "author": "mittimi",
            "title": "ComfyUI_mittimiRecalculateSize",
            "id": "comfyui_mittimi_recalculate_size",
            "reference": "https://github.com/mittimi/ComfyUI_mittimiRecalculateSize",
            "files": [
                "https://github.com/mittimi/ComfyUI_mittimiRecalculateSize"
            ],
            "install_type": "git-clone",
            "description": "This is the node that performs the magnification calculation."
        },
        {
            "author": "mittimi",
            "title": "ComfyUI_mittimiWidthHeight",
            "id": "comfyui_mittimi_width_height",
            "reference": "https://github.com/mittimi/ComfyUI_mittimiWidthHeight",
            "files": [
                "https://github.com/mittimi/ComfyUI_mittimiWidthHeight"
            ],
            "install_type": "git-clone",
            "description": "This node can easily switch between vertical and horizontal values with a single button."
        },
        {
            "author": "RodrigoSKohl",
            "title": "Panoramic Image Stitcher",
            "reference": "https://github.com/RodrigoSKohl/ComfyUI-Panoramic-ImgStitcher",
            "files": [
                "https://github.com/RodrigoSKohl/ComfyUI-Panoramic-ImgStitcher"
            ],
            "install_type": "git-clone",
            "description": "Simple Node to make panoramic images"
        },
        {
            "author": "nicehero",
            "title": "comfyui-SegGPT",
            "reference": "https://github.com/nicehero/comfyui-SegGPT",
            "files": [
                "https://github.com/nicehero/comfyui-SegGPT"
            ],
            "install_type": "git-clone",
            "description": "SegGPT model for comfyui,segmentation everything with mask prompt. Download (https://huggingface.co/BAAI/SegGPT/blob/main/seggpt_vit_large.pth) in this node path."
        },
        {
            "author": "sakura1bgx",
            "title": "ComfyUI_FlipStreamViewer",
            "reference": "https://github.com/sakura1bgx/ComfyUI_FlipStreamViewer",
            "files": [
                "https://github.com/sakura1bgx/ComfyUI_FlipStreamViewer"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI_FlipStreamViewer is a tool that provides a viewer interface for flipping images with frame interpolation, allowing you to watch high-fidelity pseudo-videos without needing AnimateDiff."
        },
        {
            "author": "ducido",
            "title": "ObjectFusion_ComfyUI_nodes",
            "id": "objectfusion-nodes",
            "reference": "https://github.com/ducido/ObjectFusion_ComfyUI_nodes",
            "files": [
                "https://github.com/ducido/ObjectFusion_ComfyUI_nodes"
            ],
            "install_type": "git-clone",
            "description": "This is a node to generate new image that combine 2 objects from different scene."
        },
        {
            "author": "DanielHabib",
            "title": "ComfyUI-Voxels",
            "reference": "https://github.com/DanielHabib/ComfyUI-Voxels",
            "files": [
                "https://github.com/DanielHabib/ComfyUI-Voxels"
            ],
            "install_type": "git-clone",
            "description": "NODES:Mesh To Voxel, Voxel Block Saver, Voxel Viewer, Voxel Block Loader, Voxel Video Viewer, Voxel Blocks Into Voxel Video, Voxel Video Preview, Voxelize Mesh, ..."
        },
        {
            "author": "jsonL",
            "title": "ComfyUI-tagger",
            "id": "comfyui-tagger",
            "reference": "https://github.com/StarMagicAI/comfyui_tagger",
            "files": [
                "https://github.com/StarMagicAI/comfyui_tagger"
            ],
            "install_type": "git-clone",
            "description": "Nodes to use Florence2 VLM for image vision tasks: object detection, captioning, segmentation and ocr"
        },
        {
            "author": "boredofnames",
            "title": "ComfyUI-ntfy",
            "reference": "https://github.com/boredofnames/ComfyUI-ntfy",
            "files": [
                "https://github.com/boredofnames/ComfyUI-ntfy"
            ],
            "install_type": "git-clone",
            "description": "NODES:Save Image and ntfy"
        },
        {
            "author": "Xclbr7",
            "title": "ComfyUI-Merlin: Magic Photo Prompter",
            "reference": "https://github.com/Xclbr7/ComfyUI-Merlin",
            "files": [
                "https://github.com/Xclbr7/ComfyUI-Merlin"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Merlin is a custom node extension for ComfyUI, introducing the Magic Photo Prompter. This powerful tool enhances your prompt engineering process by allowing users to easily construct detailed, high-quality prompts for photo-realistic image generation."
        },
        {
            "author": "microbote",
            "title": "StyledCLIPTextEncode",
            "reference": "https://github.com/microbote/ComfyUI-StyledCLIPTextEncode",
            "files": [
                "https://github.com/microbote/ComfyUI-StyledCLIPTextEncode"
            ],
            "install_type": "git-clone",
            "description": "StyledCLIPTextEncode is a node that enables you to build your prompts(both postive and negative) based on the selected style. It provides up-to 77 styles currently and has been tested on SDXL and SD1.5 checkpoints. It's ported from project [a/Style Selector for SDXL 1.0](https://mirror.ghproxy.com/https://github.com/ahgsql/StyleSelectorXL), which is only availabe on WebUI."
        },
        {
            "author": "Isi-dev",
            "title": "ComfyUI-Img2DrawingAssistants",
            "id": "Img2DrawingAssistants",
            "reference": "https://github.com/Isi-dev/ComfyUI-Img2DrawingAssistants",
            "files": [
                "https://github.com/Isi-dev/ComfyUI-Img2DrawingAssistants"
            ],
            "install_type": "git-clone",
            "description": "These are ComfyUI nodes to assist in converting an image to sketches or lineArts."
        },
        {
            "author": "tianguangliu",
            "title": "comfyui-utools",
            "id": "utools",
            "reference": "https://github.com/tianguanggliu/Utools",
            "files": [
                "https://github.com/tianguanggliu/Utools"
            ],
            "install_type": "git-clone",
            "description": "Efficiency tools, Personalized style, Other Nodes, ..."
        },
        {
            "author": "celoron",
            "title": "ComfyUI-VisualQueryTemplate",
            "reference": "https://github.com/celoron/ComfyUI-VisualQueryTemplate",
            "files": [
                "https://github.com/celoron/ComfyUI-VisualQueryTemplate"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI node for transforming images into descriptive text using templated visual question answering. Leverages Hugging Face's VQA models with transformers"
        },
        {
            "author": "Alex Genovese",
            "title": "Huggingface Api Serverless",
            "reference": "https://github.com/alexgenovese/ComfyUI_HF_Servelress_Inference",
            "files": [
                "https://github.com/alexgenovese/ComfyUI_HF_Servelress_Inference"
            ],
            "install_type": "git-clone",
            "description": "Huggingface Api Serverless request"
        },
        {
            "author": "freelifehacker",
            "title": "ComfyUI-ImgMask2PNG",
            "reference": "https://github.com/freelifehacker/ComfyUI-ImgMask2PNG",
            "files": [
                "https://github.com/freelifehacker/ComfyUI-ImgMask2PNG"
            ],
            "install_type": "git-clone",
            "description": "NODES:ImageMask2PNG"
        },
        {
            "author": "souki202",
            "title": "ComfyUI-LoadImage-Advanced",
            "reference": "https://github.com/souki202/ComfyUI-LoadImage-Advanced",
            "files": [
                "https://github.com/souki202/ComfyUI-LoadImage-Advanced"
            ],
            "install_type": "git-clone",
            "description": "This is a node that simply integrates LoadImage, Vae Encode, Upscale, Resolution factor correction, and Color Adjustment."
        },
        {
            "author": "drmbt",
            "title": "comfyui-dreambait-nodes",
            "id": "drmbt",
            "reference": "https://github.com/drmbt/comfyui-dreambait-nodes",
            "files": [
                "https://github.com/drmbt/comfyui-dreambait-nodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of forks, QoL nodes and utilities for ComfyUI"
        },
        {
            "author": "InstaSD",
            "title": "InstaSD nodes for ComfyUI",
            "reference": "https://github.com/WaddingtonHoldings/ComfyUI-InstaSD",
            "files": [
                "https://github.com/WaddingtonHoldings/ComfyUI-InstaSD"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes for use with InstaSD. These nodes will be transformed into app inputs when you deploy your ComfyUI workflow on InstaSD."
        },
        {
            "author": "Shiba-2-shiba",
            "title": "ComfyUI-color-ascii-art-node",
            "id": "comfyui-color-ascii-art-node",
            "reference": "https://github.com/Shiba-2-shiba/comfyui-color-ascii-art-node",
            "files": [
                "https://github.com/Shiba-2-shiba/comfyui-color-ascii-art-node"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node to convert png images into color ASCII art. As noted below, multiple font sizes are used in the specification. The resolution of the generated file is set to be the same as the input image."
        },
        {
            "author": "Shiba-2-shiba",
            "title": "ComfyUI_DiffusionModel_fp8_converter",
            "id": "fp8-converter",
            "reference": "https://github.com/Shiba-2-shiba/ComfyUI_DiffusionModel_fp8_converter",
            "files": [
                "https://github.com/Shiba-2-shiba/ComfyUI_DiffusionModel_fp8_converter"
            ],
            "install_type": "git-clone",
            "description": "This is a custom node to convert only the Diffusion model part or CLIP model part to fp8 in ComfyUI.\nVAE fp8 conversion is not supported.\nThe advantage of this node is that you do not need to separate unet/clip/vae in advance when converting to fp8, but can use the safetenros files that ComfyUI provides."
        },
        {
            "author": "Bao Pham",
            "title": "ComfyUI-LyraVSIH",
            "id": "lyra-vsih",
            "reference": "https://github.com/pbpbpb2705/ComfyUI-LyraVSIH",
            "files": [
                "https://github.com/pbpbpb2705/ComfyUI-LyraVSIH"
            ],
            "install_type": "git-clone",
            "description": "This extension provides a set of nodes that can be used to mask multiple object at once"
        },
        {
            "author": "AbyssBadger0",
            "title": "Kolors Awesome Prompts",
            "reference": "https://github.com/AbyssBadger0/ComfyUI_Kolors_awesome_prompts",
            "files": [
                "https://github.com/AbyssBadger0/ComfyUI_Kolors_awesome_prompts"
            ],
            "install_type": "git-clone",
            "description": "Nodes:KolorsAwesomePrompts"
        },
        {
            "author": "IuvenisSapiens",
            "title": "ComfyUI_Qwen2-VL-Instruct",
            "id": "qwen2-vl-instruct",
            "reference": "https://github.com/IuvenisSapiens/ComfyUI_Qwen2-VL-Instruct",
            "files": [
                "https://github.com/IuvenisSapiens/ComfyUI_Qwen2-VL-Instruct"
            ],
            "install_type": "git-clone",
            "description": "This is an implementation of [a/Qwen2-VL-Instruct](https://mirror.ghproxy.com/https://github.com/QwenLM/Qwen2-VL) by [a/ComfyUI](https://mirror.ghproxy.com/https://github.com/comfyanonymous/ComfyUI), which includes, but is not limited to, support for text-based queries, video queries, single-image queries, and multi-image queries to generate captions or responses."
        },
        {
            "author": "Hmily",
            "title": "ComfyUI-Light-Tool",
            "id": "comfyui-light-tool",
            "reference": "https://github.com/ihmily/ComfyUI-Light-Tool",
            "files": [
                "https://github.com/ihmily/ComfyUI-Light-Tool"
            ],
            "install_type": "git-clone",
            "description": "An awesome light image processing tool nodes for ComfyUI."
        },
        {
            "author": "k-komarov",
            "title": "comfyui-bunny-cdn-storage",
            "reference": "https://github.com/k-komarov/comfyui-bunny-cdn-storage",
            "files": [
                "https://github.com/k-komarov/comfyui-bunny-cdn-storage"
            ],
            "install_type": "git-clone",
            "description": "Save Your Image to BunnyStorage"
        },
        {
            "author": "PabloGFX",
            "title": "Head-Orientation-Node - by PabloGFX",
            "id": "head-orientation-node",
            "reference": "https://github.com/lazniak/Head-Orientation-Node-for-ComfyUI---by-PabloGFX",
            "files": [
                "https://github.com/lazniak/Head-Orientation-Node-for-ComfyUI---by-PabloGFX"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI that analyzes and sorts images based on head orientation using MediaPipe. It detects facial landmarks, calculates head pose, and intelligently sorts images for enhanced AI image processing workflows."
        },
        {
            "author": "PabloGFX",
            "title": "Google Photos Loader - by PabloGFX",
            "id": "google-photos-loader",
            "reference": "https://github.com/lazniak/comfyui-google-photos-loader",
            "files": [
                "https://github.com/lazniak/comfyui-google-photos-loader"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI custom node that integrates Google Photos into your workflows. List albums, load images from specific albums, and search photos directly within ComfyUI. Features customizable image loading options, sorting, and efficient caching for seamless integration of your Google Photos library into AI image processing pipelines."
        },
        {
            "author": "PabloGFX",
            "title": "LiquidTime - by PabloGFX",
            "id": "liquid-time-interpolation",
            "reference": "https://github.com/lazniak/LiquidTime-Interpolation",
            "files": [
                "https://github.com/lazniak/LiquidTime-Interpolation"
            ],
            "install_type": "git-clone",
            "description": "LiquidTime is a simple yet powerful frame interpolation node for ComfyUI. Just input your sequence and desired frame count - the node handles all complex calculations and generates smooth in-between frames for you. A must-have tool for AI animation and video creation that lets you shape time like liquid."
        },
        {
            "author": "45uee",
            "title": "ComfyUI-Color_Transfer",
            "reference": "https://github.com/45uee/ComfyUI-Color_Transfer",
            "files": [
                "https://github.com/45uee/ComfyUI-Color_Transfer"
            ],
            "install_type": "git-clone",
            "description": "Postprocessing nodes that implement color palette transfer for images."
        },
        {
            "author": "Phando",
            "title": "ComfyUI-PhandoNodes",
            "reference": "https://github.com/Phando/ComfyUI-PhandoNodes",
            "files": [
                "https://github.com/Phando/ComfyUI-PhandoNodes"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes to help streamline your ComfyUI workflows"
        },
        {
            "author": "geocine",
            "title": "geocine-comfyui",
            "reference": "https://github.com/geocine/geocine-comfyui",
            "files": [
                "https://github.com/geocine/geocine-comfyui"
            ],
            "install_type": "git-clone",
            "description": "NODES:Image Selector (geocine), Image Scale (geocine)"
        },
        {
            "author": "SeanScripts",
            "title": "ComfyUI-Unload-Model",
            "reference": "https://github.com/SeanScripts/ComfyUI-Unload-Model",
            "files": [
                "https://github.com/SeanScripts/ComfyUI-Unload-Model"
            ],
            "install_type": "git-clone",
            "description": "For unloading a model or all models, using the memory management that is already present in ComfyUI. Copied from [a/https://mirror.ghproxy.com/https://github.com/willblaschko/ComfyUI-Unload-Models](https://mirror.ghproxy.com/https://github.com/willblaschko/ComfyUI-Unload-Models) but without the unnecessary extra stuff."
        },
        {
            "author": "SeanScripts",
            "title": "ComfyUI-PixtralLlamaMolmoVision",
            "reference": "https://github.com/SeanScripts/ComfyUI-PixtralLlamaMolmoVision",
            "files": [
                "https://github.com/SeanScripts/ComfyUI-PixtralLlamaMolmoVision"
            ],
            "install_type": "git-clone",
            "description": "For loading and running Pixtral, Llama 3.2 Vision, and Molmo models. Put models in the models/LLM folder.\n[w/Renamed from ComfyUI-PixtralLlamaVision. Please reinstall.]"
        },
        {
            "author": "ExterminanzHS",
            "title": "Gecco Discord Autosend",
            "reference": "https://github.com/ExterminanzHS/Gecco-Discord-Autosend",
            "files": [
                "https://github.com/ExterminanzHS/Gecco-Discord-Autosend"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for ComfyUI to automatically send generated images to Discord channels."
        },
        {
            "author": "Hugo",
            "title": "ComfyUI-BiRefNet-Hugo",
            "id": "BiRefNet",
            "reference": "https://github.com/MoonHugo/ComfyUI-BiRefNet-Hugo",
            "files": [
                "https://github.com/MoonHugo/ComfyUI-BiRefNet-Hugo"
            ],
            "install_type": "git-clone",
            "description": "This repository wraps the latest BiRefNet model as ComfyUI nodes. Compared to the previous model, the latest model offers higher and better matting accuracy."
        },
        {
            "author": "MoonHugo",
            "title": "ComfyUI-FFmpeg",
            "id": "FFmpeg",
            "reference": "https://github.com/MoonHugo/ComfyUI-FFmpeg",
            "files": [
                "https://github.com/MoonHugo/ComfyUI-FFmpeg"
            ],
            "install_type": "git-clone",
            "description": "Encapsulate the commonly used functions of FFmpeg into ComfyUI nodes, making it convenient for users to perform various video processing tasks within ComfyUI."
        },
        {
            "author": "MoonHugo",
            "title": "ComfyUI-StableAudioOpen",
            "id": "stable-audio-open",
            "reference": "https://github.com/MoonHugo/ComfyUI-StableAudioOpen",
            "files": [
                "https://github.com/MoonHugo/ComfyUI-StableAudioOpen"
            ],
            "install_type": "git-clone",
            "description": "The implementation of the audio generation model stable-audio-open in ComfyUI enables ComfyUI to achieve text-to-audio functionality."
        },
        {
           "author": "GrenKain",
           "title": "PixelArt Processing Nodes",
           "id": "gk_pixelart",
           "reference": "https://github.com/GrenKain/PixelArt-Processing-Nodes-for-ComfyUI",
           "files": [
               "https://github.com/GrenKain/PixelArt-Processing-Nodes-for-ComfyUI"
           ],
           "install_type": "git-clone",
           "description": "This repository provides custom nodes for ComfyUI that enable pixel art style image processing, including downscaling, upscaling, color quantization, and resolution adjustments."
        },
        {
            "author": "Trgtuan10",
            "title": "ComfyUI_YoloSegment_Mask",
            "reference": "https://github.com/Trgtuan10/ComfyUI_YoloSegment_Mask",
            "files": [
                "https://github.com/Trgtuan10/ComfyUI_YoloSegment_Mask"
            ],
            "install_type": "git-clone",
            "description": "NODES:Object Mask.\nNOTE:push [a/yolov8x-seg.pt](https://mirror.ghproxy.com/https://github.com/ultralytics/assets/releases/download/v8.2.0/yolov8x-seg.pt) in models/yolo"
        },
        {
            "author": "lldacing",
            "title": "ComfyUI_BiRefNet_ll",
            "reference": "https://github.com/lldacing/ComfyUI_BiRefNet_ll",
            "files": [
                "https://github.com/lldacing/ComfyUI_BiRefNet_ll"
            ],
            "install_type": "git-clone",
            "description": "Sync with version of BiRefNet. NODES:AutoDownloadBiRefNetModel, LoadRembgByBiRefNetModel, RembgByBiRefNet."
        },
        {
            "author": "Tenney95",
            "title": "ComfyUI-NodeAligner",
            "reference": "https://github.com/Tenney95/ComfyUI-NodeAligner",
            "files": [
                "https://github.com/Tenney95/ComfyUI-NodeAligner"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-NodeAligner is a lightweight ComfyUI layout plugin that includes features such as node alignment, distribution, and resizing. This plugin is designed to simplify layout adjustments in visual node editors or custom UI components, making node arrangement more convenient and efficient."
        },
        {
            "author": "VykosX",
            "title": "ControlFlowUtils",
            "reference": "https://github.com/VykosX/ControlFlowUtils",
            "files": [
                "https://github.com/VykosX/ControlFlowUtils"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes to improve flow control and logic + several utilities to enhance capabilities"
        },
        {
            "author": "tachyon-beep",
            "title": "ComfyUI Simple Feed",
            "id": "simplefeed",
            "reference": "https://github.com/tachyon-beep/comfyui-simplefeed",
            "files": [
                "https://github.com/tachyon-beep/comfyui-simplefeed"
            ],
            "install_type": "git-clone",
            "description": "A lightweight image tray forked from Comfy-UI-CustomScripts with simple sorting, positioning and filtering options."
        },
        {
            "author": "alexisrolland",
            "title": "ComfyUI-Phi",
            "reference": "https://github.com/alexisrolland/ComfyUI-Phi",
            "files": [
                "https://github.com/alexisrolland/ComfyUI-Phi"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes to run microsoft/Phi models."
        },
        {
            "author": "LatentRat",
            "title": "comfy_remote_run",
            "reference": "https://github.com/LatentRat/comfy_remote_run",
            "files": [
                "https://github.com/LatentRat/comfy_remote_run"
            ],
            "install_type": "git-clone",
            "description": "Nodes to run nodes on remote ComfyUI instances."
        },
        {
            "author": "kinglord",
            "title": "Prompt Gallery",
            "id": "promptGallery",
            "reference": "https://github.com/Kinglord/ComfyUI_Prompt_Gallery",
            "files": [
                "https://github.com/Kinglord/ComfyUI_Prompt_Gallery"
            ],
            "install_type": "git-clone",
            "description": "New UI on the sidebar that allows for quick and easy navigation of images to help build styles, characters, backgrounds, etc. or even entire random prompts."
        },
        {
            "author": "kinglord",
            "title": "ComfyUI_LoRA_Sidebar",
            "reference": "https://github.com/Kinglord/ComfyUI_LoRA_Sidebar",
            "files": [
                "https://github.com/Kinglord/ComfyUI_LoRA_Sidebar"
            ],
            "install_type": "git-clone",
            "description": "A custom front-end UX node that creates a visual library of all your LoRAs. It's designed to be fast, slim, and make using LoRAs in Comfy a lot more fun for visual users - especially if you have lots of LoRAs. Should make people used to A1111 and other UI heavy platforms feel more at home. If you've got lots of LoRAs, this sidebar could be your new best friend!"
        },
        {
            "author": "alexcong",
            "title": "Qwen2-VL wrapper for ComfyUI",
            "id": "comfyui-qwen2-vl",
            "reference": "https://github.com/alexcong/ComfyUI_QwenVL",
            "files": [
                "https://github.com/alexcong/ComfyUI_QwenVL"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Qwen2-VL wrapper that supports text-based and single-image queries."
        },
        {
            "author": "Bin-sam",
            "title": "DynamicPose-ComfyUI",
            "reference": "https://github.com/Bin-sam/DynamicPose-ComfyUI",
            "files": [
                "https://github.com/Bin-sam/DynamicPose-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "NODES:pose_extraction, Load_reference_unet, Load_denoising_unet, Load_Pose_Guider, Pose_Guider_Encode, DynamicPose_Sampler, load_pose_model, align"
        },
        {
            "author": "Metal3d",
            "title": "Human Parts Detector",
            "id": "human-parts-detector",
            "reference": "https://github.com/metal3d/ComfyUI_Human_Parts",
            "files": [
                "https://github.com/metal3d/ComfyUI_Human_Parts"
            ],
            "install_type": "git-clone",
            "description": "Detect human parts using the DeepLabV3+ ResNet50 model from Keras-io. You can extract hair, arms, legs, and other parts with ease and with small memory usage."
        },
        {
            "author": "Fuwuffy",
            "title": "ComfyUI Visual Area Nodes",
            "id": "comfy-visual-area",
            "reference": "https://github.com/Fuwuffyi/ComfyUI-VisualArea-Nodes",
            "files": [
                "https://github.com/Fuwuffyi/ComfyUI-VisualArea-Nodes"
            ],
            "install_type": "git-clone",
            "description": "This is a collection of nodes created to aid when managing area conditionings."
        },
        {
            "author": "HentaiGirlfriendDotCom",
            "title": "comfyui-highlight-connections",
            "reference": "https://github.com/HentaiGirlfriendDotCom/comfyui-highlight-connections",
            "files": [
                "https://github.com/HentaiGirlfriendDotCom/comfyui-highlight-connections"
            ],
            "install_type": "git-clone",
            "description": "A node that can be dropped into a group. When a node is then clicked within that group, all nodes and connections in that group get greyed out and the connections from the clicked node go bright red."
        },
        {
            "author": "Cyber-BCat",
            "title": "ComfyUI_Auto_Caption",
            "reference": "https://github.com/Cyber-BCat/ComfyUI_Auto_Caption",
            "files": [
                "https://github.com/Cyber-BCat/ComfyUI_Auto_Caption"
            ],
            "install_type": "git-clone",
            "description": "This report contains a 'load many images' node which is going to load the image set by the number of file names from smallest to largest, and the images will no longer be loaded in the wrong order! Setting index=0 makes it load from the first small value (image flie name) image, and index=2 will load them from the second image. Another node 'load images & resize' can resize the image by the first loaded image."
        },
        {
            "author": "cr7Por",
            "title": "ComfyUI_DepthFlow",
            "reference": "https://github.com/cr7Por/ComfyUI_DepthFlow",
            "files": [
                "https://github.com/cr7Por/ComfyUI_DepthFlow"
            ],
            "install_type": "git-clone",
            "description": "comfyui custom node for depthflow\noriginal depthflow website: [a/https://mirror.ghproxy.com/https://github.com/BrokenSource/DepthFlow](https://mirror.ghproxy.com/https://github.com/BrokenSource/DepthFlow)\ncheck this for installation: [a/https://brokensrc.dev/get/](https://brokensrc.dev/get/)"
        },
        {
            "author": "aimerib",
            "title": "ComfyUI-HigherBitDepthSaveImage",
            "reference": "https://github.com/aimerib/ComfyUI_HigherBitDepthSaveImage",
            "files": [
                "https://github.com/aimerib/ComfyUI_HigherBitDepthSaveImage"
            ],
            "install_type": "git-clone",
            "description": "A comfyui node that provides save image with higher bit depth."
        },
        {
            "author": "nchenevey1",
            "title": "comfyui-gimp-nodes",
            "reference": "https://github.com/nchenevey1/comfyui-gimp-nodes",
            "files": [
                "https://github.com/nchenevey1/comfyui-gimp-nodes"
            ],
            "install_type": "git-clone",
            "description": "Provides nodes geared towards using GIMP as a frontend for ComfyUI."
        },
        {
            "author": "MetaGLM",
            "title": "ComfyUI ZhipuAI Platform",
            "id": "zhipuai-platform",
            "reference": "https://github.com/MetaGLM/ComfyUI-ZhipuAI-Platform",
            "files": [
                "https://github.com/MetaGLM/ComfyUI-ZhipuAI-Platform"
            ],
            "pip": ["zhipuai-platform-video"],
            "install_type": "git-clone",
            "description": "This platform extension provides ZhipuAI nodes, enabling you to configure a workflow for online video generation."
        },
        {
            "author": "zhiselfly",
            "title": "ComfyUI-Alimama-ControlNet-compatible",
            "reference": "https://github.com/zhiselfly/ComfyUI-Alimama-ControlNet-compatible",
            "files": [
                "https://github.com/zhiselfly/ComfyUI-Alimama-ControlNet-compatible"
            ],
            "install_type": "git-clone",
            "description": "Compatible with alimama's SD3-ControlNet Demo on ComfyUI."
        },
        {
            "author": "pydn",
            "title": "ComfyUI to Python Extension",
            "id": "comfyui-to-python-extension",
            "reference": "https://github.com/pydn/ComfyUI-to-Python-Extension",
            "files": [
                "https://github.com/pydn/ComfyUI-to-Python-Extension"
            ],
            "install_type": "git-clone",
            "description": "This custom node allows you to generate pure python code from your ComfyUI workflow with the click of a button. Great for rapid experimentation or production deployment."
        },
        {
            "author": "Dayuppy",
            "title": "Discord Webhook",
            "id": "DiscordWebhook",
            "reference": "https://github.com/Dayuppy/ComfyUI-DiscordWebhook",
            "files": [
                "https://github.com/Dayuppy/ComfyUI-DiscordWebhook"
            ],
            "install_type": "git-clone",
            "description": "A very simple Discord webhook integration node for ComfyUI that lets you post images and text."
        },
        {
            "author": "NyaamZ",
            "title": "Efficiency Nodes ExtendeD",
            "id": "efficiency-ed",
            "reference": "https://github.com/NyaamZ/efficiency-nodes-ED",
            "files": [
                "https://github.com/NyaamZ/efficiency-nodes-ED"
            ],
            "install_type": "git-clone",
            "description": "Expansion of Efficiency Nodes for ComfyUI. Significant UX improvements.[w/NOTE: This node requires [a/Efficiency Nodes for ComfyUI Version 2.0+](https://mirror.ghproxy.com/https://github.com/jags111/efficiency-nodes-comfyui) and [a/ComfyUI-Custom-Scripts](https://mirror.ghproxy.com/https://github.com/pythongosssss/ComfyUI-Custom-Scripts). Also, this node makes changes to user.css.]"
        },
        {
            "author": "NyaamZ",
            "title": "ComfyUI ImageGallery ExtendeD",
            "id": "image-gallery-ed",
            "reference": "https://github.com/NyaamZ/ComfyUI-ImageGallery-ED",
            "files": [
                "https://github.com/NyaamZ/ComfyUI-ImageGallery-ED"
            ],
            "install_type": "git-clone",
            "description": "Custom javascript extensions for better UX for ComfyUI. Double click on image to open. It's convenient for checking images."
        },
        {
            "author": "chrissy0",
            "title": "chris-comfyui-nodes",
            "reference": "https://github.com/chrissy0/chris-comfyui-nodes",
            "files": [
                "https://github.com/chrissy0/chris-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "This repository contains a custom node for ComfyUI that pads an image to be square, filling the new pixels black."
        },
        {
            "author": "amorano",
            "title": "Cozy Link Toggle",
            "id": "cozyLinkToggle",
            "reference": "https://github.com/cozy-comfyui/cozy_link_toggle",
            "files": [
                "https://github.com/cozy-comfyui/cozy_link_toggle"
            ],
            "install_type": "git-clone",
            "description": "Toggle ComfyUI Graph Links On/Off from the Menu Bar. Provides an easy example on how to register and use the ComfyUI menubar extensions."
        },
        {
            "author": "revirevy",
            "title": "Comfyui_saveimage_imgbb",
            "id": "Comfyui_saveimage_imgbb",
            "reference": "https://github.com/revirevy/Comfyui_saveimage_imgbb",
            "files": [
                "https://github.com/revirevy/Comfyui_saveimage_imgbb"
            ],
            "install_type": "git-clone",
            "description": "This custom node allow upscaling an image by a factor using a model."
        },
        {
            "author": "Kinglord",
            "title": "ComfyUI_Slider_Sidebar",
            "reference": "https://github.com/Kinglord/ComfyUI_Slider_Sidebar",
            "files": [
                "https://github.com/Kinglord/ComfyUI_Slider_Sidebar"
            ],
            "install_type": "git-clone",
            "description": "A custom node that adds a UI element to the sidebar allowing easy access, navigation, and use of a massive collection (100+) of LECO (Slider) LoRAs. LECOs are an amazing tool to generate variance in your output with a minimal impact to consistency, i.e. deviating form your prompt. They can also allow you access to control parts of your image without taking up CLIP space, saving your token weights for more valuable keywords. If you haven't used them, there's never been a better time to try!"
        },
        {
            "author": "Isi-dev",
            "title": "Image to Painting and Inspyrenet Assistant Nodes",
            "id": "ComfyUI-Img2PaintingAssistant",
            "reference": "https://github.com/Isi-dev/ComfyUI-Img2PaintingAssistant",
            "files": [
                "https://github.com/Isi-dev/ComfyUI-Img2PaintingAssistant"
            ],
            "install_type": "git-clone",
            "description": "These are ComfyUI nodes to assist in converting images to paintings and to assist the Inspyrenet Rembg node to totally remove, or replace with a color, the original background from images so that the background does not reappear in videos or in nodes that do not retain the alpha channel in rgba images."
        },
        {
            "author": "311-code",
            "title": "ComfyUI MagicClip_Strength for SDXL",
            "reference": "https://github.com/311-code/ComfyUI-MagicClip_Strength",
            "files": [
                "https://github.com/311-code/ComfyUI-MagicClip_Strength"
            ],
            "install_type": "git-clone",
            "description": "This project allows you to adjust SDXL's two text encoder's strengths individually for clip_g (ViT-bigG) and clip_l (CLIP-ViT-L) within ComfyUI. (And other adjustments)"
        },
        {
            "author": "godmt",
            "title": "ComfyUI-List-Utils",
            "reference": "https://github.com/godmt/ComfyUI-List-Utils",
            "files": [
                "https://github.com/godmt/ComfyUI-List-Utils"
            ],
            "install_type": "git-clone",
            "description": "LIST and BATCH utilities which support: create, convert, get or slice items"
        },
        {
            "author": "pedrogengo",
            "title": "ComfyUI-LumaAI-API",
            "id": "lumaai-api",
            "reference": "https://github.com/lumalabs/ComfyUI-LumaAI-API",
            "files": [
                "https://github.com/lumalabs/ComfyUI-LumaAI-API"
            ],
            "install_type": "git-clone",
            "description": "Luma Dream Machine API official ComfyUI custom node."
        },
        {
            "author": "mingsky",
            "title": "ComfyUI-MingNodes",
            "id": "ComfyUI_MingNodes_Mingsky",
            "reference": "https://github.com/mingsky-ai/ComfyUI-MingNodes",
            "files": [
                "https://github.com/mingsky-ai/ComfyUI-MingNodes"
            ],
            "install_type": "git-clone",
            "description": "Nodes: ConvertGrayChannelNode, AdjustBrightnessContrastSaturationNode, BaiduTranslateNode."
        },
        {
            "author": "blob8",
            "title": "ComfyUI_sloppy-comic",
            "reference": "https://github.com/blob8/ComfyUI_sloppy-comic",
            "files": [
                "https://github.com/blob8/ComfyUI_sloppy-comic"
            ],
            "install_type": "git-clone",
            "description": "Using IPAdapter for style consistency, the node accepts a story structured as text {prompt} text {prompt} etc. and generates a comic, saving it to /output. It also adds LLM API Request node, providing an openai compatible LLM API for generating the stories."
        },
        {
            "author": "banqingyuan",
            "title": "ComfyUI-text-replace",
            "reference": "https://github.com/banqingyuan/ComfyUI-text-replace",
            "files": [
                "https://github.com/banqingyuan/ComfyUI-text-replace"
            ],
            "install_type": "git-clone",
            "description": "NODES: OCR Location Node, Image Erase Node, Chat Overlay Node, Extract JSON Node."
        },
        {
            "author": "edelvarden",
            "title": "ComfyUI-ImageMetadataExtension",
            "reference": "https://github.com/edelvarden/ComfyUI-ImageMetadataExtension",
            "files": [
                "https://github.com/edelvarden/ComfyUI-ImageMetadataExtension"
            ],
            "install_type": "git-clone",
            "description": "Custom node for ComfyUI. It adds additional metadata for saved images, ensuring compatibility with the Civitai website."
        },
        {
            "author": "dfghsdh",
            "title": "ComfyUI_FluxPromptGen",
            "reference": "https://github.com/dfghsdh/ComfyUI_FluxPromptGen",
            "files": [
                "https://github.com/dfghsdh/ComfyUI_FluxPromptGen"
            ],
            "install_type": "git-clone",
            "description": "Flux Prompt Generator is a custom node set for ComfyUI that enhances prompt generation and image captioning capabilities. It integrates advanced language models and image captioning techniques to provide versatile and powerful prompt manipulation tools for your AI image generation workflows.\nNOTE:PORT OF [a/https://huggingface.co/Aitrepreneur/FLUX-Prompt-Generator](https://huggingface.co/Aitrepreneur/FLUX-Prompt-Generator) for COMFYUI"
        },
        {
            "author": "liushuchun",
            "title": "ComfyUI_Lora_List_With_Url_Loader",
            "reference": "https://github.com/liushuchun/ComfyUI_Lora_List_With_Url_Loader",
            "files": [
                "https://github.com/liushuchun/ComfyUI_Lora_List_With_Url_Loader"
            ],
            "install_type": "git-clone",
            "description": "Nodes:ComfyUI_Lora_List_With_Url_Loader. Load  loras from  urls and auto fetch them on web if they are missing."
        },
        {
            "author": "silveroxides",
            "title": "Model and Checkpoint Loaders for NF4 and FP4",
            "reference": "https://github.com/silveroxides/ComfyUI_bnb_nf4_fp4_Loaders",
            "files": [
                "https://github.com/silveroxides/ComfyUI_bnb_nf4_fp4_Loaders"
            ],
            "install_type": "git-clone",
            "description": "Nodes for loading both Checkpoints and UNET/Diffussion models quantized to bitsandbytes NF4 or FP4 format.\nStill under development and some limitations such as using LoRA might apply still."
        },
        {
            "author": "turkyden",
            "title": "ComfyUI-SmartCrop",
            "reference": "https://github.com/turkyden/ComfyUI-SmartCrop",
            "files": [
                "https://github.com/turkyden/ComfyUI-SmartCrop"
            ],
            "install_type": "git-clone",
            "description": "a ComfyUI Custom Node for [a/smartcrop.py](https://mirror.ghproxy.com/https://github.com/smartcrop/smartcrop.py)"
        },
        {
            "author": "DareFail",
            "title": "ComfyUI-Roboflow",
            "reference": "https://github.com/DareFail/ComfyUI-Roboflow",
            "files": [
                "https://github.com/DareFail/ComfyUI-Roboflow"
            ],
            "install_type": "git-clone",
            "description": "This is a ComfyUI node that connects with [a/Roboflow workflows](https://roboflow.com/workflows/build).\nRoboflow hosts hundreds of thousands of open source and custom object detection models."
        },
        {
            "author": "valofey",
            "title": "OpenRouter Node",
            "reference": "https://github.com/valofey/Openrouter-Node",
            "files": [
                "https://github.com/valofey/Openrouter-Node"
            ],
            "install_type": "git-clone",
            "description": "This is a node to use OpenRouter API from within ComfyUI. It supports both prompt and image+prompt requests (for multimodal LLMs)."
        },
        {
            "author": "Charlweed",
            "title": "ImageTransceiver - ComfyUI",
            "reference": "https://github.com/Charlweed/image_transceiver",
            "files": [
                "https://github.com/Charlweed/image_transceiver"
            ],
            "install_type": "git-clone",
            "description": "ImageTransceiver is a custom node that enables image generating clients to connect directly to ComfyUI, and send those images in near real-time. For example, an image manipulation program like GIMP can connect an image to a workflow in ComfyUI, and every time the image changes in GIMP, the changes are immediately made in the workflow. Cloning"
        },
        {
            "author": "tanglaoya321",
            "title": "ComfyUI-StoryMaker",
            "reference": "https://github.com/tanglaoya321/ComfyUI-StoryMaker",
            "files": [
                "https://github.com/tanglaoya321/ComfyUI-StoryMaker"
            ],
            "install_type": "git-clone",
            "description": "NODES:StoryMakerSinglePortraitNode, StoryMakerTwoPortraitNode, StoryMakerSwapClothNode."
        },
        {
            "author": "CRT",
            "title": "CRT-Nodes",
            "id": "CRT-Nodes",
            "reference": "https://github.com/plugcrypt/CRT-Nodes",
            "files": [
                "https://github.com/plugcrypt/CRT-Nodes"
            ],
            "install_type": "git-clone",
            "description": "This set includes toggle nodes for Lora Unet blocks L1/L2 and a node to remove trailing comma from string end."
        },
        {
            "author": "GiusTex",
            "title": "ComfyUI-DiffusersImageOutpaint",
            "reference": "https://github.com/GiusTex/ComfyUI-DiffusersImageOutpaint",
            "files": [
                "https://github.com/GiusTex/ComfyUI-DiffusersImageOutpaint"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for outpainting images with diffusers, based on [a/diffusers-image-outpaint](https://huggingface.co/spaces/fffiloni/diffusers-image-outpaint/tree/main) by fffiloni."
        },
        {
          "author": "CY-CHENYUE",
          "title": "ComfyUI-MiniCPM-Plus",
          "id": "minicpm-plus",
          "reference": "https://github.com/CY-CHENYUE/ComfyUI-MiniCPM-Plus",
          "files": [
            "https://github.com/CY-CHENYUE/ComfyUI-MiniCPM-Plus"
          ],
          "install_type": "git-clone",
          "description": "Custom nodes for MiniCPM language models in ComfyUI. Provides advanced text generation and image understanding functions."
        },
        {
             "author": "CY-CHENYUE",
             "title": "ComfyUI-Molmo",
             "id": "comfyui-molmo",
             "reference": "https://github.com/CY-CHENYUE/ComfyUI-Molmo",
             "files": [
               "https://github.com/CY-CHENYUE/ComfyUI-Molmo"
             ],
             "install_type": "git-clone",
             "description": "Use of the molmo model.Generate detailed image descriptions and analysis using Molmo models in ComfyUI."
        },
        {
            "author": "CY-CHENYUE",
            "title": "ComfyUI-InpaintEasy",
            "id": "ComfyUI-InpaintEasy",
            "reference": "https://github.com/CY-CHENYUE/ComfyUI-InpaintEasy",
            "files": [
               "https://github.com/CY-CHENYUE/ComfyUI-InpaintEasy"
            ],
            "description": "InpaintEasy is a set of optimized local repainting (Inpaint) nodes that provide a simpler and more powerful local repainting workflow. It makes local repainting work easier and more efficient with intelligent cropping and merging functions.",
            "tags": ["inpaint", "crop", "image"],
            "install_type": "git-clone"
        },
        {
            "author": "codecringebinge",
            "title": "ComfyUI-Arrow-Key-Canvas-Navigation",
            "id": "codecringebinge.arrow.key.canvas.navigation",
            "reference": "https://github.com/codecringebinge/ComfyUI-Arrow-Key-Canvas-Navigation",
            "files": [
                "https://github.com/codecringebinge/ComfyUI-Arrow-Key-Canvas-Navigation"
            ],
            "install_type": "git-clone",
            "description": "A ComfyUI Custom Node that enables arrow key canvas navigation with a pan speed setting."
        },
        {
            "author": "asaddi",
            "title": "ComfyUI-YALLM-node",
            "reference": "https://github.com/asaddi/ComfyUI-YALLM-node",
            "files": [
                "https://github.com/asaddi/ComfyUI-YALLM-node"
            ],
            "install_type": "git-clone",
            "description": "Yet another set of LLM nodes for ComfyUI (for local/remote OpenAI-like APIs, multi-modal models supported)"
        },
        {
            "author": "ycyy",
            "title": "ComfyUI-YCYY-LoraInfo",
            "reference": "https://github.com/ycyy/ComfyUI-YCYY-LoraInfo",
            "files": [
                "https://github.com/ycyy/ComfyUI-YCYY-LoraInfo"
            ],
            "install_type": "git-clone",
            "description": "You can use this node to get information about lora. For example trigger words, description and example images."
        },
        {
            "author": "Darth-Veitcher",
            "title": "Comfy DV",
            "id": "comfydv",
            "reference": "https://github.com/darth-veitcher/comfydv",
            "files": [
                "https://github.com/darth-veitcher/comfydv"
            ],
            "install_type": "git-clone",
            "description": "Nodes: String Formatting (f-string and jinja2), Random Choice, Model Memory management, and other quality of life improvements."
        },
        {
            "author": "ez-af",
            "title": "ComfyUI-EZ-AF-Nodes",
            "id": "ez-af",
            "reference": "https://github.com/ez-af/ComfyUI-EZ-AF-Nodes",
            "files": [
                "https://github.com/ez-af/ComfyUI-EZ-AF-Nodes"
            ],
            "install_type": "git-clone",
            "description": "This pack helps to conveniently control text in complex prompt-builder type workflows. Load/Read Prompts from .CSV; Concatenate large amounts of text; Use string input as ANY type. Requires pythongosssss custom scripts"
        },
        {
            "author": "danbochman",
            "title": "FASHN Virtual Try-On",
            "id": "fashn",
            "reference": "https://github.com/fashn-AI/ComfyUI-FASHN",
            "files": [
                "https://github.com/fashn-AI/ComfyUI-FASHN"
            ],
            "install_type": "git-clone",
            "description": "Node for the FASHN Virtual Try-On API. Requires an API Key from fashn.ai"
        },
        {
            "author": "BRIA AI",
            "title": "BRIA AI API nodes",
            "reference": "https://github.com/Bria-AI/ComfyUI-BRIA-API",
            "files": [
                "https://github.com/Bria-AI/ComfyUI-BRIA-API"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for ComfyUI using BRIA's API."
        },
        {
            "author": "L.HC",
            "title": "SimpleToolsNodes",
            "reference": "https://github.com/Mcmillian/ComfyUI-SimpleToolsNodes",
            "files": [
                "https://github.com/Mcmillian/ComfyUI-SimpleToolsNodes"
            ],
            "install_type": "git-clone",
            "description": "Two simple nodes: 1. Get the steps based on the model name, 2. Generate prompts using chatglm."
        },
        {
            "author": "creeper",
            "title": "comfyui_nai_api",
            "reference": "https://github.com/Creeper-MZ/comfyui_nai_api",
            "files": [
                "https://github.com/Creeper-MZ/comfyui_nai_api"
            ],
            "install_type": "git-clone",
            "description": "A node that can use Nai in Comfyui"
        },
        {
            "author": "syaofox",
            "title": "ComfyUI_fnodes",
            "reference": "https://github.com/syaofox/ComfyUI_fnodes",
            "files": [
                "https://github.com/syaofox/ComfyUI_fnodes"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI_fnodes is a collection of custom nodes designed for ComfyUI. These nodes provide additional functionality that can enhance your ComfyUI workflows.\nFile manipulation tools, Image resizing tools, IPAdapter tools, Image processing tools, Mask tools, Face analysis tools, Sampler tools, Miscellaneous tools"
        },
        {
            "author": "guyaton",
            "title": "guy-nodes-comfyui",
            "reference": "https://github.com/guyaton/guy-nodes-comfyui",
            "files": [
                "https://github.com/guyaton/guy-nodes-comfyui"
            ],
            "install_type": "git-clone",
            "description": "These are designed to be custom nodes i found usefulness to that hopefully others can share."
        },
        {
            "author": "thoddnn",
            "title": "ComfyUI MLX Nodes",
            "reference": "https://github.com/thoddnn/ComfyUI-MLX",
            "files": [
                "https://github.com/thoddnn/ComfyUI-MLX"
            ],
            "install_type": "git-clone",
            "description": "Faster workflows for ComfyUI users on Mac with Apple silicon"
        },
        {
            "author": "acorderob",
            "title": "Prompt PostProcessor",
            "reference": "https://github.com/acorderob/sd-webui-prompt-postprocessor",
            "files": [
                "https://github.com/acorderob/sd-webui-prompt-postprocessor"
            ],
            "install_type": "git-clone",
            "description": "The Prompt PostProcessor (PPP), formerly known as 'sd-webui-sendtonegative', is an extension designed to process the prompt, possibly after other extensions have modified it."
        },
        {
            "author": "Moooonet",
            "title": "ComfyUI-ArteMoon",
            "reference": "https://github.com/Moooonet/ComfyUI-ArteMoon",
            "files": [
                "https://github.com/Moooonet/ComfyUI-ArteMoon"
            ],
            "install_type": "git-clone",
            "description": "This plugin works with [a/IF_AI_Tools](https://mirror.ghproxy.com/https://github.com/if-ai/ComfyUI-IF_AI_tools) to build a workflow in ComfyUI that uses AI to assist in generating prompts."
        },
        {
            "author": "jetchopper",
            "title": "ComfyUI-GeneraNodes",
            "id": "genera",
            "reference": "https://github.com/evolox/ComfyUI-GeneraNodes",
            "files": [
                "https://github.com/evolox/ComfyUI-GeneraNodes"
            ],
            "install_type": "git-clone",
            "description": "Genera custom nodes and extensions"
        },
        {
            "author": "Nojahhh",
            "title": "ComfyUI GLM-4 Wrapper",
            "reference": "https://github.com/Nojahhh/ComfyUI_GLM4_Wrapper",
            "files": [
                "https://github.com/Nojahhh/ComfyUI_GLM4_Wrapper"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI GLM-4 Wrapper. This powerful tool enhances your prompt engineering process by allowing users to easily construct detailed, high-quality prompts for image/video generation based on user image and/or user prompts."
        },
        {
            "author": "nilor-corp",
            "title": "Nilor Nodes by Nilor Corp",
            "id": "nilor-nodes",
            "reference": "https://github.com/nilor-corp/nilor-nodes",
            "files": [
                "https://github.com/nilor-corp/nilor-nodes"
            ],
            "install_type": "git-clone",
            "description": "Custom utility nodes for ComfyUI by Nilor Corp. Probably not useful for most people, but contains stuff for working with lists, filenames, image batches, etc in a very specifc way."
        },
        {
            "author": "willchil",
            "title": "ComfyUI-Environment-Visualizer",
            "reference": "https://github.com/willchil/ComfyUI-Environment-Visualizer",
            "files": [
                "https://github.com/willchil/ComfyUI-Environment-Visualizer"
            ],
            "install_type": "git-clone",
            "description": "This ComfyUI node pack allows the user to take a panoramic image and a corresponding depth map, and turn them into a 3D environment, which they can view in an immersive WebXR environment."
        },
        {
            "author": "YarvixPA",
            "title": "ComfyUI-NeuralMedia",
            "reference": "https://github.com/YarvixPA/ComfyUI-NeuralMedia",
            "files": [
                "https://github.com/YarvixPA/ComfyUI-NeuralMedia"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes modified to achieve things I felt were missing."
        },
        {
            "author": "SpenserCai",
            "title": "ComfyUI-FunAudioLLM",
            "id": "funaudiollm",
            "reference": "https://github.com/SpenserCai/ComfyUI-FunAudioLLM",
            "files": [
                "https://github.com/SpenserCai/ComfyUI-FunAudioLLM"
            ],
            "install_type": "git-clone",
            "description": "Comfyui custom node for [a/FunAudioLLM](https://funaudiollm.github.io/) include [a/CosyVoice](https://mirror.ghproxy.com/https://github.com/FunAudioLLM/CosyVoice) and [a/SenseVoice](https://mirror.ghproxy.com/https://github.com/FunAudioLLM/SenseVoice)."
        },
        {
            "author": "GadzoinksOfficial",
            "title": "Gadzoinks",
            "reference": "https://github.com/GadzoinksOfficial/gadzoinks_ComfyUI",
            "files": [
                "https://github.com/GadzoinksOfficial/gadzoinks_ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Custom node for integrating with gadzoinks iPhone app"
        },
        {
            "author": "educator-art",
            "title": "ComfyUI-Load-DirectoryFiles",
            "reference": "https://github.com/educator-art/ComfyUI-Load-DirectoryFiles",
            "files": [
                "https://github.com/educator-art/ComfyUI-Load-DirectoryFiles"
            ],
            "install_type": "git-clone",
            "description": "This node loads prompts (txt) and images (png) from a specified directory. By specifying an index, it outputs the selected file."
        },
        {
            "author": "raysers",
            "title": "Mflux-ComfyUI",
            "reference": "https://github.com/raysers/Mflux-ComfyUI",
            "files": [
                "https://github.com/raysers/Mflux-ComfyUI"
            ],
            "install_type": "git-clone",
            "description": "Simple use of [a/Mflux](https://mirror.ghproxy.com/https://github.com/filipstrand/mflux) in ComfyUI, suitable for users who are not familiar with terminal usage.\nNOTE: A MLX port of FLUX based on the Huggingface Diffusers implementation."
        },
        {
            "author": "civen-cn",
            "title": "ComfyUI-PaddleOcr",
            "reference": "https://github.com/civen-cn/ComfyUI-PaddleOcr",
            "files": [
                "https://github.com/civen-cn/ComfyUI-PaddleOcr"
            ],
            "install_type": "git-clone",
            "description": "Nodes related to [a/PaddleOCR](https://paddlepaddle.github.io/PaddleOCR) OCR."
        },
        {
            "author": "rdancer",
            "title": "ComfyUI_Florence2SAM2",
            "reference": "https://github.com/rdancer/ComfyUI_Florence2SAM2",
            "files": [
                "https://github.com/rdancer/ComfyUI_Florence2SAM2"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI custom node implementing Florence 2 + Segment Anything Model 2, based on [a/SkalskiP's HuggingFace space](https://huggingface.co/spaces/SkalskiP/florence-sam)"
        },
        {
            "author": "gelasdev",
            "title": "ComfyUI-FLUX-BFL-API",
            "reference": "https://github.com/gelasdev/ComfyUI-FLUX-BFL-API",
            "files": [
                "https://github.com/gelasdev/ComfyUI-FLUX-BFL-API"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for integrating Flux models with the BFL API."
        },
        {
            "author": "ggarra13",
            "title": "ComfyUI-mrv2",
            "reference": "https://github.com/ggarra13/ComfyUI-mrv2",
            "files": [
                "https://github.com/ggarra13/ComfyUI-mrv2"
            ],
            "install_type": "git-clone",
            "description": "Nodes to interact with the mrv2 player"
        },
        {
            "author": "JichaoLiang",
            "title": "Immortal_comfyUI",
            "reference": "https://github.com/JichaoLiang/Immortal_comfyUI",
            "files": [
                "https://github.com/JichaoLiang/Immortal_comfyUI"
            ],
            "install_type": "git-clone",
            "description": "NODES:ImNewNode, ImAppendNode, MergeNode, SetProperties, SaveToDirectory, batchNodes, redirectToNode, SetEvent, ..."
        },
        {
            "author": "SSsnap",
            "title": "Snap Processing for Comfyui",
            "reference": "https://github.com/SS-snap/ComfyUI-Snap_Processing",
            "files": [
                "https://github.com/SS-snap/ComfyUI-Snap_Processing"
            ],
            "install_type": "git-clone",
            "description": "preprocessing images, presented in a visual way. It also calculates the corresponding image area."
        },
        {
            "author": "RiceRound",
            "title": "ComfyUI Compression and Encryption Node",
            "id": "cryptocat",
            "reference": "https://github.com/RiceRound/ComfyUI_CryptoCat",
            "files": [
                "https://github.com/RiceRound/ComfyUI_CryptoCat"
            ],
            "install_type": "git-clone",
            "description": "a lightweight open-source node for ComfyUI, designed to simplify workflows while providing encryption protection for them."
        },
        {
            "author": "yvann-ba",
            "title": "ComfyUI_Yvann-Nodes",
            "reference": "https://github.com/yvann-ba/ComfyUI_Yvann-Nodes",
            "files": [
                "https://github.com/yvann-ba/ComfyUI_Yvann-Nodes"
            ],
            "install_type": "git-clone",
            "description": "Audio reactivity nodes for AI animations 🔊 Analyze audio, extract drums and vocals. Generate reactive masks and weights. Create audio-driven visuals. Produce weight graphs and audio masks. Compatible with IPAdapter, ControlNets and more. Features audio scheduling and waveform analysis. Tutorials to use this pack: [a/Yvann Youtube](https://www.youtube.com/@yvann.mp4)"
        },
        {
            "author": "Playbook",
            "title": "Playbook Nodes",
            "id": "playbook-3d",
            "reference": "https://github.com/playbook3d/playbook3d-comfyui-nodes",
            "files": [
                "https://github.com/playbook3d/playbook3d-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for connecting 3D scenes and ComfyUI workflows."
        },
        {
            "author": "Elaine-chennn",
            "title": "ComfyUI Overlay Media Node",
            "reference": "https://github.com/Elaine-chennn/comfyui-overlay-media",
            "files": [
                "https://github.com/Elaine-chennn/comfyui-overlay-media"
            ],
            "install_type": "git-clone",
            "description": "This repository contains a custom ComfyUI node for overlaying media using ffmpeg."
        },
        {
            "author": "laogou666",
            "title": "ComfyUI_LG_FFT",
            "reference": "https://github.com/LAOGOU-666/ComfyUI_LG_FFT",
            "files": [
                "https://github.com/LAOGOU-666/ComfyUI_LG_FFT"
            ],
            "install_type": "git-clone",
            "description": "Implementation of Fast Fourier Transform in COMFYUI"
        },
        {
            "author": "VertexStudio",
            "title": "roblox-comfyui-nodes",
            "reference": "https://github.com/VertexStudio/roblox-comfyui-nodes",
            "files": [
                "https://github.com/VertexStudio/roblox-comfyui-nodes"
            ],
            "install_type": "git-clone",
            "description": "NODES:Scale Image Node, Switch Image Node, Switch Text Node, First Number Node, Mirror Effect Node, Text To ImageNode, Flow Nodes, Simple Save Image Node"
        },
        {
            "author": "2kpr",
            "title": "ComfyUI-PMRF",
            "id": "comfyui-pmrf",
            "reference": "https://github.com/2kpr/ComfyUI-PMRF",
            "files": [
                "https://github.com/2kpr/ComfyUI-PMRF"
            ],
            "install_type": "git-clone",
            "description": "Implementation of PMRF on ComfyUI"
        },
        {
            "author": "dionren",
            "title": "Export Workflow With Cyuai Api Available Nodes",
            "id": "comfyUI-Pro-Export-Tool",
            "reference": "https://github.com/dionren/ComfyUI-Pro-Export-Tool",
            "files": [
                "https://github.com/dionren/ComfyUI-Pro-Export-Tool"
            ],
            "install_type": "git-clone",
            "description": "This is a node to convert workflows to cyuai api available nodes."
        },
        {
            "author": "l1yongch1",
            "title": "ComfyUI_PhiCaption",
            "reference": "https://github.com/l1yongch1/ComfyUI_PhiCaption",
            "files": [
                "https://github.com/l1yongch1/ComfyUI_PhiCaption"
            ],
            "install_type": "git-clone",
            "description": "In addition to achieving conventional single-image, single-round reverse engineering, it can also achieve single-image multi-round and multi-image single-round reverse engineering. Moreover, the Phi model has a better understanding of prompts."
        },
        {
            "author": "tkreuziger",
            "title": "ComfyUI and Claude",
            "reference": "https://github.com/tkreuziger/comfyui-claude",
            "files": [
                "https://github.com/tkreuziger/comfyui-claude"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes that are using Anthropic's Claude models for describing images and transforming texts."
        },
        {
            "author": "sipie800",
            "title": "ComfyUI-PuLID-Flux-Enhanced",
            "reference": "https://github.com/sipie800/ComfyUI-PuLID-Flux-Enhanced",
            "files": [
                "https://github.com/sipie800/ComfyUI-PuLID-Flux-Enhanced"
            ],
            "install_type": "git-clone",
            "description": "adapted from [a/https://mirror.ghproxy.com/https://github.com/balazik/ComfyUI-PuLID-Flux](https://mirror.ghproxy.com/https://github.com/balazik/ComfyUI-PuLID-Flux).\ncommon fusion methods for multi-image input, some further experimental fusion methods, switch between using gray image (official) and rgb.,"
        },
        {
            "author": "EvilBT",
            "title": "JoyCaptionAlpha Two for ComfyUI",
            "reference": "https://github.com/EvilBT/ComfyUI_SLK_joy_caption_two",
            "files": [
                "https://github.com/EvilBT/ComfyUI_SLK_joy_caption_two"
            ],
            "install_type": "git-clone",
            "description": "NODES:Joy Caption Two, Joy Caption Two Advanced, Joy Caption Two Load, Joy Caption Extra Options"
        },
        {
            "author": "Q-Bug4",
            "title": "Simple JSON Parser Node for ComfyUI",
            "reference": "https://github.com/Q-Bug4/Comfyui-Simple-Json-Node",
            "files": [
                "https://github.com/Q-Bug4/Comfyui-Simple-Json-Node"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI to parse and extract data from JSON strings."
        },
        {
            "author": "Q-Bug4",
            "title": "Comfyui-Qb-Date-Nodes",
            "reference": "https://github.com/Q-Bug4/Comfyui-Qb-DateNodes",
            "files": [
                "https://github.com/Q-Bug4/Comfyui-Qb-DateNodes"
            ],
            "install_type": "git-clone",
            "description": "A custom node designed for ComfyUI, allowing users to format the current date and time based on a specified format."
        },
        {
            "author": "bartly",
            "title": "Babel Removebg Api Node for ComfyUI",
            "id": "BabelRemovebgApi",
            "reference": "https://github.com/bartly/Comfyui_babel_removebg_api",
            "files": [
                "https://github.com/bartly/Comfyui_babel_removebg_api"
            ],
            "install_type": "git-clone",
            "description": "This is a node to remove background of human picture."
        },
        {
            "author": "NumZ",
            "title": "ComfyUI-FlowChain",
            "id": "FlowChainNode",
            "reference": "https://github.com/numz/Comfyui-FlowChain",
            "files": [
                "https://github.com/numz/ComfyUI-FlowChain"
            ],
            "install_type": "git-clone",
            "description": "Convert your workflows into node and chain them."
        },
        {
            "author": "SozeInc",
            "title": "Quality of Life Nodes for ComfyUI",
            "id": "ComfyUI_Soze",
            "reference": "https://github.com/SozeInc/ComfyUI_Soze",
            "files": [
                "https://github.com/SozeInc/ComfyUI_Soze"
            ],
            "install_type": "git-clone",
            "description": "These nodes aid with batching image processing and maintaining input file names in output files and other quality of life nodes."
        },
        {
            "author": "MzMaXaM",
            "title": "ComfyUi-MzMaXaM",
            "reference": "https://github.com/MzMaXaM/ComfyUi-MzMaXaM",
            "files": [
                "https://github.com/MzMaXaM/ComfyUi-MzMaXaM"
            ],
            "install_type": "git-clone",
            "description": "A pack of nodes(only 2 for now) to make my life easier and hopefully yours ;)"
        },
        {
            "author": "robertvoy",
            "title": "ComfyUI Flux Continuum: Modular Interface",
            "reference": "https://github.com/robertvoy/ComfyUI-Flux-Continuum",
            "files": [
                "https://github.com/robertvoy/ComfyUI-Flux-Continuum"
            ],
            "install_type": "git-clone",
            "description": "NODES:Step Slider, Denoise Slider, Guidance Slider, Batch Slider, Max Shift Slider, ControlNet Slider"
        },
        {
            "author": "Lam Yan",
            "title": "ComfyUI_Lam",
            "id": "ComfyUI_Lam",
            "reference": "https://github.com/yanlang0123/ComfyUI_Lam",
            "files": [
                "https://github.com/yanlang0123/ComfyUI_Lam"
            ],
            "install_type": "git-clone",
            "description": "This extension has some useful nodes, loops, wechat public number +AI chat drawing, distributed cluster"
        },
        {
            "author": "moustafa-nasr",
            "title": "ComfyUI-SimpleLogger",
            "reference": "https://github.com/moustafa-nasr/ComfyUI-SimpleLogger",
            "files": [
                "https://github.com/moustafa-nasr/ComfyUI-SimpleLogger"
            ],
            "install_type": "git-clone",
            "description": "A simple node to save your history in html file. I saves the WorkFlow with all it's input values so you can duplicate it later."
        },
        {
            "author": "sweetndata",
            "title": "ComfyUI-googletrans",
            "reference": "https://github.com/sweetndata/ComfyUI-googletrans",
            "files": [
                "https://github.com/sweetndata/ComfyUI-googletrans"
            ],
            "install_type": "git-clone",
            "description": "NODES:Google Translate"
        },
        {
            "author": "BlackVortexAI",
            "title": "BV Nodes",
            "reference": "https://github.com/BlackVortexAI/ComfyUI-BVortexNodes",
            "files": [
                "https://github.com/BlackVortexAI/ComfyUI-BVortexNodes"
            ],
            "install_type": "git-clone",
            "description": "This repository contains a user-defined node for ComfyUI, currently there are nodes for capturing captions. But will be expanded in the future."
        },
        {
            "author": "JosephThomasParker",
            "title": "ComfyUI-DrawThingsWrapper",
            "reference": "https://github.com/JosephThomasParker/ComfyUI-DrawThingsWrapper",
            "files": [
                "https://github.com/JosephThomasParker/ComfyUI-DrawThingsWrapper"
            ],
            "install_type": "git-clone",
            "description": "These nodes provide a wrapper for calling Draw Things image generations from ComfyUI.\nWait, why? The Draw Things app has been optimized for Apple hardware and runs roughly x3 faster than ComfyUI generations. But ComfyUI is a flexible and powerful tools, and has some features - like queuing and face swapping - that haven't been implemented in Draw Things."
        },
        {
            "author": "Kesin11",
            "title": "ComfyUI-list-filter",
            "reference": "https://github.com/Kesin11/ComfyUI-list-filter",
            "files": [
                "https://github.com/Kesin11/ComfyUI-list-filter"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for convenient filtering image or string lists in ComfyUI workflow."
        },
        {
            "author": "taches-ai",
            "title": "ComfyUI Scene Composer",
            "reference": "https://github.com/taches-ai/comfyui-scene-composer",
            "files": [
                "https://github.com/taches-ai/comfyui-scene-composer"
            ],
            "install_type": "git-clone",
            "description": "A collection of nodes to facilitate the creation of scenes in ComfyUI."
        },
        {
            "author": "NguynHungNguyen",
            "title": "Segment Any Bedroom Interior",
            "reference": "https://github.com/NguynHungNguyen/Segment-Bedroom-Interior",
            "files": [
                "https://github.com/NguynHungNguyen/Segment-Bedroom-Interior"
            ],
            "install_type": "git-clone",
            "description": "Segment Any Bedroom Interior is a Python-based project designed to segment furniture and objects within a bedroom image. The segmentation process uses RGB codes to accurately differentiate between various pieces of furniture, providing a precise mask output for each segmented object. This project is integrated with ComfyUI to allow easy and intuitive usage."
        },
        {
            "author": "MyShell",
            "title": "ComfyUI-ShellAgent-Plugin",
            "id": "comfyui_shellagent_plugin",
            "reference": "https://github.com/myshell-ai/ComfyUI-ShellAgent-Plugin",
            "files": [
                "https://github.com/myshell-ai/ComfyUI-ShellAgent-Plugin"
            ],
            "install_type": "git-clone",
            "description": "This repository provides utility nodes for defining inputs and outputs in ComfyUI workflows. These nodes are essential for running ShellAgent apps with ComfyUI, but they can also be used independently to specify input/output variables and their requirements explicitly."
        },  
        {
            "author": "Vrch Studio (vrch.ai)",
            "title": "ComfyUI Web Viewer",
            "reference": "https://github.com/VrchStudio/comfyui-web-viewer",
            "files": [
                "https://github.com/VrchStudio/comfyui-web-viewer"
            ],
            "install_type": "git-clone",
            "description": "The ComfyUI Web Viewer by [a/vrch.ai](https://vrch.ai) is a custom node collection offering a real-time AI-generated interactive art framework. This utility integrates realtime streaming into ComfyUI workflows, supporting keyboard control nodes, OSC control nodes, sound input nodes, and more. Accessible from any device with a web browser, it enables real time interaction with AI-generated content, making it ideal for interactive visual projects and enhancing ComfyUI workflows with efficient content management and display."
        },
        {
            "author": "kk8bit",
            "title": "KayTool",
            "reference": "https://github.com/kk8bit/KayTool",
            "files": [
                "https://github.com/kk8bit/KayTool"
            ],
            "install_type": "git-clone",
            "description": "KayTool is a growing toolkit for ComfyUI. It includes the 'Custom Save Image' node, allowing image saving in PNG or JPG format, with support for ICC profiles like sRGB and Adobe RGB, metadata control, JPG quality adjustment."
        },
        {
            "author": "sousakujikken",
            "title": "ComfyUI-PixydustQuantizer",
            "reference": "https://github.com/sousakujikken/ComfyUI-PixydustQuantizer",
            "files": [
                "https://github.com/sousakujikken/ComfyUI-PixydustQuantizer"
            ],
            "install_type": "git-clone",
            "description": "Pixydust Quantizer is a custom node extension for ComfyUI that allows for the simplified recreation of tile patterns used in 1990s 16-color PC graphics, offering advanced color quantization and palette optimization features."
        },
        {
            "author": "hoveychen",
            "title": "ComfyUI-MusePose-Remaster",
            "id": "musepose-remaster",
            "reference": "https://github.com/hoveychen/ComfyUI-MusePose-Remaster",
            "files": [
                "https://github.com/hoveychen/ComfyUI-MusePose-Remaster"
            ],
            "install_type": "git-clone",
            "description": "MusePose Remaster is a remaster version of ComfyUI MusePose node.\nIt supports auto weights download, remove most necessary dependencies, etc."
        },
        {
            "author": "AhBumm",
            "title": "Customizable API Call Nodes by BillBum",
            "id": "billbum",
            "reference": "https://github.com/AhBumm/ComfyUI_BillBum_Nodes",
            "files": [
                "https://github.com/AhBumm/ComfyUI_BillBum_Nodes"
            ],
            "nodename_pattern": "\\(BillBum\\)$",
            "install_type": "git-clone",
            "description": "API call node for Third-party platforms both official and local. Support VLMs LLMs Dalle3 Flux-Pro SD3 etc. And some little tools: img to b64 url, b64 url to img, b64 url to b64 data, reg text to word and ',' only, etc."
        },
        {
            "author": "Scepter",
            "title": "ComfyUI-Scepter",
            "id": "scepter",
            "reference": "https://github.com/modelscope/scepter",
            "files": [
                "https://github.com/modelscope/scepter"
            ],
            "install_type": "git-clone",
            "description": "Custom nodes for various visual generation and editing tasks using Scepter."
        },
        {
            "author": "DeemosTech",
            "title": "ComfyUI-Rodin",
            "id": "rodinHyperhuamn",
            "reference": "https://github.com/Ravenmelt/ComfyUI-Rodin",
            "files": [
                "https://github.com/Ravenmelt/ComfyUI-Rodin"
            ],
            "install_type": "git-clone",
            "description": "Comfyui-rodin is a 3D generation extension based on Rodin-API. It provides many of the functionality nodes currently available in RodinAPI and It provides a 3D preview node for ComfyUI."
        },
        {
            "author": "Ardenius",
            "title": "ComfyUI-Ardenius",
            "id": "ARD",
            "reference": "https://github.com/ArdeniusAI/ComfyUI-Ardenius",
            "files": [
                "https://github.com/ArdeniusAI/ComfyUI-Ardenius"
            ],
            "install_type": "git-clone",
            "description": "ARD ComfyUI Ardenius include ARD Control Box, ARD Math nodes and other helper nodes to be added in the future. for more info https://ko-fi.com/ardenius."
        },
        {
            "author": "brayevalerien",
            "title": "ComfyUI Resynthesizer",
            "reference": "https://github.com/brayevalerien/ComfyUI-resynthesizer",
            "files": [
                "https://github.com/brayevalerien/ComfyUI-resynthesizer"
            ],
            "install_type": "git-clone",
            "description": "This repository is a quick port of [a/Resynthesizer](https://mirror.ghproxy.com/https://github.com/bootchk/resynthesizer) to ComfyUI.\nResynthesizer is the open-source implementation of a texture generation technique proposed by Paul Harrison in 2005, especially useful for removing an object from an image (inpainting), which is most likely close to what Photoshop uses to for the content aware fill feature. Note that this is not using a diffusion model to inpaint, as opposed to many techniques of today, which makes it very fast and predictable, but sometimes yields worse results."
        },
        {
            "author": "BZcreativ",
            "title": "ComfyUI-FLUX-TOGETHER-API",
            "reference": "https://github.com/BZcreativ/ComfyUI-FLUX-TOGETHER-API",
            "files": [
                "https://github.com/BZcreativ/ComfyUI-FLUX-TOGETHER-API"
            ],
            "install_type": "git-clone",
            "description": "A custom node implementation for ComfyUI that integrates with Together.ai's FLUX image generation models. This project is inspired by and adapted from [a/ComfyUI-FLUX-BFL-API](https://mirror.ghproxy.com/https://github.com/gelasdev/ComfyUI-FLUX-BFL-API) to work with the Together.ai API."
        },
        {
            "author": "stormcenter",
            "title": "ComfyUI-AutoSplitGridImage",
            "reference": "https://github.com/stormcenter/ComfyUI-AutoSplitGridImage",
            "files": [
                "https://github.com/stormcenter/ComfyUI-AutoSplitGridImage"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-AutoSplitGridImage is a custom node for ComfyUI that provides intelligent image splitting functionality. It combines edge detection for column splits and uniform division for row splits, offering a balanced approach to grid-based image segmentation."
        },
        {
            "author": "stormcenter",
            "title": "ComfyUI LivePhoto Creator",
            "reference": "https://github.com/stormcenter/ComfyUI-LivePhotoCreator",
            "files": [
                "https://github.com/stormcenter/ComfyUI-LivePhotoCreator"
            ],
            "install_type": "git-clone",
            "description": "A custom node for ComfyUI that allows you to create iPhone-compatible Live Photos from videos. This node can convert video sequences into Live Photo format, with the ability to select key frames and customize the output."
        },
        {
            "author": "AkashKarnatak",
            "title": "ComfyUI_faishme",
            "reference": "https://github.com/AkashKarnatak/ComfyUI_faishme",
            "files": [
                "https://github.com/AkashKarnatak/ComfyUI_faishme"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI nodes for our product Faishme"
        },
        {
            "author": "ARZUMATA",
            "title": "ComfyUI-ARZUMATA",
            "reference": "https://github.com/ARZUMATA/ComfyUI-ARZUMATA",
            "files": [
                "https://github.com/ARZUMATA/ComfyUI-ARZUMATA"
            ],
            "install_type": "git-clone",
            "description": "NODES:Caching CLIP Text Encode for FLUX.\nRandom nodes for ComfyUI for various purposes."
        },
        {
            "author": "Rinsanga1",
            "title": "comfyui-florence2xy",
            "reference": "https://github.com/Rinsanga1/comfyui-florence2xy",
            "files": [
                "https://github.com/Rinsanga1/comfyui-florence2xy"
            ],
            "install_type": "git-clone",
            "description": "NODES:Florence2 Coordinates (XY Split)."
        },
        {
            "author": "gt732",
            "title": "ComfyUI-DreamWaltz-G",
            "reference": "https://github.com/gt732/ComfyUI-DreamWaltz-G",
            "files": [
                "https://github.com/gt732/ComfyUI-DreamWaltz-G"
            ],
            "install_type": "git-clone",
            "description": "This repository contains custom ComfyUI nodes designed to integrate with [a/DreamWaltz-G](https://mirror.ghproxy.com/https://github.com/Yukun-Huang/DreamWaltz-G), a cutting-edge model for generating expressive 3D Gaussian avatars using skeleton-guided 2D diffusion."
        },
        {
            "author": "clhui",
            "title": "Clh Tool for ComfyUI",
            "id": "ComfyUi-clh-Tool",
            "reference": "https://github.com/clhui/ComfyUi-clh-Tool",
            "files": [
                "https://github.com/clhui/ComfyUi-clh-Tool"
            ],
            "install_type": "git-clone",
            "description": "Some mathematical calculation nodes，freedom And omnipotent, string calculation nodes, can customize the number of parameters and calculation formulas（expression）. The calculation content can also be displayed in places such as the label title of Comfy Node，String to Image Title Label"
        },
        {
            "author": "ruucm",
            "title": "Ruucm's ComfyUI Nodes",
            "id": "ruucm",
            "reference": "https://github.com/ruucm/ruucm-comfy",
            "nodename_pattern": " \\(ruucm\\)$",
            "files": [
                "https://github.com/ruucm/ruucm-comfy"
            ],
            "install_type": "git-clone",
            "description": "Nodes: Load External LoRA Model Only"
        },
        {
            "author": "Apache0ne",
            "title": "ComfyUI-EasyUrlLoader",
            "id": "easy-url-loader",
            "reference": "https://github.com/Apache0ne/ComfyUI-EasyUrlLoader",
            "files": [
                "https://github.com/Apache0ne/ComfyUI-EasyUrlLoader"
            ],
            "install_type": "git-clone",
            "description": "Simple 4k YT Downloader Through URL"
        },
        {
            "author": "TZOOTZ",
            "title": "TZOOTZ VHS Effect Node",
            "reference": "https://github.com/TZOOTZ/ComfyUI-TZOOTZ_VHS",
            "files": [
                "https://github.com/TZOOTZ/ComfyUI-TZOOTZ_VHS"
            ],
            "pip": ["numpy<2"],
            "install_type": "git-clone",
            "description": "The TZOOTZ VHS Effect Node is designed for multimedia creators who want to blend digital precision with analog imperfection ↔️. Inspired by retro VHS aesthetics, this node lets you apply grain, color bleeding, saturation adjustments, and more, giving any image a touch of analog warmth and noise."
        },
        {
            "author": "jianzhichun",
            "title": "ComfyUI-Easyai",
            "id": "comfyui-easyai",
            "reference": "https://github.com/jianzhichun/ComfyUI-Easyai",
            "files": [
                "https://github.com/jianzhichun/ComfyUI-Easyai"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI-Easyai is a powerful extension for ComfyUI that enables users to share workflows and models to easyai."
        },
        {
            "author": "Isulion",
            "title": "ComfyUI_Isulion Random Prompt Generator",
            "reference": "https://github.com/Isulion/ComfyUI_Isulion",
            "files": [
                "https://github.com/Isulion/ComfyUI_Isulion"
            ],
            "install_type": "git-clone",
            "description": "ComfyUI Nodes that generate prompts with the help of LLM from local or remote Ollama."
        },
        {
            "author": "sneccc",
            "title": "comfyui-snek-nodes",
            "reference": "https://github.com/sneccc/comfyui-snek-nodes",
            "files": [
                "https://github.com/sneccc/comfyui-snek-nodes"
            ],
            "install_type": "git-clone",
            "description": "NODES:Aesthetics, Aesthetics V2, Load AI Toolkit Latent Flux, Send_to_Eagle"
        },
        {
            "author": "mithamunda",
            "title": "OllamaKiller Node for ComfyUI",
            "reference": "https://github.com/mithamunda/ComfyUI-ollama_killer",
            "files": [
                "https://github.com/mithamunda/ComfyUI-ollama_killer"
            ],
            "install_type": "git-clone",
            "description": "OllamaKiller is a utility node for ComfyUI designed to manage VRAM usage more efficiently by automatically terminating the ollama_llama_server.exe process. This is particularly useful for users with limited VRAM, allowing them to clear up memory after running models and improve workflow performance."
        },
        {
            "author": "mithamunda",
            "title": "SD3.5 Empty Latent Size Picker",
            "reference": "https://github.com/mithamunda/ComfyUI-SD3.5-Latent-Size-Picker",
            "files": [
                "https://github.com/mithamunda/ComfyUI-SD3.5-Latent-Size-Picker"
            ],
            "install_type": "git-clone",
            "description": "A utility node for generating empty latent tensors in Stable Diffusion v3.5-compatible resolutions. This node allows for custom batch sizes, width/height overrides, and inverting aspect ratios, ensuring flexibility and compatibility in ComfyUI workflows."
        },
        {
            "author": "jeffrey2212",
            "title": "Pony Character Prompt Picker for ComfyUI",
            "reference": "https://github.com/jeffrey2212/ComfyUI-PonyCharacterPrompt",
            "files": [
                "https://github.com/jeffrey2212/ComfyUI-PonyCharacterPrompt"
            ],
            "install_type": "git-clone",
            "description": "The Pony Character Prompt Picker node reads an Excel file specified by the user, allows manual selection of a tab, and randomly picks a cell value from a specified column, starting from row 3 to the end. The selected value is output as a string to the next node in the ComfyUI workflow."
        },
        {
            "author": "Jonseed",
            "title": "ComfyUI-Detail-Daemon",
            "reference": "https://github.com/Jonseed/ComfyUI-Detail-Daemon",
            "files": [
                "https://github.com/Jonseed/ComfyUI-Detail-Daemon"
            ],
            "install_type": "git-clone",
            "description": "A port of muerrilla's [a/sd-webui-Detail-Daemon](https://mirror.ghproxy.com/https://github.com/muerrilla/sd-webui-detail-daemon) as a node for ComfyUI, to adjust sigmas that control detail."
        },
        {
            "author": "chris-arsenault",
            "title": "ComfyUI-AharaNodes",
            "reference": "https://github.com/chris-arsenault/ComfyUI-AharaNodes",
            "files": [
                "https://github.com/chris-arsenault/ComfyUI-AharaNodes"
            ],
            "install_type": "git-clone",
            "description": "NODES:Frame Segmenter, Get Frame at Index, Repeat Sampler Config, Patch Repeat Sampler Config (Model), Patch Repeat Sampler Config (Latent), KSampler (Simple Input)"
        },
        {
            "author": "mfg637",
            "title": "ComfyUI-ScheduledGuider-Ext",
            "reference": "https://github.com/mfg637/ComfyUI-ScheduledGuider-Ext",
            "files": [
                "https://github.com/mfg637/ComfyUI-ScheduledGuider-Ext"
            ],
            "install_type": "git-clone",
            "description": "NODES:ScheduledCFGGuider, PerpNegScheduledCFGGuider, CosineScheduler, Add zSNR Sigma max, InvertSigmas, ConcatSigmas, OffsetSigmas"
        },
        {
            "author": "changwook987",
            "title": "ComfyUI-Small-Utility",
            "reference": "https://github.com/changwook987/ComfyUI-Small-Utility",
            "files": [
                "https://github.com/changwook987/ComfyUI-Small-Utility"
            ],
            "install_type": "git-clone",
            "description": "Context menu extension for CLIPTextEncode (sort prompt), EmptyLatentImage (sdxl size selector)."
        },
        {
            "author": "OpalSky",
            "title": "OpalSky Nodes",
            "reference": "https://github.com/OpalSky-AI/OpalSky_Nodes",
            "files": [
                "https://github.com/OpalSky-AI/OpalSky_Nodes"
            ],
            "install_type": "git-clone",
            "description": "A set of custom nodes for ComfyUI that provides enhanced string manipulation and prompt variant generation functionality for AI workflows."
        },
        {
            "author": "JustinMatters",
            "title": "ComfyUI JMNodes",
            "id": "JMNodes",
            "reference": "https://github.com/JustinMatters/comfyUI-JMNodes",
            "files": [
                "https://github.com/JustinMatters/comfyUI-JMNodes"
            ],
            "install_type": "git-clone",
            "description": "Provides nodes to support generation of all possible combinations of a set of prompts via boolean logic"
        },
        {
            "author": "lgldlk",
            "title": "ComfyUI-PC-ding-dong",
            "reference": "https://github.com/lgldlk/ComfyUI-PC-ding-dong",
            "files": [
                "https://github.com/lgldlk/ComfyUI-PC-ding-dong"
            ],
            "install_type": "git-clone",
            "description": "Just like when your pizza is ready and the oven goes 'Ding! 🍕', this plugin lets your ComfyUI notify you when your AI creations are done baking!\nA ComfyUI custom node that sends you a friendly 'ding-dong' notification when your workflows are fully cooked and ready to serve. No more staring at the screen waiting - let the AI kitchen tell you when dinner's ready! 👨‍🍳"
        },
        {
            "author": "raspie10032",
            "title": "ComfyUI NAI Prompt Converter",
            "reference": "https://github.com/raspie10032/ComfyUI_RS_NAI_Local_Prompt_converter",
            "files": [
                "https://github.com/raspie10032/ComfyUI_RS_NAI_Local_Prompt_converter"
            ],
            "install_type": "git-clone",
            "description": "A custom node extension for ComfyUI that enables conversion between NovelAI and ComfyUI prompt formats, along with extraction of NovelAI metadata from PNG images."
        },
        {
            "author": "zmwv823",
            "title": "ComfyUI_Ctrlora",
            "reference": "https://github.com/zmwv823/ComfyUI_Ctrlora",
            "files": [
                "https://github.com/zmwv823/ComfyUI_Ctrlora"
            ],
            "install_type": "git-clone",
            "description": "Unofficial custom_node for [a/xyfJASON/ctrlora](https://mirror.ghproxy.com/https://github.com/xyfJASON/ctrlora)."
        },
        {
            "author": "Wakfull33",
            "title": "ComfyUI-SaveImageCivitAI",
            "reference": "https://github.com/Wakfull33/ComfyUI-SaveImageCivitAI",
            "files": [
                "https://github.com/Wakfull33/ComfyUI-SaveImageCivitAI"
            ],
            "install_type": "git-clone",
            "description": "A custom node allowing to save images with CIVITAI readable datas"
        },
        {
            "author": "waterminer",
            "title": "ComfyUI-tagcomplete",
            "reference": "https://github.com/waterminer/ComfyUI-tagcomplete",
            "files": [
                "https://github.com/waterminer/ComfyUI-tagcomplete"
            ],
            "install_type": "git-clone",
            "description": "This extension provides tag completion feature in textbox."
        },
        {
            "author": "grovebadger",
            "title": "Negative Wildcard Processor Node for ComfyUI",
            "id": "neg_wildcard_processor",
            "reference": "https://github.com/GrvBdgr/comfyui-negativewildcardsprocessor",
            "files": [
                "https://github.com/GrvBdgr/comfyui-negativewildcardsprocessor"
            ],
            "install_type": "git-clone",
            "description": "Node to process negative wildcard tokens (<! !>) and move them from the positive prompt to the negative."
        },
        {
            "author": "Apache0ne",
            "title": "SambaNova",
            "id": "SambaNovaAPI",
            "reference": "https://github.com/Apache0ne/SambaNova",
            "files": [
                "https://github.com/Apache0ne/SambaNova"
            ],
            "install_type": "git-clone",
            "description": "Super Fast LLM's llama3.1-405B,70B,8B and more"
        },
        {
            "author": "catboxanon",
            "title": "comfyui_stealth_pnginfo",
            "reference": "https://github.com/catboxanon/comfyui_stealth_pnginfo",
            "files": [
                "https://github.com/catboxanon/comfyui_stealth_pnginfo"
            ],
            "install_type": "git-clone",
            "description": "Fork of [a/sd_webui_stealth_pnginfo](https://mirror.ghproxy.com/https://github.com/ashen-sensored/sd_webui_stealth_pnginfo) with ComfyUI support."
        },
        {
           "author": "dafeng012",
           "title": "comfyui-imgmake",
           "reference": "https://github.com/dafeng012/comfyui-imgmake",
           "files": [
               "https://github.com/dafeng012/comfyui-imgmake"
           ],
           "install_type": "git-clone",
           "description": "This extension integrates ebsynth_utility into comfyui, and I've written some of my own nodes for secondary use."
       },
       {
           "author": "Makki_Shizu",
           "title": "ComfyUI-Prompt-Wildcards",
           "id": "Prompt-Wildcards",
           "reference": "https://github.com/MakkiShizu/ComfyUI-Prompt-Wildcards",
           "files": [
               "https://github.com/MakkiShizu/ComfyUI-Prompt-Wildcards"
           ],
           "install_type": "git-clone",
           "description": "Optional wildcards in ComfyUI"
       },
       {
           "author": "zubenelakrab",
           "title": "ComfyUI-ASV-Nodes Node",
           "id": "ComfyUI-ASV-Nodes",
           "reference": "https://github.com/zubenelakrab/ComfyUI-ASV-Nodes",
           "files": [
               "https://github.com/zubenelakrab/ComfyUI-ASV-Nodes"
           ],
           "install_type": "git-clone",
           "description": "ComfyUI-ASV-Nodes make prompting easier."
       },
       {
           "author": "zubenelakrab",
           "title": "ComfyUI Neural Nodes",
           "reference": "https://github.com/xobiomesh/ComfyUI_xObiomesh",
           "files": [
               "https://github.com/xobiomesh/ComfyUI_xObiomesh"
           ],
           "install_type": "git-clone",
           "description": "An advanced ComfyUI extension that enables multi-agent LLM conversations using Ollama models."
       },















        

        {
            "author": "Ser-Hilary",
            "title": "SDXL_sizing",
            "reference": "https://github.com/Ser-Hilary/SDXL_sizing",
            "files": [
                "https://github.com/Ser-Hilary/SDXL_sizing/raw/main/conditioning_sizing_for_SDXL.py"
            ],
            "install_type": "copy",
            "description": "Nodes:sizing_node. Size calculation node related to image size in prompts supported by SDXL."
        },
        {
            "author": "ailex000",
            "title": "Image Gallery",
            "reference": "https://github.com/ailex000/ComfyUI-Extensions",
            "js_path": "image-gallery",
            "files": [
                "https://github.com/ailex000/ComfyUI-Extensions/raw/main/image-gallery/imageGallery.js"
            ],
            "install_type": "copy",
            "description": "Custom javascript extensions for better UX for ComfyUI. Supported nodes: PreviewImage, SaveImage. Double click on image to open."
        },
        {
            "author": "rock-land",
            "title": "graphNavigator",
            "reference": "https://github.com/rock-land/graphNavigator",
            "js_path": "graphNavigator",
            "files": [
                "https://github.com/rock-land/graphNavigator/raw/main/graphNavigator/graphNavigator.js"
            ],
            "install_type": "copy",
            "description": "ComfyUI Web Extension for saving views and navigating graphs."
        },
        {
            "author": "diffus3",
            "title": "diffus3/ComfyUI-extensions",
            "reference": "https://github.com/diffus3/ComfyUI-extensions",
            "js_path": "diffus3",
            "files": [
                "https://github.com/diffus3/ComfyUI-extensions/raw/main/multiReroute/multireroute.js",
                "https://github.com/diffus3/ComfyUI-extensions/raw/main/setget/setget.js"
            ],
            "install_type": "copy",
            "description": "Extensions: subgraph, setget, multiReroute"
        },
        {
            "author": "m957ymj75urz",
            "title": "m957ymj75urz/ComfyUI-Custom-Nodes",
            "reference": "https://github.com/m957ymj75urz/ComfyUI-Custom-Nodes",
            "js_path": "m957ymj75urz",
            "files": [
                "https://github.com/m957ymj75urz/ComfyUI-Custom-Nodes/raw/main/clip-text-encode-split/clip_text_encode_split.py",
                "https://github.com/m957ymj75urz/ComfyUI-Custom-Nodes/raw/main/colors/colors.js"
            ],
            "install_type": "copy",
            "description": "Nodes: RawText, RawTextCLIPEncode, RawTextCombine, RawTextReplace, Extension: m957ymj75urz.colors"
        },
        {
            "author": "Bikecicle",
            "title": "Waveform Extensions",
            "reference": "https://github.com/Bikecicle/ComfyUI-Waveform-Extensions",
            "files": [
                "https://github.com/Bikecicle/ComfyUI-Waveform-Extensions/raw/main/EXT_AudioManipulation.py",
                "https://github.com/Bikecicle/ComfyUI-Waveform-Extensions/raw/main/EXT_VariationUtils.py"
            ],
            "install_type": "copy",
            "description": "Some additional audio utilites for use on top of Sample Diffusion ComfyUI Extension"
        },
        {
            "author": "dawangraoming",
            "title": "KSampler GPU",
            "reference": "https://github.com/dawangraoming/ComfyUI_ksampler_gpu",
            "files": [
                "https://github.com/dawangraoming/ComfyUI_ksampler_gpu/raw/main/ksampler_gpu.py"
            ],
            "install_type": "copy",
            "description": "KSampler is provided, based on GPU random noise"
        },
        {
            "author": "fitCorder",
            "title": "fcSuite",
            "reference": "https://github.com/fitCorder/fcSuite",
            "files": [
                "https://github.com/fitCorder/fcSuite/raw/main/fcSuite.py"
            ],
            "install_type": "copy",
            "description": "fcFloatMatic is a custom module, that when configured correctly will increment through the lines generating you loras at different strengths. The JSON file will load the config."
        },
        {
            "author": "lordgasmic",
            "title": "Wildcards",
            "reference": "https://github.com/lordgasmic/ComfyUI-Wildcards",
            "files": [
                "https://github.com/lordgasmic/ComfyUI-Wildcards/raw/master/wildcards.py"
            ],
            "install_type": "copy",
            "description": "Nodes:CLIPTextEncodeWithWildcards. This wildcard node is a wildcard node that operates based on the seed."
        },
        {
            "author": "throttlekitty",
            "title": "SDXLCustomAspectRatio",
            "reference": "https://github.com/throttlekitty/SDXLCustomAspectRatio",
            "files": [
                "https://raw.githubusercontent.com/throttlekitty/SDXLCustomAspectRatio/main/SDXLAspectRatio.py"
            ],
            "install_type": "copy",
            "description": "A quick and easy ComfyUI custom node for setting SDXL-friendly aspect ratios."
        },
        {
            "author": "s1dlx",
            "title": "comfy_meh",
            "reference": "https://github.com/s1dlx/comfy_meh",
            "files": [
                "https://github.com/s1dlx/comfy_meh/raw/main/meh.py"
            ],
            "install_type": "copy",
            "description": "Advanced merging methods."
        },
        {
            "author": "tudal",
            "title": "Hakkun-ComfyUI-nodes",
            "reference": "https://github.com/tudal/Hakkun-ComfyUI-nodes",
            "files": [
                "https://github.com/tudal/Hakkun-ComfyUI-nodes/raw/main/hakkun_nodes.py"
            ],
            "install_type": "copy",
            "description": "Nodes: Prompt parser. ComfyUI extra nodes. Mostly prompt parsing."
        },
        {
            "author": "SadaleNet",
            "title": "ComfyUI A1111-like Prompt Custom Node Solution",
            "reference": "https://github.com/SadaleNet/CLIPTextEncodeA1111-ComfyUI",
            "files": [
                "https://github.com/SadaleNet/CLIPTextEncodeA1111-ComfyUI/raw/master/custom_nodes/clip_text_encoder_a1111.py"
            ],
            "install_type": "copy",
            "description": "Nodes: CLIPTextEncodeA1111, RerouteTextForCLIPTextEncodeA1111."
        },
        {
            "author": "wsippel",
            "title": "SDXLResolutionPresets",
            "reference": "https://github.com/wsippel/comfyui_ws",
            "files": [
                "https://github.com/wsippel/comfyui_ws/raw/main/sdxl_utility.py"
            ],
            "install_type": "copy",
            "description": "Nodes: SDXLResolutionPresets. Easy access to the officially supported resolutions, in both horizontal and vertical formats: 1024x1024, 1152x896, 1216x832, 1344x768, 1536x640"
        },
        {
            "author": "nicolai256",
            "title": "comfyUI_Nodes_nicolai256",
            "id": "nicoali256",
            "reference": "https://github.com/nicolai256/comfyUI_Nodes_nicolai256",
            "files": [
                "https://github.com/nicolai256/comfyUI_Nodes_nicolai256/raw/main/yugioh-presets.py"
            ],
            "install_type": "copy",
            "description": "Nodes: yugioh_Presets. by Nicolai256 inspired by throttlekitty SDXLAspectRatio"
        },
        {
            "author": "Onierous",
            "title": "QRNG_Node_ComfyUI",
            "id": "qrng",
            "reference": "https://github.com/Onierous/QRNG_Node_ComfyUI",
            "files": [
                "https://github.com/Onierous/QRNG_Node_ComfyUI/raw/main/qrng_node.py"
            ],
            "install_type": "copy",
            "description": "Nodes: QRNG Node CSV. A node that takes in an array of random numbers from the ANU QRNG API and stores them locally for generating quantum random number noise_seeds in ComfyUI"
        },
        {
            "author": "ntdviet",
            "title": "ntdviet/comfyui-ext",
            "reference": "https://github.com/ntdviet/comfyui-ext",
            "files": [
                "https://github.com/ntdviet/comfyui-ext/raw/main/custom_nodes/gcLatentTunnel/gcLatentTunnel.py"
            ],
            "install_type": "copy",
            "description": "Nodes:LatentGarbageCollector. This ComfyUI custom node flushes the GPU cache and empty cuda interprocess memory. It's helpfull for low memory environment such as the free Google Colab, especially when the workflow VAE decode latents of the size above 1500x1500."
        },
        {
            "author": "alkemann",
            "title": "alkemann nodes",
            "id": "alkemann",
            "reference": "https://gist.github.com/alkemann/7361b8eb966f29c8238fd323409efb68",
            "files": [
                "https://gist.github.com/alkemann/7361b8eb966f29c8238fd323409efb68/raw/f9605be0b38d38d3e3a2988f89248ff557010076/alkemann.py"
            ],
            "install_type": "copy",
            "description": "Nodes:Int to Text, Seed With Text, Save A1 Image."
        },
        {
            "author": "catscandrive",
            "title": "Image loader with subfolders",
            "id": "imgsubfolders",
            "reference": "https://github.com/catscandrive/comfyui-imagesubfolders",
            "files": [
                "https://github.com/catscandrive/comfyui-imagesubfolders/raw/main/loadImageWithSubfolders.py"
            ],
            "install_type": "copy",
            "description": "Adds an Image Loader node that also shows images in subfolders of the default input directory"
        },
        {
            "author": "Smuzzies",
            "title": "Chatbox Overlay node for ComfyUI",
            "id": "chatbox-overlay",
            "reference": "https://github.com/Smuzzies/comfyui_chatbox_overlay",
            "files": [
                "https://github.com/Smuzzies/comfyui_chatbox_overlay/raw/main/chatbox_overlay.py"
            ],
            "install_type": "copy",
            "description": "Nodes: Chatbox Overlay. Custom node for ComfyUI to add a text box over a processed image before save node."
        },
        {
            "author": "CaptainGrock",
            "title": "ComfyUIInvisibleWatermark",
            "id": "invisible-watermark-grock",
            "reference": "https://github.com/CaptainGrock/ComfyUIInvisibleWatermark",
            "files": [
                "https://github.com/CaptainGrock/ComfyUIInvisibleWatermark/raw/main/Invisible%20Watermark.py"
            ],
            "install_type": "copy",
            "description": "Nodes:Apply Invisible Watermark, Extract Watermark. Adds up to 12 characters encoded into an image that can be extracted."
        },
        {
            "author": "LZC",
            "title": "Hayo comfyui nodes",
            "id": "lzcnodes",
            "reference": "https://github.com/1shadow1/hayo_comfyui_nodes",
            "files": [
                "https://github.com/1shadow1/hayo_comfyui_nodes/raw/main/LZCNodes.py"
            ],
            "install_type": "copy",
            "description": "Nodes:tensor_trans_pil, Make Transparent mask, MergeImages, words_generatee, load_PIL image"
        },
        {
            "author": "underclockeddev",
            "title": "Preview Subselection Node for ComfyUI",
            "id": "preview-subselection",
            "reference": "https://github.com/underclockeddev/ComfyUI-PreviewSubselection-Node",
            "files": [
                "https://github.com/underclockeddev/ComfyUI-PreviewSubselection-Node/raw/master/preview_subselection.py"
            ],
            "install_type": "copy",
            "description": "A node which takes in x, y, width, height, total width, and total height, in order to accurately represent the area of an image which is covered by area-based conditioning."
        },
        {
            "author": "underclockeddev",
            "title": "BrevImage",
            "id": "brevimage",
            "reference": "https://github.com/bkunbargi/BrevImage",
            "files": [
                "https://github.com/bkunbargi/BrevImage/raw/main/BrevLoadImage.py"
            ],
            "install_type": "copy",
            "description": "Nodes:BrevImage. ComfyUI Load Image From URL"
        },
        {
            "author": "jw782cn",
            "title": "ComfyUI-Catcat",
            "id": "catcat",
            "reference": "https://github.com/jw782cn/ComfyUI-Catcat",
            "files": [
                "https://github.com/jw782cn/ComfyUI-Catcat"
            ],
            "install_type": "copy",
            "description": "Extension to show random cat GIFs while queueing prompt."
        },
        {
            "author": "barckley75",
            "title": "comfyUI_DaVinciResolve",
            "reference": "https://github.com/barckley75/comfyUI_DaVinciResolve",
            "files": [
                "https://github.com/barckley75/comfyUI_DaVinciResolve/raw/main/custom_nodes/node_text_to_speech.py",
                "https://github.com/barckley75/comfyUI_DaVinciResolve/raw/main/custom_nodes/nodes_phy_3_contitioning.py",
                "https://github.com/barckley75/comfyUI_DaVinciResolve/raw/main/custom_nodes/save_audio_to_davinci.py",
                "https://github.com/barckley75/comfyUI_DaVinciResolve/raw/main/custom_nodes/save_image_to_davinci.py"
            ],
            "install_type": "copy",
            "description": "Nodes:TextToSpeech, phy_3_conditioning, SaveAudioToDaVinci, SaveImageToDaVinci.\nNOTE:In order to use DaVinci node you must have DaVinci Resolve Studio connected to the API. For more information check the help seciton in DaVinci Resolve Studio HELP>DOCUMENTATION>DEVELOPER. It will open a folder, search for scripting and the for README.txt file, the API documentation."
        },
        {
            "author": "Limbicnation",
            "title": "ComfyUIDepthEstimation",
            "reference": "https://github.com/Limbicnation/ComfyUIDepthEstimation",
            "files": [
                "https://github.com/Limbicnation/ComfyUIDepthEstimation/raw/main/depth_estimation_node.py"
            ],
            "pip": ["transformers"],
            "install_type": "copy",
            "description": "A custom depth estimation node for ComfyUI using transformer models. It integrates depth estimation with automatic gamma correction, contrast adjustment, and edge detection, based on the [a/TransformDepth](https://mirror.ghproxy.com/https://github.com/Limbicnation/TransformDepth) repository."
        },
        {
            "author": "seghier",
            "title": "ComfyUI_LibreTranslate",
            "reference": "https://github.com/seghier/ComfyUI_LibreTranslate",
            "files": [
                "https://github.com/seghier/ComfyUI_LibreTranslate/raw/main/translate_node.py"
            ],
            "install_type": "copy",
            "description": "Use LibreTranslation in ComfyUI [a/https://mirror.ghproxy.com/https://github.com/LibreTranslate/LibreTranslate](https://mirror.ghproxy.com/https://github.com/LibreTranslate/LibreTranslate)"
        },
        {
            "author": "ImmortalPie",
            "title": "PonySwitch Node",
            "reference": "https://github.com/ImmortalPie/ComfyUI-PonySwitch",
            "files": [
                "https://github.com/ImmortalPie/ComfyUI-PonySwitch/raw/main/PonySwitch.py"
            ],
            "install_type": "copy",
            "description": "The PonySwitch node is a custom node for ComfyUI that modifies prompts based on a toggle switch and adds configurable pony tags."
        },
        {
            "author": "ultimatech-cn",
            "title": "FaceSimilarity",
            "reference": "https://github.com/ultimatech-cn/FaceSimilarity",
            "files": [
                "https://github.com/ultimatech-cn/FaceSimilarity/raw/main/faceSimilarity.py"
            ],
            "install_type": "copy",
            "description": "A ComfyUI custom node for face comparison. This node utilizes Face++'s facial recognition and comparison algorithms by directly calling the Face++ API. Its usage in the workflow is as follows:"
        },
        {
            "author": "folkghost",
            "title": "CSV Search Node",
            "reference": "https://github.com/folkghost/comfyui_search_csv",
            "files": [
                "https://github.com/folkghost/comfyui_search_csv/raw/main/search_csv_node.py"
            ],
            "install_type": "copy",
            "description": "This repository contains a custom node for ComfyUI that allows searching for a keyword in the first column of a CSV file and returning a value from a specified column in that row. The node is designed to be modular and fit within the node-based workflow of ComfyUI."
        },
        {
            "author": "BobsBlazed",
            "title": "Bobs_FLUX_SDXL_Latent_Optimizer",
            "reference": "https://github.com/BobsBlazed/Bobs_FLUX_SDXL_Latent_Optimizer",
            "files": [
                "https://github.com/BobsBlazed/Bobs_Latent_Optimizer/raw/refs/heads/main/Bobs_Latent_Optimizer.py"
            ],
            "install_type": "copy",
            "description": "This custom node for ComfyUI is designed to optimize latent generation for use with both FLUX and SDXL modes. It provides flexible control over aspect ratios, megapixel sizes, and upscale factors, allowing users to dynamically create latents that fit specific tiling and resolution needs."
        },


        {
            "author": "theally",
            "title": "TheAlly's Custom Nodes",
            "id": "ally",
            "reference": "https://civitai.com/models/19625?modelVersionId=23296",
            "files": [
                "https://civitai.com/api/download/models/25114",
                "https://civitai.com/api/download/models/24679",
                "https://civitai.com/api/download/models/24154",
                "https://civitai.com/api/download/models/23884",
                "https://civitai.com/api/download/models/23649",
                "https://civitai.com/api/download/models/23467",
                "https://civitai.com/api/download/models/23296"
            ],
            "install_type": "unzip",
            "description": "Custom nodes for ComfyUI by TheAlly."
        },
        {
            "author": "xss",
            "title": "Custom Nodes by xss",
            "id": "xss",
            "reference": "https://civitai.com/models/24869/comfyui-custom-nodes-by-xss",
            "files": [
                "https://civitai.com/api/download/models/32717",
                "https://civitai.com/api/download/models/47776",
                "https://civitai.com/api/download/models/29772",
                "https://civitai.com/api/download/models/31618",
                "https://civitai.com/api/download/models/31591",
                "https://civitai.com/api/download/models/29773",
                "https://civitai.com/api/download/models/29774",
                "https://civitai.com/api/download/models/29755",
                "https://civitai.com/api/download/models/29750"
            ],
            "install_type": "unzip",
            "description": "Various image processing nodes."
        },
        {
            "author": "aimingfail",
            "title": "Image2Halftone Node for ComfyUI",
            "id": "img2halftone",
            "reference": "https://civitai.com/models/143293/image2halftone-node-for-comfyui",
            "files": [
                "https://civitai.com/api/download/models/158997"
            ],
            "install_type": "unzip",
            "description": "This is a node to convert an image into a CMYK Halftone dot image."
        }
    ]
}`